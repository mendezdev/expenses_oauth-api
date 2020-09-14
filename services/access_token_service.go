package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mendezdev/expenses_oauth-api/domain/users"
	"github.com/mendezdev/expenses_oauth-api/http/rest"
)

var (
	UsersService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateToken(users.UserLoginRequest) (string, error)
}

func (s *userService) CreateToken(ur users.UserLoginRequest) (string, error) {
	var err error

	atSecret := "AT_SECRET" //implment this os.Getenv("ACESS_TOKEN_SECRET")

	userReponse, userApiErr := rest.RestUsersApi.Login(ur)
	if userApiErr != nil {
		return "", userApiErr
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userReponse.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(atSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}
