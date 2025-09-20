package mapx_test

import (
	"fmt"
	"testing"

	"github.com/gophero/goal/col/mapx"
	"github.com/gophero/got"
	"github.com/stretchr/testify/assert"
)

// TestLinkedMapBasicOperations tests basic operations of LinkedMap
func TestLinkedMapBasicOperations(t *testing.T) {
	logger := got.New(t, "LinkedMap Basic Operations")

	lm := mapx.NewLinkedMap[string, int]()
	assert.NotNil(t, lm)
	assert.Equal(t, lm.Size(), 0, "size should be 0")
	assert.True(t, lm.IsEmpty(), "should be empty")

	// Test Put and Get
	lm.Put("first", 1).Put("second", 2).Put("third", 3)
	assert.Equal(t, lm.Size(), 3, "size should be 3")
	assert.False(t, lm.IsEmpty(), "should not be empty")

	// Test Get
	val, ok := lm.Get("first")
	logger.Require(ok, "key should be found")
	logger.Require(val == 1, "value should be 1")

	val, ok = lm.Get("nonexistent")
	logger.Require(!ok, "key should not be found")
	logger.Require(val == 0, "value should be zero value")

	// Test Contains
	logger.Require(lm.Contains("first"), "should contain 'first'")
	logger.Require(!lm.Contains("nonexistent"), "should not contain 'nonexistent'")
}

// TestLinkedMapInsertionOrder tests that insertion order is maintained
func TestLinkedMapInsertionOrder(t *testing.T) {
	logger := got.New(t, "LinkedMap Insertion Order")

	lm := mapx.NewLinkedMap[string, int]()
	lm.Put("z", 3).Put("a", 1).Put("m", 2)

	// Test Keys order
	keys := lm.Keys()
	expectedKeys := []string{"z", "a", "m"}
	logger.Require(eq(keys, expectedKeys), "keys should maintain insertion order: %v", keys)

	// Test Values order
	values := lm.Values()
	expectedValues := []int{3, 1, 2}
	logger.Require(eq(values, expectedValues), "values should maintain insertion order: %v", values)

	// Test Each iteration order
	var collectedKeys []string
	var collectedValues []int
	lm.Each(func(k string, v int) {
		collectedKeys = append(collectedKeys, k)
		collectedValues = append(collectedValues, v)
	})

	logger.Require(eq(collectedKeys, expectedKeys), "Each should maintain insertion order for keys: %v", collectedKeys)
	logger.Require(eq(collectedValues, expectedValues), "Each should maintain insertion order for values: %v", collectedValues)
}

// TestLinkedMapUpdateExistingKey tests updating existing keys
func TestLinkedMapUpdateExistingKey(t *testing.T) {
	logger := got.New(t, "LinkedMap Update Existing Key")

	lm := mapx.NewLinkedMap[string, int]()
	lm.Put("a", 1).Put("b", 2).Put("c", 3)

	// Update existing key
	lm.Put("b", 99)

	val, ok := lm.Get("b")
	logger.Require(ok, "key should still exist")
	logger.Require(val == 99, "value should be updated to 99")

	// Check that order is preserved
	keys := lm.Keys()
	expectedKeys := []string{"a", "b", "c"}
	logger.Require(eq(keys, expectedKeys), "order should be preserved after update: %v", keys)

	values := lm.Values()
	expectedValues := []int{1, 99, 3}
	logger.Require(eq(values, expectedValues), "values should reflect update: %v", values)
}

// TestLinkedMapDelete tests deletion operations
func TestLinkedMapDelete(t *testing.T) {
	logger := got.New(t, "LinkedMap Delete")

	lm := mapx.NewLinkedMap[string, int]()
	lm.Put("a", 1).Put("b", 2).Put("c", 3).Put("d", 4)

	// Delete existing key
	val, ok := lm.Del("b")
	logger.Require(ok, "deletion should succeed")
	logger.Require(val == 2, "deleted value should be 2")
	logger.Require(lm.Size() == 3, "size should be 3 after deletion")

	// Check remaining keys and values
	keys := lm.Keys()
	expectedKeys := []string{"a", "c", "d"}
	logger.Require(eq(keys, expectedKeys), "remaining keys should be correct: %v", keys)

	values := lm.Values()
	expectedValues := []int{1, 3, 4}
	logger.Require(eq(values, expectedValues), "remaining values should be correct: %v", values)

	// Delete non-existent key
	_, ok = lm.Del("nonexistent")
	logger.Require(!ok, "deletion should fail for non-existent key")
	logger.Require(lm.Size() == 3, "size should remain 3")

	// Delete first element
	val, ok = lm.Del("a")
	logger.Require(ok, "deletion of first element should succeed")
	logger.Require(val == 1, "deleted value should be 1")

	keys = lm.Keys()
	expectedKeys = []string{"c", "d"}
	logger.Require(eq(keys, expectedKeys), "keys after deleting first: %v", keys)

	// Delete last element
	val, ok = lm.Del("d")
	logger.Require(ok, "deletion of last element should succeed")
	logger.Require(val == 4, "deleted value should be 4")

	keys = lm.Keys()
	expectedKeys = []string{"c"}
	logger.Require(eq(keys, expectedKeys), "keys after deleting last: %v", keys)

	// Delete remaining element
	_, ok = lm.Del("c")
	logger.Require(ok, "deletion of last remaining element should succeed")
	logger.Require(lm.IsEmpty(), "map should be empty after deleting all elements")
}

// TestLinkedMapFirstLast tests First and Last methods
func TestLinkedMapFirstLast(t *testing.T) {
	logger := got.New(t, "LinkedMap First Last")

	lm := mapx.NewLinkedMap[string, int]()

	// Test empty map
	_, _, ok := lm.First()
	logger.Require(!ok, "First should return false for empty map")

	_, _, ok = lm.Last()
	logger.Require(!ok, "Last should return false for empty map")

	// Add elements
	lm.Put("first", 1).Put("middle", 2).Put("last", 3)

	// Test First
	k, v, ok := lm.First()
	logger.Require(ok, "First should return true for non-empty map")
	logger.Require(k == "first", "First key should be 'first': %s", k)
	logger.Require(v == 1, "First value should be 1: %d", v)

	// Test Last
	k, v, ok = lm.Last()
	logger.Require(ok, "Last should return true for non-empty map")
	logger.Require(k == "last", "Last key should be 'last': %s", k)
	logger.Require(v == 3, "Last value should be 3: %d", v)

	// Test single element
	lm.Clear().Put("only", 42)
	k, v, ok = lm.First()
	logger.Require(ok, "First should work for single element")
	logger.Require(k == "only", "First key should be 'only' for single element")

	k, v, ok = lm.Last()
	logger.Require(ok, "Last should work for single element")
	logger.Require(k == "only", "Last key should be 'only' for single element")
	logger.Require(v == 42, "Last value should be 42 for single element")
}

// TestLinkedMapClear tests Clear method
func TestLinkedMapClear(t *testing.T) {
	logger := got.New(t, "LinkedMap Clear")

	lm := mapx.NewLinkedMap[string, int]()
	lm.Put("a", 1).Put("b", 2).Put("c", 3)

	result := lm.Clear()
	logger.Require(result != nil, "Clear should return non-nil for chaining")
	logger.Require(lm.Size() == 0, "size should be 0 after clearing")
	logger.Require(lm.IsEmpty(), "should be empty after clearing")

	// Test that all keys are removed
	logger.Require(!lm.Contains("a"), "should not contain 'a' after clearing")
	logger.Require(!lm.Contains("b"), "should not contain 'b' after clearing")
	logger.Require(!lm.Contains("c"), "should not contain 'c' after clearing")

	// Test that we can still use the map after clearing
	lm.Put("new", 42)
	logger.Require(lm.Size() == 1, "should be able to add elements after clearing")
	logger.Require(lm.Contains("new"), "should contain new element")
}

// TestLinkedMapString tests String method
func TestLinkedMapString(t *testing.T) {
	logger := got.New(t, "LinkedMap String")

	lm := mapx.NewLinkedMap[string, int]()

	// Test empty map
	str := lm.String()
	logger.Require(str == "map[]", "empty map string should be 'map[]': %s", str)

	// Test single element
	lm.Put("a", 1)
	str = lm.String()
	logger.Require(str == "map[a:1]", "single element string should be 'map[a:1]': %s", str)

	// Test multiple elements
	lm.Put("b", 2).Put("c", 3)
	str = lm.String()
	// Order should be preserved
	logger.Require(str == "map[a:1 b:2 c:3]", "multiple elements string should maintain order: %s", str)
}

// TestLinkedMapDifferentTypes tests LinkedMap with different key and value types
func TestLinkedMapDifferentTypes(t *testing.T) {
	logger := got.New(t, "LinkedMap Different Types")

	// Test with int keys
	intMap := mapx.NewLinkedMap[int, string]()
	intMap.Put(3, "three").Put(1, "one").Put(2, "two")

	keys := intMap.Keys()
	expectedKeys := []int{3, 1, 2}
	logger.Require(eq(keys, expectedKeys), "int keys should maintain insertion order: %v", keys)

	values := intMap.Values()
	expectedValues := []string{"three", "one", "two"}
	logger.Require(eq(values, expectedValues), "string values should maintain insertion order: %v", values)

	// Test with any values
	anyMap := mapx.NewLinkedMap[string, any]()
	anyMap.Put("int", 42).Put("string", "hello").Put("float", 3.14)

	anyKeys := anyMap.Keys()
	expectedKeysStr := []string{"int", "string", "float"}
	logger.Require(eq(anyKeys, expectedKeysStr), "any values keys should maintain insertion order: %v", anyKeys)

	// Test that values are correct (can't easily compare any types)
	logger.Require(anyMap.Size() == 3, "any map should have 3 elements")
}

// TestLinkedMapMethodChaining tests method chaining
func TestLinkedMapMethodChaining(t *testing.T) {
	logger := got.New(t, "LinkedMap Method Chaining")

	lm := mapx.NewLinkedMap[string, int]()

	// Test chaining Put operations
	result := lm.Put("a", 1).Put("b", 2).Put("c", 3)
	logger.Require(result == lm, "chained Put should return the same instance")
	logger.Require(lm.Size() == 3, "size should be 3 after chained Puts")

	// Test chaining with Clear
	result = lm.Clear().Put("new", 42)
	logger.Require(result == lm, "chained Clear and Put should return the same instance")
	logger.Require(lm.Size() == 1, "size should be 1 after clear and put")
	logger.Require(lm.Contains("new"), "should contain new element")
}

// TestLinkedMapEdgeCases tests edge cases
func TestLinkedMapEdgeCases(t *testing.T) {
	logger := got.New(t, "LinkedMap Edge Cases")

	lm := mapx.NewLinkedMap[string, int]()

	// Test operations on empty map
	_, ok := lm.Get("nonexistent")
	logger.Require(!ok, "Get on empty map should return false")

	_, ok = lm.Del("nonexistent")
	logger.Require(!ok, "Del on empty map should return false")

	logger.Require(!lm.Contains("nonexistent"), "Contains on empty map should return false")

	keys := lm.Keys()
	logger.Require(len(keys) == 0, "Keys on empty map should return empty slice")

	values := lm.Values()
	logger.Require(len(values) == 0, "Values on empty map should return empty slice")

	// Test Each on empty map (should not panic)
	lm.Each(func(k string, v int) {
		t.Errorf("Each should not be called on empty map")
	})

	// Test duplicate keys
	lm.Put("a", 1).Put("a", 2)
	logger.Require(lm.Size() == 1, "duplicate keys should not increase size")

	val, ok := lm.Get("a")
	logger.Require(ok, "should be able to get value after duplicate key")
	logger.Require(val == 2, "value should be the last one set: %d", val)
}

// TestLinkedMapMoveToEnd tests MoveToEnd method
func TestLinkedMapMoveToEnd(t *testing.T) {
	logger := got.New(t, "LinkedMap MoveToEnd")

	lm := mapx.NewLinkedMap[string, int]()
	lm.Put("a", 1).Put("b", 2).Put("c", 3)

	// Move existing key to end
	moved := lm.MoveToEnd("a")
	logger.Require(moved, "MoveToEnd should succeed for existing key")

	// Check order after move
	keys := lm.Keys()
	expectedKeys := []string{"b", "c", "a"}
	logger.Require(eq(keys, expectedKeys), "order should be correct after MoveToEnd: %v", keys)

	// Test moving non-existent key
	moved = lm.MoveToEnd("nonexistent")
	logger.Require(!moved, "MoveToEnd should fail for non-existent key")

	// Test moving last element (should be no-op)
	moved = lm.MoveToEnd("a")
	logger.Require(moved, "MoveToEnd should succeed for last element")

	keys = lm.Keys()
	expectedKeys = []string{"b", "c", "a"}
	logger.Require(eq(keys, expectedKeys), "order should remain the same: %v", keys)
}

// TestLinkedMapMoveToFront tests MoveToFront method
func TestLinkedMapMoveToFront(t *testing.T) {
	logger := got.New(t, "LinkedMap MoveToFront")

	lm := mapx.NewLinkedMap[string, int]()
	lm.Put("a", 1).Put("b", 2).Put("c", 3)

	// Move existing key to front
	moved := lm.MoveToFront("c")
	logger.Require(moved, "MoveToFront should succeed for existing key")

	// Check order after move
	keys := lm.Keys()
	expectedKeys := []string{"c", "a", "b"}
	logger.Require(eq(keys, expectedKeys), "order should be correct after MoveToFront: %v", keys)

	// Test moving non-existent key
	moved = lm.MoveToFront("nonexistent")
	logger.Require(!moved, "MoveToFront should fail for non-existent key")

	// Test moving first element (should be no-op)
	moved = lm.MoveToFront("c")
	logger.Require(moved, "MoveToFront should succeed for first element")

	keys = lm.Keys()
	expectedKeys = []string{"c", "a", "b"}
	logger.Require(eq(keys, expectedKeys), "order should remain the same: %v", keys)
}

// TestLinkedMapInterfaceCompliance tests that LinkedMap[K, V] implements Map[K, V] interface
func TestLinkedMapInterfaceCompliance(t *testing.T) {
	var _ mapx.Map[string, int] = mapx.NewLinkedMap[string, int]()
	var _ mapx.Map[int, string] = mapx.NewLinkedMap[int, string]()
	var _ mapx.Map[any, any] = mapx.NewLinkedMap[any, any]()
}

// TestLinkedMapPerformance tests basic performance characteristics
func TestLinkedMapPerformance(t *testing.T) {
	logger := got.New(t, "LinkedMap Performance")

	lm := mapx.NewLinkedMap[string, int]()

	// Add many elements
	for i := 0; i < 1000; i++ {
		lm.Put(fmt.Sprintf("key%d", i), i)
	}

	logger.Require(lm.Size() == 1000, "should have 1000 elements")

	// Test random access (should be O(1))
	val, ok := lm.Get("key500")
	logger.Require(ok, "should find key500")
	logger.Require(val == 500, "value should be 500")

	// Test deletion from middle
	val, ok = lm.Del("key500")
	logger.Require(ok, "should delete key500")
	logger.Require(val == 500, "deleted value should be 500")
	logger.Require(lm.Size() == 999, "size should be 999 after deletion")

	// Test that order is maintained after deletion
	keys := lm.Keys()
	logger.Require(len(keys) == 999, "should have 999 keys")
	logger.Require(keys[499] == "key499", "key499 should be at index 499")
	logger.Require(keys[500] == "key501", "key501 should be at index 500")
}
