package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is the endpoint to check the status of the API
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
