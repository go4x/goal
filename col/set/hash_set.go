package set

import (
	"github.com/gophero/goal/col/mapx"
)

// HashSet is a set implementation using hash map for O(1) operations.
// It provides the fastest performance but does not guarantee insertion order.
//
// Time Complexity:
//   - Add: O(1) average-case
//   - Remove: O(1) average-case
//   - Contains: O(1) average-case
//   - Size/IsEmpty: O(1)
//   - Elems: O(n)
//
// Use Cases:
//   - General-purpose set operations
//   - When order doesn't matter
//   - Performance-critical applications
//   - Large datasets
//
// Example:
//
//	set := NewHashSet[string]()
//	set.Add("apple").Add("banana").Add("apple") // "apple" is only added once
//	fmt.Println(set.Size()) // Output: 2
type HashSet[T comparable] struct {
	data mapx.Map[T, struct{}]
}

// newHashSet creates a new HashSet with O(1) operations
func newHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		data: mapx.New[T, struct{}](),
	}
}

// Add adds an element to the set (no duplicates)
func (h *HashSet[T]) Add(t T) Set[T] {
	h.data.Put(t, struct{}{})
	return h
}

// Remove removes an element from the set
func (h *HashSet[T]) Remove(t T) Set[T] {
	h.data.Del(t)
	return h
}

// Size returns the number of elements in the set
func (h *HashSet[T]) Size() int {
	return h.data.Size()
}

// IsEmpty returns true if the set is empty
func (h *HashSet[T]) IsEmpty() bool {
	return h.data.IsEmpty()
}

// Contains returns true if the set contains the element
func (h *HashSet[T]) Contains(t T) bool {
	_, exists := h.data.Get(t)
	return exists
}

// Clear removes all elements from the set
func (h *HashSet[T]) Clear() Set[T] {
	h.data.Clear()
	return h
}

// Elems returns a slice containing all elements in the set
func (h *HashSet[T]) Elems() []T {
	var result []T
	h.data.Each(func(k T, v struct{}) {
		result = append(result, k)
	})
	return result
}
