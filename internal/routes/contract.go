package routes

import (
	"github.com/CeoFred/gin-boilerplate/internal/handlers"
	// "github.com/CeoFred/gin-boilerplate/internal/middleware"
	"github.com/CeoFred/gin-boilerplate/internal/repository"
	"github.com/CeoFred/gin-boilerplate/internal/validators"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterContractRoute(router *gin.RouterGroup, db *gorm.DB) {
	contract := router.Group("contract")

	contractRepository := repository.NewContractRepository(db)
	contractEventRepository := repository.NewContractEventRepository(db)

	handler := handlers.NewContractHandler(contractRepository,contractEventRepository)

	contract.POST("/", validators.NewContract, handler.NewContractIndex)
}
