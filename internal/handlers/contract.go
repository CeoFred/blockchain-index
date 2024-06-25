package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"github.com/CeoFred/gin-boilerplate/internal/helpers"
	"github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/service"
)

type ContractHandler struct {
	contractRepo      repository.ContractInterface
	contractEventRepo repository.ContractEventInterface
	eventRepo         repository.EventInterface
	eventLogRepo      repository.EventLogInterface
	blockchainService *service.BlockchainService
}

func NewContractHandler(contractRepo repository.ContractInterface,
	contractEventRepo repository.ContractEventInterface,
	eventRepo repository.EventInterface,
	eventLogRepo repository.EventLogInterface,
	blockchainService *service.BlockchainService,
) *ContractHandler {
	return &ContractHandler{
		contractRepo:      contractRepo,
		contractEventRepo: contractEventRepo,
		eventRepo:         eventRepo,
		eventLogRepo:      eventLogRepo,
		blockchainService: blockchainService,
	}
}

type ContractEventRequest struct {
	Address    string                  `json:"address" validate:"required,eth_addr"`
	Name       string                  `json:"name" validate:"required"`
	StartBlock uint                    `json:"start_block" validate:"required"`
	Category   models.ContractCategory `json:"category" validate:"required"`
	Event      []struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	} `json:"event" validate:"required"`
}

func (h *ContractHandler) NewContractIndex(c *gin.Context) {

	var input ContractEventRequest
	validatedReqBody, exists := c.Get("validatedRequestBody")

	if !exists {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf(helpers.INVALID_REQUEST_BODY), http.StatusBadRequest)
		return
	}

	input, ok := validatedReqBody.(ContractEventRequest)
	if !ok {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf(helpers.REQUEST_BODY_PARSE_ERROR), http.StatusBadRequest)
		return
	}

	switch input.Category {
	case models.ContractCategoryERC20:
	case models.ContractCategorySwap:
	default:
		helpers.ReturnError(c, "Blockchain service error", fmt.Errorf("invalid contract category"), http.StatusBadRequest)
		return
	}

	if !h.blockchainService.IsContract(input.Address) {
		helpers.ReturnError(c, "Blockchain service error", fmt.Errorf("address is not a contract"), http.StatusBadRequest)
		return
	}

	contract, err := h.contractRepo.QueryWithArgs("select * from contracts where address = ?", input.Address)

	if err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
		return
	}

	if contract == nil {
		ID, err := uuid.NewV7()
		if err != nil {
			helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
			return
		}
		contract = &models.Contract{
			ID:         ID,
			Name:       input.Name,
			Address:    input.Address,
			StartBlock: input.StartBlock,
			EndBlock:   input.StartBlock,
			Category:   input.Category,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err = h.contractRepo.Create(contract); err != nil {
			helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
			return
		}
	}

	eventSignature := map[string]string{
		"Transfer":             "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		"Approval":             "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925",
		"OwnershipTransferred": "0x8c5be1e5ebec7d5bd14f714f8100de8fa9c6a09d2323b9cfa7f5f2fa1b0cd031",
		"ApprovalForAll":       "0x17307eab39c3af08c0e4369c0d0b8d6a2a97c49ab98d7d621e49e4c0d648d1b1",
		"Mint":                 "0x7a1b0f670e6e40c8f1e61bdbfd4f0b7886b358c084c5d2c726861b5c0a4e3e3d",
		"Burn":                 "0xcc16f5dbb2b993aba1924902705498a8a94a801efed80f968f275b1c338405a9",

		"Swap": "0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67",
	}

	// create all events
	events := []*models.Event{}
	for i := 0; i < len(input.Event); i++ {
		event := input.Event[i]

		if eventSignature[event.Name] != "" {
			ID, err := uuid.NewV7()
			if err != nil {
				helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
				return
			}
			e := &models.Event{
				ID:          ID,
				Name:        event.Name,
				Signature:   eventSignature[event.Name],
				Description: event.Description,
			}
			events = append(events, e)
		}

	}

	if len(events) > 0 {
		events, err = h.eventRepo.BatchInsert(events) // updates existing records
		if err != nil {
			helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
			return
		}

		contractEvents := []*models.ContractEvent{}
		for i := 0; i < len(events); i++ {
			event := events[i]
			ID, err := uuid.NewV7()
			if err != nil {
				helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
				return
			}
			contractEvent := &models.ContractEvent{
				ID:         ID,
				ContractID: contract.ID,
				EventID:    event.ID,
				Active:     true,
				CreatedAt:  time.Now(),
			}

			contractEvents = append(contractEvents, contractEvent)
		}
		err = h.contractEventRepo.BatchInsert(contractEvents)
		if err != nil {
			helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
			return
		}
	}

	helpers.ReturnJSON(c, "OK", contract, http.StatusOK)
}
