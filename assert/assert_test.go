package assert_test

import (
	"testing"

	"github.com/gophero/goal/assert"
	"github.com/gophero/got"
)

func TestRequire(t *testing.T) {
	logger := got.Wrap(t)
	logger.Title("test assert.Require method")

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
}
func TestTrue(t *testing.T) {
	logger := got.Wrap(t)
	logger.Title("test assert.True method")

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
	logger := got.Wrap(t)
	logger.Title("test assert.Nil method")

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
	logger := got.Wrap(t)
	logger.Title("test assert.NoneNil method")

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
	logger := got.Wrap(t)
	logger.Title("test assert.Blank method")

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
	logger := got.Wrap(t)
	logger.Title("test assert.NotBlank method")

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
	logger := got.Wrap(t)
	logger.Title("test assert.HasElems method")

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
}

func TestEquals(t *testing.T) {
	logger := got.Wrap(t)
	logger.Title("test assert.Equals method")

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
	logger := got.Wrap(t)
	logger.Title("test assert.DeepEquals method")

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
