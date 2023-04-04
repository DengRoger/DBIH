package db

import (
    "fmt"
    "github.com/go-redis/redis"
    "DBIH/controller/DB"
)

var RedisClient *redis.Client

func init() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", 
        Password: "",               
        DB:       0,                
    })
}

