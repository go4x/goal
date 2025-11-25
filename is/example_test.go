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

// ExampleEqual demonstrates the Equal function
func ExampleEqual() {
	fmt.Println(is.Equal(42, 42))        // true
	fmt.Println(is.Equal("hello", "hi")) // false
	fmt.Println(is.Equal(0, 0))          // true
	// Output:
	// true
	// false
	// true
}

// ExampleDeepEqual demonstrates the DeepEqual function
func ExampleDeepEqual() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	fmt.Println(is.DeepEqual(slice1, slice2)) // true

	map1 := map[string]int{"a": 1}
	map2 := map[string]int{"a": 1}
	fmt.Println(is.DeepEqual(map1, map2)) // true
	// Output:
	// true
	// true
}
