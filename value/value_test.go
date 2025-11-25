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
