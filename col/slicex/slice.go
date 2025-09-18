package slicex

import (
	"sort"
	"strings"

	"github.com/gophero/goal/stringx"
)

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

// S is a generic slice type that provides additional methods for slice manipulation.
// It wraps a standard Go slice with enhanced functionality for filtering, mapping,
// sorting, and other operations while maintaining immutability.
type S[T comparable] []T

// New creates a new empty slice of type S[T].
//
// Returns:
//   - S[T]: A new empty slice
//
// Example:
//
//	s := slicex.New[int]()
//	fmt.Println(len(s)) // 0
func New[T comparable]() S[T] {
	return make(S[T], 0)
}

// NewSize creates a new slice of type S[T] with the specified size.
//
// Parameters:
//   - size: The initial size of the slice
//
// Returns:
//   - S[T]: A new slice with the specified size
//
// Example:
//
//	s := slicex.NewSize[int](5)
//	fmt.Println(len(s)) // 5
func NewSize[T comparable](size int) S[T] {
	return make(S[T], size)
}

// From creates a new S[T] from a standard Go slice.
//
// Parameters:
//   - s: The input slice to wrap
//
// Returns:
//   - S[T]: A new S[T] instance wrapping the input slice
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	s := slicex.From(numbers)
//	fmt.Println(s.To()) // [1 2 3 4 5]
func From[T comparable](s []T) S[T] {
	return s
}

// To converts the S[T] slice back to a standard Go slice.
//
// Returns:
//   - []T: The underlying slice
//
// Example:
//
//	s := slicex.From([]int{1, 2, 3})
//	raw := s.To()
//	fmt.Println(raw) // [1 2 3]
func (s S[T]) To() []T {
	return s
}

// Retain keeps only the elements that match the given condition.
// Elements that do not match the condition are removed from the result.
// The original slice is not modified.
//
// Parameters:
//   - cond: The condition function that determines which elements to keep
//
// Returns:
//   - S[T]: A new slice containing only the elements that match the condition
//
// Example:
//
//	numbers := slicex.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
//	even := numbers.Retain(func(x int) bool {
//	    return x%2 == 0
//	})
//	fmt.Println(even.To()) // [2 4 6 8]
func (s S[T]) Retain(cond func(a T) bool) S[T] {
	var ret []T
	for _, a := range s {
		if cond(a) { // 符合条件
			ret = append(ret, a)
		}
	}
	return From(ret)
}

// Filter removes elements that match the given condition.
// This is the opposite of Retain - it keeps elements that do NOT match the condition.
// The original slice is not modified.
//
// Parameters:
//   - cond: The condition function that determines which elements to remove
//
// Returns:
//   - S[T]: A new slice containing only the elements that do not match the condition
//
// Example:
//
//	numbers := slicex.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
//	odd := numbers.Filter(func(x int) bool {
//	    return x%2 == 0  // Remove even numbers
//	})
//	fmt.Println(odd.To()) // [1 3 5 7 9]
func (s S[T]) Filter(cond func(a T) bool) S[T] {
	var ret []T
	for _, a := range s {
		if !cond(a) { // 不符合条件
			ret = append(ret, a)
		}
	}
	return From(ret)
}

// Join concatenates all elements into a single string using the specified separator.
// Each element is converted to a string using stringx.String().
//
// Parameters:
//   - sep: The separator string to use between elements
//
// Returns:
//   - string: The joined string
//
// Example:
//
//	numbers := slicex.From([]int{1, 2, 3, 4, 5})
//	result := numbers.Join(", ")
//	fmt.Println(result) // "1, 2, 3, 4, 5"
func (s S[T]) Join(sep string) string {
	var ret []string
	for _, a := range s {
		ret = append(ret, stringx.String(a))
	}
	return strings.Join(ret, sep)
}

// Union returns all unique elements from both slices (set union).
// Duplicate elements are removed from the result.
// The original slice is not modified.
//
// Parameters:
//   - dest: The second slice to union with
//
// Returns:
//   - S[T]: A new slice containing all unique elements from both slices
//
// Example:
//
//	s1 := slicex.From([]int{1, 2, 3, 4, 5})
//	s2 := []int{4, 5, 6, 7, 8}
//	union := s1.Union(s2)
//	fmt.Println(union.To()) // [1 2 3 4 5 6 7 8]
func (s S[T]) Union(dest []T) S[T] {
	// Create a map to track unique elements
	seen := make(map[T]struct{})
	var ret []T

	// Add all elements from source slice
	for _, v := range s {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			ret = append(ret, v)
		}
	}

	// Add elements from dest slice that are not already in the result
	for _, v := range dest {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			ret = append(ret, v)
		}
	}

	return From(ret)
}

// Intersect returns elements that exist in both slices (set intersection).
// The original slice is not modified.
//
// Parameters:
//   - dest: The second slice to intersect with
//
// Returns:
//   - S[T]: A new slice containing elements that exist in both slices
//
// Example:
//
//	s1 := slicex.From([]int{1, 2, 3, 4, 5})
//	s2 := []int{4, 5, 6, 7, 8}
//	intersection := s1.Intersect(s2)
//	fmt.Println(intersection.To()) // [4 5]
func (s S[T]) Intersect(dest []T) S[T] {
	var ret []T
	for _, v := range s {
		find := false
		for _, el := range dest {
			if v == el {
				find = true
				break
			}
		}
		if find {
			ret = append(ret, v)
		}
	}
	return From(ret)
}

// Diff returns the symmetric difference between two slices.
// It returns elements that are in either slice but not in both (union - intersection).
// The original slice is not modified.
//
// Parameters:
//   - dest: The second slice to compare with
//
// Returns:
//   - S[T]: A new slice containing elements that are in either slice but not in both
//
// Example:
//
//	s1 := slicex.From([]int{1, 2, 3, 4, 5})
//	s2 := []int{4, 5, 6, 7, 8}
//	diff := s1.Diff(s2)
//	fmt.Println(diff.To()) // [1 2 3 6 7 8]
func (s S[T]) Diff(dest []T) S[T] {
	// Get union of both slices
	union := s.Union(dest)
	// Get intersection of both slices
	intersection := s.Intersect(dest)
	// Return union minus intersection (symmetric difference)
	return union.Remove(intersection.To())
}

// Remove removes all elements that exist in the dest slice from the source slice.
// The original slice is not modified.
//
// Parameters:
//   - dest: The slice containing elements to remove
//
// Returns:
//   - S[T]: A new slice with the specified elements removed
//
// Example:
//
//	s1 := slicex.From([]int{1, 2, 3, 4, 5})
//	s2 := []int{2, 4}
//	result := s1.Remove(s2)
//	fmt.Println(result.To()) // [1 3 5]
func (s S[T]) Remove(dest []T) S[T] {
	return From(s.Delete(dest...))
}

// RemoveDuplicate removes all duplicate elements from the slice, keeping only the first occurrence.
// The original slice is not modified.
//
// Returns:
//   - S[T]: A new slice with duplicate elements removed
//
// Example:
//
//	numbers := slicex.From([]int{1, 2, 2, 3, 3, 3, 4, 5})
//	unique := numbers.RemoveDuplicate()
//	fmt.Println(unique.To()) // [1 2 3 4 5]
func (s S[T]) RemoveDuplicate() S[T] {
	seen := make(map[T]struct{})
	var ret []T

	for _, v := range s {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			ret = append(ret, v)
		}
	}

	return From(ret)
}

// Contain checks if the slice contains the given element.
//
// Parameters:
//   - e: The element to search for
//
// Returns:
//   - bool: true if the element is found, false otherwise
//
// Example:
//
//	numbers := slicex.From([]int{1, 2, 3, 4, 5})
//	hasThree := numbers.Contain(3)
//	fmt.Println(hasThree) // true
func (s S[T]) Contain(e T) bool {
	for _, el := range s {
		if el == e {
			return true
		}
	}
	return false
}

// Delete removes the specified elements from the slice.
// The original slice is not modified.
//
// Parameters:
//   - elem: The elements to remove (variadic)
//
// Returns:
//   - S[T]: A new slice with the specified elements removed
//
// Example:
//
//	numbers := slicex.From([]int{1, 2, 3, 4, 5})
//	result := numbers.Delete(2, 4)
//	fmt.Println(result.To()) // [1 3 5]
func (s S[T]) Delete(elem ...T) S[T] {
	var ret []T
	for _, v := range s {
		find := false
		for _, el := range elem {
			if v == el {
				find = true
				break
			}
		}
		if !find {
			ret = append(ret, v)
		}
	}
	return From(ret)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
// This can help reduce memory usage by trimming excess capacity.
// The original slice is not modified.
//
// Returns:
//   - S[T]: A new slice with capacity equal to length
//
// Example:
//
//	s := make([]int, 3, 10) // length 3, capacity 10
//	s[0], s[1], s[2] = 1, 2, 3
//	clipped := slicex.From(s).Clip()
//	fmt.Println(len(clipped), cap(clipped)) // 3 3
func (s S[T]) Clip() S[T] {
	return s[:len(s):len(s)]
}

// sortable slice

// SortableSlice is a struct that implements sort.Interface for sorting slices.
// It wraps an S[T] slice with a comparison function to define the sort order.
type SortableSlice[T comparable] struct {
	slice S[T]
	less  func(x, y T) bool // the method to compare the elements in the slice
}

// NewSortableSlice creates a new SortableSlice instance.
//
// Parameters:
//   - slice: The slice to make sortable
//   - less: The comparison function that defines the sort order
//
// Returns:
//   - SortableSlice[T]: A new sortable slice instance
//
// Example:
//
//	numbers := slicex.From([]int{3, 1, 4, 1, 5})
//	sortable := slicex.NewSortableSlice(numbers, func(a, b int) bool {
//	    return a < b  // Sort in ascending order
//	})
func NewSortableSlice[T comparable](slice S[T], less func(x, y T) bool) SortableSlice[T] {
	return SortableSlice[T]{slice, less}
}

// Len returns the number of elements in the sortable slice.
// This method is required by sort.Interface.
func (s SortableSlice[T]) Len() int {
	return len(s.slice)
}

// Less reports whether the element with index i should sort before the element with index j.
// This method is required by sort.Interface.
func (s SortableSlice[T]) Less(i, j int) bool {
	return s.less(s.slice[i], s.slice[j])
}

// Swap swaps the elements with indexes i and j.
// This method is required by sort.Interface.
func (s SortableSlice[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

// Sort sorts the slice using the provided comparison function.
// The original slice is not modified.
//
// Parameters:
//   - less: The comparison function that defines the sort order
//
// Returns:
//   - SortableSlice[T]: A sorted slice that can be further manipulated
//
// Example:
//
//	numbers := slicex.From([]int{3, 1, 4, 1, 5})
//	sorted := numbers.Sort(func(a, b int) bool {
//	    return a < b  // Sort in ascending order
//	})
//	fmt.Println(sorted.To()) // [1 1 3 4 5]
func (s S[T]) Sort(less func(x, y T) bool) SortableSlice[T] {
	v := &SortableSlice[T]{s, less}
	sort.Sort(v)
	return *v
}

// Reverse reverses the order of elements in the sortable slice.
// This method should be called on a SortableSlice that was created by Sort().
//
// Returns:
//   - S[T]: A new slice with elements in reverse order
//
// Example:
//
//	numbers := slicex.From([]int{1, 2, 3, 4, 5})
//	sorted := numbers.Sort(func(a, b int) bool { return a < b })
//	reversed := sorted.Reverse()
//	fmt.Println(reversed.To()) // [5 4 3 2 1]
func (s SortableSlice[T]) Reverse() S[T] {
	sort.Sort(sort.Reverse(s))
	return s.slice
}

// To converts the SortableSlice back to a standard Go slice.
//
// Returns:
//   - []T: The underlying slice
//
// Example:
//
//	numbers := slicex.From([]int{3, 1, 4, 1, 5})
//	sorted := numbers.Sort(func(a, b int) bool { return a < b })
//	raw := sorted.To()
//	fmt.Println(raw) // [1 1 3 4 5]
func (s SortableSlice[T]) To() []T {
	return s.slice.To()
}
