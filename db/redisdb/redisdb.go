package redisdb

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
)

func init() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379" //change this to env var
	}
	Client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis successful connected!")
}
