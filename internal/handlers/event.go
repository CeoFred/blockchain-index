package handlers

import (
	"fmt"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/gofrs/uuid"

	"github.com/CeoFred/gin-boilerplate/internal/helpers"
	// "github.com/CeoFred/gin-boilerplate/internal/models"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/service"
)

type EventLogHandler struct {
	contractRepo      repository.ContractInterface
	contractEventRepo repository.ContractEventInterface
	eventRepo         repository.EventInterface
	eventLogRepo      repository.EventLogInterface
	blockchainService *service.BlockchainService
}

func NewEventLogHandler(contractRepo repository.ContractInterface,
	contractEventRepo repository.ContractEventInterface,
	eventRepo repository.EventInterface,
	eventLogRepo repository.EventLogInterface,
	blockchainService *service.BlockchainService,
) *EventLogHandler {
	return &EventLogHandler{
		contractRepo:      contractRepo,
		contractEventRepo: contractEventRepo,
		eventRepo:         eventRepo,
		eventLogRepo:      eventLogRepo,
		blockchainService: blockchainService,
	}
}

func (h *EventLogHandler) ContractEvents(c *gin.Context) {

	address := c.Param("address")

	event_name := c.Query("event_name")
	block_number := c.Query("block_number")
	from_block := c.Query("from_block")
	to_block := c.Query("to_block")



	if !h.blockchainService.IsContract(address) {
		helpers.ReturnError(c, "Blockchain service error", fmt.Errorf("address is not a contract"), http.StatusBadRequest)
		return
	}

	contract, err := h.contractRepo.QueryWithArgs("select * from contracts where address = ?", address)

	if err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
		return
	}

	if contract == nil {
		helpers.ReturnError(c, "Something went wrong", fmt.Errorf("contract address %s not found", address), http.StatusNotFound)
		return
	}

	query := fmt.Sprintf("select * from event_logs where contract_id = '%s' AND contract_address = '%s'",contract.ID, address)

	if event_name != "" {
		query += fmt.Sprintf(" AND event_name = '%s'", event_name)
	}

	if block_number != "" {
		query += fmt.Sprintf(" AND block_number = '%s'", block_number)
	}

	if from_block != "" {
		query += fmt.Sprintf(" AND block_number >= '%s'", from_block)
	}

	if to_block != "" {
		query += fmt.Sprintf(" AND block_number <= '%s'", to_block)
	}
	
	eventLogs, err := h.eventLogRepo.QueryRecordsWithArgs(query)

	if err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusInternalServerError)
		return
	}

	helpers.ReturnJSON(c, "OK", eventLogs, http.StatusOK)
}
