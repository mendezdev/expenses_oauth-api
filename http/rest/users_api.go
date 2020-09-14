package rest

import (
	"fmt"
	"net/http"

	"github.com/mendezdev/expenses_oauth-api/domain/users"
	"github.com/mendezdev/expenses_oauth-api/utils/api_errors"
)

const (
	// APIUsersBasePath base url for expenses_users-api
	APIUsersBasePath = "http://localhost:8080"
)

var (
	// RestUsersAPI variable that contains the logic to makes externals api calls
	RestUsersAPI restUsersApiInterface = &restUsersApi{}
)

type restUsersApi struct{}

type restUsersApiInterface interface {
	Login(users.UserLoginRequest) (*users.UserApiResponse, api_errors.RestErr)
}

func (r *restUsersApi) Login(ur users.UserLoginRequest) (*users.UserApiResponse, api_errors.RestErr) {
	resp, err := Client.R().
		SetBody(ur).
		SetResult(users.UserApiResponse{}).
		Post(APIUsersBasePath + "/users/login")

	//TODO: parse err to api_errors.RestErr
	if err != nil {
		fmt.Println("REST USERS API ERROR:", err.Error())
		return nil, api_errors.NewRestError(err.Error(), http.StatusServiceUnavailable, "service_unavailable", nil)
	}

	userApiResponse := resp.Result().(*users.UserApiResponse)
	return userApiResponse, nil
}
