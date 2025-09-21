package set

import (
	"fmt"
	"testing"

	"github.com/go4x/got"
)

func TestHashSet(t *testing.T) {
	logger := got.New(t, "HashSet")

	// Test basic operations
	set := NewHashSet[int]()
	logger.Require(set.IsEmpty(), "new set should be empty")
	logger.Require(set.Size() == 0, "new set size should be 0")

	// Test adding elements
	set.Add(1).Add(2).Add(1) // Duplicate should be ignored
	logger.Require(!set.IsEmpty(), "set should not be empty after adding")
	logger.Require(set.Size() == 2, "set should have 2 elements after adding 1,2,1")
	logger.Require(set.Contains(1), "set should contain 1")
	logger.Require(set.Contains(2), "set should contain 2")
	logger.Require(!set.Contains(3), "set should not contain 3")

	// Test removing elements
	set.Remove(1)
	logger.Require(set.Size() == 1, "set should have 1 element after removing 1")
	logger.Require(!set.Contains(1), "set should not contain 1 after removal")
	logger.Require(set.Contains(2), "set should still contain 2")

	// Test clear
	set.Clear()
	logger.Require(set.IsEmpty(), "set should be empty after clear")
	logger.Require(set.Size() == 0, "set size should be 0 after clear")
}

func TestHashSetDuplicateHandling(t *testing.T) {
	logger := got.New(t, "HashSetDuplicateHandling")

	set := NewHashSet[string]()
	set.Add("a").Add("b").Add("a").Add("c").Add("b")

	logger.Require(set.Size() == 3, "should have 3 unique elements")
	logger.Require(set.Contains("a"), "should contain 'a'")
	logger.Require(set.Contains("b"), "should contain 'b'")
	logger.Require(set.Contains("c"), "should contain 'c'")
}

func TestHashSetWithDifferentTypes(t *testing.T) {
	logger := got.New(t, "HashSetWithDifferentTypes")

	// Test with strings
	stringSet := NewHashSet[string]()
	stringSet.Add("hello").Add("world").Add("hello")
	logger.Require(stringSet.Size() == 2, "string set should have 2 elements")

	// Test with floats
	floatSet := NewHashSet[float64]()
	floatSet.Add(1.1).Add(2.2).Add(1.1).Add(3.3)
	logger.Require(floatSet.Size() == 3, "float set should have 3 elements")

	// Test with structs
	type Point struct {
		X, Y int
	}
	pointSet := NewHashSet[Point]()
	pointSet.Add(Point{1, 2}).Add(Point{3, 4}).Add(Point{1, 2})
	logger.Require(pointSet.Size() == 2, "point set should have 2 elements")
}

func TestHashSetInterfaceCompliance(t *testing.T) {
	logger := got.New(t, "HashSetInterfaceCompliance")

	// Test that HashSet implements Set interface
	var _ Set[int] = NewHashSet[int]()
	var _ Set[string] = NewHashSet[string]()
	var _ Set[float64] = NewHashSet[float64]()

	// Test polymorphic usage
	hashSet := NewHashSet[int]()
	var set Set[int] = hashSet

	set.Add(1).Add(2).Add(3)
	logger.Require(set.Size() == 3, "polymorphic usage should work")
	logger.Require(set.Contains(2), "polymorphic contains should work")
}

// ExampleNewHashSet demonstrates HashSet usage
func ExampleNewHashSet() {
	// Create a HashSet (no order guarantee)
	hashSet := NewHashSet[int]()

	// Add elements
	hashSet.Add(3).Add(1).Add(4).Add(1).Add(5)

	// Check size and contains
	fmt.Println("Size:", hashSet.Size())
	fmt.Println("Contains 1:", hashSet.Contains(1))
	fmt.Println("Contains 6:", hashSet.Contains(6))

	// Remove an element
	hashSet.Remove(1)
	fmt.Println("Size after removing 1:", hashSet.Size())
	fmt.Println("Contains 1 after removal:", hashSet.Contains(1))

	// Output:
	// Size: 4
	// Contains 1: true
	// Contains 6: false
	// Size after removing 1: 3
	// Contains 1 after removal: false
}
