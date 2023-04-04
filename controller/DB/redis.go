package db

import (
	"DBIH/controller/DB"
	"encoding/base64"
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", 
        Password: "",               
        DB:       0,                
    })
}

// insert [64]byte , [64]byte into redis
func RInsertUIDHead(UID string, key string) {
    err := RedisClient.Set(UID, key, 0).Err()
    if err != nil {
        panic(err)
    }
}

// RGetUIDHead returns the key of the UID
func RGetUIDHead(UID string) string {
    var tmp string
    key, err := RedisClient.Get(UID).Result()
    // if UID not found in redis, check if it exists in postgreSQL
    if err != nil {
        tmp = DB.PQuerryUIDHead(UID)
        // if key exists, insert into redis
        if tmp != "" {
            RInsertUIDHead(UID, tmp)
        }
        return tmp
    } else if err != nil {
        panic(err)
    }
    // if UID found in redis, decode the key
    return key
}

// Update UIDHead in redis
func RUpdateUIDHead(UID string, key string) {
    // check if UID exists in redis , if exists, update the key 
    _, err := RedisClient.Get(UID).Result()
    if err == nil {
        err := RedisClient.Set(UID, key, 0).Err()
        if err != nil {
            panic(err)
        }
    }
}

