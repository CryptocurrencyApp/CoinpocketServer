package hash

import (
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

//func ToHash(password string) string {
//	converted := sha256.Sum256([]byte(password))
//	return hex.EncodeToString(converted[:])
//}

func ToHash(pass string, salt string) string {
	usingSalt := []byte(salt)
	converted, _ := scrypt.Key([]byte(pass), usingSalt, 16384, 8, 1, 32)
	return hex.EncodeToString(converted[:])
}
