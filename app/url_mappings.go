package app

import (
	"github.com/mendezdev/expenses_oauth-api/controllers/access_token"
	"github.com/mendezdev/expenses_oauth-api/controllers/ping"
)

func mapUrls() {
	// Ping controller
	router.GET("/ping", ping.Ping)

	// AccessToken controller
	router.POST("/oauth/access_token", access_token.Create)
}
