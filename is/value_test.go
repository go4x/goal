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

// TestEqual tests the Equal function
func TestEqual(t *testing.T) {
	t.Run("equal values", func(t *testing.T) {
		if !is.Equal(42, 42) {
			t.Error("Equal should return true for equal values")
		}
		if !is.Equal("hello", "hello") {
			t.Error("Equal should return true for equal strings")
		}
	})

	t.Run("not equal values", func(t *testing.T) {
		if is.Equal(42, 43) {
			t.Error("Equal should return false for different values")
		}
		if is.Equal("hello", "world") {
			t.Error("Equal should return false for different strings")
		}
	})
}

// TestNotEqual tests the NotEqual function
func TestNotEqual(t *testing.T) {
	t.Run("equal values", func(t *testing.T) {
		if is.NotEqual(42, 42) {
			t.Error("NotEqual should return false for equal values")
		}
		if is.NotEqual("hello", "hello") {
			t.Error("NotEqual should return false for equal strings")
		}
	})

	t.Run("not equal values", func(t *testing.T) {
		if !is.NotEqual(42, 43) {
			t.Error("NotEqual should return true for different values")
		}
		if !is.NotEqual("hello", "world") {
			t.Error("NotEqual should return true for different strings")
		}
	})
}

// TestDeepEqual tests the DeepEqual function
func TestDeepEqual(t *testing.T) {
	t.Run("equal slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{1, 2, 3}
		if !is.DeepEqual(slice1, slice2) {
			t.Error("DeepEqual should return true for equal slices")
		}
	})

	t.Run("equal maps", func(t *testing.T) {
		map1 := map[string]int{"a": 1, "b": 2}
		map2 := map[string]int{"a": 1, "b": 2}
		if !is.DeepEqual(map1, map2) {
			t.Error("DeepEqual should return true for equal maps")
		}
	})

	t.Run("not equal slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{1, 2, 4}
		if is.DeepEqual(slice1, slice2) {
			t.Error("DeepEqual should return false for different slices")
		}
	})

	t.Run("equal structs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		p1 := Person{Name: "Alice", Age: 30}
		p2 := Person{Name: "Alice", Age: 30}
		if !is.DeepEqual(p1, p2) {
			t.Error("DeepEqual should return true for equal structs")
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

		if !is.DeepEqual(n1, n2) {
			t.Error("DeepEqual should return true for deeply nested equal structs")
		}
	})
}
