package mapx_test

import (
	"fmt"

	"github.com/go4x/goal/col/mapx"
)

// ExampleArrayMap demonstrates basic usage of ArrayMap
func ExampleArrayMap() {
	// Create a new sorted map
	am := mapx.NewArrayMap[string, int]()

	// Add elements in a specific order
	am.Put("third", 3)
	am.Put("first", 1)
	am.Put("second", 2)

	// Iterate in insertion order
	fmt.Println("Iteration in insertion order:")
	am.Each(func(k string, v int) {
		fmt.Printf("%s: %d\n", k, v)
	})

	// Get first and last elements
	if k, v, ok := am.First(); ok {
		fmt.Printf("First: %s = %d\n", k, v)
	}

	if k, v, ok := am.Last(); ok {
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

// ExampleArrayMap_updateExistingKey demonstrates updating existing keys
func ExampleArrayMap_updateExistingKey() {
	am := mapx.NewArrayMap[string, string]()

	// Add initial elements
	am.Put("name", "Alice")
	am.Put("age", "25")
	am.Put("city", "New York")

	fmt.Println("Initial map:")
	fmt.Println(am)

	// Update existing key - order is preserved
	am.Put("age", "26")

	fmt.Println("After updating age:")
	fmt.Println(am)

	// Output:
	// Initial map:
	// map[name:Alice age:25 city:New York]
	// After updating age:
	// map[name:Alice age:26 city:New York]
}

// ExampleArrayMap_methodChaining demonstrates method chaining
func ExampleArrayMap_methodChaining() {
	am := mapx.NewArrayMap[string, int]()

	// Chain multiple operations
	am.Put("a", 1).
		Put("b", 2).
		Put("c", 3).
		Put("b", 22) // Update existing key

	fmt.Println("After chaining operations:")
	fmt.Println(am)

	// Chain with Clear
	am.Clear().Put("new", 42)

	fmt.Println("After clear and new element:")
	fmt.Println(am)

	// Output:
	// After chaining operations:
	// map[a:1 b:22 c:3]
	// After clear and new element:
	// map[new:42]
}

// ExampleArrayMap_differentTypes demonstrates using different key and value types
func ExampleArrayMap_differentTypes() {
	// String keys, int values
	strIntMap := mapx.NewArrayMap[string, int]()
	strIntMap.Put("apple", 5).Put("banana", 3).Put("cherry", 8)

	fmt.Println("String to int map:")
	strIntMap.Each(func(k string, v int) {
		fmt.Printf("%s: %d\n", k, v)
	})

	// Int keys, string values
	intStrMap := mapx.NewArrayMap[int, string]()
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

// ExampleArrayMap_workflow demonstrates a typical workflow
func ExampleArrayMap_workflow() {
	// Create a sorted map to track user actions
	actions := mapx.NewArrayMap[string, string]()

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

	// Output:
	// User action history:
	// - login: user logged in
	// - view_profile: user viewed profile
	// - edit_settings: user edited settings
	// - logout: user logged out
	// User has edited settings
	// First action was: login (user logged in)
}
