package base64x

import (
	"math/big"
)

// Base64UrlUint provides base64url encoding and decoding for unsigned integers.
// This implementation follows RFC 7518 Section 2 for base64url encoding of unsigned integers.
// It is commonly used in JWT tokens and other cryptographic applications.
//
// The encoding uses the URL-safe base64 alphabet (A-Z, a-z, 0-9, -, _) without padding.
// Zero values are encoded as "AA" and empty strings decode to zero.
//
// Example:
//
//	bi := big.NewInt(12345)
//	encoded := base64x.Base64UrlUint.Encode(bi)
//	decoded, err := base64x.Base64UrlUint.Decode(encoded)
var Base64UrlUint base64UrlUint

// base64UrlUint implements base64url encoding and decoding for unsigned integers.
type base64UrlUint struct {
}

// Encode returns the base64url-encoded representation of the big-endian octet sequence
// as defined in RFC 7518 Section 2. The encoding uses the URL-safe base64 alphabet
// without padding and follows the RFC specification for unsigned integer encoding.
//
// Parameters:
//   - i: The big.Int to encode (must be non-negative)
//
// Returns:
//   - string: The base64url-encoded string
//
// Special cases:
//   - Zero values are encoded as "AA"
//   - Leading zeros are removed to minimize the octet sequence
//
// Example:
//
//	bi := big.NewInt(12345)
//	encoded := base64x.Base64UrlUint.Encode(bi)
//	// Result: "OTk="
//
//	// Zero value
//	zero := big.NewInt(0)
//	encoded := base64x.Base64UrlUint.Encode(zero)
//	// Result: "AA"
func (base64UrlUint) Encode(i *big.Int) string {
	// Get the big-endian bytes
	bytes := i.Bytes()

	// Handle zero case - return "AA" for zero
	if len(bytes) == 0 {
		return "AA"
	}

	// The octet sequence MUST utilize the minimum number of octets
	// needed to represent the value.
	// Remove leading zeros
	start := 0
	for start < len(bytes) && bytes[start] == 0 {
		start++
	}

	if start >= len(bytes) {
		return "AA"
	}

	return RawURLEncoding.Encode(bytes[start:])
}

// Decode returns the BigInt represented by the base64url-encoded string.
// This function decodes a base64url-encoded string back to a big.Int value,
// following the RFC 7518 Section 2 specification for unsigned integer decoding.
//
// Parameters:
//   - str: The base64url-encoded string to decode
//
// Returns:
//   - *big.Int: The decoded big.Int value
//   - error: An error if decoding fails
//
// Special cases:
//   - Empty strings decode to zero (big.Int{})
//   - Invalid base64url strings return an error
//
// Example:
//
//	decoded, err := base64x.Base64UrlUint.Decode("OTk=")
//	// Result: big.NewInt(12345), nil
//
//	// Empty string
//	decoded, err := base64x.Base64UrlUint.Decode("")
//	// Result: &big.Int{}, nil
//
//	// Invalid string
//	decoded, err := base64x.Base64UrlUint.Decode("invalid")
//	// Result: nil, error
func (base64UrlUint) Decode(str string) (*big.Int, error) {
	if str == "" {
		return &big.Int{}, nil // Return zero value for empty string
	}
	b, err := RawURLEncoding.Decode(str, true)
	if err != nil {
		return nil, err
	}
	bint := &big.Int{}
	return bint.SetBytes(b), nil
}
