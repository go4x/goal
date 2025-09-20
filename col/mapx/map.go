// Package mapx provides a generic map type with additional utility methods.
//
// This package offers a type-safe generic map implementation that extends
// Go's built-in map with convenient methods for common operations.
//
// Example:
//
//	m := mapx.New[string, int]()
//	m.Put("a", 1).Put("b", 2)
//	if val, ok := m.Get("a"); ok {
//		fmt.Printf("Value: %d\n", val)
//	}
package mapx

// M is a generic map type where K must be a comparable type (can use ==, != to compare)
// and V can be any type. This provides type safety while maintaining the flexibility
// of Go's built-in map type.
//
// Example:
//
//	stringIntMap := mapx.New[string, int]()
//	intStringMap := mapx.New[int, string]()
//	anyMap := mapx.New[string, any]()
//	m := mapx.M[string, int]{}
type M[K comparable, V any] struct {
	m map[K]V
}

// Map defines the interface for map operations.
// This interface provides a contract for map-like data structures.
type Map[K comparable, V any] interface {
	Get(k K) (V, bool)
	Put(k K, v V) Map[K, V] // Returns the concrete map type for method chaining
	Del(k K) (V, bool)
	Keys() []K
	Values() []V
	Clear() Map[K, V] // Returns the concrete map type for method chaining
	Size() int
	IsEmpty() bool
	Contains(k K) bool
	Each(fn func(K, V))
}

// New creates a new empty map with the specified key and value types.
//
// Parameters:
//   - K: The type for map keys (must be comparable)
//   - V: The type for map values (can be any type)
//
// Returns:
//   - Map[K, V]: A new empty map instance
//
// Example:
//
//	m := mapx.New[string, int]()
//	fmt.Println(m.Size()) // Output: 0
func New[K comparable, V any]() Map[K, V] {
	return &M[K, V]{m: make(map[K]V)}
}

// Put adds or updates a key-value pair in the map and returns the map for method chaining.
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
//	m := mapx.New[string, int]()
//	m.Put("a", 1).Put("b", 2) // Method chaining
func (m *M[K, V]) Put(k K, v V) Map[K, V] {
	m.m[k] = v
	return m
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
//	m := mapx.New[string, int]()
//	m.Put("a", 1)
//	val, existed := m.Del("a")
//	fmt.Printf("Deleted: %d, Existed: %t\n", val, existed) // Output: Deleted: 1, Existed: true
func (m *M[K, V]) Del(k K) (V, bool) {
	v, ok := m.m[k]
	delete(m.m, k)
	return v, ok
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
//	m := mapx.New[string, int]()
//	m.Put("a", 1)
//	if val, ok := m.Get("a"); ok {
//		fmt.Printf("Found value: %d\n", val)
//	}
func (m *M[K, V]) Get(k K) (V, bool) {
	v, ok := m.m[k]
	return v, ok
}

// Keys returns a slice containing all keys in the map.
// The order of keys is not guaranteed as Go maps are unordered.
//
// Returns:
//   - []K: A slice containing all keys in the map
//
// Example:
//
//	m := mapx.New[string, int]()
//	m.Put("a", 1).Put("b", 2).Put("c", 3)
//	keys := m.Keys()
//	fmt.Println(keys) // Output: [a b c] (order may vary)
func (m *M[K, V]) Keys() []K {
	var keys []K
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice containing all values in the map.
// The order of values is not guaranteed as Go maps are unordered.
//
// Returns:
//   - []V: A slice containing all values in the map
//
// Example:
//
//	m := mapx.New[string, int]()
//	m.Put("a", 1).Put("b", 2).Put("c", 3)
//	values := m.Values()
//	fmt.Println(values) // Output: [1 2 3] (order may vary)
func (m *M[K, V]) Values() []V {
	var vals []V
	for _, v := range m.m {
		vals = append(vals, v)
	}
	return vals
}

// Clear removes all key-value pairs from the map and returns the map for method chaining.
//
// Returns:
//   - Map[K, V]: The map itself for method chaining
//
// Example:
//
//	m := mapx.New[string, int]()
//	m.Put("a", 1).Put("b", 2)
//	m.Clear() // Map is now empty
func (m *M[K, V]) Clear() Map[K, V] {
	for k := range m.m {
		delete(m.m, k)
	}
	return m
}

// Size returns the number of key-value pairs in the map.
//
// Returns:
//   - int: The number of elements in the map
//
// Example:
//
//	m := mapx.New[string, int]()
//	m.Put("a", 1).Put("b", 2)
//	fmt.Println(m.Size()) // Output: 2
func (m *M[K, V]) Size() int {
	return len(m.m)
}

// IsEmpty returns true if the map contains no elements.
//
// Returns:
//   - bool: True if the map is empty, false otherwise
//
// Example:
//
//	m := mapx.New[string, int]()
//	fmt.Println(m.IsEmpty()) // Output: true
//	m.Put("a", 1)
//	fmt.Println(m.IsEmpty()) // Output: false
func (m *M[K, V]) IsEmpty() bool {
	return len(m.m) == 0
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
//	m := mapx.New[string, int]()
//	m.Put("a", 1)
//	fmt.Println(m.Contains("a")) // Output: true
//	fmt.Println(m.Contains("b")) // Output: false
func (m *M[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

// Each applies a function to each key-value pair in the map.
//
// Parameters:
//   - fn: The function to apply to each key-value pair
//
// Example:
//
//	m := mapx.New[string, int]()
//	m.Put("a", 1).Put("b", 2)
//	m.Each(func(k string, v int) {
//		fmt.Printf("%s: %d\n", k, v)
//	})
func (m *M[K, V]) Each(fn func(K, V)) {
	for k, v := range m.m {
		fn(k, v)
	}
}
