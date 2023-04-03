package main

import (
	"Documents/DBEE_HW/controller/module"
	"github.com/gin-gonic/gin"
)

// encrypt the listKey and store it in postgreSQL
func main() {
	// link to redis
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// link to postgreSQL
	// Create()
	// webTesting()
	router := gin.Default()
	router.Run("localhost:8443")

}
