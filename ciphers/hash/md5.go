// Package hash provides cryptographic hash functions for various use cases.
//
// This package includes:
//   - MD5: Legacy hash function (WARNING: Not cryptographically secure)
//   - SHA256: Secure hash function for data integrity
//   - BCrypt: Password hashing function with adaptive cost
//
// SECURITY WARNINGS:
//   - MD5 should NOT be used for security purposes (collision attacks)
//   - Use SHA256 for data integrity verification
//   - Use BCrypt for password hashing (not SHA256)
//
// Example usage:
//
//	// For data integrity (secure)
//	checksum := hash.SHA256("important data")
//
//	// For password hashing (secure)
//	passwordHash, err := hash.BCrypt("userPassword")
//	isValid := hash.BCryptCompare(passwordHash, "userPassword")
//
//	// Legacy compatibility only (NOT secure)
//	legacyHash := hash.MD5("legacy data")
package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 returns the MD5 hash of the input string as a hexadecimal string.
//
// WARNING: MD5 is cryptographically broken and should NOT be used for security purposes.
// MD5 is vulnerable to collision attacks and is not suitable for:
//   - Password hashing
//   - Digital signatures
//   - Security-critical applications
//
// This function is provided for compatibility with legacy systems only.
// For security purposes, use SHA256 or BCrypt instead.
//
// Parameters:
//   - str: The input string to hash
//
// Returns:
//   - string: The MD5 hash as a hexadecimal string
//
// Example:
//
//	hash := hash.MD5("hello")
//	// Result: "5d41402abc4b2a76b9719d911017c592"
func MD5(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
