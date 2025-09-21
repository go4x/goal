package assert_test

import (
	"testing"

	"github.com/go4x/goal/assert"
	"github.com/go4x/got"
)

func TestRequire(t *testing.T) {
	logger := got.New(t, "test assert.Require method")

	logger.Case("give a true condition, should not panic")
	func() {
		defer func() {
			err := recover()
			if err != nil {
				logger.Fail("should not get an error: %v", err)
			} else {
				logger.Pass("should not get an error")
			}
		}()
		assert.Require(true, "not match the condition")
	}()

	logger.Case("give a false condition, should panic and get an error")
	func() {
		msg := "not match the condition"
		defer func() {
			err := recover()
			if err == nil {
				logger.Fail("should get an error with message: " + msg)
			} else {
				logger.Pass("should get an error: %v", err)
			}
		}()
		assert.Require(false, "not match the condition")
	}()

	logger.Case("test Require with format string and arguments")
	func() {
		defer func() {
			err := recover()
			if err == nil {
				logger.Fail("should get an error with formatted message")
			} else {
				logger.Pass("should get an error: %v", err)
			}
		}()
		assert.Require(false, "error: %s, code: %d", "test error", 500)
	}()
}
func TestTrue(t *testing.T) {
	logger := got.New(t, "test assert.True method")

	logger.Case("give true, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.True(true)
	}()

	logger.Case("give false, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.True(false)
	}()
}

func TestNil(t *testing.T) {
	logger := got.New(t, "test assert.Nil method")

	logger.Case("give nil, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.Nil(nil)
	}()

	logger.Case("give not nil, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.Nil(1)
	}()
}

func TestNoneNil(t *testing.T) {
	logger := got.New(t, "test assert.NoneNil method")

	logger.Case("give not nil, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.NoneNil(1)
	}()

	logger.Case("give nil, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.NoneNil(nil)
	}()
}

func TestBlank(t *testing.T) {
	logger := got.New(t, "test assert.Blank method")

	logger.Case("give blank string, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.Blank("")
		assert.Blank("   ")
		assert.Blank("\n\t")
	}()

	logger.Case("give not blank string, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.Blank("abc")
	}()
}
func TestNotBlank(t *testing.T) {
	logger := got.New(t, "test assert.NotBlank method")

	logger.Case("give not blank string, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.NotBlank("abc")
		assert.NotBlank("  abc  ")
		assert.NotBlank("\nabc\t")
	}()

	logger.Case("give blank string, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.NotBlank("")
	}()
}

func TestHasElems(t *testing.T) {
	logger := got.New(t, "test assert.HasElems method")

	logger.Case("give non-empty slice, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.HasElems([]int{1, 2, 3})
		assert.HasElems([2]int{1, 2})
		assert.HasElems(map[string]int{"a": 1})
		ch := make(chan int, 1)
		ch <- 1
		assert.HasElems(ch)
	}()

	logger.Case("give empty slice, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.HasElems([]int{})
	}()

	logger.Case("give nil slice, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		var s []int
		assert.HasElems(s)
	}()

	logger.Case("give non-collection type, should not panic (do nothing)")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic for non-collection type, but got: %v", err)
			} else {
				logger.Pass("should not panic for non-collection type")
			}
		}()
		assert.HasElems("string")
		assert.HasElems(123)
		assert.HasElems(true)
		assert.HasElems(struct{}{})
	}()
}

func TestEquals(t *testing.T) {
	logger := got.New(t, "test assert.Equals method")

	logger.Case("give equal values, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.Equals(1, 1)
		assert.Equals("abc", "abc")
	}()

	logger.Case("give not equal values, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.Equals(1, 2)
	}()
}

func TestDeepEquals(t *testing.T) {
	logger := got.New(t, "test assert.DeepEquals method")

	logger.Case("give deeply equal values, should not panic")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		assert.DeepEquals([]int{1, 2}, []int{1, 2})
		assert.DeepEquals(map[string]int{"a": 1}, map[string]int{"a": 1})
	}()

	logger.Case("give not deeply equal values, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		assert.DeepEquals([]int{1, 2}, []int{2, 1})
	}()
}

func TestEdgeCases(t *testing.T) {
	logger := got.New(t, "test assert edge cases")

	logger.Case("test Blank with various whitespace combinations")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic for whitespace strings, but got: %v", err)
			} else {
				logger.Pass("should not panic for whitespace strings")
			}
		}()
		assert.Blank(" ")
		assert.Blank("\t")
		assert.Blank("\n")
		assert.Blank("\r")
		assert.Blank(" \t\n\r ")
	}()

	logger.Case("test NotBlank with various non-whitespace strings")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic for non-whitespace strings, but got: %v", err)
			} else {
				logger.Pass("should not panic for non-whitespace strings")
			}
		}()
		assert.NotBlank("a")
		assert.NotBlank(" a ")
		assert.NotBlank("\ta\n")
	}()

	logger.Case("test HasElems with different collection types")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic for non-empty collections, but got: %v", err)
			} else {
				logger.Pass("should not panic for non-empty collections")
			}
		}()
		// Test array
		assert.HasElems([3]int{1, 2, 3})
		// Test map
		assert.HasElems(map[string]int{"key": 1})
		// Test channel
		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		assert.HasElems(ch)
		close(ch)
	}()

	logger.Case("test Equals with different types")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic for equal values, but got: %v", err)
			} else {
				logger.Pass("should not panic for equal values")
			}
		}()
		assert.Equals(0, 0)
		assert.Equals("", "")
		assert.Equals(nil, nil)
		assert.Equals(true, true)
		assert.Equals(false, false)
	}()

	logger.Case("test DeepEquals with complex structures")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic for deeply equal structures, but got: %v", err)
			} else {
				logger.Pass("should not panic for deeply equal structures")
			}
		}()
		// Test nested slices
		assert.DeepEquals([][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}})
		// Test nested maps
		assert.DeepEquals(map[string]map[string]int{"a": {"b": 1}}, map[string]map[string]int{"a": {"b": 1}})
		// Test structs
		type Person struct {
			Name string
			Age  int
		}
		p1 := Person{Name: "Alice", Age: 30}
		p2 := Person{Name: "Alice", Age: 30}
		assert.DeepEquals(p1, p2)
	}()
}
