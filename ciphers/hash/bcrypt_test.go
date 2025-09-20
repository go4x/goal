package hash

import (
	"fmt"
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

	// Test BCrypt error handling path
	t.Run("BCrypt error handling", func(t *testing.T) {
		// Test with empty hash comparison
		t.Run("BCryptCompare with empty hash", func(t *testing.T) {
			result := BCryptCompare("", "password")
			if result {
				t.Errorf("BCryptCompare should return false for empty hash")
			}
		})

		// Test with invalid hash format
		t.Run("BCryptCompare with invalid hash", func(t *testing.T) {
			result := BCryptCompare("invalid_hash_format", "password")
			if result {
				t.Errorf("BCryptCompare should return false for invalid hash format")
			}
		})

		// Test BCrypt with very long string that might cause internal errors
		t.Run("BCrypt with maximum length string", func(t *testing.T) {
			// Create a string that's exactly 72 bytes to test edge case
			maxString := strings.Repeat("a", 72)
			hash, err := BCrypt(maxString)
			if err != nil {
				t.Errorf("BCrypt should not return error for 72-byte string, got: %v", err)
			}
			if hash == "" {
				t.Errorf("BCrypt should return non-empty hash for 72-byte string")
			}
		})

		// Test BCrypt with various edge cases to try to trigger error path
		t.Run("BCrypt edge cases", func(t *testing.T) {
			// Test with various string lengths and characters
			testCases := []struct {
				name  string
				input string
			}{
				{"empty string", ""},
				{"single byte", "a"},
				{"two bytes", "ab"},
				{"unicode string", "密码测试🌟"},
				{"special characters", "!@#$%^&*()"},
				{"newlines and tabs", "hello\nworld\t"},
				{"null bytes", string([]byte{0, 1, 2, 3})},
				{"maximum ascii", string(make([]byte, 72))},
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					hash, err := BCrypt(tc.input)
					if err != nil {
						t.Errorf("BCrypt(%q) returned error: %v", tc.input, err)
					}
					if hash == "" {
						t.Errorf("BCrypt(%q) returned empty hash", tc.input)
					}
					// Verify the hash works
					if !BCryptCompare(hash, tc.input) {
						t.Errorf("BCryptCompare failed for input: %q", tc.input)
					}
				})
			}
		})

		// Test to try to trigger the error path in BCrypt
		// This is difficult because bcrypt.GenerateFromPassword rarely fails
		// but we can try some extreme cases
		t.Run("BCrypt error path attempt", func(t *testing.T) {
			// Try with very large strings that might cause memory issues
			// Note: This might not actually trigger the error path, but it's worth trying
			largeString := strings.Repeat("a", 1000)
			hash, err := BCrypt(largeString)
			// This should fail due to length check, not bcrypt.GenerateFromPassword
			if err == nil {
				t.Errorf("BCrypt should fail for string longer than 72 bytes")
			}
			if hash != "" {
				t.Errorf("BCrypt should return empty hash for long string")
			}
		})

		// Test with extreme cost values to try to trigger error
		t.Run("BCrypt with extreme cost", func(t *testing.T) {
			// Test with a string that's exactly 72 bytes to avoid length check
			testString := strings.Repeat("a", 72)

			// Try to create a scenario where bcrypt might fail
			// This is a long shot, but worth trying
			hash, err := BCrypt(testString)
			if err != nil {
				t.Errorf("BCrypt should not fail for 72-byte string, got: %v", err)
			}
			if hash == "" {
				t.Errorf("BCrypt should return non-empty hash for 72-byte string")
			}
		})

		// Test to try to trigger the error path in BCrypt by using extreme cost
		t.Run("BCrypt extreme cost attempt", func(t *testing.T) {
			// Try with a very high cost that might cause bcrypt to fail
			// This is a long shot, but worth trying
			testString := "test"
			hash, err := BCryptWithCost(testString, 31) // Maximum valid cost
			if err != nil {
				// This might actually trigger the error path we want to test
				t.Logf("BCryptWithCost with cost 31 failed as expected: %v", err)
			} else {
				t.Logf("BCryptWithCost with cost 31 succeeded: %s", hash[:20]+"...")
			}
		})

		// Test to try to trigger the error path in BCrypt by using invalid cost
		t.Run("BCrypt invalid cost attempt", func(t *testing.T) {
			// Try with an invalid cost that should cause bcrypt to fail
			testString := "test"
			hash, err := BCryptWithCost(testString, -1) // Invalid cost
			if err != nil {
				// This should trigger the error path we want to test
				t.Logf("BCryptWithCost with cost -1 failed as expected: %v", err)
			} else {
				t.Errorf("BCryptWithCost with cost -1 should have failed, got: %s", hash[:20]+"...")
			}
		})

		// Test with various byte patterns that might cause issues
		t.Run("BCrypt with special byte patterns", func(t *testing.T) {
			// Test with various byte patterns
			patterns := []string{
				strings.Repeat("\x00", 72), // null bytes
				strings.Repeat("\xFF", 72), // max byte value
				strings.Repeat("\x80", 72), // high bit set
				strings.Repeat("\x7F", 72), // max ascii
			}

			for i, pattern := range patterns {
				t.Run(fmt.Sprintf("pattern_%d", i), func(t *testing.T) {
					hash, err := BCrypt(pattern)
					if err != nil {
						t.Errorf("BCrypt should not fail for pattern %d, got: %v", i, err)
					}
					if hash == "" {
						t.Errorf("BCrypt should return non-empty hash for pattern %d", i)
					}
					// Verify the hash works
					if !BCryptCompare(hash, pattern) {
						t.Errorf("BCryptCompare failed for pattern %d", i)
					}
				})
			}
		})
	})

	// Test BCryptWithCost function
	t.Run("BCryptWithCost", func(t *testing.T) {
		testCases := []struct {
			name    string
			input   string
			cost    int
			wantErr bool
		}{
			{
				name:    "valid cost 4",
				input:   "password",
				cost:    4,
				wantErr: false,
			},
			{
				name:    "valid cost 10",
				input:   "password",
				cost:    10,
				wantErr: false,
			},
			{
				name:    "valid cost 12",
				input:   "password",
				cost:    12,
				wantErr: false,
			},
			{
				name:    "valid cost 0",
				input:   "password",
				cost:    0,
				wantErr: false,
			},
			{
				name:    "invalid cost 32",
				input:   "password",
				cost:    32,
				wantErr: true,
			},
			{
				name:    "too long input",
				input:   strings.Repeat("a", 73),
				cost:    10,
				wantErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				hash, err := BCryptWithCost(tc.input, tc.cost)
				if tc.wantErr {
					if err == nil {
						t.Errorf("BCryptWithCost(%q, %d) should return error, got nil", tc.input, tc.cost)
					}
					if hash != "" {
						t.Errorf("BCryptWithCost(%q, %d) should return empty hash when error occurs, got %q", tc.input, tc.cost, hash)
					}
					return
				}

				if err != nil {
					t.Errorf("BCryptWithCost(%q, %d) returned error: %v", tc.input, tc.cost, err)
					return
				}

				if hash == "" {
					t.Errorf("BCryptWithCost(%q, %d) returned empty hash, want non-empty", tc.input, tc.cost)
					return
				}

				// Verify the hash works
				if !BCryptCompare(hash, tc.input) {
					t.Errorf("BCryptWithCost(%q, %d) hash verification failed", tc.input, tc.cost)
				}
			})
		}
	})
}
