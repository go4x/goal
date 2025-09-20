package mapx_test

import (
	"testing"

	"github.com/gophero/goal/col/mapx"
	"github.com/gophero/got"
	"github.com/stretchr/testify/assert"
)

// TestArrayMapBasicOperations tests basic operations of ArrayMap
func TestArrayMapBasicOperations(t *testing.T) {
	logger := got.New(t, "ArrayMap Basic Operations")

	am := mapx.NewArrayMap[string, int]()
	assert.NotNil(t, am)
	assert.Equal(t, am.Size(), 0, "size should be 0")
	assert.True(t, am.IsEmpty(), "should be empty")

	// Test Put and Get
	am.Put("first", 1).Put("second", 2).Put("third", 3)
	assert.Equal(t, am.Size(), 3, "size should be 3")
	assert.False(t, am.IsEmpty(), "should not be empty")

	// Test Get
	val, ok := am.Get("first")
	logger.Require(ok, "key should be found")
	logger.Require(val == 1, "value should be 1")

	val, ok = am.Get("nonexistent")
	logger.Require(!ok, "key should not be found")
	logger.Require(val == 0, "value should be zero value")

	// Test Contains
	logger.Require(am.Contains("first"), "should contain 'first'")
	logger.Require(!am.Contains("nonexistent"), "should not contain 'nonexistent'")
}

// TestArrayMapInsertionOrder tests that insertion order is maintained
func TestArrayMapInsertionOrder(t *testing.T) {
	logger := got.New(t, "ArrayMap Insertion Order")

	am := mapx.NewArrayMap[string, int]()
	am.Put("z", 3).Put("a", 1).Put("m", 2)

	// Test Keys order
	keys := am.Keys()
	expectedKeys := []string{"z", "a", "m"}
	logger.Require(eq(keys, expectedKeys), "keys should maintain insertion order: %v", keys)

	// Test Values order
	values := am.Values()
	expectedValues := []int{3, 1, 2}
	logger.Require(eq(values, expectedValues), "values should maintain insertion order: %v", values)

	// Test Each iteration order
	var collectedKeys []string
	var collectedValues []int
	am.Each(func(k string, v int) {
		collectedKeys = append(collectedKeys, k)
		collectedValues = append(collectedValues, v)
	})

	logger.Require(eq(collectedKeys, expectedKeys), "Each should maintain insertion order for keys: %v", collectedKeys)
	logger.Require(eq(collectedValues, expectedValues), "Each should maintain insertion order for values: %v", collectedValues)
}

// TestArrayMapUpdateExistingKey tests updating existing keys
func TestArrayMapUpdateExistingKey(t *testing.T) {
	logger := got.New(t, "ArrayMap Update Existing Key")

	am := mapx.NewArrayMap[string, int]()
	am.Put("a", 1).Put("b", 2).Put("c", 3)

	// Update existing key
	am.Put("b", 99)

	val, ok := am.Get("b")
	logger.Require(ok, "key should still exist")
	logger.Require(val == 99, "value should be updated to 99")

	// Check that order is preserved
	keys := am.Keys()
	expectedKeys := []string{"a", "b", "c"}
	logger.Require(eq(keys, expectedKeys), "order should be preserved after update: %v", keys)

	values := am.Values()
	expectedValues := []int{1, 99, 3}
	logger.Require(eq(values, expectedValues), "values should reflect update: %v", values)
}

// TestArrayMapDelete tests deletion operations
func TestArrayMapDelete(t *testing.T) {
	logger := got.New(t, "ArrayMap Delete")

	am := mapx.NewArrayMap[string, int]()
	am.Put("a", 1).Put("b", 2).Put("c", 3).Put("d", 4)

	// Delete existing key
	val, ok := am.Del("b")
	logger.Require(ok, "deletion should succeed")
	logger.Require(val == 2, "deleted value should be 2")
	logger.Require(am.Size() == 3, "size should be 3 after deletion")

	// Check remaining keys and values
	keys := am.Keys()
	expectedKeys := []string{"a", "c", "d"}
	logger.Require(eq(keys, expectedKeys), "remaining keys should be correct: %v", keys)

	values := am.Values()
	expectedValues := []int{1, 3, 4}
	logger.Require(eq(values, expectedValues), "remaining values should be correct: %v", values)

	// Delete non-existent key
	_, ok = am.Del("nonexistent")
	logger.Require(!ok, "deletion should fail for non-existent key")
	logger.Require(am.Size() == 3, "size should remain 3")

	// Delete first element
	val, ok = am.Del("a")
	logger.Require(ok, "deletion of first element should succeed")
	logger.Require(val == 1, "deleted value should be 1")

	keys = am.Keys()
	expectedKeys = []string{"c", "d"}
	logger.Require(eq(keys, expectedKeys), "keys after deleting first: %v", keys)

	// Delete last element
	val, ok = am.Del("d")
	logger.Require(ok, "deletion of last element should succeed")
	logger.Require(val == 4, "deleted value should be 4")

	keys = am.Keys()
	expectedKeys = []string{"c"}
	logger.Require(eq(keys, expectedKeys), "keys after deleting last: %v", keys)

	// Delete remaining element
	_, ok = am.Del("c")
	logger.Require(ok, "deletion of last remaining element should succeed")
	logger.Require(am.IsEmpty(), "map should be empty after deleting all elements")
}

// TestArrayMapFirstLast tests First and Last methods
func TestArrayMapFirstLast(t *testing.T) {
	logger := got.New(t, "ArrayMap First Last")

	am := mapx.NewArrayMap[string, int]()

	// Test empty map
	_, _, ok := am.First()
	logger.Require(!ok, "First should return false for empty map")

	_, _, ok = am.Last()
	logger.Require(!ok, "Last should return false for empty map")

	// Add elements
	am.Put("first", 1).Put("middle", 2).Put("last", 3)

	// Test First
	k, v, ok := am.First()
	logger.Require(ok, "First should return true for non-empty map")
	logger.Require(k == "first", "First key should be 'first': %s", k)
	logger.Require(v == 1, "First value should be 1: %d", v)

	// Test Last
	k, v, ok = am.Last()
	logger.Require(ok, "Last should return true for non-empty map")
	logger.Require(k == "last", "Last key should be 'last': %s", k)
	logger.Require(v == 3, "Last value should be 3: %d", v)

	// Test single element
	am.Clear().Put("only", 42)
	k, v, ok = am.First()
	logger.Require(ok, "First should work for single element")
	logger.Require(k == "only", "First key should be 'only' for single element")

	k, v, ok = am.Last()
	logger.Require(ok, "Last should work for single element")
	logger.Require(k == "only", "Last key should be 'only' for single element")
	logger.Require(v == 42, "Last value should be 42 for single element")
}

// TestArrayMapClear tests Clear method
func TestArrayMapClear(t *testing.T) {
	logger := got.New(t, "ArrayMap Clear")

	am := mapx.NewArrayMap[string, int]()
	am.Put("a", 1).Put("b", 2).Put("c", 3)

	result := am.Clear()
	logger.Require(result != nil, "Clear should return non-nil for chaining")
	logger.Require(am.Size() == 0, "size should be 0 after clearing")
	logger.Require(am.IsEmpty(), "should be empty after clearing")

	// Test that all keys are removed
	logger.Require(!am.Contains("a"), "should not contain 'a' after clearing")
	logger.Require(!am.Contains("b"), "should not contain 'b' after clearing")
	logger.Require(!am.Contains("c"), "should not contain 'c' after clearing")

	// Test that we can still use the map after clearing
	am.Put("new", 42)
	logger.Require(am.Size() == 1, "should be able to add elements after clearing")
	logger.Require(am.Contains("new"), "should contain new element")
}

// TestArrayMapString tests String method
func TestArrayMapString(t *testing.T) {
	logger := got.New(t, "ArrayMap String")

	am := mapx.NewArrayMap[string, int]()

	// Test empty map
	str := am.String()
	logger.Require(str == "map[]", "empty map string should be 'map[]': %s", str)

	// Test single element
	am.Put("a", 1)
	str = am.String()
	logger.Require(str == "map[a:1]", "single element string should be 'map[a:1]': %s", str)

	// Test multiple elements
	am.Put("b", 2).Put("c", 3)
	str = am.String()
	// Order should be preserved
	logger.Require(str == "map[a:1 b:2 c:3]", "multiple elements string should maintain order: %s", str)
}

// TestArrayMapDifferentTypes tests ArrayMap with different key and value types
func TestArrayMapDifferentTypes(t *testing.T) {
	logger := got.New(t, "ArrayMap Different Types")

	// Test with int keys
	intMap := mapx.NewArrayMap[int, string]()
	intMap.Put(3, "three").Put(1, "one").Put(2, "two")

	keys := intMap.Keys()
	expectedKeys := []int{3, 1, 2}
	logger.Require(eq(keys, expectedKeys), "int keys should maintain insertion order: %v", keys)

	values := intMap.Values()
	expectedValues := []string{"three", "one", "two"}
	logger.Require(eq(values, expectedValues), "string values should maintain insertion order: %v", values)

	// Test with any values
	anyMap := mapx.NewArrayMap[string, any]()
	anyMap.Put("int", 42).Put("string", "hello").Put("float", 3.14)

	anyKeys := anyMap.Keys()
	expectedKeysStr := []string{"int", "string", "float"}
	logger.Require(eq(anyKeys, expectedKeysStr), "any values keys should maintain insertion order: %v", anyKeys)

	// Test that values are correct (can't easily compare any types)
	logger.Require(anyMap.Size() == 3, "any map should have 3 elements")
}

// TestArrayMapMethodChaining tests method chaining
func TestArrayMapMethodChaining(t *testing.T) {
	logger := got.New(t, "ArrayMap Method Chaining")

	am := mapx.NewArrayMap[string, int]()

	// Test chaining Put operations
	result := am.Put("a", 1).Put("b", 2).Put("c", 3)
	logger.Require(result == am, "chained Put should return the same instance")
	logger.Require(am.Size() == 3, "size should be 3 after chained Puts")

	// Test chaining with Clear
	result = am.Clear().Put("new", 42)
	logger.Require(result == am, "chained Clear and Put should return the same instance")
	logger.Require(am.Size() == 1, "size should be 1 after clear and put")
	logger.Require(am.Contains("new"), "should contain new element")
}

// TestArrayMapInterfaceCompliance tests that ArrayMap[K, V] implements Map[K, V] interface
func TestArrayMapInterfaceCompliance(t *testing.T) {
	var _ mapx.Map[string, int] = mapx.NewArrayMap[string, int]()
	var _ mapx.Map[int, string] = mapx.NewArrayMap[int, string]()
	var _ mapx.Map[any, any] = mapx.NewArrayMap[any, any]()
}

// TestArrayMapEdgeCases tests edge cases
func TestArrayMapEdgeCases(t *testing.T) {
	logger := got.New(t, "ArrayMap Edge Cases")

	am := mapx.NewArrayMap[string, int]()

	// Test operations on empty map
	_, ok := am.Get("nonexistent")
	logger.Require(!ok, "Get on empty map should return false")

	_, ok = am.Del("nonexistent")
	logger.Require(!ok, "Del on empty map should return false")

	logger.Require(!am.Contains("nonexistent"), "Contains on empty map should return false")

	keys := am.Keys()
	logger.Require(len(keys) == 0, "Keys on empty map should return empty slice")

	values := am.Values()
	logger.Require(len(values) == 0, "Values on empty map should return empty slice")

	// Test Each on empty map (should not panic)
	am.Each(func(k string, v int) {
		t.Errorf("Each should not be called on empty map")
	})

	// Test duplicate keys
	am.Put("a", 1).Put("a", 2)
	logger.Require(am.Size() == 1, "duplicate keys should not increase size")

	val, ok := am.Get("a")
	logger.Require(ok, "should be able to get value after duplicate key")
	logger.Require(val == 2, "value should be the last one set: %d", val)
}
