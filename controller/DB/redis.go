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
func GetLists(UID string , page int) []string {
	var list []string
	if RedisClient.Exists(UID).Val() == 1 {
		list = RedisClient.LRange(UID, 0, -1).Val()
	} else {
		list = PGetPage(UID, page)
		RedisClient.RPush(UID, list)
	}
	return list
}
