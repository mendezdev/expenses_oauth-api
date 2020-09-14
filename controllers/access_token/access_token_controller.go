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

	//TODO: go to expenses_users-api an make POST to /users/login to check credentials
	token, err := services.UsersService.CreateToken(userLoginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, token)
}
