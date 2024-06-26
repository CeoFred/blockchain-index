package validators

import (
	"github.com/CeoFred/gin-boilerplate/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NewFaucet(c *gin.Context) {
	var body (handlers.NewFaucet)
	bindAndValidate(c, &body)
	c.Set("validatedRequestBody", body)
	c.Next()
}
