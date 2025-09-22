package ciphers

import (
	"fmt"
	"math"
	"testing"

	"github.com/go4x/goal/random"
	"github.com/go4x/got"
)

func TestBase36(t *testing.T) {
	tests := []got.Case{
		got.NewCase("0", 0, "0", false, nil),
		got.NewCase("1", 1, "1", false, nil),
		got.NewCase("9", 9, "9", false, nil),
		got.NewCase("10", 10, "A", false, nil),
		got.NewCase("35", 35, "Z", false, nil),
		got.NewCase("36", 36, "10", false, nil),
		got.NewCase("71", 71, "1Z", false, nil),
		got.NewCase("12345", 12345, "9IX", false, nil),
		got.NewCase("4294967295", 4294967295, "1Z141Z3", false, nil),
	}

	r := got.New(t, "Base36")
	r.Cases(tests, func(c got.Case, tt *testing.T) {
		result := Base36(uint(c.Input().(int)))
		if result != c.Want().(string) {
			r.Fail("Base36(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		} else {
			r.Pass("Base36(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		}
	})
}

func TestBase36_Random(t *testing.T) {
	m := make(map[string]int)
	g := got.New(t, "Base36_random")
	for i := 0; i < 1000000; i++ {
		n := random.Int(math.MaxInt64)
		s := Base36(uint(n))
		if _, ok := m[s]; ok {
			g.Require(m[s] == n, "Base36(%d) = %q; want %q", n, s, m[s])
		} else {
			m[s] = n
		}
	}
}

func TestBase62(t *testing.T) {
	tests := []got.Case{
		got.NewCase("0", 0, "0", false, nil),
		got.NewCase("1", 1, "1", false, nil),
		got.NewCase("9", 9, "9", false, nil),
		got.NewCase("10", 10, "a", false, nil),
		got.NewCase("35", 35, "z", false, nil),
		got.NewCase("36", 36, "A", false, nil),
		got.NewCase("61", 61, "Z", false, nil),
		got.NewCase("62", 62, "10", false, nil),
		got.NewCase("12345", 12345, "3d7", false, nil),
		got.NewCase("4294967295", 4294967295, "4GFfc3", false, nil),
	}

	r := got.New(t, "Base62")
	r.Cases(tests, func(c got.Case, tt *testing.T) {
		result := Base62(uint(c.Input().(int)))
		if result != c.Want().(string) {
			r.Fail("Base62(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		} else {
			r.Pass("Base62(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		}
	})
}

func TestBase62_Random(t *testing.T) {
	m := make(map[string]int)
	g := got.New(t, "Base62_random")
	for i := 0; i < 1000000; i++ {
		n := random.Int(math.MaxInt64)
		s := Base62(uint(n))
		if _, ok := m[s]; ok {
			g.Require(m[s] == n, "Base62(%d) = %q; want %q", n, s, m[s])
		} else {
			m[s] = n
		}
	}
}

// Edge cases tests for Base36
func TestBase36_EdgeCases(t *testing.T) {
	testCases := []struct {
		input    uint
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{35, "Z"},
		{36, "10"},
		{71, "1Z"},
		{72, "20"},
		{1295, "ZZ"},
		{1296, "100"},
		{46655, "ZZZ"},
		{46656, "1000"},
		{1679615, "ZZZZ"},
		{1679616, "10000"},
		{60466175, "ZZZZZ"},
		{60466176, "100000"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Base36_%d", tc.input), func(t *testing.T) {
			result := Base36(tc.input)
			if result != tc.expected {
				t.Errorf("Base36(%d) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

// Edge cases tests for Base62
func TestBase62_EdgeCases(t *testing.T) {
	testCases := []struct {
		input    uint
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{61, "Z"},
		{62, "10"},
		{123, "1Z"},
		{124, "20"},
		{3843, "ZZ"},
		{3844, "100"},
		{238327, "ZZZ"},
		{238328, "1000"},
		{14776335, "ZZZZ"},
		{14776336, "10000"},
		{916132831, "ZZZZZ"},
		{916132832, "100000"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Base62_%d", tc.input), func(t *testing.T) {
			result := Base62(tc.input)
			if result != tc.expected {
				t.Errorf("Base62(%d) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

// Test consistency between functions
func TestCompressionFunctions_Consistency(t *testing.T) {
	testNumbers := []uint{0, 1, 10, 100, 1000, 10000, 100000, 1000000, 10000000}

	for _, n := range testNumbers {
		t.Run(fmt.Sprintf("Consistency_%d", n), func(t *testing.T) {
			c36 := Base36(n)
			c62 := Base62(n)

			// All should be non-empty for non-zero inputs
			if n > 0 {
				if c36 == "" {
					t.Errorf("Base36(%d) returned empty string", n)
				}
				if c62 == "" {
					t.Errorf("Base62(%d) returned empty string", n)
				}
			}

			// All should return "0" for zero input
			if n == 0 {
				if c36 != "0" {
					t.Errorf("Base36(0) = %q; want \"0\"", c36)
				}
				if c62 != "0" {
					t.Errorf("Base62(0) = %q; want \"0\"", c62)
				}
			}

			// Base62 should generally be shorter than Base36 for large numbers
			if n > 1000000 {
				if len(c62) > len(c36) {
					t.Errorf("Base62 should be shorter than Base36 for large numbers: Base62(%d) = %q (len: %d), Base36(%d) = %q (len: %d)", n, c62, len(c62), n, c36, len(c36))
				}
			}
		})
	}
}

// Test compression ratio
func TestCompressionRatio(t *testing.T) {
	testNumbers := []uint{100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}

	for _, n := range testNumbers {
		t.Run(fmt.Sprintf("Ratio_%d", n), func(t *testing.T) {
			c36 := Base36(n)
			c62 := Base62(n)

			t.Logf("Number: %d", n)
			t.Logf("Base36: %s (len: %d)", c36, len(c36))
			t.Logf("Base62: %s (len: %d)", c62, len(c62))

			// Both should produce non-empty results for non-zero inputs
			if n > 0 {
				if c36 == "" {
					t.Errorf("Base36(%d) produced empty string", n)
				}
				if c62 == "" {
					t.Errorf("Base62(%d) produced empty string", n)
				}
			}

			// Base62 should be shorter than Base36 for larger numbers
			if n > 1000000 {
				if len(c62) > len(c36) {
					t.Errorf("Base62 should be shorter than Base36 for large numbers")
				}
			}
		})
	}
}

// Test large numbers
func TestLargeNumbers(t *testing.T) {
	largeNumbers := []uint{
		18446744073709551615, // uint64 max
		9999999999999999999,
		1234567890123456789,
		9876543210987654321,
	}

	for _, n := range largeNumbers {
		t.Run(fmt.Sprintf("Large_%d", n), func(t *testing.T) {
			c36 := Base36(n)
			c62 := Base62(n)

			// Both should produce non-empty results
			if c36 == "" {
				t.Errorf("Base36(%d) produced empty string", n)
			}
			if c62 == "" {
				t.Errorf("Base62(%d) produced empty string", n)
			}

			// Results should be reasonable length
			if len(c36) > 20 {
				t.Errorf("Base36(%d) produced unreasonably long string: %s (len: %d)", n, c36, len(c36))
			}
			if len(c62) > 20 {
				t.Errorf("Base62(%d) produced unreasonably long string: %s (len: %d)", n, c62, len(c62))
			}

			t.Logf("Base36(%d) = %s (len: %d)", n, c36, len(c36))
			t.Logf("Base62(%d) = %s (len: %d)", n, c62, len(c62))
		})
	}
}

// Benchmark tests for performance comparison
func BenchmarkBase36(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base36(123456789)
	}
}

func BenchmarkBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62(123456789)
	}
}

func BenchmarkBase36_Large(b *testing.B) {
	largeNum := uint(18446744073709551615)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base36(largeNum)
	}
}

func BenchmarkBase62_Large(b *testing.B) {
	largeNum := uint(18446744073709551615)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base62(largeNum)
	}
}

// Test Base36Decode function
func TestBase36Decode(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
		valid    bool
	}{
		{"0", 0, true},
		{"1", 1, true},
		{"9", 9, true},
		{"A", 10, true},
		{"Z", 35, true},
		{"10", 36, true},
		{"1Z", 71, true},
		{"9IX", 12345, true},
		{"1Z141Z3", 4294967295, true},
		{"", 0, false},
		{"a", 10, true}, // lowercase valid in base36 (case insensitive)
		{"z", 35, true}, // lowercase valid in base36 (case insensitive)
		{"@", 0, false}, // invalid character
		{" ", 0, false}, // space not valid
		{"!", 0, false}, // special char not valid
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Base36Decode_%s", tt.input), func(t *testing.T) {
			result, ok := Base36Decode(tt.input)
			if ok != tt.valid {
				t.Errorf("Base36Decode(%q) validity = %v; want %v", tt.input, ok, tt.valid)
			}
			if tt.valid && result != tt.expected {
				t.Errorf("Base36Decode(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Test Base62Decode function
func TestBase62Decode(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
		valid    bool
	}{
		{"0", 0, true},
		{"1", 1, true},
		{"9", 9, true},
		{"a", 10, true},
		{"z", 35, true},
		{"A", 36, true},
		{"Z", 61, true},
		{"10", 62, true},
		{"1Z", 123, true},
		{"3d7", 12345, true},
		{"4GFfc3", 4294967295, true},
		{"", 0, false},
		{"@", 0, false}, // invalid character
		{" ", 0, false}, // space not valid
		{"!", 0, false}, // special char not valid
		{"-", 0, false}, // dash not valid
		{"_", 0, false}, // underscore not valid
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Base62Decode_%s", tt.input), func(t *testing.T) {
			result, ok := Base62Decode(tt.input)
			if ok != tt.valid {
				t.Errorf("Base62Decode(%q) validity = %v; want %v", tt.input, ok, tt.valid)
			}
			if tt.valid && result != tt.expected {
				t.Errorf("Base62Decode(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Test round-trip encoding/decoding for Base36
func TestBase36_RoundTrip(t *testing.T) {
	testNumbers := []uint{0, 1, 9, 10, 35, 36, 71, 72, 12345, 4294967295, 999999999, 18446744073709551615}

	for _, n := range testNumbers {
		t.Run(fmt.Sprintf("RoundTrip_%d", n), func(t *testing.T) {
			encoded := Base36(n)
			decoded, ok := Base36Decode(encoded)

			if !ok {
				t.Errorf("Base36Decode failed for encoded string %q", encoded)
				return
			}

			if decoded != n {
				t.Errorf("Round trip failed: %d -> %q -> %d", n, encoded, decoded)
			}
		})
	}
}

// Test round-trip encoding/decoding for Base62
func TestBase62_RoundTrip(t *testing.T) {
	testNumbers := []uint{0, 1, 9, 10, 35, 36, 61, 62, 123, 124, 12345, 4294967295, 999999999, 18446744073709551615}

	for _, n := range testNumbers {
		t.Run(fmt.Sprintf("RoundTrip_%d", n), func(t *testing.T) {
			encoded := Base62(n)
			decoded, ok := Base62Decode(encoded)

			if !ok {
				t.Errorf("Base62Decode failed for encoded string %q", encoded)
				return
			}

			if decoded != n {
				t.Errorf("Round trip failed: %d -> %q -> %d", n, encoded, decoded)
			}
		})
	}
}

// Test round-trip with random numbers
func TestBase36_RoundTrip_Random(t *testing.T) {
	g := got.New(t, "Base36_RoundTrip_Random")
	for i := 0; i < 100000; i++ {
		n := random.Int(math.MaxInt64)
		encoded := Base36(uint(n))
		decoded, ok := Base36Decode(encoded)

		if !ok {
			g.Fail("Base36Decode failed for encoded string %q", encoded)
			continue
		}

		if decoded != uint(n) {
			g.Fail("Round trip failed: %d -> %q -> %d", n, encoded, decoded)
		}
	}
}

// Test round-trip with random numbers
func TestBase62_RoundTrip_Random(t *testing.T) {
	g := got.New(t, "Base62_RoundTrip_Random")
	for i := 0; i < 100000; i++ {
		n := random.Int(math.MaxInt64)
		encoded := Base62(uint(n))
		decoded, ok := Base62Decode(encoded)

		if !ok {
			g.Fail("Base62Decode failed for encoded string %q", encoded)
			continue
		}

		if decoded != uint(n) {
			g.Fail("Round trip failed: %d -> %q -> %d", n, encoded, decoded)
		}
	}
}

// Test edge cases for decode functions
func TestDecode_EdgeCases(t *testing.T) {
	// Test empty string
	result, ok := Base36Decode("")
	if ok || result != 0 {
		t.Errorf("Base36Decode(\"\") = (%d, %v); want (0, false)", result, ok)
	}

	result, ok = Base62Decode("")
	if ok || result != 0 {
		t.Errorf("Base62Decode(\"\") = (%d, %v); want (0, false)", result, ok)
	}

	// Test invalid characters
	invalidStrings := []string{"@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", " ", "\t", "\n"}

	for _, s := range invalidStrings {
		result, ok := Base36Decode(s)
		if ok {
			t.Errorf("Base36Decode(%q) should be invalid but got (%d, %v)", s, result, ok)
		}

		result, ok = Base62Decode(s)
		if ok {
			t.Errorf("Base62Decode(%q) should be invalid but got (%d, %v)", s, result, ok)
		}
	}

	// Test case sensitivity for Base36 (should be case insensitive)
	testCases := []struct {
		upper    string
		lower    string
		expected uint
	}{
		{"A", "a", 10},
		{"Z", "z", 35},
		{"AB", "ab", 10*36 + 11},
	}

	for _, tc := range testCases {
		// Base36 should accept both upper and lower case
		upperResult, upperOk := Base36Decode(tc.upper)
		lowerResult, lowerOk := Base36Decode(tc.lower)

		if !upperOk || upperResult != tc.expected {
			t.Errorf("Base36Decode(%q) = (%d, %v); want (%d, true)", tc.upper, upperResult, upperOk, tc.expected)
		}

		if !lowerOk || lowerResult != tc.expected {
			t.Errorf("Base36Decode(%q) = (%d, %v); want (%d, true)", tc.lower, lowerResult, lowerOk, tc.expected)
		}
	}
}

// Benchmark tests for decode functions
func BenchmarkBase36Decode(b *testing.B) {
	testString := "1Z141Z3" // 4294967295
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base36Decode(testString)
	}
}

func BenchmarkBase62Decode(b *testing.B) {
	testString := "4GFfc3" // 4294967295
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base62Decode(testString)
	}
}

func BenchmarkBase36Decode_Large(b *testing.B) {
	testString := "3W5E11264SGSF" // 18446744073709551615
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base36Decode(testString)
	}
}

func BenchmarkBase62Decode_Large(b *testing.B) {
	testString := "lYGhA16ahyf" // 18446744073709551615
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base62Decode(testString)
	}
}
