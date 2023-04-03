package Encryption

import (
	"crypto/sha256"
	_ "github.com/lib/pq"
)

func Encrypt(str string) [32]byte{
	data := []byte(str)
	return sha256.Sum256(data) 
}