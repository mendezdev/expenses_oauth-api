package access_token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mendezdev/expenses_oauth-api/domain/users"
	"github.com/mendezdev/expenses_oauth-api/services"
)

// Create validate the user credentials and create AccessToken and RefreshToken
func Create(c *gin.Context) {
	var userLoginRequest users.UserLoginRequest
	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
		return
	}

	at, createTokenErr := services.AccessTokenService.CreateToken(userLoginRequest)
	if createTokenErr != nil {
		c.JSON(createTokenErr.Status(), createTokenErr)
		return
	}

	c.JSON(http.StatusCreated, at)
}
