package user

import (
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
)

func hashPassword(password string, salt string, pepper string) string {
	var hash = sha512.New()

	hash.Write([]byte(password))
	hash.Write([]byte(salt))
	hash.Write([]byte(pepper))

	return base64.StdEncoding.EncodeToString(hash.Sum([]byte{}))
}

func authenticate(hash string, password string, salt string, pepper string) bool {
	var hash2 = hashPassword(password, salt, pepper)

	return subtle.ConstantTimeCompare([]byte(hash), []byte(hash2)) == 1
}
