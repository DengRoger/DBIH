package main

import (
	//

	db "DBIH/controller/DB"
	"DBIH/controller/module"
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
	db.RGetPage("re", 0)
	router := mux.NewRouter()
	router.HandleFunc("/modify/{uid}", module.ModifyHandler).Methods("POST")
	router.HandleFunc("/get/{uid}/{page}", module.Getpage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8006", router))

}
