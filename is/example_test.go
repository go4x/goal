package is_test

import (
	"fmt"

	"github.com/go4x/goal/is"
)

// ExampleNot demonstrates the Not function
func ExampleNot() {
	supported := true
	if is.Not(supported) {
		fmt.Println("not supported")
	} else {
		fmt.Println("supported")
	}
	// Output:
	// supported
}

// ExampleFalse demonstrates the False function
func ExampleFalse() {
	fmt.Println(is.False(true))  // false
	fmt.Println(is.False(false)) // true
	// Output:
	// false
	// true
}

// ExampleIsZero demonstrates the IsZero function
func ExampleZero() {
	fmt.Println(is.Zero(0))       // true
	fmt.Println(is.Zero(""))      // true
	fmt.Println(is.Zero(false))   // true
	fmt.Println(is.Zero(42))      // false
	fmt.Println(is.Zero("hello")) // false
	// Output:
	// true
	// true
	// true
	// false
	// false
}

// ExampleIsNil demonstrates the IsNil function
func ExampleNil() {
	var ptr *int
	fmt.Println(is.Nil(ptr))     // true
	fmt.Println(is.Nil(nil))     // true
	fmt.Println(is.Nil([]int{})) // false (empty slice is not nil)

	val := 42
	ptr = &val
	fmt.Println(is.Nil(ptr)) // false
	// Output:
	// true
	// true
	// false
	// false
}

// ExampleIsEmpty demonstrates the IsEmpty function
func ExampleEmpty() {
	fmt.Println(is.Empty(""))               // true
	fmt.Println(is.Empty([]int{}))          // true
	fmt.Println(is.Empty(map[string]int{})) // true
	fmt.Println(is.Empty(0))                // true
	fmt.Println(is.Empty("hello"))          // false
	fmt.Println(is.Empty([]int{1, 2}))      // false
	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
}

// ExampleEq demonstrates the Eq function
func ExampleEq() {
	fmt.Println(is.Eq(42, 42))        // true
	fmt.Println(is.Eq("hello", "hi")) // false
	fmt.Println(is.Eq(0, 0))          // true

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 4}
	fmt.Println(is.Eq(slice1, slice2)) // true
	fmt.Println(is.Eq(slice1, slice3)) // false

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	map3 := map[string]int{"a": 1, "b": 3}
	fmt.Println(is.Eq(map1, map2)) // true
	fmt.Println(is.Eq(map1, map3)) // false

	// Compare structs
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "John", Age: 30}
	p2 := Person{Name: "John", Age: 30}
	p3 := Person{Name: "Jane", Age: 30}
	fmt.Println(is.Eq(p1, p2)) // true
	fmt.Println(is.Eq(p1, p3)) // false

	type Func func() int
	var v1 Func = nil
	var v2 Func = nil
	fmt.Println(is.Eq(v1, v2)) // true
	f1 := func() int { return 1 }
	f2 := f1
	f3 := func() int { return 1 }
	fmt.Println(is.Eq(f1, f2)) // false
	fmt.Println(is.Eq(f1, f3)) // false

	// Output:
	// true
	// false
	// true
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// false
}
