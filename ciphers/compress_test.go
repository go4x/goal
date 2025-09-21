package ciphers

import (
	"fmt"
	"math"
	"testing"

	"github.com/go4x/goal/random"
	"github.com/go4x/got"
)

func TestC36(t *testing.T) {
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
		result := C36(uint(c.Input().(int)))
		if result != c.Want().(string) {
			r.Fail("Base36(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		} else {
			r.Pass("Base36(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		}
	})
}

func TestC36_Random(t *testing.T) {
	m := make(map[string]int)
	g := got.New(t, "Base36_random")
	for i := 0; i < 1000000; i++ {
		n := random.Int(math.MaxInt64)
		s := C36(uint(n))
		if _, ok := m[s]; ok {
			g.Require(m[s] == n, "Base36(%d) = %q; want %q", n, s, m[s])
		} else {
			m[s] = n
		}
	}
}

func TestC62(t *testing.T) {
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
		result := C62(uint(c.Input().(int)))
		if result != c.Want().(string) {
			r.Fail("Base62(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		} else {
			r.Pass("Base62(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		}
	})
}

func TestC62_Random(t *testing.T) {
	m := make(map[string]int)
	g := got.New(t, "Base62_random")
	for i := 0; i < 1000000; i++ {
		n := random.Int(math.MaxInt64)
		s := C62(uint(n))
		if _, ok := m[s]; ok {
			g.Require(m[s] == n, "Base62(%d) = %q; want %q", n, s, m[s])
		} else {
			m[s] = n
		}
	}
}

func TestC(t *testing.T) {
	tests := []got.Case{
		got.NewCase("0", 0, "0", false, nil),
		got.NewCase("1", 1, "1", false, nil),
		got.NewCase("9", 9, "9", false, nil),
		got.NewCase("10", 10, "a", false, nil),
		got.NewCase("35", 35, "z", false, nil),
		got.NewCase("36", 36, "A", false, nil),
		got.NewCase("61", 61, "Z", false, nil),
		got.NewCase("62", 62, "!", false, nil),
		got.NewCase("12345", 12345, "1D*", false, nil),
		got.NewCase("4294967295", 4294967295, "VCWD3", false, nil),
	}

	r := got.New(t, "C")
	r.Cases(tests, func(c got.Case, tt *testing.T) {
		result := C(uint(c.Input().(int)))
		if result != c.Want().(string) {
			r.Fail("C(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		} else {
			r.Pass("C(%d) = %q; want %q", c.Input().(int), result, c.Want().(string))
		}
	})
}

func TestC_Random(t *testing.T) {
	m := make(map[string]int)
	g := got.New(t, "C_random")
	for i := 0; i < 1000000; i++ {
		n := random.Int(math.MaxInt64)
		s := C(uint(n))
		// g.Logf("C(%d) = %q", n, s)
		if _, ok := m[s]; ok {
			g.Require(m[s] == n, "C(%d) = %q; want %q", n, s, m[s])
		} else {
			m[s] = n
		}
	}
}

func TestC36_EdgeCases(t *testing.T) {
	// Test edge cases for C36
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
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("C36_%d", tc.input), func(t *testing.T) {
			result := C36(tc.input)
			if result != tc.expected {
				t.Errorf("C36(%d) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestC62_EdgeCases(t *testing.T) {
	// Test edge cases for C62
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
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("C62_%d", tc.input), func(t *testing.T) {
			result := C62(tc.input)
			if result != tc.expected {
				t.Errorf("C62(%d) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestC_EdgeCases(t *testing.T) {
	// Test edge cases for C
	testCases := []struct {
		input    uint
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{95, "12"},
		{96, "13"},
		{191, "25"},
		{192, "26"},
		{9215, "168"},
		{9216, "169"},
		{884735, "19rq"},
		{884736, "19rr"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("C_%d", tc.input), func(t *testing.T) {
			result := C(tc.input)
			if result != tc.expected {
				t.Errorf("C(%d) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestCompressionFunctions_Consistency(t *testing.T) {
	// Test that all compression functions are consistent
	testNumbers := []uint{0, 1, 10, 100, 1000, 10000, 100000, 1000000, 10000000}

	for _, n := range testNumbers {
		t.Run(fmt.Sprintf("Consistency_%d", n), func(t *testing.T) {
			c36 := C36(n)
			c62 := C62(n)
			c := C(n)

			// All should be non-empty for non-zero inputs
			if n > 0 {
				if c36 == "" {
					t.Errorf("C36(%d) returned empty string", n)
				}
				if c62 == "" {
					t.Errorf("C62(%d) returned empty string", n)
				}
				if c == "" {
					t.Errorf("C(%d) returned empty string", n)
				}
			}

			// All should return "0" for zero input
			if n == 0 {
				if c36 != "0" {
					t.Errorf("C36(0) = %q; want \"0\"", c36)
				}
				if c62 != "0" {
					t.Errorf("C62(0) = %q; want \"0\"", c62)
				}
				if c != "0" {
					t.Errorf("C(0) = %q; want \"0\"", c)
				}
			}
		})
	}
}
