package db

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var pdb *sql.DB = func() *sql.DB {
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "user=postgres password=passwd dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}()

func SHA256(str string) string {
	data := []byte(str)
	myHash := sha256.Sum256(data)
	myHex := hex.EncodeToString(myHash[:])
	return myHex
}

func PInsertEntryList(UID string, list []string) {
	fmt.Println(list)
	var exists bool
	nKey := SHA256(UID)
	// check if the UID exists in the recommendations
	err := pdb.QueryRow("SELECT EXISTS(SELECT 1 FROM recommendations WHERE listKey = $1)", UID).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		var entryKey string
		pdb.QueryRow("SELECT entryKey FROM recommendations WHERE listKey = $1", UID).Scan(&entryKey)
		pdb.Exec("DELETE FROM entryList WHERE listKey = $1", entryKey)
		pdb.Exec("UPDATE recommendations SET entryKey = $1 WHERE listKey = $2", nKey, UID)
		if RedisClient.Exists(UID).Val() == 1 {
			RedisClient.Set(UID, nKey, 0)
		}
	} else {
		pdb.Exec("INSERT INTO recommendations (listKey, entryKey) VALUES ($1, $2)", UID, nKey)
	}
	_, err = pdb.Exec("INSERT INTO entryList (listKey, AID) VALUES ($1, $2)", nKey, pq.Array(list))
	if err != nil {
		log.Fatal(err)
	}
}

// query the entryKey from recommendations and use it to query the list from entryList 
func PGetPage(UID string, page string) []string {
	var list []string
	var entryKey string
	err := pdb.QueryRow("SELECT entryKey FROM recommendations WHERE listKey = $1", UID).Scan(&entryKey)
	if err != nil {
		log.Fatal(err)
	}
	err = pdb.QueryRow("SELECT AID FROM entryList WHERE listKey = $1", entryKey).Scan(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}


// use SELECT ARRAY_SLICE(AID, 1, 10) FROM entryList OFFSET $1 LIMIT $2;
// to get the first 10 elements of the list
// func PGetPage(UID string, page string) []string {
// 	var list []string
// 	err := pdb.QueryRow("SELECT ARRAY_SLICE(AID, 1, 10) FROM entryList WHERE listKey = $1", UID).Scan(&list)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return list
// }

/*CREATE TABLE recommendations (
  listKey  VARCHAR(64) NOT NULL,
  entryKey VARCHAR(64) NOT NULL,
  PRIMARY KEY (listKey)
);

CREATE TABLE entryList (
  listKey  VARCHAR(64) NOT NULL,
  AID  VARCHAR(64)[] ,
  PRIMARY KEY (listKey)
);
*/
