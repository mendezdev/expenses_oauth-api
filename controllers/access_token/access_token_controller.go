package access_token

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create validate the user credentials and create AccessToken and RefreshToken
func Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "implement me!")
}
