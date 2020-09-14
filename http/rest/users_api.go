package rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mendezdev/expenses_oauth-api/domain/users"
)

const (
	API_BASE_PATH = "http://localhost:8080"
)

var (
	client                             = resty.New()
	RestUsersApi restUsersApiInterface = &restUsersApi{}
)

type restUsersApi struct{}

type restUsersApiInterface interface {
	Login(users.UserLoginRequest) (*users.UserApiResponse, error)
}

func (r *restUsersApi) Login(ur users.UserLoginRequest) (*users.UserApiResponse, error) {
	resp, err := client.R().
		SetBody(ur).
		SetResult(users.UserApiResponse{}).
		Post(API_BASE_PATH + "/users/login")

	if err != nil {
		fmt.Println("REST USERS API ERROR:", err.Error())
		return nil, err
	}

	userApiResponse := resp.Result().(*users.UserApiResponse)
	return userApiResponse, nil
}
