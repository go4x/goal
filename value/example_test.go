package value_test

import (
	"fmt"
	"strconv"

	"github.com/go4x/goal/value"
)

// ExampleMust demonstrates the Must function
func ExampleMust() {
	// Safe to use when you know the operation will succeed
	result := value.Must(strconv.Atoi("123"))
	fmt.Println(result)
	// Output: 123
}

// ExampleIfElse demonstrates the IfElse function
func ExampleIfElse() {
	age := 20
	category := value.IfElse(age >= 18, "adult", "minor")
	fmt.Println(category)
	// Output: adult
}

// ExampleOr demonstrates the Or function
func ExampleOr() {
	// Get the first non-empty string
	result := value.Or("", "", "fallback", "ignored")
	fmt.Println(result)
	// Output: fallback
}

// ExampleOrElse demonstrates the OrElse function
func ExampleOrElse() {
	// Get the first non-empty string with a default
	result := value.OrElse("default", "", "", "fallback")
	fmt.Println(result)
	// Output: fallback
}

// ExampleIf demonstrates the If function
func ExampleIf() {
	age := 20
	result := value.If(age >= 18, "adult")
	fmt.Println(result)
	// Output: adult
}

// ExampleWhen demonstrates the When function
func ExampleWhen() {
	// Return the value if it's positive
	result := value.When(42, func(x int) bool { return x > 0 })
	fmt.Println(result)
	// Output: 42
}

// ExampleWhenElse demonstrates the WhenElse function
func ExampleWhenElse() {
	// Return different values based on condition
	result := value.WhenElse(42, func(x int) bool { return x > 0 }, 1, -1)
	fmt.Println(result)
	// Output: 1
}

// ExampleCoalesce demonstrates the Coalesce function
func ExampleCoalesce() {
	var p1, p2 *int
	val := 42
	p3 := &val
	result := value.Coalesce(p1, p2, p3)
	if result != nil {
		fmt.Println(*result)
	}
	// Output: 42
}

// ExampleCoalesceValue demonstrates the CoalesceValue function
func ExampleCoalesceValue() {
	var p1, p2 *int
	val := 42
	p3 := &val
	result := value.CoalesceValue(p1, p2, p3)
	fmt.Println(result)
	// Output: 42
}

// ExampleCoalesceValueDef demonstrates the CoalesceValueDef function
func ExampleCoalesceValueDef() {
	var p1, p2 *int
	result := value.CoalesceValueDef(999, p1, p2)
	fmt.Println(result)
	// Output: 999
}

// ExampleSafeDeref demonstrates the SafeDeref function
func ExampleSafeDeref() {
	var ptr *int
	result := value.SafeDeref(ptr)
	fmt.Println(result)

	val := 42
	ptr = &val
	result = value.SafeDeref(ptr)
	fmt.Println(result)
	// Output:
	// 0
	// 42
}

// ExampleSafeDerefDef demonstrates the SafeDerefDef function
func ExampleSafeDerefDef() {
	var ptr *int
	result := value.SafeDerefDef(ptr, 100)
	fmt.Println(result)

	val := 42
	ptr = &val
	result = value.SafeDerefDef(ptr, 100)
	fmt.Println(result)
	// Output:
	// 100
	// 42
}
