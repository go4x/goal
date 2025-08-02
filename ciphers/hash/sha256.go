package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 returns the SHA256 hash of the input string as a hexadecimal string.
func SHA256(str string) string {
	data := []byte(str)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
