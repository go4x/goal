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

// ExampleIsZero demonstrates the IsZero function
func ExampleIsZero() {
	fmt.Println(value.IsZero(0))       // true
	fmt.Println(value.IsZero(""))      // true
	fmt.Println(value.IsZero(false))   // true
	fmt.Println(value.IsZero(42))      // false
	fmt.Println(value.IsZero("hello")) // false
	// Output:
	// true
	// true
	// true
	// false
	// false
}

// ExampleIsNil demonstrates the IsNil function
func ExampleIsNil() {
	var ptr *int
	fmt.Println(value.IsNil(ptr))     // true
	fmt.Println(value.IsNil(nil))     // true
	fmt.Println(value.IsNil([]int{})) // false (empty slice is not nil)

	val := 42
	ptr = &val
	fmt.Println(value.IsNil(ptr)) // false
	// Output:
	// true
	// true
	// false
	// false
}

// ExampleIsEmpty demonstrates the IsEmpty function
func ExampleIsEmpty() {
	fmt.Println(value.IsEmpty(""))               // true
	fmt.Println(value.IsEmpty([]int{}))          // true
	fmt.Println(value.IsEmpty(map[string]int{})) // true
	fmt.Println(value.IsEmpty(0))                // true
	fmt.Println(value.IsEmpty("hello"))          // false
	fmt.Println(value.IsEmpty([]int{1, 2}))      // false
	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
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

// ExampleEqual demonstrates the Equal function
func ExampleEqual() {
	fmt.Println(value.Equal(42, 42))        // true
	fmt.Println(value.Equal("hello", "hi")) // false
	fmt.Println(value.Equal(0, 0))          // true
	// Output:
	// true
	// false
	// true
}

// ExampleDeepEqual demonstrates the DeepEqual function
func ExampleDeepEqual() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	fmt.Println(value.DeepEqual(slice1, slice2)) // true

	map1 := map[string]int{"a": 1}
	map2 := map[string]int{"a": 1}
	fmt.Println(value.DeepEqual(map1, map2)) // true
	// Output:
	// true
	// true
}
