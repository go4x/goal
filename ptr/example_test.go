package ptr_test

import (
	"fmt"
	"strings"

	ptr "github.com/go4x/goal/ptr"
)

func ExampleTo() {
	// Convert value to pointer
	value := 42
	ptr := ptr.To(value)
	fmt.Printf("Value: %d, ptr: %d\n", value, *ptr)
	// Output: Value: 42, ptr: 42
}

func ExampleFrom() {
	// Convert pointer to value
	value := 100
	p := &value
	result := ptr.From(p)
	fmt.Printf("ptr: %d, Value: %d\n", *p, result)
	// Output: ptr: 100, Value: 100
}

func ExampleIsNil() {
	// Check if pointer is nil
	var nilPtr *int
	notNilPtr := ptr.To(42)

	fmt.Printf("nilPtr is nil: %t\n", ptr.IsNil(nilPtr))
	fmt.Printf("notNilPtr is nil: %t\n", ptr.IsNil(notNilPtr))
	// Output:
	// nilPtr is nil: true
	// notNilPtr is nil: false
}

func ExampleValueOr() {
	// Safely get pointer value, return default if nil
	var nilPtr *int
	notNilPtr := ptr.To(42)

	fmt.Printf("nilPtr or default: %d\n", ptr.ValueOr(nilPtr, 100))
	fmt.Printf("notNilPtr or default: %d\n", ptr.ValueOr(notNilPtr, 100))
	// Output:
	// nilPtr or default: 100
	// notNilPtr or default: 42
}

func ExampleDeref() {
	// Safely dereference pointer
	var nilPtr *int
	notNilPtr := ptr.To(42)

	fmt.Printf("nilPtr deref: %d\n", ptr.Deref(nilPtr))
	fmt.Printf("notNilPtr deref: %d\n", ptr.Deref(notNilPtr))
	// Output:
	// nilPtr deref: 0
	// notNilPtr deref: 42
}

func ExampleEqual() {
	// Compare two pointer values
	ptr1 := ptr.To(42)
	ptr2 := ptr.To(42)
	ptr3 := ptr.To(100)
	var nilPtr *int

	fmt.Printf("ptr1 == ptr2: %t\n", ptr.Equal(ptr1, ptr2))
	fmt.Printf("ptr1 == ptr3: %t\n", ptr.Equal(ptr1, ptr3))
	fmt.Printf("ptr1 == nil: %t\n", ptr.Equal(ptr1, nilPtr))
	// Output:
	// ptr1 == ptr2: true
	// ptr1 == ptr3: false
	// ptr1 == nil: false
}

func ExampleClone() {
	// Clone the value pointed to by pointer
	original := ptr.To(42)
	cloned := ptr.Clone(original)

	fmt.Printf("Original: %d\n", *original)
	fmt.Printf("Cloned: %d\n", *cloned)
	fmt.Printf("Same ptr: %t\n", original == cloned)
	// Output:
	// Original: 42
	// Cloned: 42
	// Same ptr: false
}

func ExampleFilter() {
	// Filter pointer slice, return non-nil pointers
	slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
	filtered := ptr.Filter(slice)

	fmt.Printf("Original length: %d\n", len(slice))
	fmt.Printf("Filtered length: %d\n", len(filtered))
	fmt.Printf("Filtered values: ")
	for i, ptr := range filtered {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(*ptr)
	}
	fmt.Println()
	// Output:
	// Original length: 5
	// Filtered length: 3
	// Filtered values: 1, 3, 5
}

func ExampleFilterValues() {
	// Filter pointer slice, return non-nil pointers的值
	slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
	values := ptr.FilterValues(slice)

	fmt.Printf("Values: %v\n", values)
	// Output: Values: [1 3 5]
}

func ExampleMap() {
	// Map pointer slice using provided function
	slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
	mapped := ptr.Map(slice, func(p *int) *string {
		if p == nil {
			return nil
		}
		return ptr.To(string(rune(*p + 64)))
	})

	fmt.Printf("Mapped: ")
	for i, ptr := range mapped {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(*ptr)
	}
	fmt.Println()
	// Output: Mapped: A, B, C
}

func ExampleMapValues() {
	// Map values of pointer slice using provided function
	slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
	values := ptr.MapValues(slice, func(val int) string {
		return string(rune(val + 64))
	})

	fmt.Printf("Values: %v\n", values)
	// Output: Values: [A B C]
}

func ExampleAny() {
	// Check if there are any non-nil pointers in the slice
	slice1 := []*int{ptr.To(1), nil, ptr.To(3)}
	slice2 := []*int{nil, nil, nil}

	fmt.Printf("slice1 has any: %t\n", ptr.Any(slice1))
	fmt.Printf("slice2 has any: %t\n", ptr.Any(slice2))
	// Output:
	// slice1 has any: true
	// slice2 has any: false
}

func ExampleAll() {
	// Check if all pointers in the slice are non-nil
	slice1 := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
	slice2 := []*int{ptr.To(1), nil, ptr.To(3)}

	fmt.Printf("slice1 all non-nil: %t\n", ptr.All(slice1))
	fmt.Printf("slice2 all non-nil: %t\n", ptr.All(slice2))
	// Output:
	// slice1 all non-nil: true
	// slice2 all non-nil: false
}

func ExampleCount() {
	// Count the number of non-nil pointers in the slice
	slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
	count := ptr.Count(slice)

	fmt.Printf("Non-nil count: %d\n", count)
	// Output: Non-nil count: 3
}

func ExampleFirst() {
	// Return the first non-nil pointer in the slice
	slice := []*int{nil, ptr.To(2), ptr.To(3)}
	first := ptr.First(slice)

	if first != nil {
		fmt.Printf("First non-nil: %d\n", *first)
	} else {
		fmt.Println("No non-nil ptr found")
	}
	// Output: First non-nil: 2
}

func ExampleLast() {
	// Return the last non-nil pointer in the slice
	slice := []*int{ptr.To(1), ptr.To(2), nil, ptr.To(4)}
	last := ptr.Last(slice)

	if last != nil {
		fmt.Printf("Last non-nil: %d\n", *last)
	} else {
		fmt.Println("No non-nil ptr found")
	}
	// Output: Last non-nil: 4
}

func ExampleSet() {
	// Set the value pointed to by the pointer
	value := 42
	p := &value

	fmt.Printf("Before: %d\n", *p)
	ptr.Set(p, 100)
	fmt.Printf("After: %d\n", *p)
	// Output:
	// Before: 42
	// After: 100
}

func ExampleZero() {
	// Set the value pointed to by the pointer to zero value
	value := 42
	p := &value

	fmt.Printf("Before: %d\n", *p)
	ptr.Zero(p)
	fmt.Printf("After: %d\n", *p)
	// Output:
	// Before: 42
	// After: 0
}

func ExampleSwap() {
	// Swap the values pointed to by two pointers
	a := ptr.To(1)
	b := ptr.To(2)

	fmt.Printf("Before: a=%d, b=%d\n", *a, *b)
	ptr.Swap(a, b)
	fmt.Printf("After: a=%d, b=%d\n", *a, *b)
	// Output:
	// Before: a=1, b=2
	// After: a=2, b=1
}

func ExampleDeepEqual() {
	// Perform deep comparison of two pointer values
	type Person struct {
		Name string
		Age  int
	}

	p1 := &Person{Name: "Alice", Age: 30}
	p2 := &Person{Name: "Alice", Age: 30}
	p3 := &Person{Name: "Bob", Age: 25}

	fmt.Printf("p1 == p2: %t\n", ptr.DeepEqual(p1, p2))
	fmt.Printf("p1 == p3: %t\n", ptr.DeepEqual(p1, p3))
	// Output:
	// p1 == p2: true
	// p1 == p3: false
}

func ExampleFilter_advanced() {
	// Complex usage scenario: handling user data
	type User struct {
		ID    int
		Name  string
		Email *string // Optional field
	}

	users := []*User{
		{ID: 1, Name: "Alice", Email: ptr.To("alice@example.com")},
		{ID: 2, Name: "Bob", Email: nil}, // No email
		{ID: 3, Name: "Charlie", Email: ptr.To("charlie@example.com")},
	}

	// Filter out non-nil user pointers
	validUsers := ptr.Filter(users)

	// Get all email addresses (filter out nil emails)
	emails := []string{}
	for _, user := range validUsers {
		if user.Email != nil {
			emails = append(emails, *user.Email)
		}
	}

	// Check if all user pointers are non-nil
	allValid := ptr.All(users)

	fmt.Printf("Valid users: %d\n", len(validUsers))
	fmt.Printf("Emails: %s\n", strings.Join(emails, ", "))
	fmt.Printf("All users valid: %t\n", allValid)
	// Output:
	// Valid users: 3
	// Emails: alice@example.com, charlie@example.com
	// All users valid: true
}
