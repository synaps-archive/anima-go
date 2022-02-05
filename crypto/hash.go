package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(content []byte) string {
	h := sha256.New()
	h.Write(content)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
