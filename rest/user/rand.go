package user

import (
	cryptorand "crypto/rand"
	"encoding/base64"
	mathrand "math/rand"
)

func randBytes(n int) string {
	var (
		bytes = make([]byte, n)
		err   error
	)

	// try to use cryptographically random number
	_, err = cryptorand.Read(bytes)
	if err != nil {
		for i := 0; i < n; i++ {
			// if error, use pseudo random numbers
			bytes[i] = byte(mathrand.Int31n(128))
		}
	}

	return base64.StdEncoding.EncodeToString(bytes)
}
