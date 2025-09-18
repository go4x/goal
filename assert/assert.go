// Package assert provides assertion functions that panic on failure.
//
// This package is designed for testing and debugging purposes only. All functions
// in this package will panic when assertions fail, making them unsuitable for
// production code.
//
// WARNING: This package is designed for testing and debugging purposes only.
// All functions in this package will panic when assertions fail.
//
// IMPORTANT USAGE GUIDELINES:
//   - DO NOT use in production code
//   - DO NOT use in goroutines (panic will terminate the entire program)
//   - DO NOT use for input validation in HTTP handlers or similar contexts
//   - ONLY use in test functions or during development/debugging
//
// For production code, use proper error handling instead:
//
//	if err != nil {
//	    return fmt.Errorf("operation failed: %w", err)
//	}
//
// For testing, consider using testing.T methods:
//
//	if got != want {
//	    t.Fatalf("got %v, want %v", got, want)
//	}
//
// Example usage in tests:
//
//	func TestSomething(t *testing.T) {
//	    assert.True(condition)
//	    assert.Nil(err)
//	    assert.NotBlank(name)
//	}
//
// The package provides the following assertion functions:
//   - Require: Basic assertion with custom message
//   - True: Assert boolean value is true
//   - Nil: Assert value is nil
//   - NoneNil: Assert value is not nil
//   - Blank: Assert string is blank (empty or whitespace only)
//   - NotBlank: Assert string is not blank
//   - HasElems: Assert collection has elements
//   - Equals: Assert two values are equal
//   - DeepEquals: Assert two values are deeply equal
//
// All functions follow the same pattern: they panic with a descriptive message
// when the assertion fails, and do nothing when the assertion passes.
package assert

import (
	"fmt"
	"reflect"
	"strings"
)

// basic methods

// Require is the basic method to assert a condition.
// If the condition is false, it will panic with the given format and arguments.
//
// This is the fundamental assertion function that all other assertion functions
// use internally. It provides the most flexibility for custom assertion messages.
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - cond: The condition to assert (must be true for assertion to pass)
//   - format: A format string for the panic message (supports fmt.Printf formatting)
//   - v: Arguments to format the message (optional)
//
// Panics:
//   - When cond is false, panics with the formatted message
//
// Example:
//
//	assert.Require(x > 0, "x must be positive, got %d", x)
//	assert.Require(err == nil, "unexpected error: %v", err)
//	assert.Require(len(slice) > 0, "slice cannot be empty")
func Require(cond bool, format string, v ...any) {
	if !cond {
		panic(fmt.Sprintf(format, v...))
	}
}

// convenient methods

// True asserts that the given boolean value is true.
//
// This is a convenience function for asserting boolean conditions. It's equivalent
// to calling Require(b, "assert failed, expects [true] but found [false]").
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - b: The boolean value to assert as true
//
// Panics:
//   - When b is false, panics with a descriptive message
//
// Example:
//
//	assert.True(x > 0)
//	assert.True(err == nil)
//	assert.True(len(items) > 0)
func True(b bool) {
	Require(b, "assert failed, expects [%t] but found [%t]", true, b)
}

// Nil asserts that the given value is nil.
//
// This function is commonly used to assert that error values are nil, or that
// pointer values are nil. It's equivalent to calling Require(t == nil, "...").
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - t: The value to assert as nil
//
// Panics:
//   - When t is not nil, panics with a descriptive message
//
// Example:
//
//	assert.Nil(err)
//	assert.Nil(ptr)
//	assert.Nil(interface{}(nil))
func Nil(t any) {
	Require(t == nil, "assert failed, expects [nil] but found [not nil]: %v", t)
}

// NoneNil asserts that the given value is not nil.
//
// This function is commonly used to assert that objects have been properly
// initialized and are not nil. It's equivalent to calling Require(t != nil, "...").
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - t: The value to assert as non-nil
//
// Panics:
//   - When t is nil, panics with a descriptive message
//
// Example:
//
//	assert.NoneNil(user)
//	assert.NoneNil(config)
//	assert.NoneNil(service)
func NoneNil(t any) {
	Require(t != nil, "assert failed, expects [not nil] but found [nil]")
}

// Blank asserts that the given string is blank (empty or only contains whitespace).
//
// This function trims whitespace from the string before checking if it's empty.
// A string is considered blank if it's empty or contains only whitespace characters
// (spaces, tabs, newlines, etc.).
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - s: The string to assert as blank
//
// Panics:
//   - When s is not blank (contains non-whitespace characters), panics with a descriptive message
//
// Example:
//
//	assert.Blank("")
//	assert.Blank("   ")
//	assert.Blank("\n\t")
func Blank(s string) {
	s = strings.TrimSpace(s)
	Require(s == "", "assert failed, expects [\"\"] but found [%s]", s)
}

// NotBlank asserts that the given string is not blank (not empty or not only contains whitespace).
//
// This function trims whitespace from the string before checking if it's empty.
// A string is considered not blank if it contains at least one non-whitespace character.
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - s: The string to assert as non-blank
//
// Panics:
//   - When s is blank (empty or only whitespace), panics with a descriptive message
//
// Example:
//
//	assert.NotBlank("hello")
//	assert.NotBlank("  hello  ")
//	assert.NotBlank("\thello\n")
func NotBlank(s string) {
	s = strings.TrimSpace(s)
	Require(s != "", "assert failed, expects [not empty] but found [\"\"]")
}

// HasElems asserts that the given collection is not nil and has elements.
//
// The parameter c should be a collection type such as array, map, slice, or channel.
// If c is not a collection type, this function does nothing (no panic).
// This function uses reflection to determine the type and length of the collection.
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - c: The collection to assert as non-nil and non-empty
//
// Panics:
//   - When c is nil or empty, panics with a descriptive message
//
// Example:
//
//	assert.HasElems([]int{1, 2, 3})
//	assert.HasElems(map[string]int{"key": 1})
//	assert.HasElems([3]int{1, 2, 3})
//	assert.HasElems(make(chan int, 1))
func HasElems(c any) {
	typ := reflect.TypeOf(c).Kind()
	if typ == reflect.Array || typ == reflect.Map || typ == reflect.Slice || typ == reflect.Chan {
		n := reflect.ValueOf(c).Len()
		Require(c != nil && n > 0, "assert failed, expects collection is none nil and has elements")
	}
}

// Equals asserts that the given two values are equal.
//
// This function performs a simple equality comparison using the == operator.
// For complex types (slices, maps, structs), use DeepEquals instead.
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - t1: The first value to compare
//   - t2: The second value to compare
//
// Panics:
//   - When t1 and t2 are not equal, panics with a descriptive message
//
// Example:
//
//	assert.Equals(42, 42)
//	assert.Equals("hello", "hello")
//	assert.Equals(nil, nil)
func Equals(t1 any, t2 any) {
	Require(t1 == t2, "assert failed, expecting t1 equals t2 but not")
}

// DeepEquals asserts that the given two values are deeply equal.
//
// It uses reflect.DeepEqual to compare the two values, which is suitable for
// complex types like slices, maps, structs, and nested data structures.
// This is more comprehensive than Equals for complex types.
//
// WARNING: This function will panic on failure. Do not use in production code or goroutines.
//
// Parameters:
//   - t1: The first value to compare
//   - t2: The second value to compare
//
// Panics:
//   - When t1 and t2 are not deeply equal, panics with a descriptive message
//
// Example:
//
//	assert.DeepEquals([]int{1, 2, 3}, []int{1, 2, 3})
//	assert.DeepEquals(map[string]int{"a": 1}, map[string]int{"a": 1})
//	assert.DeepEquals(struct{Name string}{"Alice"}, struct{Name string}{"Alice"})
func DeepEquals(t1 any, t2 any) {
	Require(reflect.DeepEqual(t1, t2), "assert failed, expects %v deep equals %v, but not", t1, t2)
}
