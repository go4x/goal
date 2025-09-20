package mapx

import (
	"fmt"
)

// node represents a node in the doubly linked list
type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

// LinkedMap is an ordered map that maintains insertion order using a doubly linked list
// combined with a hash map for O(1) average-case performance.
//
// Implementation Details:
// The implementation uses a hybrid approach:
// - A doubly linked list maintains insertion order
// - A hash map provides O(1) access to nodes
// - Each node contains key, value, and pointers to prev/next nodes
//
// Time Complexity:
//   - Get: O(1) average-case (hash map lookup)
//   - Put: O(1) average-case (hash map + linked list insertion)
//   - Del: O(1) average-case (hash map + linked list deletion)
//   - Keys/Values/Each: O(n) - linked list traversal
//   - MoveToEnd/MoveToFront: O(1) - direct pointer manipulation
//   - Memory: O(n) - hash map + linked list nodes with pointers
//
// Advantages:
//   - O(1) average-case performance for all basic operations
//   - Deterministic iteration order
//   - Efficient for large datasets
//   - Supports advanced operations (MoveToEnd, MoveToFront)
//   - Ideal for LRU cache implementations
//   - Good performance for frequent lookups and deletions
//
// Disadvantages:
//   - Higher memory overhead (pointers + hash map)
//   - More complex implementation
//   - Potential for pointer chasing (cache misses)
//   - Hash map overhead for small datasets
//   - More memory allocations per operation
//
// Use Cases:
//   - Large datasets (> 1000 elements)
//   - Frequent lookups, insertions, or deletions
//   - LRU cache implementations
//   - Performance-critical applications
//   - When you need O(1) operations
//   - High-frequency access patterns
//   - When order matters and performance is critical
//
// When NOT to use:
//   - Small datasets (< 100 elements) - ArrayMap might be better
//   - Memory-constrained environments
//   - Simple applications where ArrayMap is sufficient
//   - When you rarely access elements (mostly iteration)
//
// Special Features:
//   - MoveToEnd: Move a key to the end (useful for LRU caches)
//   - MoveToFront: Move a key to the front
//   - First/Last: Get first/last key-value pairs
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("first", 1)
//	lm.Put("second", 2)
//	lm.Put("third", 3)
//
//	// Iteration will always be in insertion order
//	lm.Each(func(k string, v int) {
//		fmt.Printf("%s: %d\n", k, v) // first: 1, second: 2, third: 3
//	})
//
//	// LRU cache example
//	lm.MoveToEnd("first") // Move "first" to end (most recently used)
type LinkedMap[K comparable, V any] struct {
	head *node[K, V]       // head of the linked list (oldest element)
	tail *node[K, V]       // tail of the linked list (newest element)
	data map[K]*node[K, V] // map for O(1) access to nodes
	size int
}

// NewLinkedMap creates a new empty linked map.
//
// Returns:
//   - *LinkedMap[K, V]: A new empty linked map instance
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	fmt.Println(lm.Size()) // Output: 0
func NewLinkedMap[K comparable, V any]() *LinkedMap[K, V] {
	return &LinkedMap[K, V]{
		data: make(map[K]*node[K, V]),
		size: 0,
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
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1).Put("b", 2) // Method chaining
func (lm *LinkedMap[K, V]) Put(k K, v V) Map[K, V] {
	if existingNode, exists := lm.data[k]; exists {
		// Key exists, update value
		existingNode.value = v
	} else {
		// New key, create new node and add to tail
		newNode := &node[K, V]{
			key:   k,
			value: v,
			prev:  lm.tail,
			next:  nil,
		}

		if lm.tail != nil {
			lm.tail.next = newNode
		}
		lm.tail = newNode

		if lm.head == nil {
			lm.head = newNode
		}

		lm.data[k] = newNode
		lm.size++
	}
	return lm
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
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1)
//	if val, ok := lm.Get("a"); ok {
//		fmt.Printf("Found value: %d\n", val)
//	}
func (lm *LinkedMap[K, V]) Get(k K) (V, bool) {
	if node, exists := lm.data[k]; exists {
		return node.value, true
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
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1)
//	val, existed := lm.Del("a")
//	fmt.Printf("Deleted: %d, Existed: %t\n", val, existed) // Output: Deleted: 1, Existed: true
func (lm *LinkedMap[K, V]) Del(k K) (V, bool) {
	node, exists := lm.data[k]
	if !exists {
		var zero V
		return zero, false
	}

	val := node.value

	// Remove from linked list
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// This is the head node
		lm.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// This is the tail node
		lm.tail = node.prev
	}

	// Remove from map
	delete(lm.data, k)
	lm.size--

	return val, true
}

// Keys returns a slice containing all keys in insertion order.
//
// Returns:
//   - []K: A slice containing all keys in insertion order
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("first", 1).Put("second", 2)
//	keys := lm.Keys()
//	fmt.Println(keys) // Output: [first second]
func (lm *LinkedMap[K, V]) Keys() []K {
	keys := make([]K, 0, lm.size)
	for current := lm.head; current != nil; current = current.next {
		keys = append(keys, current.key)
	}
	return keys
}

// Values returns a slice containing all values in insertion order.
//
// Returns:
//   - []V: A slice containing all values in insertion order
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("first", 1).Put("second", 2)
//	values := lm.Values()
//	fmt.Println(values) // Output: [1 2]
func (lm *LinkedMap[K, V]) Values() []V {
	values := make([]V, 0, lm.size)
	for current := lm.head; current != nil; current = current.next {
		values = append(values, current.value)
	}
	return values
}

// Clear removes all key-value pairs from the map.
//
// Returns:
//   - Map[K, V]: The map itself for method chaining
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1).Put("b", 2)
//	lm.Clear() // Map is now empty
func (lm *LinkedMap[K, V]) Clear() Map[K, V] {
	lm.head = nil
	lm.tail = nil
	lm.data = make(map[K]*node[K, V])
	lm.size = 0
	return lm
}

// Size returns the number of key-value pairs in the map.
//
// Returns:
//   - int: The number of elements in the map
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1).Put("b", 2)
//	fmt.Println(lm.Size()) // Output: 2
func (lm *LinkedMap[K, V]) Size() int {
	return lm.size
}

// IsEmpty returns true if the map contains no elements.
//
// Returns:
//   - bool: True if the map is empty, false otherwise
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	fmt.Println(lm.IsEmpty()) // Output: true
//	lm.Put("a", 1)
//	fmt.Println(lm.IsEmpty()) // Output: false
func (lm *LinkedMap[K, V]) IsEmpty() bool {
	return lm.size == 0
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
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1)
//	fmt.Println(lm.Contains("a")) // Output: true
//	fmt.Println(lm.Contains("b")) // Output: false
func (lm *LinkedMap[K, V]) Contains(k K) bool {
	_, exists := lm.data[k]
	return exists
}

// Each applies a function to each key-value pair in insertion order.
//
// Parameters:
//   - fn: The function to apply to each key-value pair
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("first", 1).Put("second", 2)
//	lm.Each(func(k string, v int) {
//		fmt.Printf("%s: %d\n", k, v) // first: 1, second: 2
//	})
func (lm *LinkedMap[K, V]) Each(fn func(K, V)) {
	for current := lm.head; current != nil; current = current.next {
		fn(current.key, current.value)
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
//	lm := NewLinkedMap[string, int]()
//	lm.Put("first", 1).Put("second", 2)
//	if k, v, ok := lm.First(); ok {
//		fmt.Printf("First: %s = %d\n", k, v) // Output: First: first = 1
//	}
func (lm *LinkedMap[K, V]) First() (K, V, bool) {
	if lm.head == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return lm.head.key, lm.head.value, true
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
//	lm := NewLinkedMap[string, int]()
//	lm.Put("first", 1).Put("second", 2)
//	if k, v, ok := lm.Last(); ok {
//		fmt.Printf("Last: %s = %d\n", k, v) // Output: Last: second = 2
//	}
func (lm *LinkedMap[K, V]) Last() (K, V, bool) {
	if lm.tail == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return lm.tail.key, lm.tail.value, true
}

// String returns a string representation of the map.
//
// Returns:
//   - string: A string representation showing all key-value pairs in insertion order
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1).Put("b", 2)
//	fmt.Println(lm) // Output: map[a:1 b:2]
func (lm *LinkedMap[K, V]) String() string {
	if lm.size == 0 {
		return "map[]"
	}

	var result = "map["
	count := 0
	for current := lm.head; current != nil; current = current.next {
		if count > 0 {
			result += " "
		}
		result += fmt.Sprintf("%v:%v", current.key, current.value)
		count++
	}
	result += "]"
	return result
}

// MoveToEnd moves an existing key to the end of the map (makes it the newest).
// This is useful for implementing LRU cache behavior.
//
// Parameters:
//   - k: The key to move to the end
//
// Returns:
//   - bool: True if the key existed and was moved, false otherwise
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1).Put("b", 2).Put("c", 3)
//	lm.MoveToEnd("a") // "a" is now the newest element
//	// Order is now: b, c, a
func (lm *LinkedMap[K, V]) MoveToEnd(k K) bool {
	node, exists := lm.data[k]
	if !exists || node == lm.tail {
		return exists
	}

	// Remove from current position
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lm.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	// Add to end
	node.prev = lm.tail
	node.next = nil
	lm.tail.next = node
	lm.tail = node

	return true
}

// MoveToFront moves an existing key to the front of the map (makes it the oldest).
//
// Parameters:
//   - k: The key to move to the front
//
// Returns:
//   - bool: True if the key existed and was moved, false otherwise
//
// Example:
//
//	lm := NewLinkedMap[string, int]()
//	lm.Put("a", 1).Put("b", 2).Put("c", 3)
//	lm.MoveToFront("c") // "c" is now the oldest element
//	// Order is now: c, a, b
func (lm *LinkedMap[K, V]) MoveToFront(k K) bool {
	node, exists := lm.data[k]
	if !exists || node == lm.head {
		return exists
	}

	// Remove from current position
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lm.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lm.tail = node.prev
	}

	// Add to front
	node.prev = nil
	node.next = lm.head
	lm.head.prev = node
	lm.head = node

	return true
}
