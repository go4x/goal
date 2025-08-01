package ciphers

import (
	"testing"
)

func TestMD5(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			input:    "hello",
			expected: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			input:    "The quick brown fox jumps over the lazy dog",
			expected: "9e107d9d372bb6826bd81d3542a419d6",
		},
		{
			input:    "The quick brown fox jumps over the lazy dog.",
			expected: "e4d909c290d0fb1ca068ffaddf22cbd0",
		},
	}

	for _, tt := range tests {
		result := MD5(tt.input)
		if result != tt.expected {
			t.Errorf("MD5(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
