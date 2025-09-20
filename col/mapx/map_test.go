package mapx_test

import (
	"testing"

	"github.com/gophero/goal/col/mapx"
	"github.com/gophero/got"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	logger := got.New(t, "Map")

	// Test direct struct creation
	m := mapx.New[string, int]()
	assert.NotNil(t, m)
	assert.Equal(t, m.Size(), 0, "size of map should be 0")

	intMap := mapx.New[string, int]()
	intMap.Put("a", 1).Put("b", 2).Put("c", 3)
	v, ok := intMap.Get("a")
	logger.Require(ok, "key should be found")
	logger.Require(v == 1, "value of map should be %d", v)
	keys1 := intMap.Keys()
	logger.Require(eq(keys1, []string{"a", "b", "c"}), "keys of map should be equals to %v", keys1)
	val1 := intMap.Values()
	logger.Require(eq(val1, []int{1, 2, 3}), "values of map should be equals to %v", val1)
	del1, ok := intMap.Del("a")
	logger.Require(ok, "key should be found")
	logger.Require(del1 == 1, "del value of map should be %v", del1)
	intMap.Put("b", 3)
	upv1, ok := intMap.Get("b")
	logger.Require(ok, "key should be found")
	logger.Require(upv1 == 3, "update value of map should be %v", upv1)

	intMap1 := mapx.New[int, int]()
	intMap1.Put(1, 1).Put(2, 2).Put(3, 3)
	vv, ok := intMap1.Get(1)
	logger.Require(ok, "key should be found")
	logger.Require(vv == 1, "value of map should be %d", vv)
	keys2 := intMap1.Keys()
	logger.Require(eq(keys2, []int{1, 2, 3}), "keys of map should be equals to %v", keys2)
	val2 := intMap1.Values()
	logger.Require(eq(val2, []int{1, 2, 3}), "values of map should be equals to %v", val2)

	stringMap := mapx.New[string, string]()
	stringMap.Put("a", "1").Put("b", "2").Put("c", "3")
	v1, ok := stringMap.Get("a")
	logger.Require(ok, "key should be found")
	logger.Require(v1 == "1", "value of map should be %s", v1)
	keys3 := stringMap.Keys()
	logger.Require(eq(keys3, []string{"a", "b", "c"}), "keys of map should be be equals to %v", keys3)
	val3 := stringMap.Values()
	logger.Require(eq(val3, []string{"1", "2", "3"}), "values of map should be equals to %v", val3)

	anyMap := mapx.New[string, any]()
	anyMap.Put("a", 1).Put("b", "b").Put("c", 3.14)
	v2, ok := anyMap.Get("a")
	logger.Require(ok, "key should be found")
	logger.Require(v2 == 1, "value of map should be %s", v2)
	keys4 := anyMap.Keys()
	logger.Require(eq(keys4, []string{"a", "b", "c"}), "keys of map should be be equals to %v", keys4)
	val4 := anyMap.Values()
	// Equal(val4, []any{1, "b", 3.14}) 这里不能用==比较了，any没有实现comparable
	logger.Require(len(val4) == 3, "values of map should be equals to %v", val4)
}

// TestInterfaceCompliance tests that M[K, V] implements Map[K, V] interface
func TestInterfaceCompliance(t *testing.T) {
	var _ = mapx.New[string, int]()
	var _ = mapx.New[int, string]()
	var _ = mapx.New[any, any]()
}

// TestContains tests the Contains method
func TestContains(t *testing.T) {
	logger := got.New(t, "Contains")

	m := mapx.New[string, int]()
	m.Put("a", 1).Put("b", 2)

	// Test existing keys
	logger.Require(m.Contains("a"), "Contains should return true for existing key")
	logger.Require(m.Contains("b"), "Contains should return true for existing key")

	// Test non-existing key
	logger.Require(!m.Contains("c"), "Contains should return false for non-existing key")
}

// TestForEach tests the ForEach method
func TestForEach(t *testing.T) {
	logger := got.New(t, "ForEach")

	m := mapx.New[string, int]()
	m.Put("a", 1).Put("b", 2).Put("c", 3)

	// Collect keys and values during iteration
	var keys []string
	var values []int

	m.Each(func(k string, v int) {
		keys = append(keys, k)
		values = append(values, v)
	})

	logger.Require(len(keys) == 3, "ForEach should iterate over all 3 keys")
	logger.Require(len(values) == 3, "ForEach should iterate over all 3 values")

	// Check that all expected keys and values were collected
	expectedKeys := []string{"a", "b", "c"}
	expectedValues := []int{1, 2, 3}

	for _, expectedKey := range expectedKeys {
		found := false
		for _, key := range keys {
			if key == expectedKey {
				found = true
				break
			}
		}
		logger.Require(found, "Key %s should be found during ForEach", expectedKey)
	}

	for _, expectedValue := range expectedValues {
		found := false
		for _, value := range values {
			if value == expectedValue {
				found = true
				break
			}
		}
		logger.Require(found, "Value %d should be found during ForEach", expectedValue)
	}
}

// TestClear tests the Clear method
func TestClear(t *testing.T) {
	logger := got.New(t, "Clear")

	m := mapx.New[string, int]()
	m.Put("a", 1).Put("b", 2).Put("c", 3)

	// Test clearing
	result := m.Clear()
	logger.Require(result != nil, "Clear should return a non-nil map for chaining")
	logger.Require(m.Size() == 0, "Map should be empty after clearing")
	logger.Require(m.IsEmpty(), "Map should be empty after clearing")

	// Test that all keys are removed
	logger.Require(!m.Contains("a"), "Map should not contain 'a' after clearing")
	logger.Require(!m.Contains("b"), "Map should not contain 'b' after clearing")
	logger.Require(!m.Contains("c"), "Map should not contain 'c' after clearing")
}

func eq[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	// Create a map to count occurrences of each element in b
	count := make(map[T]int)
	for _, item := range b {
		count[item]++
	}

	// Check if each element in a exists in b with the same count
	for _, item := range a {
		if count[item] == 0 {
			return false
		}
		count[item]--
	}

	return true
}
