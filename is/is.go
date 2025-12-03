package is

import (
	"cmp"
	"reflect"
)

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
//	is.Zero(0)        // true
//	is.Zero("")        // true
//	is.Zero(false)     // true
//	is.Zero(42)        // false
//	is.Zero("hello")   // false
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
//	is.NotZero(42)     // true
//	is.NotZero("hello") // true
//	is.NotZero(0)      // false
//	is.NotZero("")     // false
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
//	is.Nil(ptr)        // true
//	is.Nil(nil)        // true
//	is.Nil([]int{})    // false (empty slice is not nil)
func Nil(v any) bool {
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
//	is.NotNil(ptr)     // true
//	is.NotNil([]int{}) // true (empty slice is not nil)
//	is.NotNil(nil)     // false
func NotNil(v any) bool {
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
//	is.Empty("")           // true
//	is.Empty([]int{})      // true
//	is.Empty(map[string]int{}) // true
//	is.Empty(0)            // true
//	is.Empty("hello")      // false
//	is.Empty([]int{1, 2})  // false
func Empty(v any) bool {
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
//	is.NotEmpty("hello")     // true
//	is.NotEmpty([]int{1, 2}) // true
//	is.NotEmpty(42)          // true
//	is.NotEmpty("")          // false
//	is.NotEmpty([]int{})     // false
func NotEmpty(v any) bool {
	return !Empty(v)
}

// Eq checks if two values are equal.
// For basic types, it uses the == operator.
// For pointers, it compares the values they point to (not the pointer addresses).
// For structs and other composite types, it performs a deep comparison.
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
//	// Basic types
//	is.Eq(42, 42)        // true
//	is.Eq("hello", "hi") // false
//
//	// Pointers - compares values, not addresses
//	val1, val2 := 42, 42
//	ptr1, ptr2 := &val1, &val2
//	is.Eq(ptr1, ptr2)    // true (values are equal, even though pointers differ)
//
//	// Structs
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	p1 := Person{Name: "John", Age: 30}
//	p2 := Person{Name: "John", Age: 30}
//	is.Eq(p1, p2)        // true
//
// For functions, it uses reflect.DeepEqual, only return true if both are nil, otherwise return false
// Example:
//
//	type Func func() int
//	var v1 Func = nil
//	var v2 Func = nil
//	is.Eq(v1, v2) // true
//	v1 := func() int { return 1 }
//	v2 := v1
//	v3 := func() int { return 1 }
//	is.Eq(v1, v2) // false
//	is.Eq(v1, v3) // false
func Eq(a, b any) bool {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	// Handle invalid values (nil interfaces)
	if !av.IsValid() && !bv.IsValid() {
		return true
	}
	if !av.IsValid() || !bv.IsValid() {
		return false
	}

	switch av.Kind() {
	case reflect.Interface:
		if av.IsNil() && bv.IsNil() {
			return true
		}
		if av.IsNil() || bv.IsNil() {
			return false
		}
		// Compare the concrete values stored in the interfaces
		return Eq(av.Elem().Interface(), bv.Elem().Interface())
	case reflect.Ptr:
		if av.IsNil() && bv.IsNil() {
			return true
		}
		if av.IsNil() || bv.IsNil() {
			return false
		}
		// Recursively compare the values pointed to
		// This handles nested pointers correctly
		return Eq(av.Elem().Interface(), bv.Elem().Interface())
	case reflect.Chan:
		// Channels compare by reference, not by content
		// This is the standard Go behavior
		return av.Interface() == bv.Interface()
	case reflect.Func:
		// for functions, use deep equality, only return true if both are nil, otherwise return false
		return reflect.DeepEqual(a, b)
	}

	// For comparable types (including structs with all comparable fields, arrays), use == operator
	if av.Type().Comparable() && bv.Type().Comparable() {
		return av.Interface() == bv.Interface()
	}

	// For non-comparable types (slices, maps, etc.), use deep equality
	return reflect.DeepEqual(a, b)
}

// Same checks if two values are the same.
// This is an alias for Eq, using the == operator for comparison.
// Same emphasizes that the values are identical, not just equal.
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if the values are the same, false otherwise
//
// Example:
//
//	is.Same(42, 42)        // true
//	is.Same("hello", "hi") // false
//	is.Same(0, 0)          // true
//
//	var ptr1, ptr2 *int
//	val := 42
//	ptr1 = &val
//	ptr2 = &val
//	is.Same(ptr1, ptr2)    // true (same pointer)
func Same[T comparable](a, b T) bool {
	return a == b
}

// NotSame checks if two values are not the same.
// This is the logical negation of Same.
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if the values are not the same, false otherwise
//
// Example:
//
//	is.NotSame(42, 43)     // true
//	is.NotSame("hello", "hello") // false
//	is.NotSame(0, 1)       // true
func NotSame[T comparable](a, b T) bool {
	return Not(Same(a, b))
}

// Neq checks if two values are not equal.
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
//	is.Neq(42, 43)     // true
//	is.Neq("hello", "hello") // false
//	is.Neq(0, 1)       // true
func Neq[T comparable](a, b T) bool {
	return Not(Eq(a, b))
}

// Gt checks if a is greater than b.
// This function requires types that support ordering (numbers, strings, etc.).
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if a > b, false otherwise
//
// Example:
//
//	is.Gt(42, 10)        // true
//	is.Gt("z", "a")      // true
//	is.Gt(5, 10)         // false
func Gt[T cmp.Ordered](a, b T) bool {
	return a > b
}

// Gte checks if a is greater than or equal to b.
// This function requires types that support ordering (numbers, strings, etc.).
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if a >= b, false otherwise
//
// Example:
//
//	is.Gte(42, 42)       // true
//	is.Gte(42, 10)       // true
//	is.Gte(5, 10)        // false
func Gte[T cmp.Ordered](a, b T) bool {
	return a >= b
}

// Lt checks if a is less than b.
// This function requires types that support ordering (numbers, strings, etc.).
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if a < b, false otherwise
//
// Example:
//
//	is.Lt(5, 10)         // true
//	is.Lt("a", "z")      // true
//	is.Lt(42, 10)        // false
func Lt[T cmp.Ordered](a, b T) bool {
	return a < b
}

// Lte checks if a is less than or equal to b.
// This function requires types that support ordering (numbers, strings, etc.).
//
// Parameters:
//   - a: The first value to compare
//   - b: The second value to compare
//
// Returns:
//   - bool: true if a <= b, false otherwise
//
// Example:
//
//	is.Lte(42, 42)       // true
//	is.Lte(5, 10)        // true
//	is.Lte(42, 10)       // false
func Lte[T cmp.Ordered](a, b T) bool {
	return a <= b
}
