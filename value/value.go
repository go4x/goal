package value

import (
	"reflect"
)

// Must returns the value if error is nil, otherwise panics.
// This is useful for cases where you're certain the operation will succeed.
// Use with caution as it will panic if an error occurs.
//
// Parameters:
//   - t: The value to return if no error
//   - err: The error to check
//
// Returns:
//   - T: The value if err is nil
//
// Panics:
//   - If err is not nil
//
// Example:
//
//	// Safe to use when you know the operation will succeed
//	result := value.Must(strconv.Atoi("123"))
//	// result is 123
//
//	// This will panic if the string is not a valid integer
//	// result := value.Must(strconv.Atoi("invalid"))
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// IfElse returns v1 if condition is true, otherwise returns v2.
// This is a generic conditional operator that provides a functional approach to conditional logic.
//
// Parameters:
//   - b: The boolean condition to evaluate
//   - v1: The value to return if condition is true
//   - v2: The value to return if condition is false
//
// Returns:
//   - T: v1 if condition is true, v2 otherwise
//
// Example:
//
//	result := value.IfElse(age >= 18, "adult", "minor")
//	// result is "adult" if age >= 18, "minor" otherwise
func IfElse[T any](b bool, v1 T, v2 T) T {
	if b {
		return v1
	}
	return v2
}

// Or returns the first non-zero value from the given values.
// If all values are zero, returns the zero value of type T.
// This is useful for providing fallback values in a chain.
//
// Parameters:
//   - values: Variable number of values to check
//
// Returns:
//   - T: The first non-zero value, or zero value if all are zero
//
// Example:
//
//	result := value.Or("", "", "fallback", "ignored")
//	// result is "fallback"
//
//	empty := value.Or("", "", "")
//	// empty is "" (zero value for string)
func Or[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

// OrElse returns the first non-zero value from the given values.
// If all values are zero, returns the default value.
// This provides a guaranteed non-zero result when a default is provided.
//
// Parameters:
//   - defaultValue: The value to return if all values are zero
//   - values: Variable number of values to check
//
// Returns:
//   - T: The first non-zero value, or defaultValue if all are zero
//
// Example:
//
//	result := value.OrElse("default", "", "", "fallback")
//	// result is "fallback"
//
//	defaultResult := value.OrElse("default", "", "", "")
//	// defaultResult is "default"
func OrElse[T comparable](defaultValue T, values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return defaultValue
}

// IsZero checks if the value is the zero value for its type.
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
//	value.IsZero(0)        // true
//	value.IsZero("")        // true
//	value.IsZero(false)     // true
//	value.IsZero(42)        // false
//	value.IsZero("hello")   // false
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

// IsNotZero checks if the value is not the zero value for its type.
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
//	value.IsNotZero(42)     // true
//	value.IsNotZero("hello") // true
//	value.IsNotZero(0)      // false
//	value.IsNotZero("")     // false
func IsNotZero[T comparable](v T) bool {
	return !IsZero(v)
}

// IsNil checks if the value is nil (for pointer, slice, map, channel, function, interface).
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
//	value.IsNil(ptr)        // true
//	value.IsNil(nil)        // true
//	value.IsNil([]int{})    // false (empty slice is not nil)
//	value.IsNil((*int)(nil)) // true
func IsNil(v interface{}) bool {
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

// IsNotNil checks if the value is not nil.
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
//	value.IsNotNil(ptr)     // true
//	value.IsNotNil([]int{}) // true (empty slice is not nil)
//	value.IsNotNil(nil)     // false
func IsNotNil(v interface{}) bool {
	return !IsNil(v)
}

// IsEmpty checks if the value is empty (zero value or empty string/slice/map).
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
//	value.IsEmpty("")           // true
//	value.IsEmpty([]int{})      // true
//	value.IsEmpty(map[string]int{}) // true
//	value.IsEmpty(0)            // true
//	value.IsEmpty("hello")      // false
//	value.IsEmpty([]int{1, 2})  // false
func IsEmpty(v interface{}) bool {
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

// IsNotEmpty checks if the value is not empty.
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
//	value.IsNotEmpty("hello")     // true
//	value.IsNotEmpty([]int{1, 2}) // true
//	value.IsNotEmpty(42)          // true
//	value.IsNotEmpty("")          // false
//	value.IsNotEmpty([]int{})     // false
func IsNotEmpty(v interface{}) bool {
	return !IsEmpty(v)
}

// If returns the value if condition is true, otherwise returns the zero value.
// This is a simple conditional function that returns either the value or zero.
//
// Parameters:
//   - condition: The boolean condition to evaluate
//   - value: The value to return if condition is true
//
// Returns:
//   - T: The value if condition is true, zero value otherwise
//
// Example:
//
//	result := value.If(age >= 18, "adult")
//	// result is "adult" if age >= 18, "" otherwise
func If[T any](condition bool, value T) T {
	if condition {
		return value
	}
	var zero T
	return zero
}

// When returns the value if the predicate function returns true, otherwise returns zero value.
// This allows for more complex conditional logic using a predicate function.
//
// Parameters:
//   - value: The value to potentially return
//   - predicate: A function that takes the value and returns a boolean
//
// Returns:
//   - T: The value if predicate returns true, zero value otherwise
//
// Example:
//
//	result := value.When(42, func(x int) bool { return x > 0 })
//	// result is 42 if x > 0, 0 otherwise
func When[T any](value T, predicate func(T) bool) T {
	if predicate(value) {
		return value
	}
	var zero T
	return zero
}

// WhenElse returns value1 if the predicate function returns true, otherwise returns value2.
// This provides a conditional with both true and false branches.
//
// Parameters:
//   - value: The value to test with the predicate
//   - predicate: A function that takes the value and returns a boolean
//   - value1: The value to return if predicate returns true
//   - value2: The value to return if predicate returns false
//
// Returns:
//   - T: value1 if predicate returns true, value2 otherwise
//
// Example:
//
//	result := value.WhenElse(42, func(x int) bool { return x > 0 }, "positive", "negative")
//	// result is "positive" if x > 0, "negative" otherwise
func WhenElse[T any](value T, predicate func(T) bool, value1, value2 T) T {
	if predicate(value) {
		return value1
	}
	return value2
}

// Coalesce returns the first non-nil pointer from the given pointers.
// This is useful for providing fallback pointers in a chain.
//
// Parameters:
//   - pointers: Variable number of pointers to check
//
// Returns:
//   - *T: The first non-nil pointer, or nil if all are nil
//
// Example:
//
//	var p1, p2 *int
//	p3 := &42
//	result := value.Coalesce(p1, p2, p3)
//	// result is p3 (points to 42)
func Coalesce[T any](pointers ...*T) *T {
	for _, p := range pointers {
		if p != nil {
			return p
		}
	}
	return nil
}

// CoalesceValue returns the first non-nil pointer's value from the given pointers.
// This dereferences the first non-nil pointer and returns its value.
//
// Parameters:
//   - pointers: Variable number of pointers to check
//
// Returns:
//   - T: The value of the first non-nil pointer, or zero value if all are nil
//
// Example:
//
//	var p1, p2 *int
//	p3 := &42
//	result := value.CoalesceValue(p1, p2, p3)
//	// result is 42
func CoalesceValue[T any](pointers ...*T) T {
	for _, p := range pointers {
		if p != nil {
			return *p
		}
	}
	var zero T
	return zero
}

// CoalesceValueDef returns the first non-nil pointer's value, or default value if all are nil.
// This provides a guaranteed non-zero result when a default is provided.
//
// Parameters:
//   - defaultValue: The value to return if all pointers are nil
//   - pointers: Variable number of pointers to check
//
// Returns:
//   - T: The value of the first non-nil pointer, or defaultValue if all are nil
//
// Example:
//
//	var p1, p2 *int
//	result := value.CoalesceValueDef(0, p1, p2)
//	// result is 0 (default value)
func CoalesceValueDef[T any](defaultValue T, pointers ...*T) T {
	for _, p := range pointers {
		if p != nil {
			return *p
		}
	}
	return defaultValue
}

// SafeDeref safely dereferences a pointer, returning the zero value if nil.
// This prevents panics when dereferencing nil pointers.
//
// Parameters:
//   - ptr: The pointer to dereference
//
// Returns:
//   - T: The value pointed to by ptr, or zero value if ptr is nil
//
// Example:
//
//	var ptr *int
//	result := value.SafeDeref(ptr)
//	// result is 0 (zero value for int)
//
//	ptr = &42
//	result = value.SafeDeref(ptr)
//	// result is 42
func SafeDeref[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// SafeDerefDef safely dereferences a pointer, returning the default value if nil.
// This prevents panics when dereferencing nil pointers and provides a fallback value.
//
// Parameters:
//   - ptr: The pointer to dereference
//   - defaultValue: The value to return if ptr is nil
//
// Returns:
//   - T: The value pointed to by ptr, or defaultValue if ptr is nil
//
// Example:
//
//	var ptr *int
//	result := value.SafeDerefDef(ptr, 100)
//	// result is 100 (default value)
//
//	ptr = &42
//	result = value.SafeDerefDef(ptr, 100)
//	// result is 42
func SafeDerefDef[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// Value returns the value from a pointer, or zero value if nil.
// This is an alias for SafeDeref for consistency.
//
// Parameters:
//   - ptr: The pointer to dereference
//
// Returns:
//   - T: The value pointed to by ptr, or zero value if ptr is nil
//
// Example:
//
//	var ptr *int
//	result := value.Value(ptr)
//	// result is 0 (zero value for int)
func Value[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// Def returns the value from a pointer, or default value if nil.
// This is an alias for SafeDerefDef for consistency.
//
// Parameters:
//   - ptr: The pointer to dereference
//   - defaultValue: The value to return if ptr is nil
//
// Returns:
//   - T: The value pointed to by ptr, or defaultValue if ptr is nil
//
// Example:
//
//	var ptr *int
//	result := value.Def(ptr, 100)
//	// result is 100 (default value)
func Def[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
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
