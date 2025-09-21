// Package set provides generic set implementations that maintain unique elements.
// It offers set semantics similar to Java's Set interface with multiple implementation
// options optimized for different use cases.
//
// This package offers three set implementations:
//
// 1. HashSet:
//   - Uses hash map for O(1) operations
//   - No insertion order guarantee
//   - Best for general-purpose set operations
//
// 2. ArraySet:
//   - Uses array-based map for ordered elements
//   - Maintains insertion order
//   - O(n) operations, good for small datasets
//
// 3. LinkedSet:
//   - Uses linked map for O(1) operations with order
//   - Maintains insertion order
//   - Best for large datasets requiring order
//
// All implementations implement the Set[T] interface for polymorphic usage.
//
// Quick Decision Guide:
// - Need O(1) performance and don't care about order? → Use NewHashSet()
// - Small dataset (<1000) and need order? → Use NewArraySet()
// - Large dataset and need order? → Use NewLinkedSet()
// - Building LRU cache? → Use NewLinkedSet()
//
// Example:
//
//	// Hash set (fastest, no order)
//	hashSet := set.NewHashSet[string]()
//	hashSet.Add("a").Add("b").Add("a") // "a" is only added once
//
//	// Array set (ordered, good for small datasets)
//	arraySet := set.NewArraySet[string]()
//	arraySet.Add("a").Add("b").Add("a")
//
//	// Linked set (ordered, good for large datasets)
//	linkedSet := set.NewLinkedSet[string]()
//	linkedSet.Add("a").Add("b").Add("a")
package set

// Set is a generic set interface that maintains unique elements.
type Set[T any] interface {
	// Add adds an element to the set (no duplicates)
	Add(t T) Set[T]

	// Remove removes an element from the set
	Remove(t T) Set[T]

	// Size returns the number of elements in the set
	Size() int

	// IsEmpty returns true if the set is empty
	IsEmpty() bool

	// Contains returns true if the set contains the element
	Contains(t T) bool

	// Clear removes all elements from the set
	Clear() Set[T]

	// Elems returns a slice containing all elements in the set
	Elems() []T
}

// Convenience constructors for different set types

// New creates a HashSet by default (most common use case)
func New[T comparable]() Set[T] {
	return NewHashSet[T]()
}

// NewHashSet creates a new HashSet with O(1) operations
func NewHashSet[T comparable]() Set[T] {
	return newHashSet[T]()
}

// NewArraySet creates a new ArraySet that maintains insertion order
func NewArraySet[T comparable]() Set[T] {
	return newArraySet[T]()
}

// NewLinkedSet creates a new LinkedSet with O(1) operations and insertion order
func NewLinkedSet[T comparable]() Set[T] {
	return newLinkedSet[T]()
}
