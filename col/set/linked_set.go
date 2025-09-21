package set

import (
	"github.com/gophero/goal/col/mapx"
)

// LinkedSet is a set implementation using linked map for O(1) operations with order.
// It maintains insertion order and provides O(1) average-case performance for all operations.
//
// Time Complexity:
//   - Add: O(1) average-case
//   - Remove: O(1) average-case
//   - Contains: O(1) average-case
//   - Size/IsEmpty: O(1)
//   - Elems: O(n)
//   - MoveToEnd/MoveToFront: O(1)
//
// Use Cases:
//   - Large datasets requiring order
//   - LRU cache implementations
//   - Performance-critical applications with order requirements
//   - When you need O(1) operations and insertion order
//
// Example:
//
//	set := NewLinkedSet[string]()
//	set.Add("first").Add("second").Add("third")
//	elems := set.Elems() // Returns ["first", "second", "third"] in insertion order
//
//	// For LRU cache operations
//	linkedSet := set.(*LinkedSet[string])
//	linkedSet.MoveToEnd("first") // Move "first" to end (most recently used)
type LinkedSet[T comparable] struct {
	data mapx.Map[T, struct{}]
}

// newLinkedSet creates a new LinkedSet with O(1) operations and insertion order
func newLinkedSet[T comparable]() *LinkedSet[T] {
	return &LinkedSet[T]{
		data: mapx.NewLinkedMap[T, struct{}](),
	}
}

// Add adds an element to the set (maintains insertion order)
func (l *LinkedSet[T]) Add(t T) Set[T] {
	l.data.Put(t, struct{}{})
	return l
}

// Remove removes an element from the set
func (l *LinkedSet[T]) Remove(t T) Set[T] {
	l.data.Del(t)
	return l
}

// Size returns the number of elements in the set
func (l *LinkedSet[T]) Size() int {
	return l.data.Size()
}

// IsEmpty returns true if the set is empty
func (l *LinkedSet[T]) IsEmpty() bool {
	return l.data.IsEmpty()
}

// Contains returns true if the set contains the element
func (l *LinkedSet[T]) Contains(t T) bool {
	_, exists := l.data.Get(t)
	return exists
}

// Clear removes all elements from the set
func (l *LinkedSet[T]) Clear() Set[T] {
	l.data.Clear()
	return l
}

// Elems returns a slice containing all elements in insertion order
func (l *LinkedSet[T]) Elems() []T {
	var result []T
	l.data.Each(func(k T, v struct{}) {
		result = append(result, k)
	})
	return result
}

// MoveToEnd moves an element to the end (useful for LRU cache operations)
func (l *LinkedSet[T]) MoveToEnd(t T) bool {
	if linkedMap, ok := l.data.(*mapx.LinkedMap[T, struct{}]); ok {
		return linkedMap.MoveToEnd(t)
	}
	return false
}

// MoveToFront moves an element to the front
func (l *LinkedSet[T]) MoveToFront(t T) bool {
	if linkedMap, ok := l.data.(*mapx.LinkedMap[T, struct{}]); ok {
		return linkedMap.MoveToFront(t)
	}
	return false
}
