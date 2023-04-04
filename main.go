package main

import (
	//
	module "DBIH/controller/module"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.HandleFunc("/modify/{UID}", module.ModifyHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8006", router))
}
