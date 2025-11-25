# is

A comprehensive value checking and comparison package for Go that provides utilities for boolean operations, zero/nil/empty checks, and value comparisons.

## Features

- **Boolean Operations**: Logical negation and boolean value checks
- **Zero Value Checks**: Check if values are zero for their type
- **Nil Checks**: Comprehensive nil checking for reference types
- **Empty Checks**: Check for empty values across different types
- **Value Comparison**: Equality and deep equality checks
- **Type Safety**: Full compile-time type checking with generics

## Installation

```bash
go get github.com/go4x/goal/is
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/is"
)

func main() {
    // Zero value checks
    if is.Zero(0) {
        fmt.Println("Value is zero")
    }
    
    // Nil checks
    var ptr *int
    if is.Nil(ptr) {
        fmt.Println("Pointer is nil")
    }
    
    // Empty checks
    if is.Empty("") {
        fmt.Println("String is empty")
    }
    
    // Value comparison
    if is.Equal(42, 42) {
        fmt.Println("Values are equal")
    }
}
```

## Core Functions

### Boolean Operations

#### Not - Logical Negation

```go
import "github.com/go4x/goal/is"

// Negate a boolean value
fmt.Println(is.Not(true))   // false
fmt.Println(is.Not(false))  // true

// Use in conditions
supported := true
if is.Not(supported) {
    fmt.Println("not supported")
} else {
    fmt.Println("supported")
}
```

#### True - Check if True

```go
import "github.com/go4x/goal/is"

// Check if value is true
fmt.Println(is.True(true))   // true
fmt.Println(is.True(false))  // false
```

#### False - Check if False

```go
import "github.com/go4x/goal/is"

// Check if value is false
fmt.Println(is.False(true))   // false
fmt.Println(is.False(false))  // true
```

### Zero Value Checks

#### Zero - Zero Value Check

```go
import "github.com/go4x/goal/is"

// Check for zero values
fmt.Println(is.Zero(0))        // true
fmt.Println(is.Zero(""))      // true
fmt.Println(is.Zero(false))   // true
fmt.Println(is.Zero(42))      // false
fmt.Println(is.Zero("hello")) // false

// With different types
fmt.Println(is.Zero(0.0))     // true
fmt.Println(is.Zero(0.0))     // true
fmt.Println(is.Zero([]int{}))  // false (empty slice is not zero)
```

#### NotZero - Non-Zero Value Check

```go
import "github.com/go4x/goal/is"

// Check for non-zero values
fmt.Println(is.NotZero(42))     // true
fmt.Println(is.NotZero("hello")) // true
fmt.Println(is.NotZero(0))      // false
fmt.Println(is.NotZero(""))     // false
```

### Nil Checks

#### Nil - Nil Check

```go
import "github.com/go4x/goal/is"

var ptr *int
fmt.Println(is.Nil(ptr))        // true
fmt.Println(is.Nil(nil))        // true
fmt.Println(is.Nil([]int{}))    // false (empty slice is not nil)
fmt.Println(is.Nil((*int)(nil))) // true

// With different reference types
var m map[string]int
fmt.Println(is.Nil(m))          // true

var ch chan int
fmt.Println(is.Nil(ch))         // true
```

#### NotNil - Non-Nil Check

```go
import "github.com/go4x/goal/is"

// Check for non-nil values
ptr := &42
fmt.Println(is.NotNil(ptr))     // true
fmt.Println(is.NotNil([]int{})) // true (empty slice is not nil)
fmt.Println(is.NotNil(nil))     // false

// With maps and channels
m := make(map[string]int)
fmt.Println(is.NotNil(m))       // true

ch := make(chan int)
fmt.Println(is.NotNil(ch))      // true
```

### Empty Checks

#### Empty - Empty Value Check

```go
import "github.com/go4x/goal/is"

// Check for empty values
fmt.Println(is.Empty(""))                    // true
fmt.Println(is.Empty([]int{}))              // true
fmt.Println(is.Empty(map[string]int{}))     // true
fmt.Println(is.Empty(0))                    // true
fmt.Println(is.Empty("hello"))              // false
fmt.Println(is.Empty([]int{1, 2}))          // false

// With different types
fmt.Println(is.Empty(nil))                  // true
var ptr *int
fmt.Println(is.Empty(ptr))                  // true
```

#### NotEmpty - Non-Empty Value Check

```go
import "github.com/go4x/goal/is"

// Check for non-empty values
fmt.Println(is.NotEmpty("hello"))           // true
fmt.Println(is.NotEmpty([]int{1, 2}))       // true
fmt.Println(is.NotEmpty(42))                // true
fmt.Println(is.NotEmpty(""))                // false
fmt.Println(is.NotEmpty([]int{}))           // false
```

### Value Comparison

#### Equal - Equality Check

```go
import "github.com/go4x/goal/is"

// Compare values
fmt.Println(is.Equal(42, 42))              // true
fmt.Println(is.Equal("hello", "hello"))    // true
fmt.Println(is.Equal(42, 43))              // false
fmt.Println(is.Equal("hello", "hi"))      // false

// With different types
fmt.Println(is.Equal(true, true))          // true
fmt.Println(is.Equal(3.14, 3.14))          // true
```

#### NotEqual - Inequality Check

```go
import "github.com/go4x/goal/is"

// Check inequality
fmt.Println(is.NotEqual(42, 43))          // true
fmt.Println(is.NotEqual("hello", "hi"))   // true
fmt.Println(is.NotEqual(42, 42))          // false
fmt.Println(is.NotEqual("hello", "hello")) // false
```

#### DeepEqual - Deep Equality Check

```go
import "github.com/go4x/goal/is"

// Deep comparison using reflection
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{1, 2, 4}

fmt.Println(is.DeepEqual(slice1, slice2)) // true
fmt.Println(is.DeepEqual(slice1, slice3)) // false

// Compare maps
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"a": 1, "b": 2}
map3 := map[string]int{"a": 1, "b": 3}

fmt.Println(is.DeepEqual(map1, map2)) // true
fmt.Println(is.DeepEqual(map1, map3)) // false

// Compare structs
type Person struct {
    Name string
    Age  int
}

p1 := Person{Name: "John", Age: 30}
p2 := Person{Name: "John", Age: 30}
p3 := Person{Name: "Jane", Age: 30}

fmt.Println(is.DeepEqual(p1, p2)) // true
fmt.Println(is.DeepEqual(p1, p3)) // false
```

## Advanced Usage

### Data Validation

```go
import "github.com/go4x/goal/is"

// Validate user input
func ValidateUser(user User) error {
    if is.Empty(user.Name) {
        return errors.New("name is required")
    }
    
    if is.Zero(user.Age) || user.Age < 0 {
        return errors.New("age must be positive")
    }
    
    if is.Nil(user.Email) || is.Empty(*user.Email) {
        return errors.New("email is required")
    }
    
    return nil
}

type User struct {
    Name  string
    Age   int
    Email *string
}
```

### Configuration Validation

```go
import "github.com/go4x/goal/is"

// Validate configuration
func ValidateConfig(config Config) error {
    if is.Empty(config.Host) {
        return errors.New("host is required")
    }
    
    if is.Zero(config.Port) || config.Port <= 0 {
        return errors.New("port must be positive")
    }
    
    if is.Nil(config.Database) {
        return errors.New("database config is required")
    }
    
    return nil
}
```

### API Response Processing

```go
import "github.com/go4x/goal/is"

// Process API responses
func ProcessResponse(response *APIResponse) error {
    if is.Nil(response) {
        return errors.New("response is nil")
    }
    
    if is.Empty(response.Data) {
        return errors.New("response data is empty")
    }
    
    // Check if response is successful
    if is.NotEqual(response.Status, "success") {
        return fmt.Errorf("unexpected status: %s", response.Status)
    }
    
    return nil
}
```

### Functional Programming

```go
import "github.com/go4x/goal/is"

// Filter items based on checks
func FilterValidItems(items []Item) []Item {
    var valid []Item
    for _, item := range items {
        if is.NotEmpty(item.Name) && is.NotZero(item.Price) {
            valid = append(valid, item)
        }
    }
    return valid
}

// Compare collections
func AreCollectionsEqual(a, b []int) bool {
    return is.DeepEqual(a, b)
}
```

## Performance Considerations

### Zero Value Checks

- **Zero/NotZero**: O(1) - Direct comparison with zero value
- **Empty/NotEmpty**: O(1) for most types, O(n) for slices/maps (length check)
- **Nil/NotNil**: O(1) - Reflection-based nil check

### Comparison Operations

- **Equal/NotEqual**: O(1) - Direct equality check
- **DeepEqual**: O(n) - Reflection-based deep comparison

### Best Practices

1. **Use Zero/NotZero** for simple zero-value checks with comparable types
2. **Use Empty/NotEmpty** for comprehensive empty checks across all types
3. **Use Nil/NotNil** for reference type nil checks
4. **Use Equal/NotEqual** for simple value comparisons
5. **Use DeepEqual** for complex nested structure comparisons

## Thread Safety

⚠️ **Important**: All functions in this package are **thread-safe** and can be called concurrently from multiple goroutines. However, the underlying data being checked must be thread-safe if accessed concurrently.

## API Reference

### Boolean Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `Not(v)` | Return logical negation of boolean | O(1) |
| `True(v)` | Return true if value is true | O(1) |
| `False(v)` | Return true if value is false | O(1) |

### Check Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `Zero(value)` | Check if value is zero | O(1) |
| `NotZero(value)` | Check if value is not zero | O(1) |
| `Nil(value)` | Check if value is nil | O(1) |
| `NotNil(value)` | Check if value is not nil | O(1) |
| `Empty(value)` | Check if value is empty | O(1) |
| `NotEmpty(value)` | Check if value is not empty | O(1) |

### Comparison Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `Equal(a, b)` | Check if values are equal | O(1) |
| `NotEqual(a, b)` | Check if values are not equal | O(1) |
| `DeepEqual(a, b)` | Deep comparison using reflection | O(n) |

## Use Cases

### 1. Input Validation

```go
import "github.com/go4x/goal/is"

// Validate input data
func ValidateInput(data InputData) error {
    if is.Empty(data.Name) {
        return errors.New("name is required")
    }
    
    if is.Zero(data.Age) || data.Age < 0 {
        return errors.New("age must be positive")
    }
    
    if is.Nil(data.Email) || is.Empty(*data.Email) {
        return errors.New("email is required")
    }
    
    return nil
}
```

### 2. Configuration Checking

```go
import "github.com/go4x/goal/is"

// Check configuration values
func CheckConfig(config Config) error {
    if is.Empty(config.Host) {
        return errors.New("host cannot be empty")
    }
    
    if is.Zero(config.Port) {
        return errors.New("port must be specified")
    }
    
    return nil
}
```

### 3. Data Comparison

```go
import "github.com/go4x/goal/is"

// Compare data structures
func CompareData(old, new Data) bool {
    return is.DeepEqual(old, new)
}

// Check if values changed
func HasChanged(old, new string) bool {
    return is.NotEqual(old, new)
}
```

### 4. Conditional Logic

```go
import "github.com/go4x/goal/is"

// Use in conditional logic
func ProcessData(data string) error {
    if is.Empty(data) {
        return errors.New("data is empty")
    }
    
    if is.NotEmpty(data) {
        // Process non-empty data
        return process(data)
    }
    
    return nil
}
```

## License

This package is part of the goal project and follows the same license terms.

