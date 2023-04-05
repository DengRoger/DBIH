package db

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"time"

	"github.com/lib/pq"
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
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	str = str + timeStr
	data := []byte(str)
	myHash := sha256.Sum256(data)
	myHex := hex.EncodeToString(myHash[:])
	return myHex
}

func PInsertEntryList(UID string, list []string) {
	var exists bool
	nKey := SHA256(UID)
	// check if the UID exists in the recommendations
	err := pdb.QueryRow("SELECT EXISTS(SELECT 1 FROM recommendations WHERE UID = $1)", UID).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		_, err = pdb.Exec("UPDATE recommendations SET listKey = $1, AID = $2 WHERE UID = $3", nKey, pq.Array(list), UID)
		if err != nil {
			log.Fatal(err)
		}
		if RedisClient.Exists(UID).Val() == 1 {
			RedisClient.Set(UID, nKey, 0)
		}
	} else {
		// insert into recommendations
		_, err = pdb.Exec("INSERT INTO recommendations (UID, listKey , AID ) VALUES ($1, $2, $3)", UID, nKey, pq.Array(list))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetPage(UID string, page int) []string {
	offset := page
	limit := 10
	rows, err := pdb.Query("SELECT unnest(AID) FROM recommendations WHERE UID=$1 OFFSET $2 LIMIT $3", UID, offset, limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var aids []string
	for rows.Next() {
		var aid string
		if err := rows.Scan(&aid); err != nil {
			panic(err)
		}
		aids = append(aids, aid)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return aids
}

func PGetListKey(UID string) string {
	var listKey string
	err := pdb.QueryRow("SELECT listKey FROM recommendations WHERE UID = $1", UID).Scan(&listKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		} else {
			log.Fatal(err)
		}
	}
	return listKey
}

