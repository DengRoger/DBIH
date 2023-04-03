package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

func getList() []byte {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	exists, err := client.Exists(fmt.Sprintf("%v", context.Background()), "mykey").Result()
	if exists == 1 {
		val, err := client.Get("mykey").Result()
		if err != nil {panic(err)}
		var list []string
		for i := 0; i < len(val); i += 64 {
			list = append(list, val[i:i+64])
		}
		jsonList, err := json.Marshal(list)
		if err != nil {panic(err)}
		return jsonList
	}
	if err != nil {
		panic(err)
	}
	return nil
}

