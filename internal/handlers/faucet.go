package handlers

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"github.com/CeoFred/gin-boilerplate/internal/helpers"
	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/service"
)

type FaucetHandler struct {
	contractRepo       repository.ContractInterface
	contractEventRepo  repository.ContractEventInterface
	eventRepo          repository.EventInterface
	eventLogRepo       repository.EventLogInterface
	blockchainService  *service.BlockchainService
	faucetRepo         repository.FaucetInterface
	faucetTransferRepo repository.FaucetTransferInterface
	userActionRepo     repository.UserActionInterface
	userRepo           repository.UserInterface
}

func NewFaucetHandler(contractRepo repository.ContractInterface,
	contractEventRepo repository.ContractEventInterface,
	eventRepo repository.EventInterface,
	eventLogRepo repository.EventLogInterface,
	blockchainService *service.BlockchainService,
	faucetRepo repository.FaucetInterface,
	faucetTransferRepo repository.FaucetTransferInterface,
	userActionRepo repository.UserActionInterface,
	userRepo repository.UserInterface,
) *FaucetHandler {
	return &FaucetHandler{
		contractRepo:       contractRepo,
		contractEventRepo:  contractEventRepo,
		eventRepo:          eventRepo,
		eventLogRepo:       eventLogRepo,
		blockchainService:  blockchainService,
		faucetRepo:         faucetRepo,
		faucetTransferRepo: faucetTransferRepo,
		userActionRepo:     userActionRepo,
		userRepo:           userRepo,
	}
}

type NewFaucet struct {
	Asset      models.FaucetAsset `json:"asset" validate:"required"`
	Address    string             `json:"address" validate:"required,eth_addr"`
	StartBlock uint               `json:"start_block" validate:"required"`
}

func (h *FaucetHandler) NewFaucet(c *gin.Context) {

	var input NewFaucet
	validatedReqBody, exists := c.Get("validatedRequestBody")

	if !exists {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf(helpers.INVALID_REQUEST_BODY), http.StatusBadRequest)
		return
	}

	input, ok := validatedReqBody.(NewFaucet)
	if !ok {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf(helpers.REQUEST_BODY_PARSE_ERROR), http.StatusBadRequest)
		return
	}

	switch input.Asset {
	case models.FaucetAssetRWA:
	case models.FaucetAssetUSDT:
	default:
		helpers.ReturnError(c, "faucet service error", fmt.Errorf("asset type not supported: %s", input.Asset), http.StatusBadRequest)
		return
	}
	faucet, err := h.faucetRepo.QueryWithArgs("select * from faucets where address = ?", input.Address)

	if err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
		return
	}

	if faucet == nil {
		ID, err := uuid.NewV7()
		if err != nil {
			helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
			return
		}
		faucet = &models.Faucet{
			ID:         ID,
			Asset:      input.Asset,
			Address:    input.Address,
			StartBlock: input.StartBlock,
			EndBlock:   input.StartBlock,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err = h.faucetRepo.Create(faucet); err != nil {
			helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
			return
		}
	}

	helpers.ReturnJSON(c, "Faucet added", faucet, http.StatusCreated)
}

func (h *FaucetHandler) ProcessBlock() error {

	return nil
}

func (h *FaucetHandler) PoolFaucetTransfer() error {
	faucets, err := h.faucetRepo.QueryRecordsWithArgs("SELECT * FROM faucets WHERE  deleted_at IS NULL")
	if err != nil {
		panic(err)
	}

	latestBlockNumber, err := h.blockchainService.GetLatestBlockNumber()
	if err != nil {
		return fmt.Errorf("failed to get latest block number: %w", err)
	}
	latestBlock := uint(latestBlockNumber.Uint64())

	for _, f := range faucets {
		startBlock := f.EndBlock
		address := f.Address

		if startBlock == latestBlock {
			continue
		}

		for blockNumber := startBlock; blockNumber <= latestBlock; blockNumber++ {
			fmt.Println("block number \n", blockNumber)
			logs, err := h.blockchainService.BlockTransaction(blockNumber)
			if err != nil {
				return fmt.Errorf("failed to poll events for contract %s: %w", address, err)
			}

			block, err := h.blockchainService.BlockByNumber(uint64(blockNumber))
			if err != nil {
				return fmt.Errorf("failed to get block %d: %w", blockNumber, err)
			}

			for _, tx := range logs {
				if tx.To() != nil && tx.Value().Cmp(big.NewInt(0)) > 0 && len(tx.Data()) == 0 {
					value := tx.Value().String()
					hash := tx.Hash().Hex()
					recipient := tx.To().Hex()

					log.Printf("sent to: %s, amount: %s, hash: %s \n", recipient, value, hash)

					ID, err := uuid.NewV7()
					if err != nil {
						panic(err)
					}
					timestamp := time.Unix(int64(block.Time()), 0)

					// Insert the transaction details into the database
					faucetTransfer := &models.FaucetTransfer{
						ID:              ID,
						FaucetID:        f.ID,
						Recipient:       recipient,
						Value:           uint(tx.Value().Uint64()),
						TransactionHash: hash,
						BlockNumber:     fmt.Sprintf("%d", blockNumber),
						Timestamp:       timestamp,
						CreatedAt:       time.Now(),
					}

					if err := h.faucetTransferRepo.Create(faucetTransfer); err != nil {
						log.Printf("failed to save faucet transfer: %v", err)
					}

					// create user
					ID, err = uuid.NewV7()
					if err != nil {
						panic(err)
					}

					user := &models.User{
						Address:   faucetTransfer.Recipient,
						ID:        ID,
						CreatedAt: time.Now(),
					}
					user, err = h.userRepo.Create(user)
					if err != nil {
						panic(err)
					}

					ID, err = uuid.NewV7()
					if err != nil {
						panic(err)
					}
					userAction := &models.UserAction{
						ID:              ID,
						TransactionHash: hash,
						Point:           1,
						Action:          "RWA Faucet",
						CreatedAt:       time.Now(),
						EventLogID:      faucetTransfer.ID,
						UserID:          user.ID,
					}

					err = h.userActionRepo.Create(userAction)
					if err != nil {
						panic(err)
					}

					time.Sleep(1000 * time.Millisecond)
				}
			}

			// update faucet end block number
			f.EndBlock = blockNumber + 1
			_, err = h.faucetRepo.Save(f)
			if err != nil {
				panic(err)
			}
			log.Printf("done processing transaction in  block %d \n", blockNumber)

		}
	}
	return nil
}
