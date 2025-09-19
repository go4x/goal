package valuex

import (
	"errors"
	"testing"

	"github.com/gophero/got"
)

func TestMust(t *testing.T) {
	logger := got.New(t, "test valuex.Must method")

	logger.Case("give valid value and nil error, should return value")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		result := Must(42, nil)
		logger.Require(result == 42, "result should be 42, got %v", result)
	}()

	logger.Case("give valid value and nil error with string, should return value")
	func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Fail("should not panic, but got: %v", err)
			} else {
				logger.Pass("should not panic")
			}
		}()
		result := Must("hello", nil)
		logger.Require(result == "hello", "result should be 'hello', got %v", result)
	}()

	logger.Case("give error, should panic")
	func() {
		defer func() {
			if err := recover(); err == nil {
				logger.Fail("should panic, but not")
			} else {
				logger.Pass("should panic: %v", err)
			}
		}()
		Must(42, errors.New("test error"))
	}()
}

func TestDef(t *testing.T) {
	logger := got.New(t, "test valuex.Def method")

	logger.Case("give true condition, should return first value")
	result1 := Def(true, "first", "second")
	logger.Require(result1 == "first", "result should be 'first', got %v", result1)

	logger.Case("give false condition, should return second value")
	result2 := Def(false, "first", "second")
	logger.Require(result2 == "second", "result should be 'second', got %v", result2)

	logger.Case("test with integers")
	result3 := Def(true, 10, 20)
	logger.Require(result3 == 10, "result should be 10, got %v", result3)

	result4 := Def(false, 10, 20)
	logger.Require(result4 == 20, "result should be 20, got %v", result4)

	logger.Case("test with floats")
	result5 := Def(true, 1.5, 2.5)
	logger.Require(result5 == 1.5, "result should be 1.5, got %v", result5)

	result6 := Def(false, 1.5, 2.5)
	logger.Require(result6 == 2.5, "result should be 2.5, got %v", result6)
}

func TestMin(t *testing.T) {
	logger := got.New(t, "test valuex.Min method")

	logger.Case("test with integers")
	result1 := Min(5, 10)
	logger.Require(result1 == 5, "result should be 5, got %v", result1)

	result2 := Min(10, 5)
	logger.Require(result2 == 5, "result should be 5, got %v", result2)

	result3 := Min(5, 5)
	logger.Require(result3 == 5, "result should be 5, got %v", result3)

	logger.Case("test with floats")
	result4 := Min(1.5, 2.5)
	logger.Require(result4 == 1.5, "result should be 1.5, got %v", result4)

	result5 := Min(2.5, 1.5)
	logger.Require(result5 == 1.5, "result should be 1.5, got %v", result5)

	logger.Case("test with strings")
	result6 := Min("apple", "banana")
	logger.Require(result6 == "apple", "result should be 'apple', got %v", result6)

	result7 := Min("banana", "apple")
	logger.Require(result7 == "apple", "result should be 'apple', got %v", result7)
}

func TestMax(t *testing.T) {
	logger := got.New(t, "test valuex.Max method")

	logger.Case("test with integers")
	result1 := Max(5, 10)
	logger.Require(result1 == 10, "result should be 10, got %v", result1)

	result2 := Max(10, 5)
	logger.Require(result2 == 10, "result should be 10, got %v", result2)

	result3 := Max(5, 5)
	logger.Require(result3 == 5, "result should be 5, got %v", result3)

	logger.Case("test with floats")
	result4 := Max(1.5, 2.5)
	logger.Require(result4 == 2.5, "result should be 2.5, got %v", result4)

	result5 := Max(2.5, 1.5)
	logger.Require(result5 == 2.5, "result should be 2.5, got %v", result5)

	logger.Case("test with strings")
	result6 := Max("apple", "banana")
	logger.Require(result6 == "banana", "result should be 'banana', got %v", result6)

	result7 := Max("banana", "apple")
	logger.Require(result7 == "banana", "result should be 'banana', got %v", result7)
}

func TestEdgeCases(t *testing.T) {
	logger := got.New(t, "test edge cases")

	logger.Case("test with negative numbers")
	result1 := Min(-5, -10)
	logger.Require(result1 == -10, "result should be -10, got %v", result1)

	result2 := Max(-5, -10)
	logger.Require(result2 == -5, "result should be -5, got %v", result2)

	logger.Case("test with zero")
	result3 := Min(0, 5)
	logger.Require(result3 == 0, "result should be 0, got %v", result3)

	result4 := Max(0, 5)
	logger.Require(result4 == 5, "result should be 5, got %v", result4)

	logger.Case("test with empty strings")
	result5 := Min("", "hello")
	logger.Require(result5 == "", "result should be empty string, got %v", result5)

	result6 := Max("", "hello")
	logger.Require(result6 == "hello", "result should be 'hello', got %v", result6)
}
