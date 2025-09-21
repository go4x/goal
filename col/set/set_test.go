package set

import (
	"fmt"
	"testing"

	"github.com/go4x/got"
)

func TestSetInterface(t *testing.T) {
	logger := got.New(t, "SetInterface")

	// Test that all implementations implement Set interface
	var _ Set[int] = NewHashSet[int]()
	var _ Set[int] = NewArraySet[int]()
	var _ Set[int] = NewLinkedSet[int]()

	// Test polymorphic usage
	sets := []Set[string]{
		NewHashSet[string](),
		NewArraySet[string](),
		NewLinkedSet[string](),
	}

	for i, set := range sets {
		set.Add("test").Add("duplicate").Add("test")
		logger.Require(set.Size() == 2, "set %d should have 2 elements", i)
		logger.Require(set.Contains("test"), "set %d should contain 'test'", i)
	}
}

func TestSetOperations(t *testing.T) {
	logger := got.New(t, "SetOperations")

	// Test with different types
	set1 := New[int]()
	set1.Add(1).Add(2).Add(3)

	set2 := New[string]()
	set2.Add("a").Add("b").Add("c")

	logger.Require(set1.Size() == 3, "int set should have 3 elements")
	logger.Require(set2.Size() == 3, "string set should have 3 elements")

	// Test duplicate handling
	set3 := New[float64]()
	set3.Add(1.1).Add(1.1).Add(2.2).Add(1.1)
	logger.Require(set3.Size() == 2, "float set should have 2 unique elements")
}

func TestSetChaining(t *testing.T) {
	logger := got.New(t, "SetChaining")

	// Test method chaining
	set := New[int]()
	result := set.Add(1).Add(2).Add(3).Remove(2).Clear().Add(4).Add(5)

	logger.Require(result.Size() == 2, "chained operations should work correctly")
	logger.Require(result.Contains(4), "should contain 4")
	logger.Require(result.Contains(5), "should contain 5")
	logger.Require(!result.Contains(1), "should not contain 1 after clear")
}

func TestSetEmptyOperations(t *testing.T) {
	logger := got.New(t, "SetEmptyOperations")

	set := New[int]()

	// Test operations on empty set
	logger.Require(set.IsEmpty(), "new set should be empty")
	logger.Require(set.Size() == 0, "new set size should be 0")
	logger.Require(!set.Contains(1), "empty set should not contain any elements")

	// Test removing from empty set
	set.Remove(1)
	logger.Require(set.IsEmpty(), "removing from empty set should still be empty")

	// Test clearing empty set
	set.Clear()
	logger.Require(set.IsEmpty(), "clearing empty set should still be empty")
}

// ExampleNew demonstrates basic Set usage
func ExampleNew() {
	// Create a new set (HashSet by default)
	set := New[string]()

	// Add elements (duplicates are automatically ignored)
	set.Add("apple").Add("banana").Add("apple").Add("cherry")

	fmt.Println("Size:", set.Size())
	fmt.Println("Contains 'apple':", set.Contains("apple"))
	fmt.Println("Contains 'grape':", set.Contains("grape"))

	// Output:
	// Size: 3
	// Contains 'apple': true
	// Contains 'grape': false
}

// ExampleSet_polymorphic demonstrates polymorphic usage
func ExampleSet_polymorphic() {
	// Create different set types
	sets := []Set[string]{
		NewHashSet[string](),
		NewArraySet[string](),
		NewLinkedSet[string](),
	}

	// Add same elements to all sets
	for _, set := range sets {
		set.Add("a").Add("b").Add("a").Add("c")
	}

	// All sets should have same size (duplicates ignored)
	for i, set := range sets {
		fmt.Printf("Set %d size: %d\n", i+1, set.Size())
	}

	// Output:
	// Set 1 size: 3
	// Set 2 size: 3
	// Set 3 size: 3
}
