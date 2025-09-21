package set

import (
	"github.com/go4x/goal/col/mapx"
)

// ArraySet is a set implementation using array-based map for ordered elements.
// It maintains insertion order but has O(n) operations, making it suitable for small datasets.
//
// Time Complexity:
//   - Add: O(n) if element exists (linear search), O(1) if new element (append)
//   - Remove: O(n) (linear search and slice manipulation)
//   - Contains: O(n) (linear search)
//   - Size/IsEmpty: O(1)
//   - Elems: O(n) (direct slice iteration)
//
// Use Cases:
//   - Small datasets (< 1000 elements)
//   - When insertion order is important
//   - Memory-constrained environments
//   - Simple applications where O(n) performance is acceptable
//
// Example:
//
//	set := NewArraySet[string]()
//	set.Add("first").Add("second").Add("third")
//	elems := set.Elems() // Returns ["first", "second", "third"] in insertion order
type ArraySet[T comparable] struct {
	data mapx.Map[T, struct{}]
}

// NewArraySet creates a new ArraySet that maintains insertion order
func newArraySet[T comparable]() *ArraySet[T] {
	return &ArraySet[T]{
		data: mapx.NewArrayMap[T, struct{}](),
	}
}

// Add adds an element to the set (maintains insertion order)
func (a *ArraySet[T]) Add(t T) Set[T] {
	a.data.Put(t, struct{}{})
	return a
}

// Remove removes an element from the set
func (a *ArraySet[T]) Remove(t T) Set[T] {
	a.data.Del(t)
	return a
}

// Size returns the number of elements in the set
func (a *ArraySet[T]) Size() int {
	return a.data.Size()
}

// IsEmpty returns true if the set is empty
func (a *ArraySet[T]) IsEmpty() bool {
	return a.data.IsEmpty()
}

// Contains returns true if the set contains the element
func (a *ArraySet[T]) Contains(t T) bool {
	_, exists := a.data.Get(t)
	return exists
}

// Clear removes all elements from the set
func (a *ArraySet[T]) Clear() Set[T] {
	a.data.Clear()
	return a
}

// Elems returns a slice containing all elements in insertion order
func (a *ArraySet[T]) Elems() []T {
	var result []T
	a.data.Each(func(k T, v struct{}) {
		result = append(result, k)
	})
	return result
}
