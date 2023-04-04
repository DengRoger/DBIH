package controller

import (
	"encoding/json"
	"fmt"
	"DBIH/controller/encryption"
	"DBIH/controller/DB"
)

func ModifyList(list []string, UID string) {
	Key := encryption.Encrypt(UID)
	
}
