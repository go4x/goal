package slicex_test

import (
	"fmt"
	"testing"

	"github.com/go4x/goal/col/slicex"
)

// TestEach tests the Each function
func TestEach(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := 0
		slicex.Each(numbers, func(x int) {
			sum += x
		})
		if sum != 15 {
			t.Errorf("Each should iterate over all elements, expected 15, got %d", sum)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		empty := []int{}
		count := 0
		slicex.Each(empty, func(x int) {
			count++
		})
		if count != 0 {
			t.Errorf("Each on empty slice should not call function, expected 0, got %d", count)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		words := []string{"hello", "world", "test"}
		concatenated := ""
		slicex.Each(words, func(s string) {
			concatenated += s
		})
		if concatenated != "helloworldtest" {
			t.Errorf("Each should work with strings, expected 'helloworldtest', got '%s'", concatenated)
		}
	})
}

// TestEachv tests the Eachv function
func TestEachv(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := slicex.Eachv(numbers, func(x int) string {
			return fmt.Sprintf("num_%d", x)
		})
		expected := []string{"num_1", "num_2", "num_3", "num_4", "num_5"}
		if !slicex.Equal(result, expected) {
			t.Errorf("Eachv should transform elements, expected %v, got %v", expected, result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		empty := []int{}
		result := slicex.Eachv(empty, func(x int) string {
			return fmt.Sprintf("num_%d", x)
		})
		if len(result) != 0 {
			t.Errorf("Eachv on empty slice should return empty, expected length 0, got %d", len(result))
		}
	})

	t.Run("different types", func(t *testing.T) {
		strings := []string{"a", "b", "c"}
		lengths := slicex.Eachv(strings, func(s string) int {
			return len(s)
		})
		expectedLengths := []int{1, 1, 1}
		if !slicex.Equal(lengths, expectedLengths) {
			t.Errorf("Eachv should work with different types, expected %v, got %v", expectedLengths, lengths)
		}
	})
}

// TestGroup tests the Group function
func TestGroup(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := slicex.Group(numbers, func(x int) string {
			return fmt.Sprintf("num_%d", x)
		})

		expected := map[string]int{
			"num_1": 1,
			"num_2": 2,
			"num_3": 3,
			"num_4": 4,
			"num_5": 5,
		}

		if len(result) != len(expected) {
			t.Errorf("Group should create correct number of entries, expected %d, got %d", len(expected), len(result))
		}
		for k, v := range result {
			if expected[k] != v {
				t.Errorf("Group should have correct key-value pairs, expected %d for key %s, got %d", expected[k], k, v)
			}
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		empty := []int{}
		result := slicex.Group(empty, func(x int) string {
			return fmt.Sprintf("num_%d", x)
		})
		if len(result) != 0 {
			t.Errorf("Group on empty slice should return empty map, expected length 0, got %d", len(result))
		}
	})

	t.Run("duplicate keys", func(t *testing.T) {
		duplicates := []int{1, 2, 1, 3, 2}
		result := slicex.Group(duplicates, func(x int) string {
			return "same_key"
		})
		if len(result) != 1 {
			t.Errorf("Group should handle duplicate keys, expected 1 entry, got %d", len(result))
		}
		if result["same_key"] != 2 {
			t.Errorf("Group should use last value for duplicate keys, expected 2, got %d", result["same_key"])
		}
	})
}

// TestGroupTo tests the GroupTo function
func TestGroupTo(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := slicex.GroupTo(numbers, func(x int) (string, int) {
			return fmt.Sprintf("num_%d", x), x * 2
		})

		expected := map[string]int{
			"num_1": 2,
			"num_2": 4,
			"num_3": 6,
			"num_4": 8,
			"num_5": 10,
		}

		if len(result) != len(expected) {
			t.Errorf("GroupTo should create correct number of entries, expected %d, got %d", len(expected), len(result))
		}
		for k, v := range result {
			if expected[k] != v {
				t.Errorf("GroupTo should have correct key-value pairs, expected %d for key %s, got %d", expected[k], k, v)
			}
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		empty := []int{}
		result := slicex.GroupTo(empty, func(x int) (string, int) {
			return fmt.Sprintf("num_%d", x), x * 2
		})
		if len(result) != 0 {
			t.Errorf("GroupTo on empty slice should return empty map, expected length 0, got %d", len(result))
		}
	})

	t.Run("struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}
		result := slicex.GroupTo(people, func(p Person) (string, int) {
			return p.Name, p.Age
		})

		expected := map[string]int{
			"Alice": 30,
			"Bob":   25,
		}

		if len(result) != len(expected) {
			t.Errorf("GroupTo should work with structs, expected %d entries, got %d", len(expected), len(result))
		}
		for k, v := range result {
			if expected[k] != v {
				t.Errorf("GroupTo should have correct struct key-value pairs, expected %d for key %s, got %d", expected[k], k, v)
			}
		}
	})
}
