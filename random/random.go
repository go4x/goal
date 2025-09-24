package random

import (
	"crypto/rand"
	"encoding/binary"
	"time"

	mathrand "math/rand"
)

func New() *mathrand.Rand {
	return mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
}

func NewSecure() *mathrand.Rand {
	return mathrand.New(mathrand.NewSource(getSecureSeed()))
}

func NewWithSeed(seed int64) *mathrand.Rand {
	return mathrand.New(mathrand.NewSource(seed))
}

// getSecureSeed generates a cryptographically secure seed
func getSecureSeed() int64 {
	// Try to use crypto/rand for better security
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err == nil {
		return int64(binary.BigEndian.Uint64(b))
	}

	// Fallback to time-based seed if crypto/rand fails
	return time.Now().UnixNano()
}
