package mapx_test

import (
	"fmt"

	"github.com/go4x/goal/col/mapx"
)

// ExampleLinkedMap demonstrates basic usage of LinkedMap
func ExampleLinkedMap() {
	// Create a new linked map
	lm := mapx.NewLinkedMap[string, int]()

	// Add elements in a specific order
	lm.Put("third", 3)
	lm.Put("first", 1)
	lm.Put("second", 2)

	// Iterate in insertion order
	fmt.Println("Iteration in insertion order:")
	lm.Each(func(k string, v int) {
		fmt.Printf("%s: %d\n", k, v)
	})

	// Get first and last elements
	if k, v, ok := lm.First(); ok {
		fmt.Printf("First: %s = %d\n", k, v)
	}

	if k, v, ok := lm.Last(); ok {
		fmt.Printf("Last: %s = %d\n", k, v)
	}

	// Output:
	// Iteration in insertion order:
	// third: 3
	// first: 1
	// second: 2
	// First: third = 3
	// Last: second = 2
}

// ExampleLinkedMap_updateExistingKey demonstrates updating existing keys
func ExampleLinkedMap_updateExistingKey() {
	lm := mapx.NewLinkedMap[string, string]()

	// Add initial elements
	lm.Put("name", "Alice")
	lm.Put("age", "25")
	lm.Put("city", "New York")

	fmt.Println("Initial map:")
	fmt.Println(lm)

	// Update existing key - order is preserved
	lm.Put("age", "26")

	fmt.Println("After updating age:")
	fmt.Println(lm)

	// Output:
	// Initial map:
	// map[name:Alice age:25 city:New York]
	// After updating age:
	// map[name:Alice age:26 city:New York]
}

// ExampleLinkedMap_methodChaining demonstrates method chaining
func ExampleLinkedMap_methodChaining() {
	lm := mapx.NewLinkedMap[string, int]()

	// Chain multiple operations
	lm.Put("a", 1).
		Put("b", 2).
		Put("c", 3).
		Put("b", 22) // Update existing key

	fmt.Println("After chaining operations:")
	fmt.Println(lm)

	// Chain with Clear
	lm.Clear().Put("new", 42)

	fmt.Println("After clear and new element:")
	fmt.Println(lm)

	// Output:
	// After chaining operations:
	// map[a:1 b:22 c:3]
	// After clear and new element:
	// map[new:42]
}

// ExampleLinkedMap_moveOperations demonstrates MoveToEnd and MoveToFront operations
func ExampleLinkedMap_moveOperations() {
	lm := mapx.NewLinkedMap[string, int]()

	// Add elements
	lm.Put("a", 1).Put("b", 2).Put("c", 3).Put("d", 4)

	fmt.Println("Initial order:")
	fmt.Println(lm)

	// Move 'b' to the end (makes it newest)
	lm.MoveToEnd("b")
	fmt.Println("After moving 'b' to end:")
	fmt.Println(lm)

	// Move 'c' to the front (makes it oldest)
	lm.MoveToFront("c")
	fmt.Println("After moving 'c' to front:")
	fmt.Println(lm)

	// Output:
	// Initial order:
	// map[a:1 b:2 c:3 d:4]
	// After moving 'b' to end:
	// map[a:1 c:3 d:4 b:2]
	// After moving 'c' to front:
	// map[c:3 a:1 d:4 b:2]
}

// ExampleLinkedMap_lruCache demonstrates using LinkedMap as an LRU cache
func ExampleLinkedMap_lruCache() {
	// Simple LRU cache implementation using LinkedMap
	cache := mapx.NewLinkedMap[string, string]()

	// Cache operations
	setCache := func(key, value string) {
		// If key exists, move it to end (mark as recently used)
		if cache.Contains(key) {
			cache.MoveToEnd(key)
			cache.Put(key, value) // Update value
		} else {
			cache.Put(key, value) // Add new key (automatically goes to end)
		}
	}

	getCache := func(key string) (string, bool) {
		if value, exists := cache.Get(key); exists {
			// Move to end to mark as recently used
			cache.MoveToEnd(key)
			return value, true
		}
		return "", false
	}

	// Simulate cache usage
	setCache("user1", "Alice")
	setCache("user2", "Bob")
	setCache("user3", "Charlie")

	fmt.Println("Cache after adding users:")
	fmt.Println(cache)

	// Access user1 (should move it to end)
	if value, ok := getCache("user1"); ok {
		fmt.Printf("Retrieved user1: %s\n", value)
	}

	fmt.Println("Cache after accessing user1:")
	fmt.Println(cache)

	// Output:
	// Cache after adding users:
	// map[user1:Alice user2:Bob user3:Charlie]
	// Retrieved user1: Alice
	// Cache after accessing user1:
	// map[user2:Bob user3:Charlie user1:Alice]
}

// ExampleLinkedMap_differentTypes demonstrates using different key and value types
func ExampleLinkedMap_differentTypes() {
	// String keys, int values
	strIntMap := mapx.NewLinkedMap[string, int]()
	strIntMap.Put("apple", 5).Put("banana", 3).Put("cherry", 8)

	fmt.Println("String to int map:")
	strIntMap.Each(func(k string, v int) {
		fmt.Printf("%s: %d\n", k, v)
	})

	// Int keys, string values
	intStrMap := mapx.NewLinkedMap[int, string]()
	intStrMap.Put(3, "three").Put(1, "one").Put(2, "two")

	fmt.Println("Int to string map:")
	intStrMap.Each(func(k int, v string) {
		fmt.Printf("%d: %s\n", k, v)
	})

	// Output:
	// String to int map:
	// apple: 5
	// banana: 3
	// cherry: 8
	// Int to string map:
	// 3: three
	// 1: one
	// 2: two
}

// ExampleLinkedMap_workflow demonstrates a typical workflow
func ExampleLinkedMap_workflow() {
	// Create a linked map to track user actions
	actions := mapx.NewLinkedMap[string, string]()

	// Simulate user actions in order
	actions.Put("login", "user logged in")
	actions.Put("view_profile", "user viewed profile")
	actions.Put("edit_settings", "user edited settings")
	actions.Put("logout", "user logged out")

	fmt.Println("User action history:")
	actions.Each(func(action, description string) {
		fmt.Printf("- %s: %s\n", action, description)
	})

	// Check if user performed a specific action
	if actions.Contains("edit_settings") {
		fmt.Println("User has edited settings")
	}

	// Get the first action
	if action, desc, ok := actions.First(); ok {
		fmt.Printf("First action was: %s (%s)\n", action, desc)
	}

	// Get the last action
	if action, desc, ok := actions.Last(); ok {
		fmt.Printf("Last action was: %s (%s)\n", action, desc)
	}

	// Output:
	// User action history:
	// - login: user logged in
	// - view_profile: user viewed profile
	// - edit_settings: user edited settings
	// - logout: user logged out
	// User has edited settings
	// First action was: login (user logged in)
	// Last action was: logout (user logged out)
}
