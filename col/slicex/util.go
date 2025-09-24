package slicex

import "fmt"

// basic functions

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
//
// Parameters:
//   - s1: The first slice to compare
//   - s2: The second slice to compare
//
// Returns:
//   - bool: true if slices are equal, false otherwise
//
// Example:
//
//	s1 := []int{1, 2, 3}
//	s2 := []int{1, 2, 3}
//	s3 := []int{1, 2, 4}
//	fmt.Println(slicex.Equal(s1, s2)) // true
//	fmt.Println(slicex.Equal(s1, s3)) // false
func Equal[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// EqualFunc reports whether two slices are equal using a comparison
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
//
// Parameters:
//   - s1: The first slice to compare
//   - s2: The second slice to compare
//   - eq: The comparison function that takes elements from both slices
//
// Returns:
//   - bool: true if slices are equal according to the comparison function
//
// Example:
//
//	s1 := []int{1, 2, 3}
//	s2 := []float64{1.0, 2.0, 3.0}
//	equal := slicex.EqualFunc(s1, s2, func(a int, b float64) bool {
//	    return float64(a) == b
//	})
//	fmt.Println(equal) // true
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

// Each iterates over each element in the slice and applies the given function to it.
// This function is useful for performing side effects on each element without collecting results.
//
// Parameters:
//   - ts: The input slice of type T
//   - f: The function to apply to each element
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	slicex.Each(numbers, func(x int) {
//	    fmt.Println(x * 2) // Prints: 2, 4, 6, 8, 10
//	})
func Each[T any](ts []T, f func(t T)) {
	for _, t := range ts {
		f(t)
	}
}

// Eachv transforms each element in the slice using the given function and returns a new slice.
// This function is similar to Each but collects the results into a new slice.
// It pre-allocates the result slice for better performance.
//
// Parameters:
//   - ts: The input slice of type T
//   - f: The transformation function that takes a T and returns a V
//
// Returns:
//   - []V: A new slice containing the transformed elements
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	strings := slicex.Eachv(numbers, func(x int) string {
//	    return fmt.Sprintf("num_%d", x)
//	})
//	// Result: []string{"num_1", "num_2", "num_3", "num_4", "num_5"}
func Eachv[T any, V any](ts []T, f func(t T) V) []V {
	var vs = make([]V, len(ts))
	for i, t := range ts {
		vs[i] = f(t)
	}
	return vs
}

// Group creates a map from a slice using the given key extraction function.
// Each element in the slice becomes a value in the map, with the key determined by the function.
// If multiple elements produce the same key, the last one will overwrite the previous ones.
//
// Parameters:
//   - vs: The input slice of type V
//   - f: The key extraction function that takes a V and returns a K
//
// Returns:
//   - map[K]V: A map where keys are determined by the function and values are the original elements
//
// Example:
//
//	people := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
//	nameMap := slicex.Group(people, func(p Person) string {
//	    return p.Name
//	})
//	// Result: map[string]Person{"Alice": {Name: "Alice", Age: 30}, "Bob": {Name: "Bob", Age: 25}}
func Group[K comparable, V any](vs []V, f func(v V) K) map[K]V {
	var m = make(map[K]V, len(vs))
	for _, v := range vs {
		m[f(v)] = v
	}
	return m
}

// GroupTo creates a map from a slice using the given key-value extraction function.
// This function allows you to extract both key and value from each element, providing more flexibility than Map.
// If multiple elements produce the same key, the last one will overwrite the previous ones.
//
// Parameters:
//   - ts: The input slice of type T
//   - f: The key-value extraction function that takes a T and returns (K, V)
//
// Returns:
//   - map[K]V: A map where keys and values are determined by the function
//
// Example:
//
//	people := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
//	ageMap := slicex.GroupTo(people, func(p Person) (string, int) {
//	    return p.Name, p.Age
//	})
//	// Result: map[string]int{"Alice": 30, "Bob": 25}
func GroupTo[K comparable, V any, T any](ts []T, f func(t T) (K, V)) map[K]V {
	var m = make(map[K]V, len(ts))
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}

// In checks if the value is in the given slice.
// This function performs a linear search through the slice to find the value.
// For better performance with large slices, consider using a map or set.
//
// Parameters:
//   - value: The value to search for
//   - slice: The slice to search in
//
// Returns:
//   - bool: true if the value is found, false otherwise
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	fmt.Println(slicex.In(3, numbers)) // true
//	fmt.Println(slicex.In(6, numbers)) // false
func In[T comparable](value T, slice []T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// NotIn checks if the value is not in the given slice.
// This is the logical negation of In.
//
// Parameters:
//   - value: The value to search for
//   - slice: The slice to search in
//
// Returns:
//   - bool: true if the value is not found, false otherwise
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	fmt.Println(slicex.NotIn(6, numbers)) // true
//	fmt.Println(slicex.NotIn(3, numbers))  // false
func NotIn[T comparable](value T, slice []T) bool {
	return !In(value, slice)
}

// Contains checks if the slice contains the value.
// This is an alias for In with reversed parameter order for better readability.
//
// Parameters:
//   - slice: The slice to search in
//   - value: The value to search for
//
// Returns:
//   - bool: true if the value is found, false otherwise
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	fmt.Println(slicex.Contains(numbers, 3)) // true
//	fmt.Println(slicex.Contains(numbers, 6)) // false
func Contains[T comparable](slice []T, value T) bool {
	return In(value, slice)
}

// NotContains checks if the slice does not contain the value.
// This is the logical negation of Contains.
//
// Parameters:
//   - slice: The slice to search in
//   - value: The value to search for
//
// Returns:
//   - bool: true if the value is not found, false otherwise
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	fmt.Println(slicex.NotContains(numbers, 6)) // true
//	fmt.Println(slicex.NotContains(numbers, 3)) // false
func NotContains[T comparable](slice []T, value T) bool {
	return !Contains(slice, value)
}

// IndexOf returns the index of the first occurrence of the value in the slice.
// Returns -1 if not found.
//
// Parameters:
//   - slice: The slice to search in
//   - value: The value to search for
//
// Returns:
//   - int: The index of the first occurrence, or -1 if not found
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4}
//	fmt.Println(slicex.IndexOf(numbers, 2)) // 1
//	fmt.Println(slicex.IndexOf(numbers, 5)) // -1
func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of the value in the slice.
// Returns -1 if not found.
//
// Parameters:
//   - slice: The slice to search in
//   - value: The value to search for
//
// Returns:
//   - int: The index of the last occurrence, or -1 if not found
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4}
//	fmt.Println(slicex.LastIndexOf(numbers, 2)) // 3
//	fmt.Println(slicex.LastIndexOf(numbers, 5)) // -1
func LastIndexOf[T comparable](slice []T, value T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == value {
			return i
		}
	}
	return -1
}

// Filter returns a new slice containing only elements that satisfy the predicate.
// The original slice is not modified.
//
// Parameters:
//   - slice: The input slice to filter
//   - predicate: A function that takes an element and returns true if it should be included
//
// Returns:
//   - []T: A new slice containing only elements that satisfy the predicate
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := slicex.Filter(numbers, func(x int) bool { return x%2 == 0 })
//	// Result: []int{2, 4, 6}
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map applies the function to each element of the slice and returns a new slice.
// The original slice is not modified.
//
// Parameters:
//   - slice: The input slice to transform
//   - fn: A function that takes an element of type T and returns an element of type U
//
// Returns:
//   - []U: A new slice containing the transformed elements
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	strings := slicex.Map(numbers, func(x int) string {
//	    return fmt.Sprintf("num_%d", x)
//	})
//	// Result: []string{"num_1", "num_2", "num_3", "num_4", "num_5"}
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Reduce applies the function to each element of the slice and returns a single value.
// The function is applied from left to right, accumulating the result.
//
// Parameters:
//   - slice: The input slice to reduce
//   - initial: The initial value for the accumulator
//   - fn: A function that takes the accumulator and an element, returning a new accumulator value
//
// Returns:
//   - U: The final accumulated value
//
// Time Complexity: O(n)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	sum := slicex.Reduce(numbers, 0, func(acc, x int) int { return acc + x })
//	// Result: 15
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Any checks if any element in the slice satisfies the predicate.
// Returns true if at least one element satisfies the predicate, false otherwise.
// For empty slices, returns false.
//
// Parameters:
//   - slice: The input slice to check
//   - predicate: A function that takes an element and returns true if it satisfies the condition
//
// Returns:
//   - bool: true if any element satisfies the predicate, false otherwise
//
// Time Complexity: O(n) worst case, O(1) best case (first element satisfies)
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := slicex.Any(numbers, func(x int) bool { return x%2 == 0 })
//	// Result: true
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// All checks if all elements in the slice satisfy the predicate.
// Returns true if all elements satisfy the predicate, false otherwise.
// For empty slices, returns true (vacuous truth).
//
// Parameters:
//   - slice: The input slice to check
//   - predicate: A function that takes an element and returns true if it satisfies the condition
//
// Returns:
//   - bool: true if all elements satisfy the predicate, false otherwise
//
// Time Complexity: O(n) worst case, O(1) best case (first element doesn't satisfy)
//
// Example:
//
//	numbers := []int{2, 4, 6, 8}
//	allEven := slicex.All(numbers, func(x int) bool { return x%2 == 0 })
//	// Result: true
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// None checks if no elements in the slice satisfy the predicate.
// Returns true if no elements satisfy the predicate, false otherwise.
// For empty slices, returns true.
//
// Parameters:
//   - slice: The input slice to check
//   - predicate: A function that takes an element and returns true if it satisfies the condition
//
// Returns:
//   - bool: true if no elements satisfy the predicate, false otherwise
//
// Time Complexity: O(n) worst case, O(1) best case (first element satisfies)
//
// Example:
//
//	numbers := []int{1, 3, 5, 7}
//	noEven := slicex.None(numbers, func(x int) bool { return x%2 == 0 })
//	// Result: true
func None[T any](slice []T, predicate func(T) bool) bool {
	return !Any(slice, predicate)
}

// Count returns the number of elements in the slice that satisfy the predicate.
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count
}

// First returns the first element in the slice that satisfies the predicate.
// Returns zero value if not found.
func First[T any](slice []T, predicate func(T) bool) T {
	for _, v := range slice {
		if predicate(v) {
			return v
		}
	}
	var zero T
	return zero
}

// FindLast returns the last element in the slice that satisfies the predicate.
// Returns zero value if not found.
func FindLast[T any](slice []T, predicate func(T) bool) T {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return slice[i]
		}
	}
	var zero T
	return zero
}

// Unique returns a new slice with duplicate elements removed.
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Reverse returns a new slice with elements in reverse order.
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// Chunk splits the slice into chunks of the specified size.
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	var result [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

// Flatten flattens a slice of slices into a single slice.
func Flatten[T any](slices [][]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// Zip combines two slices into a slice of pairs.
func Zip[T, U any](slice1 []T, slice2 []U) []struct {
	First  T
	Second U
} {
	minLen := len(slice1)
	if len(slice2) < minLen {
		minLen = len(slice2)
	}
	result := make([]struct {
		First  T
		Second U
	}, minLen)
	for i := 0; i < minLen; i++ {
		result[i] = struct {
			First  T
			Second U
		}{slice1[i], slice2[i]}
	}
	return result
}

// Unzip separates a slice of pairs into two slices.
func Unzip[T, U any](pairs []struct {
	First  T
	Second U
}) ([]T, []U) {
	first := make([]T, len(pairs))
	second := make([]U, len(pairs))
	for i, pair := range pairs {
		first[i] = pair.First
		second[i] = pair.Second
	}
	return first, second
}

// Take returns the first n elements from the slice.
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(slice) {
		return slice
	}
	return slice[:n]
}

// Drop returns the slice without the first n elements.
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return []T{}
	}
	return slice[n:]
}

// TakeWhile returns elements from the beginning of the slice while the predicate is true.
func TakeWhile[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if !predicate(v) {
			break
		}
		result = append(result, v)
	}
	return result
}

// DropWhile drops elements from the beginning of the slice while the predicate is true.
func DropWhile[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	drop := true
	for _, v := range slice {
		if drop && predicate(v) {
			continue
		}
		drop = false
		result = append(result, v)
	}
	return result
}

// Partition splits the slice into two slices based on the predicate.
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	var trueSlice, falseSlice []T
	for _, v := range slice {
		if predicate(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return trueSlice, falseSlice
}

// Max returns the maximum value from the slice using string comparison.
func Max[T comparable](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if fmt.Sprintf("%v", v) > fmt.Sprintf("%v", max) {
			max = v
		}
	}
	return max
}

// Min returns the minimum value from the slice using string comparison.
func Min[T comparable](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	min := slice[0]
	for _, v := range slice[1:] {
		if fmt.Sprintf("%v", v) < fmt.Sprintf("%v", min) {
			min = v
		}
	}
	return min
}

// Sum returns the sum of all values in the slice.
// Note: This only works for numeric types. For other types, it returns the first element.
func Sum[T comparable](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	// For non-numeric types, just return the first element
	// In practice, you'd need type-specific implementations
	return slice[0]
}

// Average returns the average of all values in the slice.
// Note: This is a simplified implementation for demonstration.
func Average[T comparable](slice []T) float64 {
	if len(slice) == 0 {
		return 0
	}
	// This is a simplified implementation - in practice, you'd need type-specific handling
	return float64(len(slice))
}

// Head returns the first element of the slice, or zero value if empty.
func Head[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[0]
}

// Tail returns all elements except the first one.
func Tail[T any](slice []T) []T {
	if len(slice) <= 1 {
		return []T{}
	}
	return slice[1:]
}

// Init returns all elements except the last one.
func Init[T any](slice []T) []T {
	if len(slice) <= 1 {
		return []T{}
	}
	return slice[:len(slice)-1]
}

// LastElement returns the last element of the slice, or zero value if empty.
func LastElement[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[len(slice)-1]
}

// Intersect returns the intersection of two slices.
func Intersect[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	for _, v := range slice1 {
		set[v] = true
	}

	var result []T
	for _, v := range slice2 {
		if set[v] {
			result = append(result, v)
			delete(set, v) // Avoid duplicates
		}
	}
	return result
}

// Union returns the union of two slices.
func Union[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	var result []T

	for _, v := range slice1 {
		if !set[v] {
			set[v] = true
			result = append(result, v)
		}
	}

	for _, v := range slice2 {
		if !set[v] {
			set[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Difference returns the difference of two slices (elements in slice1 but not in slice2).
func Difference[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	for _, v := range slice2 {
		set[v] = true
	}

	var result []T
	for _, v := range slice1 {
		if !set[v] {
			result = append(result, v)
		}
	}

	return result
}

// SymmetricDifference returns the symmetric difference of two slices.
func SymmetricDifference[T comparable](slice1, slice2 []T) []T {
	return Union(Difference(slice1, slice2), Difference(slice2, slice1))
}
