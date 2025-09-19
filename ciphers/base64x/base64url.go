package base64x

import "encoding/base64"

// URLEncoding provides URL-safe base64 encoding and decoding.
// This uses the URL-safe base64 alphabet (A-Z, a-z, 0-9, -, _) with padding characters (=).
// URL-safe encoding is suitable for use in URLs and filenames.
//
// Example:
//
//	encoded := base64x.URLEncoding.Encode([]byte("hello world"))
//	decoded, err := base64x.URLEncoding.Decode(encoded)
var URLEncoding base64Url

// RawURLEncoding provides URL-safe base64 encoding and decoding without padding.
// This uses the URL-safe base64 alphabet (A-Z, a-z, 0-9, -, _) but omits padding characters (=).
// Raw URL-safe encoding is suitable for use in URLs and filenames where padding is not desired.
//
// Example:
//
//	encoded := base64x.RawURLEncoding.Encode([]byte("hello world"))
//	decoded, err := base64x.RawURLEncoding.Decode(encoded)
var RawURLEncoding base64RawUrl

// base64Url implements URL-safe base64 encoding and decoding.
type base64Url struct {
}

// Encode encodes the given byte slice into a URL-safe base64-encoded string.
// The result uses the URL-safe base64 alphabet (A-Z, a-z, 0-9, -, _) with padding characters (=).
// This encoding is suitable for use in URLs and filenames.
//
// Parameters:
//   - b: The byte slice to encode
//
// Returns:
//   - string: The URL-safe base64-encoded string
//
// Example:
//
//	encoded := base64x.URLEncoding.Encode([]byte("hello world"))
//	// Result: "aGVsbG8gd29ybGQ="
func (base64Url) Encode(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

// Decode decodes a URL-safe base64-encoded string.
// If the optional 'strict' parameter is true, it uses strict decoding mode (rejects invalid characters).
// Returns the decoded bytes and an error if decoding fails.
//
// Parameters:
//   - str: The URL-safe base64-encoded string to decode
//   - strict: Optional parameter to enable strict decoding mode (default: false)
//
// Returns:
//   - []byte: The decoded byte slice
//   - error: An error if decoding fails
//
// Example:
//
//	decoded, err := base64x.URLEncoding.Decode("aGVsbG8gd29ybGQ=")
//	// Result: []byte("hello world"), nil
//
//	// With strict mode
//	decoded, err := base64x.URLEncoding.Decode("aGVsbG8gd29ybGQ=", true)
func (base64Url) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.URLEncoding.Strict().DecodeString(str)
	} else {
		return base64.URLEncoding.DecodeString(str)
	}
}

// base64RawUrl implements URL-safe base64 encoding and decoding without padding.
type base64RawUrl struct {
}

// Encode encodes the given byte slice into a URL-safe base64-encoded string without padding.
// The result uses the URL-safe base64 alphabet (A-Z, a-z, 0-9, -, _) but omits padding characters (=).
// This encoding is suitable for use in URLs and filenames where padding is not desired.
//
// Parameters:
//   - b: The byte slice to encode
//
// Returns:
//   - string: The URL-safe base64-encoded string without padding
//
// Example:
//
//	encoded := base64x.RawURLEncoding.Encode([]byte("hello world"))
//	// Result: "aGVsbG8gd29ybGQ"
func (base64RawUrl) Encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

// Decode decodes a URL-safe base64-encoded string without padding.
// If the optional 'strict' parameter is true, it uses strict decoding mode (rejects invalid characters).
// Returns the decoded bytes and an error if decoding fails.
//
// Parameters:
//   - str: The URL-safe base64-encoded string to decode (without padding)
//   - strict: Optional parameter to enable strict decoding mode (default: false)
//
// Returns:
//   - []byte: The decoded byte slice
//   - error: An error if decoding fails
//
// Example:
//
//	decoded, err := base64x.RawURLEncoding.Decode("aGVsbG8gd29ybGQ")
//	// Result: []byte("hello world"), nil
//
//	// With strict mode
//	decoded, err := base64x.RawURLEncoding.Decode("aGVsbG8gd29ybGQ", true)
func (base64RawUrl) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.RawURLEncoding.Strict().DecodeString(str)
	} else {
		return base64.RawURLEncoding.DecodeString(str)
	}
}
