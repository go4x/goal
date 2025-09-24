package mathx_test

import (
	"fmt"

	"github.com/go4x/goal/mathx"
)

func ExampleAdd() {
	// Basic addition with precision
	result := mathx.Add(0.1, 0.2)
	fmt.Printf("0.1 + 0.2 = %.10f\n", result.Float64())
	// Output: 0.1 + 0.2 = 0.3000000000
}

func ExampleAdd_chainable() {
	// Chainable operations
	result := mathx.Add(0.1, 0.2).
		Mul(10).
		Div(3, 2).
		Round(2).
		ToStringFixed(2)
	fmt.Printf("Result: %s\n", result)
	// Output: Result: 1.00
}

func ExampleMul() {
	// Money formatting with thousands separator
	price := mathx.Mul(99.99, 1.15).
		Round(2).
		FormatMoney(2)
	fmt.Printf("Price: $%s\n", price)
	// Output: Price: $114.99
}

func ExampleAverage() {
	// Statistical calculations
	scores := []float64{85, 92, 78, 96, 88}

	avg := mathx.Average(scores...)
	median := mathx.Median(scores...)
	std := mathx.StandardDeviation(scores...)

	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Printf("Standard Deviation: %.2f\n", std)
	// Output:
	// Average: 87.80
	// Median: 88.00
	// Standard Deviation: 6.87
}

func ExamplePercentage() {
	// Calculate tip
	bill := 50.00
	tipPercent := 18.0
	tip := mathx.Percentage(bill, tipPercent)

	// Calculate total
	total := mathx.Add(bill, tip)

	fmt.Printf("Bill: $%.2f\n", bill)
	fmt.Printf("Tip (%.0f%%): $%.2f\n", tipPercent, tip)
	fmt.Printf("Total: $%.2f\n", total.Float64())
	// Output:
	// Bill: $50.00
	// Tip (18%): $9.00
	// Total: $59.00
}

func ExampleAdd_clean() {
	// Remove trailing zeros
	result := mathx.Add(0.1, 0.2).
		Clean().
		ToString()
	fmt.Printf("Clean result: %s\n", result)
	// Output: Clean result: 0.3
}

func ExampleToString() {
	// Various string conversion methods
	value := 3.14159

	fmt.Printf("ToString: %s\n", mathx.ToString(value))
	fmt.Printf("ToStringFixed(2): %s\n", mathx.ToStringFixed(value, 2))
	fmt.Printf("ToStringBank(2): %s\n", mathx.ToStringBank(value, 2))
	// Output:
	// ToString: 3.14159
	// ToStringFixed(2): 3.14
	// ToStringBank(2): 3.14
}

func ExampleAbs() {
	// Mathematical functions
	fmt.Printf("Abs(-3.14): %.2f\n", mathx.Abs(-3.14))
	fmt.Printf("Ceil(3.14): %.0f\n", mathx.Ceil(3.14))
	fmt.Printf("Floor(3.14): %.0f\n", mathx.Floor(3.14))
	fmt.Printf("Pow(2, 3): %.0f\n", mathx.Pow(2, 3))
	fmt.Printf("Sqrt(16): %.0f\n", mathx.Sqrt(16))
	// Output:
	// Abs(-3.14): 3.14
	// Ceil(3.14): 4
	// Floor(3.14): 3
	// Pow(2, 3): 8
	// Sqrt(16): 4
}

func ExampleClamp() {
	// Utility functions
	value := 15.0
	min := 0.0
	max := 10.0

	clamped := mathx.Clamp(value, min, max)
	fmt.Printf("Clamp(%.0f, %.0f, %.0f): %.0f\n", value, min, max, clamped)

	// Linear interpolation
	lerp := mathx.Lerp(0.0, 10.0, 0.5)
	fmt.Printf("Lerp(0, 10, 0.5): %.1f\n", lerp)
	// Output:
	// Clamp(15, 0, 10): 10
	// Lerp(0, 10, 0.5): 5.0
}

func ExampleParseFloat() {
	// Safe string parsing
	value, err := mathx.ParseFloat("3.14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Parsed value: %.2f\n", value)
	// Output: Parsed value: 3.14
}
