package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.RouterGroup, db *gorm.DB) {

	RegisterContractRoute(r, db)
	RegisterEventLogRoute(r, db)

}
