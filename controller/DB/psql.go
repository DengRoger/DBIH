package db

import (
	"DBIH/controller/encryption"
	"database/sql"
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

func PInsertEntryList(UID string, list []string) {
	// check if UID exists in recommendation table
	var tmp, eKey string
	err := pdb.QueryRow("SELECT listKey FROM recommendations WHERE listKey = $1", UID).Scan(&tmp)
	if err != nil { //
		eKey := encryption.Encrypt(UID)
		_, err := pdb.Exec("INSERT INTO recommendations (listKey,entryKey) VALUES ($1,$2)", eKey, UID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// if UID exists , exchange eKey with UID
		eKey = encryption.Encrypt(UID)
		_, err := pdb.Exec("UPDATE recommendations SET listKey = $1 WHERE listKey = $2", eKey, UID)
		if err != nil {
			log.Fatal(err)
		}
		// delete the old list
		_, err = pdb.Exec("DELETE FROM entryList WHERE listKey = $1", UID)
		if err != nil {
			log.Fatal(err)
		}
		// check if UID exists in redis
		_, err = RedisClient.Get(UID).Result()
		if err == nil {
			// if UID exists in redis, update the key
			err := RedisClient.Set(UID, eKey, 0).Err()
			if err != nil {
				panic(err)
			}
		}
	}
	// insert the new list
	_, err = pdb.Exec("INSERT INTO entryList (listKey,AID) VALUES ($1,$2)", eKey, list)
	if err != nil {
		log.Fatal(err)
	}
}


// use SELECT ARRAY_SLICE(AID, 1, 10) FROM entryList OFFSET $1 LIMIT $2;
// to get the first 10 elements of the list
func PGetPage(UID string, page string) []string {
    var list []string
    err := pdb.QueryRow("SELECT ARRAY_SLICE(AID, 1, 10) FROM entryList WHERE listKey = $1", UID).Scan(&list)
    if err != nil {
        log.Fatal(err)
    }
    return list
}

/*CREATE TABLE recommendations (
  listKey  VARCHAR(64) NOT NULL,
  entryKey VARCHAR(64) NOT NULL,
  PRIMARY KEY (listKey, entryKey)
);

CREATE TABLE entryList (
  listKey  VARCHAR(64) NOT NULL,
  AID  VARCHAR(64)[] ,
  PRIMARY KEY (listKey)
);
*/
