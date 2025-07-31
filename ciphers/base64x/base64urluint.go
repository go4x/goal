package base64x

import (
	"errors"
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

	// The octet sequence MUST utilize the minimum number of octets
	// needed to represent the value.
	for i, val := range bytes {
		if val > 0 {
			return RawURLEncoding.Encode(bytes[i:])
		}
	}

	return RawURLEncoding.Encode([]byte{0})
}

// Decode returns the BigInt represented by the base64url-encoded string.
func (base64UrlUint) Decode(str string) (*big.Int, error) {
	if str == "" {
		return nil, errors.New("input is empty")
	}
	b, err := RawURLEncoding.Decode(str, true)
	if err != nil {
		return nil, err
	}
	bint := &big.Int{}
	return bint.SetBytes(b), nil
}
