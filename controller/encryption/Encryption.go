package encryption

import (
	"crypto/sha256"
	"encoding/hex"

	_ "github.com/lib/pq"
)

// using sha256 encrypt any string into a string with 64 bytes
// redis does not support []byte as key or value so we need to convert it to string
// [32]byte can't supported by UTF-8
func Encrypt(str string) string {
	data := []byte(str)
	myHash := sha256.Sum256(data)
	myHex := hex.EncodeToString(myHash[:])
	return myHex
}
