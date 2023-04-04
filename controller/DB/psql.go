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

// insert recommendation list in postgreSQL
func PInsertList(UID string, list []string) {
    pKey := encryption.Encrypt(UID)
    _, err := pdb.Exec("INSERT INTO entryList (listKey,AID) VALUES ($1,$2)", pKey, list)
    if err != nil {
        log.Fatal(err)
    }
}


/*
CREATE TABLE recommendations (
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