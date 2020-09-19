package services

import (
	"time"

	"github.com/mendezdev/expenses_oauth-api/db/redisdb"
	"github.com/mendezdev/expenses_oauth-api/domain/access_token"
	"github.com/mendezdev/expenses_oauth-api/utils/api_errors"
)

var (
	// RedisService variable to interact with de redis store
	RedisService redisServiceInterface = &redisService{}
)

type redisService struct{}

type redisServiceInterface interface {
	CreateAuth(string, *access_token.TokenDetails) api_errors.RestErr
	FetchAuth(*access_token.AccessDetails) (*string, api_errors.RestErr)
}

// CreateAuth will persist data in redis store
func (s *redisService) CreateAuth(userID string, td *access_token.TokenDetails) api_errors.RestErr {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := redisdb.Client.Set(td.AccessUuid, userID, at.Sub(now)).Err()
	if errAccess != nil {
		return api_errors.NewInternalServerError("error to store AT", errAccess)
	}

	errRefresh := redisdb.Client.Set(td.RefreshUuid, userID, rt.Sub(now)).Err()
	if errRefresh != nil {
		return api_errors.NewInternalServerError("error to store RT", errRefresh)
	}
	return nil
}

// FetchAuth will get the item stored with the given key
func (s *redisService) FetchAuth(ad *access_token.AccessDetails) (*string, api_errors.RestErr) {
	userID, err := redisdb.Client.Get(ad.AccessUuid).Result()
	if err != nil {
		return nil, api_errors.NewNotFoundError("access token not found by the given key")
	}
	return &userID, nil
}
