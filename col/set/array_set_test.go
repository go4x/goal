package set

import (
	"fmt"
	"testing"

	"github.com/go4x/got"
)

func TestArraySet(t *testing.T) {
	logger := got.New(t, "ArraySet")

	// Test ordered insertion
	set := NewArraySet[string]()
	set.Add("first").Add("second").Add("third").Add("first") // Duplicate

	elems := set.Elems()
	logger.Require(len(elems) == 3, "should have 3 elements")
	logger.Require(elems[0] == "first", "first element should be 'first'")
	logger.Require(elems[1] == "second", "second element should be 'second'")
	logger.Require(elems[2] == "third", "third element should be 'third'")

	// Test removal maintains order
	set.Remove("second")
	elems = set.Elems()
	logger.Require(len(elems) == 2, "should have 2 elements after removal")
	logger.Require(elems[0] == "first", "first element should still be 'first'")
	logger.Require(elems[1] == "third", "second element should be 'third'")
}

func TestArraySetOrderPreservation(t *testing.T) {
	logger := got.New(t, "ArraySetOrderPreservation")

	set := NewArraySet[int]()
	set.Add(3).Add(1).Add(4).Add(1).Add(5).Add(9).Add(2).Add(6)

	elems := set.Elems()
	expected := []int{3, 1, 4, 5, 9, 2, 6} // Duplicates removed, order preserved
	logger.Require(len(elems) == len(expected), "should have correct number of elements")

	for i, expectedElem := range expected {
		logger.Require(elems[i] == expectedElem, "element %d should be %d", i, expectedElem)
	}
}

func TestArraySetPerformanceCharacteristics(t *testing.T) {
	logger := got.New(t, "ArraySetPerformanceCharacteristics")

	// Test with small dataset (ArraySet's sweet spot)
	set := NewArraySet[int]()

	// Add elements
	for i := 0; i < 100; i++ {
		set.Add(i)
	}

	logger.Require(set.Size() == 100, "should have 100 elements")
	logger.Require(set.Contains(50), "should contain element 50")

	// Test removal
	set.Remove(50)
	logger.Require(!set.Contains(50), "should not contain element 50 after removal")
	logger.Require(set.Size() == 99, "should have 99 elements after removal")
}

func TestArraySetInterfaceCompliance(t *testing.T) {
	logger := got.New(t, "ArraySetInterfaceCompliance")

	// Test that ArraySet implements Set interface
	var _ Set[int] = NewArraySet[int]()
	var _ Set[string] = NewArraySet[string]()

	// Test polymorphic usage
	arraySet := NewArraySet[int]()
	var set Set[int] = arraySet

	set.Add(1).Add(2).Add(3)
	logger.Require(set.Size() == 3, "polymorphic usage should work")
	logger.Require(set.Contains(2), "polymorphic contains should work")
}

// ExampleNewArraySet demonstrates ArraySet usage
func ExampleNewArraySet() {
	// Create an ArraySet (maintains insertion order)
	arraySet := NewArraySet[string]()

	// Add elements
	arraySet.Add("first").Add("second").Add("third").Add("first")

	// Get all elements (in insertion order)
	elems := arraySet.Elems()
	fmt.Println("Elements in order:", elems)

	// Remove middle element
	arraySet.Remove("second")
	fmt.Println("After removing 'second':", arraySet.Elems())

	// Output:
	// Elements in order: [first second third]
	// After removing 'second': [first third]
}
