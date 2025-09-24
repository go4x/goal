# ptr - Pointer Utilities for Go

A comprehensive Go package providing safe and efficient pointer operations with full nil-safety support.

## Features

- **Type-safe**: Built with Go generics for complete type safety
- **Nil-safe**: All functions safely handle nil pointers without panics
- **Comprehensive**: Complete set of pointer operations for common use cases
- **Performance-optimized**: Efficient implementations with minimal allocations
- **Easy to use**: Simple, intuitive API design

## Installation

```bash
go get github.com/go4x/goal/ptr
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/ptr"
)

func main() {
    // Basic pointer operations
    value := 42
    p := ptr.To(value)        // Convert value to pointer
    result := ptr.From(p)     // Convert pointer to value
    
    // Safe dereferencing
    var nilPtr *int
    safeValue := ptr.Deref(nilPtr)  // Returns 0, no panic
    
    // Pointer validation
    if ptr.IsNil(nilPtr) {
        fmt.Println("Pointer is nil")
    }
    
    fmt.Printf("Value: %d, Result: %d, Safe: %d\n", value, result, safeValue)
}
```

## API Reference

### Basic Operations

#### `To[T any](v T) *T`
Converts a value to a pointer.

```go
ptr := ptr.To(42)  // *int with value 42
```

#### `From[T any](v *T) T`
Converts a pointer to a value. **Note**: This will panic if the pointer is nil.

```go
value := ptr.From(&someInt)  // int value
```

### Safe Operations

#### `IsNil[T any](v *T) bool`
Checks if a pointer is nil.

```go
var p *int
if ptr.IsNil(p) {
    fmt.Println("Pointer is nil")
}
```

#### `IsNotNil[T any](v *T) bool`
Checks if a pointer is not nil.

```go
p := ptr.To(42)
if ptr.IsNotNil(p) {
    fmt.Println("Pointer is not nil")
}
```

#### `Deref[T any](v *T) T`
Safely dereferences a pointer, returns zero value if nil.

```go
var nilPtr *int
value := ptr.Deref(nilPtr)  // Returns 0, no panic
```

#### `DerefOr[T any](v *T, defaultValue T) T`
Safely dereferences a pointer, returns specified value if nil.

```go
var nilPtr *int
value := ptr.DerefOr(nilPtr, 100)  // Returns 100
```

#### `ValueOr[T any](v *T, defaultValue T) T`
Returns the value if pointer is not nil, otherwise returns the default value.

```go
var nilPtr *int
value := ptr.ValueOr(nilPtr, 100)  // Returns 100
```

#### `ValueOrDefault[T any](v *T) T`
Returns the value if pointer is not nil, otherwise returns the zero value.

```go
var nilPtr *int
value := ptr.ValueOrDefault(nilPtr)  // Returns 0
```

### Comparison Operations

#### `Equal[T comparable](a, b *T) bool`
Compares two pointer values for equality (handles nil cases).

```go
p1 := ptr.To(42)
p2 := ptr.To(42)
p3 := ptr.To(100)
var nilPtr *int

fmt.Println(ptr.Equal(p1, p2))    // true
fmt.Println(ptr.Equal(p1, p3))    // false
fmt.Println(ptr.Equal(p1, nilPtr)) // false
```

#### `DeepEqual[T any](a, b *T) bool`
Performs deep comparison of two pointer values.

```go
type Person struct {
    Name string
    Age  int
}

p1 := &Person{Name: "Alice", Age: 30}
p2 := &Person{Name: "Alice", Age: 30}
p3 := &Person{Name: "Bob", Age: 25}

fmt.Println(ptr.DeepEqual(p1, p2))  // true
fmt.Println(ptr.DeepEqual(p1, p3))  // false
```

### Cloning Operations

#### `Clone[T any](v *T) *T`
Clones the value pointed to by the pointer.

```go
original := ptr.To(42)
cloned := ptr.Clone(original)
// cloned is a new pointer with the same value
```

#### `CloneSlice[T any](v []*T) []*T`
Clones a slice of pointers.

```go
slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
cloned := ptr.CloneSlice(slice)
// cloned is a new slice with cloned pointers
```

### Slice Operations

#### `ToSlice[T any](v []T) []*T`
Converts a slice of values to a slice of pointers.

```go
values := []int{1, 2, 3}
pointers := ptr.ToSlice(values)  // []*int
```

#### `FromSlice[T any](v []*T) []T`
Converts a slice of pointers to a slice of values.

```go
pointers := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
values := ptr.FromSlice(pointers)  // []int{1, 2, 3}
```

#### `Filter[T any](v []*T) []*T`
Filters a slice of pointers, returns slice of non-nil pointers.

```go
slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
filtered := ptr.Filter(slice)  // []*int with non-nil pointers
```

#### `FilterValues[T any](v []*T) []T`
Filters a slice of pointers, returns slice of values from non-nil pointers.

```go
slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
values := ptr.FilterValues(slice)  // []int{1, 3, 5}
```

#### `Map[T, U any](v []*T, fn func(*T) *U) []*U`
Maps a slice of pointers using the provided function.

```go
slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
strings := ptr.Map(slice, func(ptr *int) *string {
    if ptr == nil {
        return nil
    }
    return ptr.To(strconv.Itoa(*ptr))
})
```

#### `MapValues[T, U any](v []*T, fn func(T) U) []U`
Maps the values of a slice of pointers using the provided function.

```go
slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
strings := ptr.MapValues(slice, func(val int) string {
    return strconv.Itoa(val)
})  // []string{"1", "2", "3"}
```

### Query Operations

#### `Any[T any](v []*T) bool`
Checks if there are any non-nil pointers in the slice.

```go
slice := []*int{ptr.To(1), nil, ptr.To(3)}
hasAny := ptr.Any(slice)  // true
```

#### `All[T any](v []*T) bool`
Checks if all pointers in the slice are non-nil.

```go
slice1 := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
slice2 := []*int{ptr.To(1), nil, ptr.To(3)}

all1 := ptr.All(slice1)  // true
all2 := ptr.All(slice2)  // false
```

#### `Count[T any](v []*T) int`
Counts the number of non-nil pointers in the slice.

```go
slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
count := ptr.Count(slice)  // 3
```

#### `First[T any](v []*T) *T`
Returns the first non-nil pointer in the slice.

```go
slice := []*int{nil, ptr.To(2), ptr.To(3)}
first := ptr.First(slice)  // *int pointing to 2
```

#### `Last[T any](v []*T) *T`
Returns the last non-nil pointer in the slice.

```go
slice := []*int{ptr.To(1), ptr.To(2), nil, ptr.To(4)}
last := ptr.Last(slice)  // *int pointing to 4
```

### Modification Operations

#### `Set[T any](v *T, value T)`
Sets the value pointed to by the pointer.

```go
value := 42
ptr := &value
ptr.Set(ptr, 100)  // *ptr is now 100
```

#### `Zero[T any](v *T)`
Sets the value pointed to by the pointer to its zero value.

```go
value := 42
ptr := &value
ptr.Zero(ptr)  // *ptr is now 0
```

#### `Swap[T any](a, b *T)`
Swaps the values pointed to by two pointers.

```go
a := ptr.To(1)
b := ptr.To(2)
ptr.Swap(a, b)  // *a is now 2, *b is now 1
```

## Usage Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/ptr"
)

func main() {
    // Convert between values and pointers
    value := 42
    p := ptr.To(value)
    result := ptr.From(p)
    
    // Safe operations
    var nilPtr *int
    safeValue := ptr.Deref(nilPtr)  // 0, no panic
    
    fmt.Printf("Original: %d, Result: %d, Safe: %d\n", value, result, safeValue)
}
```

### Working with Slices

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/ptr"
)

func main() {
    // Create slice of pointers
    values := []int{1, 2, 3, 4, 5}
    pointers := ptr.ToSlice(values)
    
    // Filter out nil pointers (if any)
    valid := ptr.Filter(pointers)
    
    // Count non-nil pointers
    count := ptr.Count(pointers)
    
    // Get first non-nil pointer
    first := ptr.First(pointers)
    
    fmt.Printf("Valid pointers: %d, Count: %d, First: %d\n", 
        len(valid), count, ptr.Deref(first))
}
```

### Complex Usage with Structs

```go
package main

import (
    "fmt"
    "strings"
    "github.com/go4x/goal/ptr"
)

type User struct {
    ID    int
    Name  string
    Email *string // Optional field
}

func main() {
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
        if ptr.IsNotNil(user.Email) {
            emails = append(emails, ptr.Deref(user.Email))
        }
    }
    
    // Check if all user pointers are non-nil
    allValid := ptr.All(users)
    
    fmt.Printf("Valid users: %d\n", len(validUsers))
    fmt.Printf("Emails: %s\n", strings.Join(emails, ", "))
    fmt.Printf("All users valid: %t\n", allValid)
}
```

## Performance

The `ptr` package is designed for performance with minimal allocations:

- **Zero-copy operations**: Most operations don't allocate new memory
- **Efficient filtering**: Optimized algorithms for slice operations
- **Type-safe**: Compile-time type checking with generics
- **Nil-safe**: No runtime panics from nil pointer dereferences

## Testing

Run the tests:

```bash
go test ./ptr
```

Run with coverage:

```bash
go test ./ptr -cover
```

Run examples:

```bash
go test ./ptr -run Example
```

## License

This package is part of the `goal` project and follows the same license terms.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
