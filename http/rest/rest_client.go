package rest

import "github.com/go-resty/resty/v2"

var (
	// Client this is the client to make request to and external APIs
	Client *resty.Client = resty.New()
)
