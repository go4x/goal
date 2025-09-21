package slicex_test

import (
	"fmt"
	"testing"

	"github.com/gophero/goal/col/slicex"
)

// ExampleS_Sort_immutable demonstrates the immutable Sort method
func ExampleS_Sort_immutable() {
	// Create a slice
	numbers := slicex.From([]int{3, 1, 4, 1, 5, 9, 2, 6})
	fmt.Println("Original:", numbers.To())

	// Sort without modifying the original
	sorted := numbers.Sort(func(a, b int) bool {
		return a < b // Ascending order
	})

	fmt.Println("Sorted result:", sorted.To())
	fmt.Println("Original unchanged:", numbers.To())

	// Output:
	// Original: [3 1 4 1 5 9 2 6]
	// Sorted result: [1 1 2 3 4 5 6 9]
	// Original unchanged: [3 1 4 1 5 9 2 6]
}

// ExampleS_SortInPlace_mutable demonstrates the mutable SortInPlace method
func ExampleS_SortInPlace_mutable() {
	// Create a slice
	numbers := slicex.From([]int{3, 1, 4, 1, 5, 9, 2, 6})
	fmt.Println("Original:", numbers.To())

	// Sort in place (modifies the original)
	sorted := numbers.SortInPlace(func(a, b int) bool {
		return a < b // Ascending order
	})

	fmt.Println("Sorted result:", sorted.To())
	fmt.Println("Original modified:", numbers.To())

	// Output:
	// Original: [3 1 4 1 5 9 2 6]
	// Sorted result: [1 1 2 3 4 5 6 9]
	// Original modified: [1 1 2 3 4 5 6 9]
}

// TestSortComparison demonstrates the difference between Sort and SortInPlace
func TestSortComparison(t *testing.T) {
	fmt.Println("=== Sort vs SortInPlace Comparison ===")

	// Test immutable Sort
	fmt.Println("\n1. Immutable Sort (recommended for most cases):")
	original1 := slicex.From([]int{3, 1, 4, 1, 5})
	fmt.Printf("   Before Sort: %v\n", original1.To())

	sorted1 := original1.Sort(func(a, b int) bool { return a < b })
	fmt.Printf("   After Sort: %v\n", sorted1.To())
	fmt.Printf("   Original: %v (unchanged)\n", original1.To())

	// Test mutable SortInPlace
	fmt.Println("\n2. Mutable SortInPlace (for performance-critical scenarios):")
	original2 := slicex.From([]int{3, 1, 4, 1, 5})
	fmt.Printf("   Before SortInPlace: %v\n", original2.To())

	sorted2 := original2.SortInPlace(func(a, b int) bool { return a < b })
	fmt.Printf("   After SortInPlace: %v\n", sorted2.To())
	fmt.Printf("   Original: %v (modified!)\n", original2.To())

	fmt.Println("\n=== When to use which method? ===")
	fmt.Println("• Use Sort() when:")
	fmt.Println("  - You need to preserve the original data")
	fmt.Println("  - Working with concurrent code")
	fmt.Println("  - Following functional programming principles")
	fmt.Println("  - Memory usage is not critical")

	fmt.Println("\n• Use SortInPlace() when:")
	fmt.Println("  - Memory efficiency is critical")
	fmt.Println("  - Working with large datasets")
	fmt.Println("  - You don't need the original order")
	fmt.Println("  - Performance is more important than safety")
}
