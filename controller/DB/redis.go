package db

import (
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// userGet pages are cached in redis
