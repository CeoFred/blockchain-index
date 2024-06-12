package validators

import (
	"net/http"

	"github.com/CeoFred/gin-boilerplate/internal/handlers"
	"github.com/CeoFred/gin-boilerplate/internal/helpers"
	"github.com/CeoFred/gin-boilerplate/validator"

	"github.com/gin-gonic/gin"
)

func NewContract(c *gin.Context) {
	var body handlers.ContractEventRequest
	bindAndValidate(c, &body)
	c.Set("validated_body", body)
	c.Next()
}

func bindAndValidate(c *gin.Context, body interface{}) {
	if err := c.ShouldBindJSON(body); err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusBadRequest)
		c.Abort()
		return
	}

	if err := validator.Validate(body); err != nil {
		helpers.ReturnError(c, "Something went wrong", err, http.StatusBadRequest)
		c.Abort()
		return
	}
}
