package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func hashBlock(b block) string {
	j, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}

	h := hash(j)

	return hex.EncodeToString(h)
}

func hash(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}
