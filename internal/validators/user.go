package validators

import (
	"net/http"

	"github.com/CeoFred/gin-boilerplate/internal/handlers"
	"github.com/CeoFred/gin-boilerplate/internal/helpers"
	"github.com/CeoFred/gin-boilerplate/validator"

	"github.com/gin-gonic/gin"
)



func ValidateAccountResetScheme(c *gin.Context) {
	var body helpers.AccountReset
	bindAndValidate(c, &body)
	c.Set("validatedRequestBody", body)
	c.Next()
}

func ValidateOTPVerifySchema(c *gin.Context) {
	var body helpers.OtpVerify
	bindAndValidate(c, &body)
	c.Set("validatedRequestBody", body)
	c.Next()
}



func ValidateUpdateUserProfile(c *gin.Context) {
	var body handlers.UpdateUserProfileInput
	bindAndValidate(c, &body)
	c.Set("validatedRequestBody", body)
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
