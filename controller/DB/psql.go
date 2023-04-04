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
func PQuerryUIDHead(UID string) string {
    var key string
    err := db.QueryRow("SELECT key FROM users WHERE UID = $1", UID).Scan(&key)
    if err != nil {
        return ""
    }
    return key
}

// query the database for a list's key and nextKey
func PQuerryList(Key string) (string , string) {
    var AIDList string
    var NextKey string
    err := db.QueryRow("SELECT AIDList, NextKey FROM lists WHERE Key = $1", Key).Scan(&AIDList, &NextKey)
    if err != nil {
        return "" , ""
    }
    return AIDList, NextKey
}

// psql update a user's key
func PUpdateUIDHead(UID string, key string) {
    err := db.QueryRow("UPDATE users SET key = $1 WHERE UID = $2", key, UID)
    if err != nil {
        log.Fatal(err)
    }
}

// psql update a list's key and nextKey
func PUpdateList(Key string, AIDList string , NextKey string) {
    err := db.QueryRow("UPDATE lists SET AIDList = $1, NextKey = $2 WHERE Key = $3", AIDList, NextKey, Key)
    if err != nil {
        log.Fatal(err)
    }
}