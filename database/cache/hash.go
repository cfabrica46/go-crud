package cache

import (
	"crypto/sha256"
	"encoding/hex"
)

func getTokenHash(token string) (hash string, err error) {
	h := sha256.New()
	_, err = h.Write([]byte(token))
	if err != nil {
		return
	}

	hash = hex.EncodeToString(h.Sum(nil))
	return
}
