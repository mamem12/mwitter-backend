package common

import (
	"crypto/sha256"
	"encoding/hex"
)

func StrToHash(str string) (string, error) {

	hash := sha256.New()
	_, err := hash.Write([]byte(str))

	if err != nil {
		return "", err
	}

	md := hash.Sum(nil)
	hashStr := hex.EncodeToString(md)

	return hashStr, nil
}
