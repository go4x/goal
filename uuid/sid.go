// Package uuid provides utilities for generating unique identifiers.
//
// This package offers two main approaches for ID generation:
//  1. Sonyflake-based distributed ID generation (Sid struct)
//  2. Standard UUID generation (UUID functions)
//
// The Sid struct provides distributed ID generation using Sonyflake algorithm,
// which generates unique IDs across multiple machines without coordination.
// Generated IDs are converted to base62 format for compact string representation.
//
// Example usage:
//
//	// Sonyflake-based ID generation
//	sid := uuid.NewSid()
//	id, err := sid.GenString() // Returns base62 string
//	rawID, err := sid.GenUint64() // Returns raw uint64
//
//	// Standard UUID generation
//	uuidStr := uuid.UUID() // Returns standard UUID string
//	uuid32 := uuid.UUID32() // Returns UUID without hyphens
package uuid

import (
	"fmt"

	"github.com/sony/sonyflake"
)

// Sid represents a Sonyflake-based distributed ID generator.
// Sonyflake is a distributed unique ID generator inspired by Twitter's Snowflake.
// It generates 64-bit unique IDs that are roughly sortable by time.
//
// The generated IDs are guaranteed to be unique across multiple machines
// without requiring coordination between them.
type Sid struct {
	sf *sonyflake.Sonyflake
}

// NewSid creates a new Sonyflake ID generator instance.
// This function will panic if the Sonyflake instance cannot be created,
// which typically happens when the machine ID cannot be determined.
//
// Example:
//
//	sid := uuid.NewSid()
//	id, err := sid.GenString()
func NewSid() *Sid {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	return &Sid{sf}
}

// GenString generates a new unique ID and returns it as a base62 string.
// The returned string is more compact than the raw numeric representation
// and is safe for use in URLs and filenames.
//
// Example:
//
//	sid := uuid.NewSid()
//	id, err := sid.GenString()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(id) // Output: "1Z3k9X2m"
func (s Sid) GenString() (string, error) {
	id, err := s.sf.NextID()
	if err != nil {
		return "", fmt.Errorf("failed to generate sonyflake ID: %w", err)
	}
	return intToBase62(int(id)), nil
}

// GenUint64 generates a new unique ID and returns it as a raw uint64.
// This is the underlying numeric ID before base62 conversion.
//
// Example:
//
//	sid := uuid.NewSid()
//	id, err := sid.GenUint64()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(id) // Output: 1234567890123456789
func (s Sid) GenUint64() (uint64, error) {
	return s.sf.NextID()
}

// intToBase62 converts an integer to base62 string representation.
// Base62 uses characters: 0-9, a-z, A-Z (62 characters total).
// This provides a more compact string representation than base10.
//
// Example:
//
//	intToBase62(0)    // returns "0"
//	intToBase62(61)   // returns "Z"
//	intToBase62(62)   // returns "10"
//	intToBase62(123)  // returns "1Z"
func intToBase62(n int) string {
	const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if n == 0 {
		return string(base62[0])
	}

	// Handle negative numbers
	negative := n < 0
	if negative {
		n = -n
	}

	var result []byte
	for n > 0 {
		result = append(result, base62[n%62])
		n /= 62
	}

	// Reverse the string
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	if negative {
		return "-" + string(result)
	}

	return string(result)
}
