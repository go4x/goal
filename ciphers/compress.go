// Package ciphers provides data compression functions for converting numbers to compact string representations.
//
// This package includes compression and decompression functions:
//   - Base36/Base36Decode: Base36 encoding/decoding (0-9, A-Z) - 36 chars, case insensitive
//   - Base62/Base62Decode: Base62 encoding/decoding (0-9, a-z, A-Z) - 62 chars, case sensitive
//
// Compression ratio comparison (for 12345):
//   - Base36: "9IX" (3 chars)
//   - Base62: "3d7" (3 chars)
//
// These functions are useful for:
//   - URL shortening
//   - Database ID compression
//   - Token generation
//   - Compact number representation
//
// Example usage:
//
//	// Encoding
//	compressed := ciphers.Base36(12345) // "9IX"
//	compressed := ciphers.Base62(12345) // "3d7"
//
//	// Decoding
//	number, ok := ciphers.Base36Decode("9IX") // 12345, true
//	number, ok := ciphers.Base62Decode("3d7") // 12345, true
package ciphers

import "strings"

const (
	base36chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	base36len = len(base36chars)
	base62len = len(base62chars)
)

// Base36 converts a uint to a base36 string, only support positive number.
// This function is used to compress a large number to a short string which is case insensitive.
// The result string is uppercase.
func Base36(n uint) string {
	if n == 0 {
		return base36chars[0:1]
	}
	ss := make([]string, 0)
	for n != 0 {
		m := n % uint(base36len)
		ss = append(ss, base36chars[m:m+1])
		n /= uint(base36len)
	}
	// reverse
	b := strings.Builder{}
	for i := len(ss) - 1; i >= 0; i-- {
		b.WriteString(ss[i])
	}
	return b.String()
}

// Base36Decode converts a base36 string back to a uint.
// This function is the inverse of Base36.
// Returns 0 and false if the input string is invalid.
// Base36 is case insensitive, so both upper and lower case letters are accepted.
func Base36Decode(s string) (uint, bool) {
	if s == "" {
		return 0, false
	}

	result := uint(0)
	for _, char := range s {
		// Convert to uppercase for case insensitive comparison
		upperChar := strings.ToUpper(string(char))
		index := strings.IndexRune(base36chars, []rune(upperChar)[0])
		if index == -1 {
			return 0, false
		}
		result = result*uint(base36len) + uint(index)
	}
	return result, true
}

// Base62 converts a uint to a base62 string, only support positive number.
// This function is used to compress a large number to a short string which is case sensitive.
func Base62(n uint) string {
	if n == 0 {
		return base62chars[0:1]
	}
	ss := make([]string, 0)
	for n != 0 {
		m := n % uint(base62len)
		ss = append(ss, base62chars[m:m+1])
		n /= uint(base62len)
	}
	b := strings.Builder{}
	for i := len(ss) - 1; i >= 0; i-- {
		b.WriteString(ss[i])
	}
	return b.String()
}

// Base62Decode converts a base62 string back to a uint.
// This function is the inverse of Base62.
// Returns 0 and false if the input string is invalid.
func Base62Decode(s string) (uint, bool) {
	if s == "" {
		return 0, false
	}

	result := uint(0)
	for _, char := range s {
		index := strings.IndexRune(base62chars, char)
		if index == -1 {
			return 0, false
		}
		result = result*uint(base62len) + uint(index)
	}
	return result, true
}
