package set

import (
	"fmt"
	"testing"

	"github.com/gophero/got"
)

func TestLinkedSet(t *testing.T) {
	logger := got.New(t, "LinkedSet")

	// Test ordered insertion
	set := NewLinkedSet[int]()
	set.Add(1).Add(2).Add(3).Add(1) // Duplicate

	elems := set.Elems()
	logger.Require(len(elems) == 3, "should have 3 elements")
	logger.Require(elems[0] == 1, "first element should be 1")
	logger.Require(elems[1] == 2, "second element should be 2")
	logger.Require(elems[2] == 3, "third element should be 3")

	// Test move operations
	linkedSet := set.(*LinkedSet[int])
	linkedSet.MoveToEnd(1)
	elems = set.Elems()
	logger.Require(elems[2] == 1, "1 should be moved to end")

	linkedSet.MoveToFront(3)
	elems = set.Elems()
	logger.Require(elems[0] == 3, "3 should be moved to front")
}

func TestLinkedSetLRUCache(t *testing.T) {
	logger := got.New(t, "LinkedSetLRUCache")

	// Simulate LRU cache behavior
	cache := NewLinkedSet[string]()

	// Add items
	cache.Add("item1").Add("item2").Add("item3")

	// Access item1 (move to end - most recently used)
	linkedSet := cache.(*LinkedSet[string])
	linkedSet.MoveToEnd("item1")

	elems := cache.Elems()
	logger.Require(elems[2] == "item1", "item1 should be at end after access")

	// Add new item (should be added at end)
	cache.Add("item4")
	elems = cache.Elems()
	logger.Require(elems[3] == "item4", "item4 should be at end")

	// Access item2 (move to end)
	linkedSet.MoveToEnd("item2")
	elems = cache.Elems()
	logger.Require(elems[3] == "item2", "item2 should be at end after access")
}

func TestLinkedSetMoveOperations(t *testing.T) {
	logger := got.New(t, "LinkedSetMoveOperations")

	set := NewLinkedSet[int]()
	set.Add(1).Add(2).Add(3).Add(4).Add(5)
	linkedSet := set.(*LinkedSet[int])

	// Test MoveToFront
	linkedSet.MoveToFront(3)
	elems := set.Elems()
	logger.Require(elems[0] == 3, "3 should be moved to front")

	// Test MoveToEnd
	linkedSet.MoveToEnd(1)
	elems = set.Elems()
	logger.Require(elems[4] == 1, "1 should be moved to end")

	// Test moving non-existent element
	result := linkedSet.MoveToFront(999)
	logger.Require(!result, "moving non-existent element should return false")
}

func TestLinkedSetLargeDataset(t *testing.T) {
	logger := got.New(t, "LinkedSetLargeDataset")

	// Test with larger dataset to verify O(1) performance
	set := NewLinkedSet[int]()

	// Add many elements
	for i := 0; i < 1000; i++ {
		set.Add(i)
	}

	logger.Require(set.Size() == 1000, "should have 1000 elements")

	// Test contains (should be O(1))
	logger.Require(set.Contains(500), "should contain element 500")
	logger.Require(!set.Contains(1000), "should not contain element 1000")

	// Test removal (should be O(1))
	set.Remove(500)
	logger.Require(!set.Contains(500), "should not contain element 500 after removal")
	logger.Require(set.Size() == 999, "should have 999 elements after removal")
}

func TestLinkedSetInterfaceCompliance(t *testing.T) {
	logger := got.New(t, "LinkedSetInterfaceCompliance")

	// Test that LinkedSet implements Set interface
	var _ Set[int] = NewLinkedSet[int]()
	var _ Set[string] = NewLinkedSet[string]()

	// Test polymorphic usage
	linkedSet := NewLinkedSet[int]()
	var set Set[int] = linkedSet

	set.Add(1).Add(2).Add(3)
	logger.Require(set.Size() == 3, "polymorphic usage should work")
	logger.Require(set.Contains(2), "polymorphic contains should work")
}

// ExampleNewLinkedSet demonstrates LinkedSet usage
func ExampleNewLinkedSet() {
	// Create a LinkedSet (maintains insertion order with O(1) operations)
	linkedSet := NewLinkedSet[int]()

	// Add elements
	linkedSet.Add(1).Add(2).Add(3).Add(1)

	// Get all elements (in insertion order)
	elems := linkedSet.Elems()
	fmt.Println("Elements in order:", elems)

	// Move element to end (useful for LRU cache)
	linkedSetImpl := linkedSet.(*LinkedSet[int])
	linkedSetImpl.MoveToEnd(1)
	fmt.Println("After moving 1 to end:", linkedSet.Elems())

	// Output:
	// Elements in order: [1 2 3]
	// After moving 1 to end: [2 3 1]
}
