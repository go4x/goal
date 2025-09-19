// Package base64x provides enhanced base64 encoding and decoding functionality.
//
// This package extends the standard Go base64 package with additional encoding
// variants and specialized functions for different use cases. It includes:
//
//   - Standard base64 encoding (with and without padding)
//   - URL-safe base64 encoding (with and without padding)
//   - Base64URL encoding for unsigned integers (RFC 7518 compliant)
//
// The package provides a consistent API across all encoding variants, making it
// easy to switch between different base64 encodings as needed.
//
// Example usage:
//
//	// Standard base64 encoding
//	encoded := base64x.StdEncoding.Encode([]byte("hello world"))
//	decoded, err := base64x.StdEncoding.Decode(encoded)
//
//	// URL-safe base64 encoding
//	urlEncoded := base64x.URLEncoding.Encode([]byte("hello world"))
//	urlDecoded, err := base64x.URLEncoding.Decode(urlEncoded)
//
//	// Base64URL encoding for big integers
//	bi := big.NewInt(12345)
//	uintEncoded := base64x.Base64UrlUint.Encode(bi)
//	uintDecoded, err := base64x.Base64UrlUint.Decode(uintEncoded)
package base64x

import "encoding/base64"

// StdEncoding provides standard base64 encoding and decoding.
// This uses the standard base64 alphabet with padding characters.
//
// Example:
//
//	encoded := base64x.StdEncoding.Encode([]byte("hello"))
//	decoded, err := base64x.StdEncoding.Decode(encoded)
var StdEncoding base64std

// RawStdEncoding provides standard base64 encoding and decoding without padding.
// This uses the standard base64 alphabet but omits padding characters.
//
// Example:
//
//	encoded := base64x.RawStdEncoding.Encode([]byte("hello"))
//	decoded, err := base64x.RawStdEncoding.Decode(encoded)
var RawStdEncoding base64raw

// base64std implements standard base64 encoding and decoding.
type base64std struct {
}

// Encode encodes the given byte slice into a base64-encoded string using standard base64 encoding.
// The result uses the standard base64 alphabet (A-Z, a-z, 0-9, +, /) with padding characters (=).
//
// Parameters:
//   - b: The byte slice to encode
//
// Returns:
//   - string: The base64-encoded string
//
// Example:
//
//	encoded := base64x.StdEncoding.Encode([]byte("hello world"))
//	// Result: "aGVsbG8gd29ybGQ="
func (base64std) Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Decode decodes a base64-encoded string using standard base64 decoding.
// If the optional 'strict' parameter is true, it uses strict decoding mode (rejects invalid characters).
// Returns the decoded bytes and an error if decoding fails.
//
// Parameters:
//   - str: The base64-encoded string to decode
//   - strict: Optional parameter to enable strict decoding mode (default: false)
//
// Returns:
//   - []byte: The decoded byte slice
//   - error: An error if decoding fails
//
// Example:
//
//	decoded, err := base64x.StdEncoding.Decode("aGVsbG8gd29ybGQ=")
//	// Result: []byte("hello world"), nil
//
//	// With strict mode
//	decoded, err := base64x.StdEncoding.Decode("aGVsbG8gd29ybGQ=", true)
func (base64std) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.StdEncoding.Strict().DecodeString(str)
	} else {
		return base64.StdEncoding.DecodeString(str)
	}
}

// base64raw implements standard base64 encoding and decoding without padding.
type base64raw struct {
}

// Encode encodes the given byte slice into a base64-encoded string using standard base64 encoding without padding.
// The result uses the standard base64 alphabet (A-Z, a-z, 0-9, +, /) but omits padding characters (=).
//
// Parameters:
//   - b: The byte slice to encode
//
// Returns:
//   - string: The base64-encoded string without padding
//
// Example:
//
//	encoded := base64x.RawStdEncoding.Encode([]byte("hello world"))
//	// Result: "aGVsbG8gd29ybGQ"
func (base64raw) Encode(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}

// Decode decodes a base64-encoded string using standard base64 decoding without padding.
// If the optional 'strict' parameter is true, it uses strict decoding mode (rejects invalid characters).
// Returns the decoded bytes and an error if decoding fails.
//
// Parameters:
//   - str: The base64-encoded string to decode (without padding)
//   - strict: Optional parameter to enable strict decoding mode (default: false)
//
// Returns:
//   - []byte: The decoded byte slice
//   - error: An error if decoding fails
//
// Example:
//
//	decoded, err := base64x.RawStdEncoding.Decode("aGVsbG8gd29ybGQ")
//	// Result: []byte("hello world"), nil
//
//	// With strict mode
//	decoded, err := base64x.RawStdEncoding.Decode("aGVsbG8gd29ybGQ", true)
func (base64raw) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.RawStdEncoding.Strict().DecodeString(str)
	} else {
		return base64.RawStdEncoding.DecodeString(str)
	}
}
