package is_test

import (
	"testing"

	"github.com/go4x/goal/is"
)

// TestTrue tests the True function
func TestTrue(t *testing.T) {
	t.Run("true values", func(t *testing.T) {
		if !is.True(true) {
			t.Error("True should return true for true")
		}
	})
	t.Run("false values", func(t *testing.T) {
		if is.True(false) {
			t.Error("True should return false for false")
		}
	})
}

// TestFalse tests the False function
func TestFalse(t *testing.T) {
	t.Run("true values", func(t *testing.T) {
		if is.False(true) {
			t.Error("False should return false for true")
		}
	})
	t.Run("false values", func(t *testing.T) {
		if !is.False(false) {
			t.Error("False should return true for false")
		}
	})
}

// TestIsZero tests the IsZero function
func TestIsZero(t *testing.T) {
	t.Run("zero values", func(t *testing.T) {
		if !is.Zero(0) {
			t.Error("IsZero should return true for 0")
		}
		if !is.Zero("") {
			t.Error("IsZero should return true for empty string")
		}
		if !is.Zero(false) {
			t.Error("IsZero should return true for false")
		}
	})

	t.Run("non-zero values", func(t *testing.T) {
		if is.Zero(42) {
			t.Error("IsZero should return false for 42")
		}
		if is.Zero("hello") {
			t.Error("IsZero should return false for 'hello'")
		}
		if is.Zero(true) {
			t.Error("IsZero should return false for true")
		}
	})
}

// TestIsNotZero tests the IsNotZero function
func TestIsNotZero(t *testing.T) {
	t.Run("zero values", func(t *testing.T) {
		if is.NotZero(0) {
			t.Error("IsNotZero should return false for 0")
		}
		if is.NotZero("") {
			t.Error("IsNotZero should return false for empty string")
		}
	})

	t.Run("non-zero values", func(t *testing.T) {
		if !is.NotZero(42) {
			t.Error("IsNotZero should return true for 42")
		}
		if !is.NotZero("hello") {
			t.Error("IsNotZero should return true for 'hello'")
		}
	})
}

// TestIsNil tests the IsNil function
func TestIsNil(t *testing.T) {
	t.Run("nil values", func(t *testing.T) {
		var ptr *int
		if !is.Nil(ptr) {
			t.Error("IsNil should return true for nil pointer")
		}
		if !is.Nil(nil) {
			t.Error("IsNil should return true for nil")
		}
		if !is.Nil((*int)(nil)) {
			t.Error("IsNil should return true for nil pointer cast")
		}
	})

	t.Run("non-nil values", func(t *testing.T) {
		val := 42
		ptr := &val
		if is.Nil(ptr) {
			t.Error("IsNil should return false for non-nil pointer")
		}
		if is.Nil([]int{}) {
			t.Error("IsNil should return false for empty slice")
		}
		if is.Nil(map[string]int{}) {
			t.Error("IsNil should return false for empty map")
		}
	})
}

// TestIsNotNil tests the IsNotNil function
func TestIsNotNil(t *testing.T) {
	t.Run("nil values", func(t *testing.T) {
		var ptr *int
		if is.NotNil(ptr) {
			t.Error("IsNotNil should return false for nil pointer")
		}
		if is.NotNil(nil) {
			t.Error("IsNotNil should return false for nil")
		}
	})

	t.Run("non-nil values", func(t *testing.T) {
		val := 42
		ptr := &val
		if !is.NotNil(ptr) {
			t.Error("IsNotNil should return true for non-nil pointer")
		}
		if !is.NotNil([]int{}) {
			t.Error("IsNotNil should return true for empty slice")
		}
	})
}

// TestIsEmpty tests the IsEmpty function
func TestIsEmpty(t *testing.T) {
	t.Run("empty values", func(t *testing.T) {
		if !is.Empty("") {
			t.Error("IsEmpty should return true for empty string")
		}
		if !is.Empty([]int{}) {
			t.Error("IsEmpty should return true for empty slice")
		}
		if !is.Empty(map[string]int{}) {
			t.Error("IsEmpty should return true for empty map")
		}
		if !is.Empty(0) {
			t.Error("IsEmpty should return true for 0")
		}
		if !is.Empty(nil) {
			t.Error("IsEmpty should return true for nil")
		}
	})

	t.Run("non-empty values", func(t *testing.T) {
		if is.Empty("hello") {
			t.Error("IsEmpty should return false for 'hello'")
		}
		if is.Empty([]int{1, 2, 3}) {
			t.Error("IsEmpty should return false for non-empty slice")
		}
		if is.Empty(map[string]int{"a": 1}) {
			t.Error("IsEmpty should return false for non-empty map")
		}
		if is.Empty(42) {
			t.Error("IsEmpty should return false for 42")
		}
	})
}

// TestIsNotEmpty tests the IsNotEmpty function
func TestIsNotEmpty(t *testing.T) {
	t.Run("empty values", func(t *testing.T) {
		if is.NotEmpty("") {
			t.Error("IsNotEmpty should return false for empty string")
		}
		if is.NotEmpty([]int{}) {
			t.Error("IsNotEmpty should return false for empty slice")
		}
		if is.NotEmpty(0) {
			t.Error("IsNotEmpty should return false for 0")
		}
	})

	t.Run("non-empty values", func(t *testing.T) {
		if !is.NotEmpty("hello") {
			t.Error("IsNotEmpty should return true for 'hello'")
		}
		if !is.NotEmpty([]int{1, 2, 3}) {
			t.Error("IsNotEmpty should return true for non-empty slice")
		}
		if !is.NotEmpty(42) {
			t.Error("IsNotEmpty should return true for 42")
		}
	})
}

// TestEq tests the Eq function
func TestEq(t *testing.T) {
	t.Run("basic equal values", func(t *testing.T) {
		if !is.Eq(42, 42) {
			t.Error("Eq should return true for equal values")
		}
		if !is.Eq("hello", "hello") {
			t.Error("Eq should return true for equal strings")
		}
	})

	t.Run("basic not equal values", func(t *testing.T) {
		if is.Eq(42, 43) {
			t.Error("Eq should return false for different values")
		}
		if is.Eq("hello", "world") {
			t.Error("Eq should return false for different strings")
		}
	})

	t.Run("struct equality using == path", func(t *testing.T) {
		type simple struct {
			A int
			B string
		}
		a := simple{A: 1, B: "x"}
		b := simple{A: 1, B: "x"}
		c := simple{A: 2, B: "x"}

		if !is.Eq(a, b) {
			t.Error("Eq should return true for equal structs")
		}
		if is.Eq(a, c) {
			t.Error("Eq should return false for different structs")
		}
	})

	t.Run("pointer equality by value", func(t *testing.T) {
		x1, x2 := 10, 10
		p1, p2 := &x1, &x2
		p3 := &x1

		if !is.Eq(p1, p2) {
			t.Error("Eq should treat pointers to equal values as equal")
		}
		if !is.Eq(p1, p3) {
			t.Error("Eq should treat pointers to same value as equal")
		}

		var pNil *int
		if !is.Eq(pNil, (*int)(nil)) {
			t.Error("Eq should treat two nil pointers as equal")
		}
		if is.Eq(pNil, p1) {
			t.Error("Eq should treat nil and non-nil pointers as not equal")
		}
	})

	t.Run("interface values", func(t *testing.T) {
		var a interface{} = 42
		var b interface{} = 42
		var c interface{} = 43

		if !is.Eq(a, b) {
			t.Error("Eq should return true for equal interface values")
		}
		if is.Eq(a, c) {
			t.Error("Eq should return false for different interface values")
		}

		var iNil1 interface{}
		var iNil2 interface{}
		if !is.Eq(iNil1, iNil2) {
			t.Error("Eq should treat two nil interfaces as equal")
		}
		if is.Eq(iNil1, a) {
			t.Error("Eq should treat nil and non-nil interfaces as not equal")
		}
	})

	t.Run("slice and map deep equality", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{1, 2, 3}
		s3 := []int{1, 2, 4}

		if !is.Eq(s1, s2) {
			t.Error("Eq should return true for equal slices (deep comparison)")
		}
		if is.Eq(s1, s3) {
			t.Error("Eq should return false for different slices (deep comparison)")
		}

		m1 := map[string]int{"a": 1}
		m2 := map[string]int{"a": 1}
		m3 := map[string]int{"a": 2}

		if !is.Eq(m1, m2) {
			t.Error("Eq should return true for equal maps (deep comparison)")
		}
		if is.Eq(m1, m3) {
			t.Error("Eq should return false for different maps (deep comparison)")
		}
	})

	t.Run("channels", func(t *testing.T) {
		ch1 := make(chan int)
		ch2 := ch1
		ch3 := make(chan int)

		if !is.Eq(ch1, ch2) {
			t.Error("Eq should return true for same channel reference")
		}
		if is.Eq(ch1, ch3) {
			t.Error("Eq should return false for different channel references")
		}
	})

	t.Run("functions", func(t *testing.T) {
		type Func func() int
		var v1 Func = nil
		var v2 Func = nil
		if !is.Eq(v1, v2) {
			t.Error("Eq should return true for nil functions")
		}

		f1 := func() int { return 1 }
		f2 := f1
		f3 := func() int { return 1 }

		if is.Eq(f1, f2) {
			t.Error("Eq should return false for same function reference because they are not nil")
		}
		if is.Eq(f1, f3) {
			t.Error("Eq should return false for different function values")
		}
	})

	type Person struct {
		Name string
		Age  int
	}
	p1 := Person{Name: "John", Age: 30}
	p2 := Person{Name: "John", Age: 30}
	p3 := Person{Name: "Jane", Age: 30}
	if !is.Eq(p1, p2) {
		t.Error("Eq should return true for equal structs")
	}
	if is.Eq(p1, p3) {
		t.Error("Eq should return false for different structs")
	}
}

// TestSame tests the Same function (alias of Eq)
func TestSame(t *testing.T) {
	if !is.Same(42, 42) {
		t.Error("Same should return true for equal values")
	}
	if is.Same(42, 43) {
		t.Error("Same should return false for different values")
	}
}

// TestNeq tests the Neq function
func TestNeq(t *testing.T) {
	t.Run("equal values", func(t *testing.T) {
		if is.Neq(42, 42) {
			t.Error("Neq should return false for equal values")
		}
		if is.Neq("hello", "hello") {
			t.Error("Neq should return false for equal strings")
		}
	})

	t.Run("not equal values", func(t *testing.T) {
		if !is.Neq(42, 43) {
			t.Error("Neq should return true for different values")
		}
		if !is.Neq("hello", "world") {
			t.Error("Neq should return true for different strings")
		}
	})
}

// TestEdgeCases tests edge cases and error conditions
func TestEdgeCases(t *testing.T) {
	t.Run("nil interface values", func(t *testing.T) {
		var i interface{}
		if !is.Nil(i) {
			t.Error("IsNil should return true for nil interface")
		}
		if !is.Empty(i) {
			t.Error("IsEmpty should return true for nil interface")
		}
	})

	t.Run("zero value structs", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		var p Point
		if !is.Zero(p) {
			t.Error("IsZero should return true for zero value struct")
		}
		if !is.Empty(p) {
			t.Error("IsEmpty should return true for zero value struct")
		}
	})
}

// TestPerformance tests performance characteristics
func TestPerformance(t *testing.T) {
	t.Run("large slice operations", func(t *testing.T) {
		largeSlice := make([]int, 1000)
		for i := range largeSlice {
			largeSlice[i] = i
		}

		if is.Empty(largeSlice) {
			t.Error("IsEmpty should return false for large slice")
		}
		if !is.NotEmpty(largeSlice) {
			t.Error("IsNotEmpty should return true for large slice")
		}
	})

	t.Run("deep nesting", func(t *testing.T) {
		type Nested struct {
			Level1 struct {
				Level2 struct {
					Value int
				}
			}
		}

		n1 := Nested{}
		n1.Level1.Level2.Value = 42
		n2 := Nested{}
		n2.Level1.Level2.Value = 42

		if !is.Eq(n1, n2) {
			t.Error("EqDeep should return true for deeply nested equal structs")
		}
	})
}

// TestOrderComparisons tests Gt/Gte/Lt/Lte functions
func TestOrderComparisons(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		if !is.Gt(2, 1) {
			t.Error("Gt should return true for 2 > 1")
		}
		if is.Gt(1, 2) {
			t.Error("Gt should return false for 1 > 2")
		}

		if !is.Gte(2, 2) || !is.Gte(3, 2) {
			t.Error("Gte should return true for a >= b")
		}
		if is.Gte(1, 2) {
			t.Error("Gte should return false for 1 >= 2")
		}

		if !is.Lt(1, 2) {
			t.Error("Lt should return true for 1 < 2")
		}
		if is.Lt(2, 1) {
			t.Error("Lt should return false for 2 < 1")
		}

		if !is.Lte(2, 2) || !is.Lte(1, 2) {
			t.Error("Lte should return true for a <= b")
		}
		if is.Lte(3, 2) {
			t.Error("Lte should return false for 3 <= 2")
		}
	})

	t.Run("strings", func(t *testing.T) {
		if !is.Gt("b", "a") {
			t.Error("Gt should return true for \"b\" > \"a\"")
		}
		if !is.Lt("a", "b") {
			t.Error("Lt should return true for \"a\" < \"b\"")
		}
		if !is.Gte("b", "b") || !is.Lte("b", "b") {
			t.Error("Gte/Lte should return true for equal strings")
		}
	})
}
