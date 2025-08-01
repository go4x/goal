package ciphers

import (
	"math"
	"testing"

	"github.com/gophero/goal/random"
	"github.com/gophero/got"
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
		g.Logf("C(%d) = %q", n, s)
		if _, ok := m[s]; ok {
			g.Require(m[s] == n, "C(%d) = %q; want %q", n, s, m[s])
		} else {
			m[s] = n
		}
	}
}
