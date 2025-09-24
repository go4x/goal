package value_test

import (
	"errors"
	"testing"

	"github.com/go4x/goal/value"
)

// TestMust tests the Must function
func TestMust(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		result := value.Must(42, nil)
		if result != 42 {
			t.Errorf("Must should return 42, got %d", result)
		}
	})

	t.Run("panic on error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Must should panic when error is not nil")
			}
		}()
		value.Must(0, errors.New("test error"))
	})
}

// TestIfElse tests the IfElse function
func TestIfElse(t *testing.T) {
	t.Run("true condition", func(t *testing.T) {
		result := value.IfElse(true, "yes", "no")
		if result != "yes" {
			t.Errorf("IfElse should return 'yes' for true condition, got %s", result)
		}
	})

	t.Run("false condition", func(t *testing.T) {
		result := value.IfElse(false, "yes", "no")
		if result != "no" {
			t.Errorf("IfElse should return 'no' for false condition, got %s", result)
		}
	})

	t.Run("with integers", func(t *testing.T) {
		result := value.IfElse(5 > 3, 100, 200)
		if result != 100 {
			t.Errorf("IfElse should return 100 for true condition, got %d", result)
		}
	})
}

// TestOr tests the Or function
func TestOr(t *testing.T) {
	t.Run("first non-zero value", func(t *testing.T) {
		result := value.Or("", "", "fallback", "ignored")
		if result != "fallback" {
			t.Errorf("Or should return 'fallback', got %s", result)
		}
	})

	t.Run("all zero values", func(t *testing.T) {
		result := value.Or("", "", "")
		if result != "" {
			t.Errorf("Or should return empty string for all zero values, got %s", result)
		}
	})

	t.Run("with integers", func(t *testing.T) {
		result := value.Or(0, 0, 42, 100)
		if result != 42 {
			t.Errorf("Or should return 42, got %d", result)
		}
	})

	t.Run("single value", func(t *testing.T) {
		result := value.Or("single")
		if result != "single" {
			t.Errorf("Or should return 'single', got %s", result)
		}
	})
}

// TestOrElse tests the OrElse function
func TestOrElse(t *testing.T) {
	t.Run("first non-zero value", func(t *testing.T) {
		result := value.OrElse("default", "", "", "fallback")
		if result != "fallback" {
			t.Errorf("OrElse should return 'fallback', got %s", result)
		}
	})

	t.Run("all zero values", func(t *testing.T) {
		result := value.OrElse("default", "", "", "")
		if result != "default" {
			t.Errorf("OrElse should return 'default' for all zero values, got %s", result)
		}
	})

	t.Run("with integers", func(t *testing.T) {
		result := value.OrElse(999, 0, 0, 42)
		if result != 42 {
			t.Errorf("OrElse should return 42, got %d", result)
		}
	})
}

// TestIsZero tests the IsZero function
func TestIsZero(t *testing.T) {
	t.Run("zero values", func(t *testing.T) {
		if !value.IsZero(0) {
			t.Error("IsZero should return true for 0")
		}
		if !value.IsZero("") {
			t.Error("IsZero should return true for empty string")
		}
		if !value.IsZero(false) {
			t.Error("IsZero should return true for false")
		}
	})

	t.Run("non-zero values", func(t *testing.T) {
		if value.IsZero(42) {
			t.Error("IsZero should return false for 42")
		}
		if value.IsZero("hello") {
			t.Error("IsZero should return false for 'hello'")
		}
		if value.IsZero(true) {
			t.Error("IsZero should return false for true")
		}
	})
}

// TestIsNotZero tests the IsNotZero function
func TestIsNotZero(t *testing.T) {
	t.Run("zero values", func(t *testing.T) {
		if value.IsNotZero(0) {
			t.Error("IsNotZero should return false for 0")
		}
		if value.IsNotZero("") {
			t.Error("IsNotZero should return false for empty string")
		}
	})

	t.Run("non-zero values", func(t *testing.T) {
		if !value.IsNotZero(42) {
			t.Error("IsNotZero should return true for 42")
		}
		if !value.IsNotZero("hello") {
			t.Error("IsNotZero should return true for 'hello'")
		}
	})
}

// TestIsNil tests the IsNil function
func TestIsNil(t *testing.T) {
	t.Run("nil values", func(t *testing.T) {
		var ptr *int
		if !value.IsNil(ptr) {
			t.Error("IsNil should return true for nil pointer")
		}
		if !value.IsNil(nil) {
			t.Error("IsNil should return true for nil")
		}
		if !value.IsNil((*int)(nil)) {
			t.Error("IsNil should return true for nil pointer cast")
		}
	})

	t.Run("non-nil values", func(t *testing.T) {
		val := 42
		ptr := &val
		if value.IsNil(ptr) {
			t.Error("IsNil should return false for non-nil pointer")
		}
		if value.IsNil([]int{}) {
			t.Error("IsNil should return false for empty slice")
		}
		if value.IsNil(map[string]int{}) {
			t.Error("IsNil should return false for empty map")
		}
	})
}

// TestIsNotNil tests the IsNotNil function
func TestIsNotNil(t *testing.T) {
	t.Run("nil values", func(t *testing.T) {
		var ptr *int
		if value.IsNotNil(ptr) {
			t.Error("IsNotNil should return false for nil pointer")
		}
		if value.IsNotNil(nil) {
			t.Error("IsNotNil should return false for nil")
		}
	})

	t.Run("non-nil values", func(t *testing.T) {
		val := 42
		ptr := &val
		if !value.IsNotNil(ptr) {
			t.Error("IsNotNil should return true for non-nil pointer")
		}
		if !value.IsNotNil([]int{}) {
			t.Error("IsNotNil should return true for empty slice")
		}
	})
}

// TestIsEmpty tests the IsEmpty function
func TestIsEmpty(t *testing.T) {
	t.Run("empty values", func(t *testing.T) {
		if !value.IsEmpty("") {
			t.Error("IsEmpty should return true for empty string")
		}
		if !value.IsEmpty([]int{}) {
			t.Error("IsEmpty should return true for empty slice")
		}
		if !value.IsEmpty(map[string]int{}) {
			t.Error("IsEmpty should return true for empty map")
		}
		if !value.IsEmpty(0) {
			t.Error("IsEmpty should return true for 0")
		}
		if !value.IsEmpty(nil) {
			t.Error("IsEmpty should return true for nil")
		}
	})

	t.Run("non-empty values", func(t *testing.T) {
		if value.IsEmpty("hello") {
			t.Error("IsEmpty should return false for 'hello'")
		}
		if value.IsEmpty([]int{1, 2, 3}) {
			t.Error("IsEmpty should return false for non-empty slice")
		}
		if value.IsEmpty(map[string]int{"a": 1}) {
			t.Error("IsEmpty should return false for non-empty map")
		}
		if value.IsEmpty(42) {
			t.Error("IsEmpty should return false for 42")
		}
	})
}

// TestIsNotEmpty tests the IsNotEmpty function
func TestIsNotEmpty(t *testing.T) {
	t.Run("empty values", func(t *testing.T) {
		if value.IsNotEmpty("") {
			t.Error("IsNotEmpty should return false for empty string")
		}
		if value.IsNotEmpty([]int{}) {
			t.Error("IsNotEmpty should return false for empty slice")
		}
		if value.IsNotEmpty(0) {
			t.Error("IsNotEmpty should return false for 0")
		}
	})

	t.Run("non-empty values", func(t *testing.T) {
		if !value.IsNotEmpty("hello") {
			t.Error("IsNotEmpty should return true for 'hello'")
		}
		if !value.IsNotEmpty([]int{1, 2, 3}) {
			t.Error("IsNotEmpty should return true for non-empty slice")
		}
		if !value.IsNotEmpty(42) {
			t.Error("IsNotEmpty should return true for 42")
		}
	})
}

// TestIf tests the If function
func TestIf(t *testing.T) {
	t.Run("true condition", func(t *testing.T) {
		result := value.If(true, "yes")
		if result != "yes" {
			t.Errorf("If should return 'yes' for true condition, got %s", result)
		}
	})

	t.Run("false condition", func(t *testing.T) {
		result := value.If(false, "yes")
		if result != "" {
			t.Errorf("If should return empty string for false condition, got %s", result)
		}
	})

	t.Run("with integers", func(t *testing.T) {
		result := value.If(5 > 3, 100)
		if result != 100 {
			t.Errorf("If should return 100 for true condition, got %d", result)
		}
	})
}

// TestWhen tests the When function
func TestWhen(t *testing.T) {
	t.Run("predicate returns true", func(t *testing.T) {
		result := value.When(42, func(x int) bool { return x > 0 })
		if result != 42 {
			t.Errorf("When should return 42 for positive number, got %d", result)
		}
	})

	t.Run("predicate returns false", func(t *testing.T) {
		result := value.When(42, func(x int) bool { return x < 0 })
		if result != 0 {
			t.Errorf("When should return 0 for negative predicate, got %d", result)
		}
	})

	t.Run("with strings", func(t *testing.T) {
		result := value.When("hello", func(s string) bool { return len(s) > 3 })
		if result != "hello" {
			t.Errorf("When should return 'hello' for long string, got %s", result)
		}
	})
}

// TestWhenElse tests the WhenElse function
func TestWhenElse(t *testing.T) {
	t.Run("predicate returns true", func(t *testing.T) {
		result := value.WhenElse(42, func(x int) bool { return x > 0 }, 1, -1)
		if result != 1 {
			t.Errorf("WhenElse should return 1 for positive number, got %d", result)
		}
	})

	t.Run("predicate returns false", func(t *testing.T) {
		result := value.WhenElse(-5, func(x int) bool { return x > 0 }, 1, -1)
		if result != -1 {
			t.Errorf("WhenElse should return -1 for negative number, got %d", result)
		}
	})
}

// TestCoalesce tests the Coalesce function
func TestCoalesce(t *testing.T) {
	t.Run("first non-nil pointer", func(t *testing.T) {
		var p1, p2 *int
		val := 42
		p3 := &val
		result := value.Coalesce(p1, p2, p3)
		if result != p3 {
			t.Errorf("Coalesce should return p3, got %v", result)
		}
	})

	t.Run("all nil pointers", func(t *testing.T) {
		var p1, p2, p3 *int
		result := value.Coalesce(p1, p2, p3)
		if result != nil {
			t.Errorf("Coalesce should return nil for all nil pointers, got %v", result)
		}
	})

	t.Run("single pointer", func(t *testing.T) {
		val := 42
		p := &val
		result := value.Coalesce(p)
		if result != p {
			t.Errorf("Coalesce should return p, got %v", result)
		}
	})
}

// TestCoalesceValue tests the CoalesceValue function
func TestCoalesceValue(t *testing.T) {
	t.Run("first non-nil pointer", func(t *testing.T) {
		var p1, p2 *int
		val := 42
		p3 := &val
		result := value.CoalesceValue(p1, p2, p3)
		if result != 42 {
			t.Errorf("CoalesceValue should return 42, got %d", result)
		}
	})

	t.Run("all nil pointers", func(t *testing.T) {
		var p1, p2, p3 *int
		result := value.CoalesceValue(p1, p2, p3)
		if result != 0 {
			t.Errorf("CoalesceValue should return 0 for all nil pointers, got %d", result)
		}
	})
}

// TestCoalesceValueDef tests the CoalesceValueDef function
func TestCoalesceValueDef(t *testing.T) {
	t.Run("first non-nil pointer", func(t *testing.T) {
		var p1, p2 *int
		val := 42
		p3 := &val
		result := value.CoalesceValueDef(999, p1, p2, p3)
		if result != 42 {
			t.Errorf("CoalesceValueDef should return 42, got %d", result)
		}
	})

	t.Run("all nil pointers", func(t *testing.T) {
		var p1, p2, p3 *int
		result := value.CoalesceValueDef(999, p1, p2, p3)
		if result != 999 {
			t.Errorf("CoalesceValueDef should return 999 for all nil pointers, got %d", result)
		}
	})
}

// TestSafeDeref tests the SafeDeref function
func TestSafeDeref(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ptr *int
		result := value.SafeDeref(ptr)
		if result != 0 {
			t.Errorf("SafeDeref should return 0 for nil pointer, got %d", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		val := 42
		ptr := &val
		result := value.SafeDeref(ptr)
		if result != 42 {
			t.Errorf("SafeDeref should return 42 for non-nil pointer, got %d", result)
		}
	})
}

// TestSafeDerefDef tests the SafeDerefDef function
func TestSafeDerefDef(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ptr *int
		result := value.SafeDerefDef(ptr, 100)
		if result != 100 {
			t.Errorf("SafeDerefDef should return 100 for nil pointer, got %d", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		val := 42
		ptr := &val
		result := value.SafeDerefDef(ptr, 100)
		if result != 42 {
			t.Errorf("SafeDerefDef should return 42 for non-nil pointer, got %d", result)
		}
	})
}

// TestValue tests the Value function
func TestValue(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ptr *int
		result := value.Value(ptr)
		if result != 0 {
			t.Errorf("Value should return 0 for nil pointer, got %d", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		val := 42
		ptr := &val
		result := value.Value(ptr)
		if result != 42 {
			t.Errorf("Value should return 42 for non-nil pointer, got %d", result)
		}
	})
}

// TestDef tests the Def function
func TestDef(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		var ptr *int
		result := value.Def(ptr, 100)
		if result != 100 {
			t.Errorf("Def should return 100 for nil pointer, got %d", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		val := 42
		ptr := &val
		result := value.Def(ptr, 100)
		if result != 42 {
			t.Errorf("Def should return 42 for non-nil pointer, got %d", result)
		}
	})
}

// TestEqual tests the Equal function
func TestEqual(t *testing.T) {
	t.Run("equal values", func(t *testing.T) {
		if !value.Equal(42, 42) {
			t.Error("Equal should return true for equal values")
		}
		if !value.Equal("hello", "hello") {
			t.Error("Equal should return true for equal strings")
		}
	})

	t.Run("not equal values", func(t *testing.T) {
		if value.Equal(42, 43) {
			t.Error("Equal should return false for different values")
		}
		if value.Equal("hello", "world") {
			t.Error("Equal should return false for different strings")
		}
	})
}

// TestNotEqual tests the NotEqual function
func TestNotEqual(t *testing.T) {
	t.Run("equal values", func(t *testing.T) {
		if value.NotEqual(42, 42) {
			t.Error("NotEqual should return false for equal values")
		}
		if value.NotEqual("hello", "hello") {
			t.Error("NotEqual should return false for equal strings")
		}
	})

	t.Run("not equal values", func(t *testing.T) {
		if !value.NotEqual(42, 43) {
			t.Error("NotEqual should return true for different values")
		}
		if !value.NotEqual("hello", "world") {
			t.Error("NotEqual should return true for different strings")
		}
	})
}

// TestDeepEqual tests the DeepEqual function
func TestDeepEqual(t *testing.T) {
	t.Run("equal slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{1, 2, 3}
		if !value.DeepEqual(slice1, slice2) {
			t.Error("DeepEqual should return true for equal slices")
		}
	})

	t.Run("equal maps", func(t *testing.T) {
		map1 := map[string]int{"a": 1, "b": 2}
		map2 := map[string]int{"a": 1, "b": 2}
		if !value.DeepEqual(map1, map2) {
			t.Error("DeepEqual should return true for equal maps")
		}
	})

	t.Run("not equal slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{1, 2, 4}
		if value.DeepEqual(slice1, slice2) {
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
		if !value.DeepEqual(p1, p2) {
			t.Error("DeepEqual should return true for equal structs")
		}
	})
}

// TestEdgeCases tests edge cases and error conditions
func TestEdgeCases(t *testing.T) {
	t.Run("empty variadic arguments", func(t *testing.T) {
		result := value.Or[int]()
		if result != 0 {
			t.Errorf("Or with no arguments should return 0, got %d", result)
		}

		result2 := value.OrElse(42)
		if result2 != 42 {
			t.Errorf("OrElse with no arguments should return default, got %d", result2)
		}
	})

	t.Run("nil interface values", func(t *testing.T) {
		var i interface{}
		if !value.IsNil(i) {
			t.Error("IsNil should return true for nil interface")
		}
		if !value.IsEmpty(i) {
			t.Error("IsEmpty should return true for nil interface")
		}
	})

	t.Run("zero value structs", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		var p Point
		if !value.IsZero(p) {
			t.Error("IsZero should return true for zero value struct")
		}
		if !value.IsEmpty(p) {
			t.Error("IsEmpty should return true for zero value struct")
		}
	})
}

// TestComplexTypes tests with complex types
func TestComplexTypes(t *testing.T) {
	t.Run("slice pointers", func(t *testing.T) {
		var s1, s2 *[]int
		s3 := &[]int{1, 2, 3}
		result := value.Coalesce(s1, s2, s3)
		if result != s3 {
			t.Errorf("Coalesce should return s3, got %v", result)
		}
	})

	t.Run("map pointers", func(t *testing.T) {
		var m1, m2 *map[string]int
		m3 := &map[string]int{"a": 1}
		result := value.Coalesce(m1, m2, m3)
		if result != m3 {
			t.Errorf("Coalesce should return m3, got %v", result)
		}
	})

	t.Run("function pointers", func(t *testing.T) {
		var f1, f2 *func() int
		f := func() int { return 42 }
		f3 := &f
		result := value.Coalesce(f1, f2, f3)
		if result != f3 {
			t.Errorf("Coalesce should return f3, got %v", result)
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

		if value.IsEmpty(largeSlice) {
			t.Error("IsEmpty should return false for large slice")
		}
		if !value.IsNotEmpty(largeSlice) {
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

		if !value.DeepEqual(n1, n2) {
			t.Error("DeepEqual should return true for deeply nested equal structs")
		}
	})
}
