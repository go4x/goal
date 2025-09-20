package mapx

import (
	"fmt"
)

// ArrayMap is an ordered map that maintains insertion order.
// It uses two slices to store keys and values, ensuring that iteration
// always follows the insertion order.
//
// Implementation Details:
// The implementation uses two parallel slices:
// - keys: stores the keys in insertion order
// - vals: stores the corresponding values
//
// Time Complexity:
//   - Get: O(n) - requires linear search through keys
//   - Put: O(n) if key exists (linear search), O(1) if new key (append)
//   - Del: O(n) - requires linear search and slice manipulation
//   - Keys/Values/Each: O(n) - direct slice iteration
//   - Memory: O(n) - two slices with minimal overhead
//
// Advantages:
//   - Simple implementation with minimal memory overhead
//   - Deterministic iteration order
//   - No pointer chasing (cache-friendly for small datasets)
//   - Easy to understand and debug
//   - Memory-efficient (no extra pointers or metadata)
//
// Disadvantages:
//   - O(n) lookup and deletion operations (not suitable for large datasets)
//   - Poor performance for frequent lookups or deletions
//   - Linear search becomes expensive as map size grows
//   - No O(1) operations except appending new keys
//
// Use Cases:
//   - Small to medium datasets (< 1000 elements)
//   - Infrequent lookups, mostly insertion and iteration
//   - Configuration maps where order matters
//   - Prototyping or simple applications
//   - Memory-constrained environments
//   - When you need predictable iteration order
//
// When NOT to use:
//   - Large datasets (> 1000 elements)
//   - Frequent lookups or deletions
//   - Performance-critical applications
//   - When you need O(1) operations
//   - High-frequency access patterns
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("first", 1)
//	am.Put("second", 2)
//	am.Put("third", 3)
//
//	// Iteration will always be in insertion order
//	am.Each(func(k string, v int) {
//		fmt.Printf("%s: %d\n", k, v) // first: 1, second: 2, third: 3
//	})
type ArrayMap[K comparable, V any] struct {
	keys []K
	vals []V
}

// NewArrayMap creates a new empty sorted map.
//
// Returns:
//   - *ArrayMap[K, V]: A new empty sorted map instance
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	fmt.Println(am.Size()) // Output: 0
func NewArrayMap[K comparable, V any]() *ArrayMap[K, V] {
	return &ArrayMap[K, V]{
		keys: make([]K, 0),
		vals: make([]V, 0),
	}
}

// Put adds or updates a key-value pair in the map.
// If the key already exists, its value is updated and the insertion order is preserved.
// If the key is new, it's added to the end of the map.
//
// Parameters:
//   - k: The key to add or update
//   - v: The value to associate with the key
//
// Returns:
//   - Map[K, V]: The map itself for method chaining
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("a", 1).Put("b", 2) // Method chaining
func (am *ArrayMap[K, V]) Put(k K, v V) Map[K, V] {
	if idx := am.findIndex(k); idx >= 0 {
		// Key exists, update value
		am.vals[idx] = v
	} else {
		// New key, append to end
		am.keys = append(am.keys, k)
		am.vals = append(am.vals, v)
	}
	return am
}

// Get retrieves the value associated with a key and indicates whether the key exists.
//
// Parameters:
//   - k: The key to look up
//
// Returns:
//   - V: The value associated with the key (zero value if key doesn't exist)
//   - bool: True if the key exists, false otherwise
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("a", 1)
//	if val, ok := am.Get("a"); ok {
//		fmt.Printf("Found value: %d\n", val)
//	}
func (am *ArrayMap[K, V]) Get(k K) (V, bool) {
	if idx := am.findIndex(k); idx >= 0 {
		return am.vals[idx], true
	}
	var zero V
	return zero, false
}

// Del removes a key from the map and returns the associated value and whether the key existed.
//
// Parameters:
//   - k: The key to remove
//
// Returns:
//   - V: The value associated with the key (zero value if key didn't exist)
//   - bool: True if the key existed, false otherwise
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("a", 1)
//	val, existed := am.Del("a")
//	fmt.Printf("Deleted: %d, Existed: %t\n", val, existed) // Output: Deleted: 1, Existed: true
func (am *ArrayMap[K, V]) Del(k K) (V, bool) {
	if idx := am.findIndex(k); idx >= 0 {
		val := am.vals[idx]

		// Remove element by slicing
		am.keys = append(am.keys[:idx], am.keys[idx+1:]...)
		am.vals = append(am.vals[:idx], am.vals[idx+1:]...)

		return val, true
	}
	var zero V
	return zero, false
}

// Keys returns a slice containing all keys in insertion order.
//
// Returns:
//   - []K: A slice containing all keys in insertion order
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("first", 1).Put("second", 2)
//	keys := am.Keys()
//	fmt.Println(keys) // Output: [first second]
func (am *ArrayMap[K, V]) Keys() []K {
	result := make([]K, len(am.keys))
	copy(result, am.keys)
	return result
}

// Values returns a slice containing all values in insertion order.
//
// Returns:
//   - []V: A slice containing all values in insertion order
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("first", 1).Put("second", 2)
//	values := am.Values()
//	fmt.Println(values) // Output: [1 2]
func (am *ArrayMap[K, V]) Values() []V {
	result := make([]V, len(am.vals))
	copy(result, am.vals)
	return result
}

// Clear removes all key-value pairs from the map.
//
// Returns:
//   - Map[K, V]: The map itself for method chaining
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("a", 1).Put("b", 2)
//	am.Clear() // Map is now empty
func (am *ArrayMap[K, V]) Clear() Map[K, V] {
	am.keys = am.keys[:0]
	am.vals = am.vals[:0]
	return am
}

// Size returns the number of key-value pairs in the map.
//
// Returns:
//   - int: The number of elements in the map
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("a", 1).Put("b", 2)
//	fmt.Println(am.Size()) // Output: 2
func (am *ArrayMap[K, V]) Size() int {
	return len(am.keys)
}

// IsEmpty returns true if the map contains no elements.
//
// Returns:
//   - bool: True if the map is empty, false otherwise
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	fmt.Println(am.IsEmpty()) // Output: true
//	am.Put("a", 1)
//	fmt.Println(am.IsEmpty()) // Output: false
func (am *ArrayMap[K, V]) IsEmpty() bool {
	return len(am.keys) == 0
}

// Contains checks if the map contains the specified key.
//
// Parameters:
//   - k: The key to check for
//
// Returns:
//   - bool: True if the key exists, false otherwise
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("a", 1)
//	fmt.Println(am.Contains("a")) // Output: true
//	fmt.Println(am.Contains("b")) // Output: false
func (am *ArrayMap[K, V]) Contains(k K) bool {
	return am.findIndex(k) >= 0
}

// Each applies a function to each key-value pair in insertion order.
//
// Parameters:
//   - fn: The function to apply to each key-value pair
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("first", 1).Put("second", 2)
//	am.Each(func(k string, v int) {
//		fmt.Printf("%s: %d\n", k, v) // first: 1, second: 2
//	})
func (am *ArrayMap[K, V]) Each(fn func(K, V)) {
	for i := range am.keys {
		fn(am.keys[i], am.vals[i])
	}
}

// First returns the first key-value pair in insertion order.
//
// Returns:
//   - K: The first key (zero value if map is empty)
//   - V: The first value (zero value if map is empty)
//   - bool: True if the map is not empty, false otherwise
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("first", 1).Put("second", 2)
//	if k, v, ok := am.First(); ok {
//		fmt.Printf("First: %s = %d\n", k, v) // Output: First: first = 1
//	}
func (am *ArrayMap[K, V]) First() (K, V, bool) {
	if len(am.keys) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return am.keys[0], am.vals[0], true
}

// Last returns the last key-value pair in insertion order.
//
// Returns:
//   - K: The last key (zero value if map is empty)
//   - V: The last value (zero value if map is empty)
//   - bool: True if the map is not empty, false otherwise
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("first", 1).Put("second", 2)
//	if k, v, ok := am.Last(); ok {
//		fmt.Printf("Last: %s = %d\n", k, v) // Output: Last: second = 2
//	}
func (am *ArrayMap[K, V]) Last() (K, V, bool) {
	if len(am.keys) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	lastIdx := len(am.keys) - 1
	return am.keys[lastIdx], am.vals[lastIdx], true
}

// String returns a string representation of the map.
//
// Returns:
//   - string: A string representation showing all key-value pairs in insertion order
//
// Example:
//
//	am := NewArrayMap[string, int]()
//	am.Put("a", 1).Put("b", 2)
//	fmt.Println(am) // Output: map[a:1 b:2]
func (am *ArrayMap[K, V]) String() string {
	if len(am.keys) == 0 {
		return "map[]"
	}

	var result = "map["
	for i := range am.keys {
		if i > 0 {
			result += " "
		}
		result += fmt.Sprintf("%v:%v", am.keys[i], am.vals[i])
	}
	result += "]"
	return result
}

// findIndex finds the index of a key in the keys slice.
// Returns -1 if the key is not found.
func (am *ArrayMap[K, V]) findIndex(k K) int {
	for i, key := range am.keys {
		if key == k {
			return i
		}
	}
	return -1
}
