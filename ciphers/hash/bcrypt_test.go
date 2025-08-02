package hash

import (
	"strings"
	"testing"
)

func TestBCrypt(t *testing.T) {
	// test cases
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "empty string",
			input:   "",
			wantErr: false,
		},
		{
			name:    "simple password",
			input:   "password123",
			wantErr: false,
		},
		{
			name:    "unicode password",
			input:   "密码测试🌟",
			wantErr: false,
		},
		{
			name:    "single character",
			input:   "a",
			wantErr: false,
		},
		{
			name:    "exactly 72 bytes",
			input:   strings.Repeat("a", 72),
			wantErr: false,
		},
		{
			name:    "73 bytes - should fail",
			input:   strings.Repeat("a", 73),
			wantErr: true,
		},
		{
			name:    "very long string - should fail",
			input:   strings.Repeat("a", 1000),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := BCrypt(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("BCrypt(%q) should return error, got nil", tt.input)
				}
				if hash != "" {
					t.Errorf("BCrypt(%q) should return empty hash when error occurs, got %q", tt.input, hash)
				}
				return
			}

			if err != nil {
				t.Errorf("BCrypt(%q) returned error: %v", tt.input, err)
				return
			}

			if hash == "" {
				t.Errorf("BCrypt(%q) returned empty hash, want non-empty", tt.input)
				return
			}

			// verify the hash is correct
			if !BCryptCompare(hash, tt.input) {
				t.Errorf("BCrypt(%q) hash verification failed", tt.input)
			}
		})
	}

	// edge case: multiple hashes for the same input should be different
	t.Run("hashes should be different for same input", func(t *testing.T) {
		input := "repeatPassword"
		hash1, err1 := BCrypt(input)
		hash2, err2 := BCrypt(input)

		if err1 != nil || err2 != nil {
			t.Errorf("BCrypt returned error for input %q: %v, %v", input, err1, err2)
			return
		}

		if hash1 == "" || hash2 == "" {
			t.Errorf("BCrypt returned empty hash for input %q", input)
			return
		}

		if hash1 == hash2 {
			t.Errorf("BCrypt generated identical hashes for the same input, want different hashes")
		}
	})

	// test BCryptCompare with invalid inputs
	t.Run("BCryptCompare with long input", func(t *testing.T) {
		hash, err := BCrypt("valid")
		if err != nil {
			t.Fatalf("Failed to create valid hash: %v", err)
		}

		longInput := strings.Repeat("a", 100)
		// Should return false for long input
		if BCryptCompare(hash, longInput) {
			t.Errorf("BCryptCompare should return false for input longer than 72 bytes")
		}
	})
}
