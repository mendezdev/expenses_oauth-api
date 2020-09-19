package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mendezdev/expenses_oauth-api/domain/access_token"
	"github.com/mendezdev/expenses_oauth-api/domain/users"
	"github.com/mendezdev/expenses_oauth-api/http/rest"
	"github.com/mendezdev/expenses_oauth-api/utils/api_errors"
)

var (
	AccessTokenService accessTokenServiceInterface = &accessTokenService{}
)

type accessTokenService struct{}

type accessTokenServiceInterface interface {
	CreateToken(users.UserLoginRequest) (*access_token.AccessToken, error) // change error to api_errors.RestErr
	TokenValid(string) api_errors.RestErr
	ExtractTokenMetadata(string) (*access_token.AccessDetails, api_errors.RestErr)
}

func (s *accessTokenService) CreateToken(ur users.UserLoginRequest) (*access_token.AccessToken, error) {
	var err error
	atSecret := "AT_SECRET" //implment this os.Getenv("ACESS_TOKEN_SECRET")
	rtSecret := "RT_SECRET" //implment this os.Getenv("ACESS_TOKEN_SECRET")

	td := &access_token.TokenDetails{
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		AccessUuid:  uuid.New().String(),
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUuid: uuid.New().String(),
	}

	userResponse, userAPIErr := rest.RestUsersAPI.Login(ur)
	if userAPIErr != nil {
		return nil, userAPIErr
	}

	//Creating acccess token
	atClaims := buildAccessTokenClaims(userResponse.ID, td)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(atSecret))
	if err != nil {
		return nil, err
	}

	//Creating refresh token
	rtClaims := buildRefreshTokenClaims(userResponse.ID, td)
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(rtSecret))
	if err != nil {
		return nil, err
	}

	savedAuthErr := RedisService.CreateAuth(userResponse.ID, td)
	if savedAuthErr != nil {
		return nil, savedAuthErr
	}

	accessToken := &access_token.AccessToken{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}
	return accessToken, nil
}

func (s *accessTokenService) verifyToken(at string) (*jwt.Token, api_errors.RestErr) {
	atSecret := "AT_SECRET" //implment this os.Getenv("ACESS_TOKEN_SECRET")
	token, err := jwt.Parse(at, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(atSecret), nil
	})

	if err != nil {
		return nil, api_errors.NewInternalServerError("error trying to verify token", err)
	}
	return token, nil
}

func (s *accessTokenService) TokenValid(at string) api_errors.RestErr {
	token, err := s.verifyToken(at)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return api_errors.NewInternalServerError("error trying to validate access token", nil)
	}

	return nil
}

func (s *accessTokenService) ExtractTokenMetadata(at string) (*access_token.AccessDetails, api_errors.RestErr) {
	token, err := s.verifyToken(at)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, api_errors.NewInternalServerError("error trying to parse access_uuid", nil)
		}
		userID := claims["user_id"].(string)

		return &access_token.AccessDetails{
			AccessUuid: accessUuid,
			UserID:     userID,
		}, nil
	}

	return nil, api_errors.NewInternalServerError("error trying to extract token metadata", nil)
}

func buildAccessTokenClaims(userID string, td *access_token.TokenDetails) jwt.MapClaims {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires
	return atClaims
}

func buildRefreshTokenClaims(userID string, td *access_token.TokenDetails) jwt.MapClaims {
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	return rtClaims
}
