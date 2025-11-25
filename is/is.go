package is

import "reflect"

// Not returns the logical negation of the value. Is is an alias for False.
//
// Parameters:
//   - v: The boolean value to negate
//
// Returns:
//   - bool: true if the value is false, false if the value is true
//
// Example:
//
//	is.Not(true)     // false
//	is.Not(false)    // true
func Not(v bool) bool {
	return False(v)
}

// True returns true if the value is true.
//
// Parameters:
//   - v: The boolean value to check
//
// Returns:
//   - bool: true if the value is true
//
// Example:
//
//	is.True(true)     // true
//	is.True(false)    // false
func True(v bool) bool {
	return v
}

// False returns true if the value is false.
//
// Parameters:
//   - v: The boolean value to check
//
// Returns:
//   - bool: true if the value is false
//
// Example:
//
//	is.False(true)     // false
//	is.False(false)    // true
func False(v bool) bool {
	return !v
}

// Zero checks if the value is the zero value for its type.
// Zero values are: 0 for numeric types, "" for strings, false for bools, nil for pointers/slices/maps/channels/functions/interfaces.
//
// Parameters:
//   - v: The value to check
//
// Returns:
//   - bool: true if the value is the zero value for its type
//
// Example:
//
//	value.Zero(0)        // true
//	value.Zero("")        // true
//	value.Zero(false)     // true
//	value.Zero(42)        // false
//	value.Zero("hello")   // false
func Zero[T comparable](v T) bool {
	var zero T
	return v == zero
}

// NotZero checks if the value is not the zero value for its type.
// This is the logical negation of IsZero.
//
// Parameters:
//   - v: The value to check
//
// Returns:
//   - bool: true if the value is not the zero value for its type
//
// Example:
//
//	value.NotZero(42)     // true
//	value.NotZero("hello") // true
//	value.NotZero(0)      // false
//	value.NotZero("")     // false
func NotZero[T comparable](v T) bool {
	return !Zero(v)
}

// Nil checks if the value is nil (for pointer, slice, map, channel, function, interface).
// This function uses reflection to properly check nil values for reference types.
//
// Parameters:
//   - v: The value to check (can be any type)
//
// Returns:
//   - bool: true if the value is nil
//
// Example:
//
//	var ptr *int
//	value.Nil(ptr)        // true
//	value.Nil(nil)        // true
//	value.Nil([]int{})    // false (empty slice is not nil)
//	value.Nil((*int)(nil)) // true
func Nil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	}
	return false
}

// NotNil checks if the value is not nil.
// This is the logical negation of IsNil.
//
// Parameters:
//   - v: The value to check (can be any type)
//
// Returns:
//   - bool: true if the value is not nil
//
// Example:
//
//	ptr := &42
//	value.NotNil(ptr)     // true
//	value.NotNil([]int{}) // true (empty slice is not nil)
//	value.NotNil(nil)     // false
func NotNil(v interface{}) bool {
	return !Nil(v)
}

// Empty checks if the value is empty (zero value or empty string/slice/map).
// This function provides a comprehensive check for "empty" values across different types.
//
// Parameters:
//   - v: The value to check (can be any type)
//
// Returns:
//   - bool: true if the value is considered empty
//
// Empty conditions:
//   - nil values
//   - empty strings ("")
//   - empty slices, maps, arrays (length 0)
//   - nil pointers and interfaces
//   - zero values for other types
//
// Example:
//
//	value.Empty("")           // true
//	value.Empty([]int{})      // true
//	value.Empty(map[string]int{}) // true
//	value.Empty(0)            // true
//	value.Empty("hello")      // false
//	value.Empty([]int{1, 2})  // false
func Empty(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Slice, reflect.Map, reflect.Array:
		return rv.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return rv.IsNil()
	}
	// For other types, check if it's the zero value
	return rv.IsZero()
}

// NotEmpty checks if the value is not empty.
// This is the logical negation of IsEmpty.
//
// Parameters:
//   - v: The value to check (can be any type)
//
// Returns:
//   - bool: true if the value is not empty
//
// Example:
//
//	value.NotEmpty("hello")     // true
//	value.NotEmpty([]int{1, 2}) // true
//	value.NotEmpty(42)          // true
//	value.NotEmpty("")          // false
//	value.NotEmpty([]int{})     // false
func NotEmpty(v interface{}) bool {
	return !Empty(v)
}

// Equal checks if two values are equal.
// This is a simple equality check using the == operator.
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if the values are equal, false otherwise
//
// Example:
//
//	value.Equal(42, 42)        // true
//	value.Equal("hello", "hi") // false
//	value.Equal(0, 0)          // true
func Equal[T comparable](a, b T) bool {
	return a == b
}

// NotEqual checks if two values are not equal.
// This is the logical negation of Equal.
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if the values are not equal, false otherwise
//
// Example:
//
//	value.NotEqual(42, 43)     // true
//	value.NotEqual("hello", "hello") // false
//	value.NotEqual(0, 1)       // true
func NotEqual[T comparable](a, b T) bool {
	return a != b
}

// DeepEqual checks if two values are deeply equal using reflection.
// This performs a deep comparison that can handle complex nested structures.
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if the values are deeply equal, false otherwise
//
// Example:
//
//	slice1 := []int{1, 2, 3}
//	slice2 := []int{1, 2, 3}
//	value.DeepEqual(slice1, slice2) // true
//
//	map1 := map[string]int{"a": 1}
//	map2 := map[string]int{"a": 1}
//	value.DeepEqual(map1, map2) // true
func DeepEqual[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}
