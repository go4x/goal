package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	// BCryptMaxLength is the maximum length of input string for bcrypt
	BCryptMaxLength = 72
)

var (
	// ErrBCryptInputTooLong is returned when input string exceeds maximum length
	ErrBCryptInputTooLong = errors.New("bcrypt input string too long (max 72 bytes)")
)

// BCrypt returns the bcrypt hash of the input string.
// Returns empty string and error if input length exceeds 72 bytes.
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

// BCryptCompare compares the input string with the bcrypt hash.
// Returns false if input length exceeds 72 bytes or comparison fails.
func BCryptCompare(hash string, str string) bool {
	if len(str) > BCryptMaxLength {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(str)) == nil
}
