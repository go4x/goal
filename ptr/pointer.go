package ptr

import (
	"reflect"
)

func To[T any](v T) *T {
	return &v
}

func From[T any](v *T) T {
	return *v
}

func ToSlice[T any](v []T) []*T {
	var ret []*T
	for _, v := range v {
		ret = append(ret, To(v))
	}
	return ret
}

func FromSlice[T any](v []*T) []T {
	var ret []T
	for _, v := range v {
		ret = append(ret, From(v))
	}
	return ret
}

// IsNil checks if a pointer is nil
func IsNil[T any](v *T) bool {
	return v == nil
}

// IsNotNil checks if a pointer is not nil
func IsNotNil[T any](v *T) bool {
	return v != nil
}

// ValueOr returns the value if pointer is not nil, otherwise returns the default value
func ValueOr[T any](v *T, defaultValue T) T {
	if v == nil {
		return defaultValue
	}
	return *v
}

// ValueOrDefault returns the value if pointer is not nil, otherwise returns the zero value
func ValueOrDefault[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

// Deref safely dereferences a pointer, returns zero value if nil
func Deref[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

// DerefOr safely dereferences a pointer, returns the specified value if nil
func DerefOr[T any](v *T, defaultValue T) T {
	if v == nil {
		return defaultValue
	}
	return *v
}

// Equal compares two pointer values for equality (handles nil cases)
func Equal[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// Clone clones the value pointed to by the pointer
func Clone[T any](v *T) *T {
	if v == nil {
		return nil
	}
	cloned := *v
	return &cloned
}

// CloneSlice clones a slice of pointers
func CloneSlice[T any](v []*T) []*T {
	if v == nil {
		return nil
	}
	result := make([]*T, len(v))
	for i, ptr := range v {
		if ptr != nil {
			cloned := *ptr
			result[i] = &cloned
		}
	}
	return result
}

// Filter filters a slice of pointers, returns slice of non-nil pointers
func Filter[T any](v []*T) []*T {
	if v == nil {
		return nil
	}
	var result []*T
	for _, ptr := range v {
		if ptr != nil {
			result = append(result, ptr)
		}
	}
	return result
}

// FilterValues filters a slice of pointers, returns slice of values from non-nil pointers
func FilterValues[T any](v []*T) []T {
	if v == nil {
		return nil
	}
	var result []T
	for _, ptr := range v {
		if ptr != nil {
			result = append(result, *ptr)
		}
	}
	return result
}

// Map maps a slice of pointers using the provided function
func Map[T, U any](v []*T, fn func(*T) *U) []*U {
	if v == nil {
		return nil
	}
	result := make([]*U, len(v))
	for i, ptr := range v {
		result[i] = fn(ptr)
	}
	return result
}

// MapValues maps the values of a slice of pointers using the provided function
func MapValues[T, U any](v []*T, fn func(T) U) []U {
	if v == nil {
		return nil
	}
	result := make([]U, len(v))
	for i, ptr := range v {
		if ptr != nil {
			result[i] = fn(*ptr)
		} else {
			var zero U
			result[i] = zero
		}
	}
	return result
}

// Any checks if there are any non-nil pointers in the slice
func Any[T any](v []*T) bool {
	if v == nil {
		return false
	}
	for _, ptr := range v {
		if ptr != nil {
			return true
		}
	}
	return false
}

// All checks if all pointers in the slice are non-nil
func All[T any](v []*T) bool {
	if v == nil {
		return true
	}
	for _, ptr := range v {
		if ptr == nil {
			return false
		}
	}
	return true
}

// Count counts the number of non-nil pointers in the slice
func Count[T any](v []*T) int {
	if v == nil {
		return 0
	}
	count := 0
	for _, ptr := range v {
		if ptr != nil {
			count++
		}
	}
	return count
}

// First returns the first non-nil pointer in the slice
func First[T any](v []*T) *T {
	if v == nil {
		return nil
	}
	for _, ptr := range v {
		if ptr != nil {
			return ptr
		}
	}
	return nil
}

// Last returns the last non-nil pointer in the slice
func Last[T any](v []*T) *T {
	if v == nil {
		return nil
	}
	for i := len(v) - 1; i >= 0; i-- {
		if v[i] != nil {
			return v[i]
		}
	}
	return nil
}

// Set sets the value pointed to by the pointer
func Set[T any](v *T, value T) {
	if v != nil {
		*v = value
	}
}

// Zero sets the value pointed to by the pointer to its zero value
func Zero[T any](v *T) {
	if v != nil {
		var zero T
		*v = zero
	}
}

// Swap swaps the values pointed to by two pointers
func Swap[T any](a, b *T) {
	if a != nil && b != nil {
		*a, *b = *b, *a
	}
}

// DeepEqual performs deep comparison of two pointer values
func DeepEqual[T any](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return reflect.DeepEqual(*a, *b)
}
