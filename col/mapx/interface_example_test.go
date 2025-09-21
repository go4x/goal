package mapx_test

import (
	"fmt"

	"github.com/go4x/goal/col/mapx"
)

// ExampleMap demonstrates that M, ArrayMap, and LinkedMap all implement the Map interface
func ExampleMap() {
	// M, ArrayMap, and LinkedMap all implement the Map interface
	// This allows for polymorphic usage

	// Create different map implementations
	regularMap := mapx.New[string, int]()
	arrayMap := mapx.NewArrayMap[string, int]()
	linkedMap := mapx.NewLinkedMap[string, int]()

	// Store them as Map interface
	var maps []mapx.Map[string, int]
	maps = append(maps, regularMap)
	maps = append(maps, arrayMap)
	maps = append(maps, linkedMap)

	// Use them polymorphically
	for i, m := range maps {
		m.Put("key1", 10+i).Put("key2", 20+i).Put("key3", 30+i)

		fmt.Printf("Map %d:\n", i+1)
		m.Each(func(k string, v int) {
			fmt.Printf("  %s: %d\n", k, v)
		})
		fmt.Printf("  Size: %d\n", m.Size())
		fmt.Println()
	}

	// Output:
	// Map 1:
	//   key1: 10
	//   key2: 20
	//   key3: 30
	//   Size: 3
	//
	// Map 2:
	//   key1: 11
	//   key2: 21
	//   key3: 31
	//   Size: 3
	//
	// Map 3:
	//   key1: 12
	//   key2: 22
	//   key3: 32
	//   Size: 3
}

// ExampleMap_operations demonstrates common operations on Map interface
func ExampleMap_operations() {
	// Create a LinkedMap but use it as Map interface
	var m mapx.Map[string, string] = mapx.NewLinkedMap[string, string]()

	// Basic operations
	m.Put("name", "Alice")
	m.Put("age", "25")
	m.Put("city", "New York")

	// Check if key exists
	if m.Contains("name") {
		if name, ok := m.Get("name"); ok {
			fmt.Printf("Name: %s\n", name)
		}
	}

	// Get all keys and values
	keys := m.Keys()
	values := m.Values()
	fmt.Printf("Keys: %v\n", keys)
	fmt.Printf("Values: %v\n", values)

	// Iterate through the map
	fmt.Println("Iterating:")
	m.Each(func(k, v string) {
		fmt.Printf("  %s = %s\n", k, v)
	})

	// Delete a key
	if value, ok := m.Del("age"); ok {
		fmt.Printf("Deleted age: %s\n", value)
	}

	fmt.Printf("Final size: %d\n", m.Size())

	// Output:
	// Name: Alice
	// Keys: [name age city]
	// Values: [Alice 25 New York]
	// Iterating:
	//   name = Alice
	//   age = 25
	//   city = New York
	// Deleted age: 25
	// Final size: 2
}

// ExampleMap_chaining demonstrates method chaining with Map interface
func ExampleMap_chaining() {
	// Create a map using the interface
	var m mapx.Map[string, int] = mapx.NewLinkedMap[string, int]()

	// Method chaining works with the interface
	result := m.Put("a", 1).Put("b", 2).Put("c", 3)

	// The result is still a Map interface
	fmt.Printf("After chaining, size: %d\n", result.Size())

	// Clear and chain again
	result = result.Clear().Put("new", 42)
	fmt.Printf("After clear and put, size: %d\n", result.Size())

	// Iterate to show the result
	result.Each(func(k string, v int) {
		fmt.Printf("%s: %d\n", k, v)
	})

	// Output:
	// After chaining, size: 3
	// After clear and put, size: 1
	// new: 42
}

// ExampleMap_typeAssertion demonstrates type assertion to access specific methods
func ExampleMap_typeAssertion() {
	// Create a LinkedMap but use it as Map interface
	var m mapx.Map[string, int] = mapx.NewLinkedMap[string, int]()

	// Use basic Map interface methods
	m.Put("a", 1).Put("b", 2).Put("c", 3)

	// Type assert to access LinkedMap-specific methods
	if linkedMap, ok := m.(*mapx.LinkedMap[string, int]); ok {
		// Now we can use LinkedMap-specific methods
		linkedMap.MoveToEnd("a")
		fmt.Println("After moving 'a' to end:")
		linkedMap.Each(func(k string, v int) {
			fmt.Printf("%s: %d\n", k, v)
		})

		// Get first and last elements
		if k, v, ok := linkedMap.First(); ok {
			fmt.Printf("First: %s = %d\n", k, v)
		}
		if k, v, ok := linkedMap.Last(); ok {
			fmt.Printf("Last: %s = %d\n", k, v)
		}
	}

	// Output:
	// After moving 'a' to end:
	// b: 2
	// c: 3
	// a: 1
	// First: b = 2
	// Last: a = 1
}

// ExampleMap_factory demonstrates a factory pattern using Map interface
func ExampleMap_factory() {
	// Factory function that returns different Map implementations
	createMap := func(mapType string) mapx.Map[string, int] {
		switch mapType {
		case "linked":
			return mapx.NewLinkedMap[string, int]()
		case "array":
			return mapx.NewArrayMap[string, int]()
		default:
			return mapx.New[string, int]()
		}
	}

	// Create different map types
	arrayMap := createMap("array")
	linkedMap := createMap("linked")

	// Use them identically through the interface
	// Note: ArrayMap and LinkedMap guarantee insertion order, regular Map does not
	orderedMaps := []mapx.Map[string, int]{arrayMap, linkedMap}
	for _, m := range orderedMaps {
		m.Put("x", 1).Put("y", 2).Put("z", 3)

		// All operations work the same way
		fmt.Printf("Size: %d\n", m.Size())
		fmt.Printf("Contains 'y': %t\n", m.Contains("y"))

		if val, ok := m.Get("y"); ok {
			fmt.Printf("Value of 'y': %d\n", val)
		}

		fmt.Println("Keys:", m.Keys())
		fmt.Println()
	}

	// Output:
	// Size: 3
	// Contains 'y': true
	// Value of 'y': 2
	// Keys: [x y z]
	//
	// Size: 3
	// Contains 'y': true
	// Value of 'y': 2
	// Keys: [x y z]
}
