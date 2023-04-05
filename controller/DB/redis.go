package db

import (
	"strconv"

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

func RGetPage(uid string, page int) []string {
	key := uid + strconv.Itoa(page)
	// []string is the type of listKey
	listKey := []string{}
	// check if the key exists in redis
	if RedisClient.Exists(key).Val() == 1 {
		return RedisClient.LRange(key, 0, -1).Val()
	} else { // if not, get the listKey from postgreSQL
		listKey = GetPage(uid, page)
	}
	return listKey
}

// userGet pages are cached in redis
