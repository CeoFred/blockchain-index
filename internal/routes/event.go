package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/CeoFred/gin-boilerplate/constants"
	"github.com/CeoFred/gin-boilerplate/internal/handlers"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/service"
)

func RegisterEventLogRoute(router *gin.RouterGroup, db *gorm.DB) {
	env := constants.New()

	contractRepository := repository.NewContractRepository(db)
	contractEventRepository := repository.NewContractEventRepository(db)
	eventRepository := repository.NewEventRepository(db)
	eventLogRepository := repository.NewEventLogRepository(db)

	blockchainService, err := service.NewBlockchainService(env.RPC)
	if err != nil {
		log.Fatalf("Failed to initialize blockchain service with URL %s: %v", env.RPC, err)
		return
	}

	eventHandler := handlers.NewEventLogHandler(contractRepository, contractEventRepository, eventRepository, eventLogRepository, blockchainService)

	contract := router.Group("/logs")
	{
		contract.GET("/:address", eventHandler.ContractEvents)
	}

}
