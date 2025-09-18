package slicex_test

import (
	"fmt"
	"testing"

	"github.com/gophero/goal/col/slicex"
	"github.com/gophero/got"
)

// TestEach tests the Each function
func TestEach(t *testing.T) {
	logger := got.New(t, "Each")

	// Test with int slice
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	slicex.Each(numbers, func(x int) {
		sum += x
	})
	logger.Require(sum == 15, "Each should iterate over all elements")

	// Test with empty slice
	empty := []int{}
	count := 0
	slicex.Each(empty, func(x int) {
		count++
	})
	logger.Require(count == 0, "Each on empty slice should not call function")

	// Test with string slice
	words := []string{"hello", "world", "test"}
	concatenated := ""
	slicex.Each(words, func(s string) {
		concatenated += s
	})
	logger.Require(concatenated == "helloworldtest", "Each should work with strings")
}

// TestEachv tests the Eachv function
func TestEachv(t *testing.T) {
	logger := got.New(t, "Eachv")

	// Test with int slice
	numbers := []int{1, 2, 3, 4, 5}
	result := slicex.Eachv(numbers, func(x int) string {
		return fmt.Sprintf("num_%d", x)
	})
	expected := []string{"num_1", "num_2", "num_3", "num_4", "num_5"}
	logger.Require(slicex.Equal(result, expected), "Eachv should transform elements")

	// Test with empty slice
	empty := []int{}
	result2 := slicex.Eachv(empty, func(x int) string {
		return fmt.Sprintf("num_%d", x)
	})
	logger.Require(len(result2) == 0, "Eachv on empty slice should return empty")

	// Test with different types
	strings := []string{"a", "b", "c"}
	lengths := slicex.Eachv(strings, func(s string) int {
		return len(s)
	})
	expectedLengths := []int{1, 1, 1}
	logger.Require(slicex.Equal(lengths, expectedLengths), "Eachv should work with different types")
}

// TestMap tests the Map function
func TestMap(t *testing.T) {
	logger := got.New(t, "Map")

	// Test with int slice
	numbers := []int{1, 2, 3, 4, 5}
	result := slicex.Map(numbers, func(x int) string {
		return fmt.Sprintf("num_%d", x)
	})

	expected := map[string]int{
		"num_1": 1,
		"num_2": 2,
		"num_3": 3,
		"num_4": 4,
		"num_5": 5,
	}

	logger.Require(len(result) == len(expected), "Map should create correct number of entries")
	for k, v := range result {
		logger.Require(expected[k] == v, "Map should have correct key-value pairs")
	}

	// Test with empty slice
	empty := []int{}
	result2 := slicex.Map(empty, func(x int) string {
		return fmt.Sprintf("num_%d", x)
	})
	logger.Require(len(result2) == 0, "Map on empty slice should return empty map")

	// Test with duplicate keys (last value should win)
	duplicates := []int{1, 2, 1, 3, 2}
	result3 := slicex.Map(duplicates, func(x int) string {
		return "same_key"
	})
	logger.Require(len(result3) == 1, "Map should handle duplicate keys")
	logger.Require(result3["same_key"] == 2, "Map should use last value for duplicate keys")
}

// TestMapv tests the Mapv function
func TestMapv(t *testing.T) {
	logger := got.New(t, "Mapv")

	// Test with int slice
	numbers := []int{1, 2, 3, 4, 5}
	result := slicex.Mapv(numbers, func(x int) (string, int) {
		return fmt.Sprintf("num_%d", x), x * 2
	})

	expected := map[string]int{
		"num_1": 2,
		"num_2": 4,
		"num_3": 6,
		"num_4": 8,
		"num_5": 10,
	}

	logger.Require(len(result) == len(expected), "Mapv should create correct number of entries")
	for k, v := range result {
		logger.Require(expected[k] == v, "Mapv should have correct key-value pairs")
	}

	// Test with empty slice
	empty := []int{}
	result2 := slicex.Mapv(empty, func(x int) (string, int) {
		return fmt.Sprintf("num_%d", x), x * 2
	})
	logger.Require(len(result2) == 0, "Mapv on empty slice should return empty map")

	// Test with struct
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	result3 := slicex.Mapv(people, func(p Person) (string, int) {
		return p.Name, p.Age
	})

	expected3 := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	logger.Require(len(result3) == len(expected3), "Mapv should work with structs")
	for k, v := range result3 {
		logger.Require(expected3[k] == v, "Mapv should have correct struct key-value pairs")
	}
}
