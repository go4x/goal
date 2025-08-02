package base64x

import (
	"math/big"
)

var Base64UrlUint base64UrlUint

type base64UrlUint struct {
}

// Encode returns the base64url-encoded representation
// of the big-endian octet sequence as defined in
// [RFC 7518 2](https://www.rfc-editor.org/rfc/rfc7518.html#section-2)
func (base64UrlUint) Encode(i *big.Int) string {
	// Get the big-endian bytes
	bytes := i.Bytes()

	// Handle zero case - return "AA" for zero
	if len(bytes) == 0 || (len(bytes) == 1 && bytes[0] == 0) {
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
func (base64UrlUint) Decode(str string) (*big.Int, error) {
	if str == "" {
		return nil, nil // Return nil for empty string as expected by tests
	}
	b, err := RawURLEncoding.Decode(str, true)
	if err != nil {
		return nil, err
	}
	bint := &big.Int{}
	return bint.SetBytes(b), nil
}
