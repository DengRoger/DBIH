package main

import (
	//

	db "DBIH/controller/DB"
	"fmt"
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
	// router := mux.NewRouter()
	// router.HandleFunc("/modify/{UID}", module.ModifyHandler).Methods("POST")
	// router.HandleFunc("/get/{UID}/{page}", module.Getpage).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8006", router))
	// module.Getpage()

	fmt.Println(db.RGetPage("rew", 10))
	fmt.Println(db.RGetPage("re", 21))
}
