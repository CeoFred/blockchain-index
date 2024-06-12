package handlers

import (
	// "context"
	// "fmt"
	// "mime/multipart"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"

	// "github.com/CeoFred/gin-boilerplate/constants"
	"github.com/CeoFred/gin-boilerplate/internal/helpers"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
)

var ()

type ContractHandler struct {
	contractRepo repository.ContractInterface
}

func NewContractHandler(contractRepo repository.ContractInterface,
) *ContractHandler {
	return &ContractHandler{
		contractRepo: contractRepo,
	}
}

type ContractEventRequest struct {
	Address string `json:"address" validate:"address"`
	Name    string `json:"name" validate:"required"`
	Event   struct {
		Signature   string `json:"event_signature"`
		Name        string `json:"event_name"`
		Description string `json:"event_description"`
	} `json:"event" validate:"required"`
}

func (h *ContractHandler) NewContractIndex(c *gin.Context) {
	helpers.ReturnJSON(c, "OK", nil, http.StatusOK)
}
