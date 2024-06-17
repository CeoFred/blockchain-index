package routes

import (
	"github.com/CeoFred/gin-boilerplate/constants"
	"github.com/CeoFred/gin-boilerplate/internal/handlers"
	"github.com/CeoFred/gin-boilerplate/internal/service"

	// "github.com/CeoFred/gin-boilerplate/internal/middleware"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/validators"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterContractRoute(router *gin.RouterGroup, db *gorm.DB) {
	contract := router.Group("contract")
	env := constants.New()

	contractRepository := repository.NewContractRepository(db)
	contractEventRepository := repository.NewContractEventRepository(db)
	eventRepository := repository.NewEventRepository(db)
	eventLogRepository := repository.NewEventLogRepository(db)
	blockchainService, err := service.NewBlockchainService(env.RPC)

	if err != nil {
		panic(err)
	}

	handler := handlers.NewContractHandler(contractRepository, contractEventRepository, eventRepository, eventLogRepository, blockchainService)

	contract.POST("", validators.NewContract, handler.NewContractIndex)
}
