package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	// BCryptMaxLength is the maximum length of input string for bcrypt.
	// BCrypt has a 72-byte limit on input strings. Longer strings are truncated.
	BCryptMaxLength = 72
)

var (
	// ErrBCryptInputTooLong is returned when input string exceeds maximum length.
	// BCrypt has a hard limit of 72 bytes for input strings.
	ErrBCryptInputTooLong = errors.New("bcrypt input string too long (max 72 bytes)")
)

// BCrypt returns the bcrypt hash of the input string.
// BCrypt is a password hashing function designed to be slow and resistant to
// brute force attacks. It uses an adaptive cost parameter that can be increased
// as hardware becomes faster.
//
// This function is specifically designed for password hashing and should NOT be
// used for general-purpose hashing (use SHA256 instead).
//
// Parameters:
//   - str: The input string to hash (max 72 bytes)
//
// Returns:
//   - string: The bcrypt hash string
//   - error: An error if input is too long or hashing fails
//
// Example:
//
//	hash, err := hash.BCrypt("myPassword123")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Result: "$2a$10$N9qo8uLOickgx2ZMRZoMye..."
func BCrypt(str string) (string, error) {
	if len(str) > BCryptMaxLength {
		return "", ErrBCryptInputTooLong
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// BCryptWithCost returns the bcrypt hash of the input string with a custom cost.
// This function allows specifying a custom cost parameter for bcrypt hashing.
// Higher cost values make the hashing slower but more secure against brute force attacks.
//
// Parameters:
//   - str: The input string to hash (max 72 bytes)
//   - cost: The bcrypt cost parameter (4-31, recommended: 10-12)
//
// Returns:
//   - string: The bcrypt hash string
//   - error: An error if input is too long, cost is invalid, or hashing fails
//
// Example:
//
//	hash, err := hash.BCryptWithCost("myPassword123", 12)
//	if err != nil {
//	    log.Fatal(err)
//	}
func BCryptWithCost(str string, cost int) (string, error) {
	if len(str) > BCryptMaxLength {
		return "", ErrBCryptInputTooLong
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(str), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// BCryptCompare compares the input string with the bcrypt hash.
// This function verifies that a plaintext password matches a previously
// hashed password. It handles the salt and cost parameters automatically.
//
// Parameters:
//   - hash: The bcrypt hash string to compare against
//   - str: The plaintext string to verify (max 72 bytes)
//
// Returns:
//   - bool: true if the password matches, false otherwise
//
// Example:
//
//	isValid := hash.BCryptCompare(storedHash, "myPassword123")
//	if isValid {
//	    // Password is correct
//	}
func BCryptCompare(hash string, str string) bool {
	if len(str) > BCryptMaxLength {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(str)) == nil
}
