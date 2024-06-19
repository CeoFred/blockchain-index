package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/CeoFred/gin-boilerplate/constants"
	"github.com/CeoFred/gin-boilerplate/internal/handlers"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/service"
	"github.com/CeoFred/gin-boilerplate/internal/validators"
)

func RegisterContractRoute(router *gin.RouterGroup, db *gorm.DB) {
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

	contractHandler := handlers.NewContractHandler(contractRepository, contractEventRepository, eventRepository, eventLogRepository, blockchainService)
	contractEventHandler := handlers.NewContractEventHandler(contractRepository, contractEventRepository, eventRepository, eventLogRepository, blockchainService)

	contract := router.Group("/contract")
	{
		contract.POST("", validators.NewContract, contractHandler.NewContractIndex)
	}

	contractEvent := router.Group("/contract-event")
	{
		contractEvent.GET("/:address", contractEventHandler.GetContractEvents)
	}

	// Start Event Subscription in a Goroutine
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if err := contractEventHandler.StartPolling(); err != nil {
				log.Printf("Failed to poll contract events: %v", err)
			}
		}
	}()
}
