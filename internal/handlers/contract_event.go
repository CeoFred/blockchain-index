package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/CeoFred/gin-boilerplate/internal/helpers"
	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/service"
	"github.com/gofrs/uuid"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type ContractEventHandler struct {
	contractRepo      repository.ContractInterface
	contractEventRepo repository.ContractEventInterface
	eventRepo         repository.EventInterface
	eventLogRepo      repository.EventLogInterface
	userRepo          repository.UserInterface
	userActionRepo    repository.UserActionInterface
	blockchainService *service.BlockchainService
}

func NewContractEventHandler(contractRepo repository.ContractInterface,
	contractEventRepo repository.ContractEventInterface,
	eventRepo repository.EventInterface,
	eventLogRepo repository.EventLogInterface,
	blockchainService *service.BlockchainService,
	userRepo repository.UserInterface,
	userActionRepo repository.UserActionInterface,
) *ContractEventHandler {
	return &ContractEventHandler{
		contractRepo:      contractRepo,
		contractEventRepo: contractEventRepo,
		eventRepo:         eventRepo,
		eventLogRepo:      eventLogRepo,
		userRepo:          userRepo,
		userActionRepo:    userActionRepo,
		blockchainService: blockchainService,
	}
}

func (h *ContractEventHandler) StartPolling() error {
	// get active contract events
	contract_events, err := h.contractEventRepo.QueryRecordsWithArgs("SELECT * FROM contract_events WHERE active = ?", models.ContractEventStatusActive)

	if err != nil {
		panic(err)
	}

	contractEventMap := map[string][]*models.ContractEvent{}
	contractLastBlock := map[string]uint{}
	contractMap := map[string]*models.Contract{}

	for _, v := range contract_events {
		contract_address := v.Contract.Address
		contractEventMap[contract_address] = append(contractEventMap[contract_address], v)
		contractLastBlock[contract_address] = v.Contract.EndBlock
		contractMap[contract_address] = v.Contract
	}

	// get latest block for contract
	latestBlockNumber, err := h.blockchainService.GetLatestBlockNumber()
	if err != nil {
		return fmt.Errorf("failed to get latest block number: %w", err)
	}
	latestBlock := uint(latestBlockNumber.Uint64())

	// Iterate over each contract and poll events from the next block
	for contractAddress, events := range contractEventMap {

		startBlock := contractLastBlock[contractAddress] + 1
		if startBlock > latestBlock {
			log.Printf("\n no new block, contract: %s \n", contractAddress)
			continue // No new blocks to poll
		}

		// Poll events from startBlock to latestBlock
		logs, err := h.blockchainService.QueryLogs(contractAddress, startBlock, latestBlock)

		if err != nil {
			return fmt.Errorf("failed to poll events for contract %s: %w", contractAddress, err)
		}

		logCount := len(logs)

		log.Printf("total logs found %d", logCount)
		var eventLogs []*models.EventLog
		var userEvents map[string][]*models.UserAction
		contract_category := contractMap[contractAddress].Category

		switch contract_category {
		case models.ContractCategoryERC20:
			eventLogs, userEvents = h.blockchainService.ProcessERCTokenLogs(logs, events)
		case models.ContractCategorySwap:
			eventLogs, userEvents = h.blockchainService.ProcessPoolSwapLogs(logs, events)
		case models.ContractCategoryBridge:
			eventLogs, userEvents = h.blockchainService.ProcessTokenBridge(logs, events)
		default:
			panic("unknown contract category")
		}

		eventLogs, err = h.eventLogRepo.BatchInsert(eventLogs)
		if err != nil {
			panic(err)
		}

		for address, events := range userEvents {
			// try creating a new user or updating
			ID, err := uuid.NewV7()
			if err != nil {
				return err
			}

			user := &models.User{
				ID:        ID,
				Address:   address,
				CreatedAt: time.Now(),
			}

			user, err = h.userRepo.Create(user)
			if err != nil {
				return err
			}

			actions := []*models.UserAction{}
			for _, user_action := range events {
				user_action.UserID = user.ID
				user_action.CreatedAt = time.Now()

				for _, log := range eventLogs {
					// get original event log id because it's possible that during insert the new event log id assigned at the blokchain service layer was ignored .. this is to avoid reference errors
					if log.TransactionHash == user_action.TransactionHash {
						user_action.EventLogID = log.ID
					}
				}
				actions = append(actions, user_action)
			}
			err = h.userActionRepo.BatchInsert(actions)
			if err != nil {
				return err
			}

		}

		// Update the last polled block for the contract
		contractLastBlock[contractAddress] = latestBlock
		contract, err := h.contractRepo.QueryWithArgs("select * from contracts where address = ?", contractAddress)
		if err != nil {
			panic(err)
		}

		if contract != nil {
			contract.UpdatedAt = time.Now()
			contract.EndBlock = latestBlock
			_, err = h.contractRepo.Save(contract)
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

func (h *ContractEventHandler) HandleEvent(v *models.ContractEvent) {
	log.Printf("contract: %s \n", v.ContractID)
	log.Printf("event: %s \n", v.Event.Name)

	logs := make(chan types.Log)
	err := h.blockchainService.Subscribe(v.Contract.Address, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

}

func (h *ContractEventHandler) GetContractEvents(c *gin.Context) {
	address := c.Param("address")

	if address == "" {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf("address is required"), http.StatusBadRequest)
		return
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	isAddress := re.MatchString(address)

	if !isAddress {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf("address is not a valid ethereum address"), http.StatusBadRequest)
		return
	}

	contract, err := h.contractRepo.QueryWithArgs("select * from contracts where address = ?", address)

	if err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
		return
	}

	if contract == nil {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf("contract with address %s not registered", address), http.StatusNotFound)
		return
	}

	// get contract events
	contract_events, err := h.contractEventRepo.Where("contract_id", contract.ID.String())

	if err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
		return
	}

	helpers.ReturnJSON(c, "success", contract_events, http.StatusOK)
}
