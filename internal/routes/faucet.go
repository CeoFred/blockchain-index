package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/CeoFred/gin-boilerplate/constants"
	"github.com/CeoFred/gin-boilerplate/internal/handlers"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/service"
	"github.com/CeoFred/gin-boilerplate/internal/validators"
)

func RegisterFaucetRoute(router *gin.RouterGroup, db *gorm.DB) {
	env := constants.New()

	contractRepository := repository.NewContractRepository(db)
	contractEventRepository := repository.NewContractEventRepository(db)
	eventRepository := repository.NewEventRepository(db)
	eventLogRepository := repository.NewEventLogRepository(db)
	faucetRepository := repository.NewFaucetRepository(db)
	faucetTransferRepository := repository.NewFaucetTransferRepository(db)
	userActionRepository := repository.NewUserActionRepository(db)
	userRepository := repository.NewUserRepository(db)

	blockchainService, err := service.NewBlockchainService(env.RPC)
	if err != nil {
		log.Fatalf("Failed to initialize blockchain service with URL %s: %v", env.RPC, err)
		return
	}

	handler := handlers.NewFaucetHandler(contractRepository, contractEventRepository, eventRepository, eventLogRepository, blockchainService, faucetRepository, faucetTransferRepository, userActionRepository, userRepository)

	faucet := router.Group("/faucet")
	{
		faucet.POST("", validators.NewFaucet, handler.NewFaucet)
	}

	// Start Event Subscription in a Goroutine
	go func() {
		if err := handler.PoolFaucetTransfer(); err != nil {
			log.Printf("Failed to poll faucet transfer events: %v", err)
		}
	}()

}
