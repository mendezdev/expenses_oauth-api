package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mendezdev/expenses_oauth-api/domain/access_token"
	"github.com/mendezdev/expenses_oauth-api/domain/users"
	"github.com/mendezdev/expenses_oauth-api/http/rest"
)

var (
	UsersService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateToken(users.UserLoginRequest) (*access_token.AccessToken, error)
}

func (s *userService) CreateToken(ur users.UserLoginRequest) (*access_token.AccessToken, error) {
	var err error

	atSecret := "AT_SECRET" //implment this os.Getenv("ACESS_TOKEN_SECRET")

	userReponse, userAPIErr := rest.RestUsersAPI.Login(ur)
	if userAPIErr != nil {
		return nil, userAPIErr
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userReponse.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(atSecret))
	if err != nil {
		return nil, err
	}

	return &access_token.AccessToken{AccessToken: token}, nil
}
