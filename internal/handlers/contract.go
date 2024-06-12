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
)

type ContractHandler struct {
	contractRepo      repository.ContractInterface
	contractEventRepo repository.ContractEventInterface
	eventRepo         repository.EventInterface
	eventLogRepo      repository.EventLogInterface
}

func NewContractHandler(contractRepo repository.ContractInterface,
	contractEventRepo repository.ContractEventInterface,
	eventRepo repository.EventInterface,
	eventLogRepo repository.EventLogInterface,
) *ContractHandler {
	return &ContractHandler{
		contractRepo:      contractRepo,
		contractEventRepo: contractEventRepo,
		eventRepo:         eventRepo,
		eventLogRepo:      eventLogRepo,
	}
}

type ContractEventRequest struct {
	Address string `json:"address" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Event   struct {
		Signature   string `json:"signature" validate:"required"`
		Name        string `json:"event_name" validate:"required"`
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
			ID:        ID,
			Name:      input.Name,
			Address:   input.Address,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err = h.contractRepo.Create(contract); err != nil {
			helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
			return
		}

		//TODO: check if event already exists or create

		// TODO: create contract event, it's a new contact registration so the event cannot exist 
		// create contract event if not exist
		// ID, err = uuid.NewV7()
		// if err != nil {
		// 	helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
		// 	return
		// }
		// contract_event := &models.ContractEvent{
		// 	ID:ID,
		// 	ContractID: contract.ID,
		// 	Active: true,
		// 	CreatedAt: time.Now(),
		// }
	}

	helpers.ReturnJSON(c, "OK", contract, http.StatusOK)
}
