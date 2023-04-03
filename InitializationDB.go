package main
import (
	// "crypto/sha256"
	// "database/sql"
	// "encoding/hex"
	// "encoding/json"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
func Create() {

	db, err := sql.Open("postgres", "user=postgres password=passwd host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// _, err = db.Exec("CREATE DATABASE mydatabase")
	// recovery
	fmt.Println("Database created successfully")

	db, err = sql.Open("postgres", "user=postgres password=passwd dbname=mydatabase sslmode=disable")
	if err != nil {
		panic(err)
	}
	// UID-Head Table : { UID , HeadKey }
	// Linked List Table : { Key , listKey , nextPersionalListKey }
	// listKey Table : {Key , List} // 10 datas in a page with sha256 (CHAR(64))
	// AID Table : { AID , Aritcle }
	_, err = db.Exec(`
		CREATE TABLE UID-Head (
			UID CHAR(64) PRIMARY KEY,
			HeadKey CHAR(64)
		);
	`)
	_, err = db.Exec(`
		CREATE TABLE Link (
			LID CHAR(64) PRIMARY KEY, 
			listKey CHAR(64) , 
			NextKey CHAR(64) 
		);
	`)
	_, err = db.Exec(`
		CREATE TABLE List (
			LID CHAR(64) PRIMARY KEY , 
			AIDList CHAR(640) 
		);
	`)
	_, err = db.Exec(`
		CREATE TABLE AID (
			AID CHAR(64) PRIMARY KEY ,
			Topic   STRING,
			Aritcle STRING
		);
	`)
	fmt.Println("Table created successfully")
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'")
	defer rows.Close()
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}

		fmt.Println(tableName)
	}
}
