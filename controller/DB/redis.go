package db

import (
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

type Response struct {
	ListContent []string `json:"listContent"`
}

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func RGetPage(uid string, page int) []string {
	var listKey string
	page /= 10
	lexists, err := RedisClient.Exists(uid).Result()
	if err != nil {
		panic(err)
	}
	if lexists == 0 {
		listKey = PGetListKey(uid)
		if listKey == "" {
			return []string{}
		}
		RedisClient.Set(uid, listKey, 0)
	} else {
		listKey, err = RedisClient.Get(uid).Result()
		if err != nil {
			panic(err)
		}
	}

	listKey += strconv.Itoa(page)
	check, err := RedisClient.Exists(listKey).Result()
	if err != nil {
		panic(err)
	}
	if check == 0 {
		tmp := GetPage(uid, page)
		jsonString, err := json.Marshal(tmp)
		if err != nil {
			panic(err)
		}
		err = RedisClient.Set(listKey, jsonString, 0).Err()
		if err != nil {
			panic(err)
		}
		return tmp
	} else {
		tmp, _ := RedisClient.Get(listKey).Result()
		var getInfoResult []string
		err := json.Unmarshal([]byte(tmp), &getInfoResult)
		if err != nil {
			panic(err)
		}
		return getInfoResult
	}
}
