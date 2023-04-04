package db
import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

var db *sql.DB = func() *sql.DB {
    var db *sql.DB
    var err error
    db, err = sql.Open("postgres", "user=postgres password=passwd dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    return db
}()


// insert a new user into the database UID , [32]byte
// P = Postgres 
func PInsertUIDHead(UID [64]byte, key [64]byte) {
	_, err := db.Exec("INSERT INTO users (UID, key) VALUES ($1, $2)", UID, key)
	if err != nil {
		log.Fatal(err)
	}
}

// insert a new list into the database LID CHAR(32) , AIDList CHAR(320)
func PInsertList(Key [64]byte , AIDList [640]byte , NextKey [64]byte) {
    _, err := db.Exec("INSERT INTO lists (Key, AIDList, NextKey) VALUES ($1, $2, $3)", Key, AIDList, NextKey)
    if err != nil {
        log.Fatal(err)
    }
}

// query the database for a user's key
func PQuerryUIDHead(UID string) [64]byte {
    var key [64]byte
    err := db.QueryRow("SELECT key FROM users WHERE UID = $1", UID).Scan(&key)
    if err != nil {
        log.Fatal(err)
    }
    return key
}

// query the database for a list's key and nextKey
func PQuerryList(Key string) ([640]byte , [64]byte) {
    var AIDList [640]byte
    var NextKey [64]byte
    err := db.QueryRow("SELECT AIDList, NextKey FROM lists WHERE Key = $1", Key).Scan(&AIDList, &NextKey)
    if err != nil {
        log.Fatal(err)
    }
    return AIDList, NextKey
}