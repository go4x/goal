package ciphers

import "strings"

const (
	base36chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fullchars   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=[]{}|;:'\"<>,./?`~"

	base36len    = len(base36chars)
	base62len    = len(base62chars)
	fullcharslen = len(fullchars)
)

// C36 converts a uint to a base36 string, only support positive number.
// This function is used to compress a large number to a short string which is case insensitive.
// The result string is uppercase.
func C36(n uint) string {
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

// C62 converts a uint to a base62 string, only support positive number.
// This function is used to compress a large number to a short string which is case sensitive.
func C62(n uint) string {
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

// C converts a uint to a full string, only support positive number.
// This function is used to compress a large number to a short string which is case sensitive.
// The result string is case sensitive.
func C(n uint) string {
	if n == 0 {
		return fullchars[0:1]
	}
	ss := make([]string, 0)
	for n != 0 {
		m := n % uint(fullcharslen)
		ss = append(ss, fullchars[m:m+1])
		n /= uint(fullcharslen)
	}
	b := strings.Builder{}
	for i := len(ss) - 1; i >= 0; i-- {
		b.WriteString(ss[i])
	}
	return b.String()
}
