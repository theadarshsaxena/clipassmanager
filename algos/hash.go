package algos

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Hashalgo(value string) string {
	if len(value) < 32 {
		inputBytes := []byte(fmt.Sprint(value))
		hash := sha256.New()
		hash.Write(inputBytes)
		hashBytes := hash.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)
		return hashString[:(32-len(value))] + value
	}
	return value
}
