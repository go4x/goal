package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 returns the SHA256 hash of the input string as a hexadecimal string.
// SHA256 is a cryptographically secure hash function suitable for:
//   - Data integrity verification
//   - Digital signatures
//   - Checksums
//   - Non-password hashing applications
//
// For password hashing, use BCrypt instead as it provides better security
// against brute force attacks through its adaptive cost parameter.
//
// Parameters:
//   - str: The input string to hash
//
// Returns:
//   - string: The SHA256 hash as a hexadecimal string (64 characters)
//
// Example:
//
//	hash := hash.SHA256("hello world")
//	// Result: "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
func SHA256(str string) string {
	data := []byte(str)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
