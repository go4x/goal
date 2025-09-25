# value

A comprehensive generic value handling package for Go that provides utilities for value manipulation, conditional logic, null/empty checks, and safe operations.

## Features

- **Generic Type Support**: Works with any comparable type
- **Null/Empty Handling**: Comprehensive checks for nil, empty, and zero values
- **Conditional Logic**: Functional approach to conditional operations
- **Safe Operations**: Panic-safe operations with proper error handling
- **Pointer Operations**: Safe dereferencing and pointer manipulation
- **Value Coalescing**: Fallback value operations
- **Type Safety**: Full compile-time type checking

## Installation

```bash
go get github.com/go4x/goal/value
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/value"
)

func main() {
    // Conditional logic
    result := value.IfElse(age >= 18, "adult", "minor")
    
    // Null/empty checks
    if value.IsNotEmpty(data) {
        fmt.Println("Data is not empty")
    }
    
    // Safe operations
    safeValue := value.Must(strconv.Atoi("123"))
    
    // Value coalescing
    fallback := value.Or("", "", "default")
}
```

## Core Functions

### Conditional Logic

#### IfElse - Conditional Operator

```go
import "github.com/go4x/goal/value"

// Basic conditional logic
result := value.IfElse(age >= 18, "adult", "minor")
fmt.Println(result) // "adult" if age >= 18, "minor" otherwise

// With different types
max := value.IfElse(a > b, a, b)
status := value.IfElse(isActive, "active", "inactive")
```

#### If - Single Condition

```go
import "github.com/go4x/goal/value"

// Return value if condition is true, otherwise zero value
result := value.If(age >= 18, "adult")
fmt.Println(result) // "adult" if age >= 18, "" otherwise

// With numbers
positive := value.If(x > 0, x)
fmt.Println(positive) // x if x > 0, 0 otherwise
```

#### When - Conditional with Function

```go
import "github.com/go4x/goal/value"

// Execute function if condition is true
result := value.When(age >= 18, func() string {
    return "You are an adult"
})
fmt.Println(result) // Function result if age >= 18, "" otherwise
```

#### WhenElse - Conditional with Two Functions

```go
import "github.com/go4x/goal/value"

// Execute different functions based on condition
result := value.WhenElse(age >= 18, 
    func() string { return "adult" },
    func() string { return "minor" },
)
fmt.Println(result) // "adult" or "minor" based on condition
```

### Value Coalescing

#### Or - First Non-Zero Value

```go
import "github.com/go4x/goal/value"

// Get first non-zero value
result := value.Or("", "", "fallback", "ignored")
fmt.Println(result) // "fallback"

// With numbers
number := value.Or(0, 0, 42, 100)
fmt.Println(number) // 42

// All zero values
empty := value.Or("", "", "")
fmt.Println(empty) // "" (zero value for string)
```

#### OrElse - First Non-Zero with Default

```go
import "github.com/go4x/goal/value"

// Get first non-zero value with default fallback
result := value.OrElse("default", "", "", "fallback")
fmt.Println(result) // "fallback"

// All zero values return default
defaultResult := value.OrElse("default", "", "", "")
fmt.Println(defaultResult) // "default"
```

#### Coalesce - Multiple Value Coalescing

```go
import "github.com/go4x/goal/value"

// Coalesce multiple values
result := value.Coalesce("", "first", "second")
fmt.Println(result) // "first"

// With different types
var p1, p2 *int
val := 42
p3 := &val
number := value.Coalesce(p1, p2, p3)
fmt.Println(number) // 42
```

### Null/Empty Checks

#### IsZero/IsNotZero - Zero Value Checks

```go
import "github.com/go4x/goal/value"

// Check for zero values
fmt.Println(value.IsZero(0))        // true
fmt.Println(value.IsZero(""))      // true
fmt.Println(value.IsZero(false))  // true
fmt.Println(value.IsZero(42))     // false
fmt.Println(value.IsZero("hello")) // false

// Check for non-zero values
fmt.Println(value.IsNotZero(42))     // true
fmt.Println(value.IsNotZero("hello")) // true
fmt.Println(value.IsNotZero(0))      // false
```

#### IsNil/IsNotNil - Nil Checks

```go
import "github.com/go4x/goal/value"

var ptr *int
fmt.Println(value.IsNil(ptr))        // true
fmt.Println(value.IsNil(nil))       // true
fmt.Println(value.IsNil([]int{}))   // false (empty slice is not nil)
fmt.Println(value.IsNil((*int)(nil))) // true

// Check for non-nil values
ptr = &42
fmt.Println(value.IsNotNil(ptr))     // true
fmt.Println(value.IsNotNil([]int{})) // true (empty slice is not nil)
fmt.Println(value.IsNotNil(nil))     // false
```

#### IsEmpty/IsNotEmpty - Empty Value Checks

```go
import "github.com/go4x/goal/value"

// Check for empty values
fmt.Println(value.IsEmpty(""))                    // true
fmt.Println(value.IsEmpty([]int{}))              // true
fmt.Println(value.IsEmpty(map[string]int{}))     // true
fmt.Println(value.IsEmpty(0))                   // true
fmt.Println(value.IsEmpty("hello"))              // false
fmt.Println(value.IsEmpty([]int{1, 2}))         // false

// Check for non-empty values
fmt.Println(value.IsNotEmpty("hello"))           // true
fmt.Println(value.IsNotEmpty([]int{1, 2}))      // true
fmt.Println(value.IsNotEmpty(""))                // false
```

### Safe Operations

#### Must - Panic on Error

```go
import (
    "strconv"
    "github.com/go4x/goal/value"
)

// Safe to use when you know the operation will succeed
result := value.Must(strconv.Atoi("123"))
fmt.Println(result) // 123

// This will panic if the string is not a valid integer
// result := value.Must(strconv.Atoi("invalid")) // panic!
```

#### SafeDeref - Safe Pointer Dereferencing

```go
import "github.com/go4x/goal/value"

var ptr *int
var val int = 42
ptr = &val

// Safe dereferencing
result := value.SafeDeref(ptr)
fmt.Println(result) // 42 true

// Safe dereferencing with default
result = value.SafeDerefDef(ptr, 0)
fmt.Println(result) // 42

// Nil pointer handling
var nilPtr *int
result = value.SafeDeref(nilPtr)
fmt.Println(result) // 0 false

result = value.SafeDerefDef(nilPtr, -1)
fmt.Println(result) // -1
```

### Value Operations

#### Value - Extract Value from Interface

```go
import "github.com/go4x/goal/value"

// Extract value from interface{}
var data interface{} = "hello"
result := value.Value(data)
fmt.Println(result) // "hello" true

// Handle nil interface
var nilData interface{}
result = value.Value(nilData)
fmt.Println(result) // nil false
```

#### Def - Default Value

```go
import "github.com/go4x/goal/value"

// Provide default value
result := value.Def("", "default")
fmt.Println(result) // "default"

// Non-empty value
result = value.Def("hello", "default")
fmt.Println(result) // "hello"
```

### Comparison Operations

#### Equal/NotEqual - Value Comparison

```go
import "github.com/go4x/goal/value"

// Compare values
fmt.Println(value.Equal(42, 42))     // true
fmt.Println(value.Equal("hello", "hello")) // true
fmt.Println(value.Equal(42, 43))    // false

// Check inequality
fmt.Println(value.NotEqual(42, 43))  // true
fmt.Println(value.NotEqual(42, 42))  // false
```

#### DeepEqual - Deep Value Comparison

```go
import "github.com/go4x/goal/value"

// Deep comparison using reflection
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{1, 2, 4}

fmt.Println(value.DeepEqual(slice1, slice2)) // true
fmt.Println(value.DeepEqual(slice1, slice3)) // false

// Compare structs
type Person struct {
    Name string
    Age  int
}

p1 := Person{Name: "John", Age: 30}
p2 := Person{Name: "John", Age: 30}
p3 := Person{Name: "Jane", Age: 30}

fmt.Println(value.DeepEqual(p1, p2)) // true
fmt.Println(value.DeepEqual(p1, p3)) // false
```

## Advanced Usage

### Configuration Handling

```go
import "github.com/go4x/goal/value"

// Configuration with fallbacks
type Config struct {
    Host     string
    Port     int
    Timeout  int
    Debug    bool
}

func LoadConfig() Config {
    return Config{
        Host:    value.OrElse("localhost", os.Getenv("HOST"), ""),
        Port:    value.OrElse(8080, parseInt(os.Getenv("PORT")), 0),
        Timeout: value.OrElse(30, parseInt(os.Getenv("TIMEOUT")), 0),
        Debug:   value.IfElse(os.Getenv("DEBUG") == "true", true, false),
    }
}

func parseInt(s string) int {
    if value.IsEmpty(s) {
        return 0
    }
    return value.Must(strconv.Atoi(s))
}
```

### Error Handling

```go
import "github.com/go4x/goal/value"

// Safe error handling
func ProcessData(data string) (string, error) {
    if value.IsEmpty(data) {
        return "", errors.New("data is empty")
    }
    
    // Process data safely
    result := value.Must(transformData(data))
    return result, nil
}

func transformData(data string) (string, error) {
    if value.IsEmpty(data) {
        return "", errors.New("empty data")
    }
    return strings.ToUpper(data), nil
}
```

### Data Validation

```go
import "github.com/go4x/goal/value"

// Validate user input
func ValidateUser(user User) error {
    if value.IsEmpty(user.Name) {
        return errors.New("name is required")
    }
    
    if value.IsZero(user.Age) || user.Age < 0 {
        return errors.New("age must be positive")
    }
    
    if value.IsNil(user.Email) || value.IsEmpty(*user.Email) {
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

### API Response Handling

```go
import "github.com/go4x/goal/value"

// Handle API responses with fallbacks
func ProcessAPIResponse(response *APIResponse) string {
    // Use coalescing for fallback values
    message := value.Coalesce(
        response.Message,
        response.Error,
        "No message available",
    )
    
    // Safe dereferencing
    status := value.SafeDerefDef(response.Status, "unknown")
    
    // Conditional formatting
    return value.IfElse(
        value.IsNotEmpty(message),
        fmt.Sprintf("[%s] %s", status, message),
        fmt.Sprintf("[%s] No message", status),
    )
}

type APIResponse struct {
    Message *string
    Error   *string
    Status  *string
}
```

### Functional Programming

```go
import "github.com/go4x/goal/value"

// Functional approach to data processing
func ProcessItems(items []Item) []Item {
    return slicex.From(items).
        Filter(func(item Item) bool {
            return value.IsNotEmpty(item.Name) && value.IsNotZero(item.Price)
        }).
        Map(func(item Item) Item {
            return Item{
                Name:  value.OrElse("Unknown", item.Name, ""),
                Price: value.OrElse(0.0, item.Price, 0.0),
                Category: value.IfElse(
                    value.IsNotEmpty(item.Category),
                    item.Category,
                    "uncategorized",
                ),
            }
        }).
        To()
}

type Item struct {
    Name     string
    Price    float64
    Category string
}
```

## Performance Considerations

### Zero-Value Checks

- **IsZero/IsNotZero**: O(1) - Direct comparison with zero value
- **IsEmpty/IsNotEmpty**: O(1) for most types, O(n) for slices/maps (length check)
- **IsNil/IsNotNil**: O(1) - Reflection-based nil check

### Coalescing Operations

- **Or/OrElse**: O(n) - Linear scan through values
- **Coalesce**: O(n) - Linear scan through values
- **SafeDeref**: O(1) - Direct pointer dereferencing

### Conditional Operations

- **IfElse/If**: O(1) - Direct conditional evaluation
- **When/WhenElse**: O(1) - Function call overhead
- **Must**: O(1) - Error check and panic if needed

### Best Practices

1. **Use IsZero/IsNotZero** for simple zero-value checks
2. **Use IsEmpty/IsNotEmpty** for comprehensive empty checks
3. **Use Or/OrElse** for fallback value chains
4. **Use SafeDeref** for safe pointer operations
5. **Use Must** only when you're certain operations will succeed

## Thread Safety

⚠️ **Important**: All functions in this package are **thread-safe** and can be called concurrently from multiple goroutines. However, the underlying data being operated on must be thread-safe if accessed concurrently.

## API Reference

### Conditional Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `IfElse(condition, v1, v2)` | Return v1 if condition is true, v2 otherwise | O(1) |
| `If(condition, value)` | Return value if condition is true, zero value otherwise | O(1) |
| `When(condition, func)` | Execute function if condition is true | O(1) |
| `WhenElse(condition, func1, func2)` | Execute func1 if condition is true, func2 otherwise | O(1) |

### Coalescing Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `Or(values...)` | Return first non-zero value | O(n) |
| `OrElse(default, values...)` | Return first non-zero value or default | O(n) |
| `Coalesce(values...)` | Return first non-zero value | O(n) |
| `CoalesceValue(values...)` | Return first non-zero value | O(n) |
| `CoalesceValueDef(default, values...)` | Return first non-zero value or default | O(n) |

### Check Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `IsZero(value)` | Check if value is zero | O(1) |
| `IsNotZero(value)` | Check if value is not zero | O(1) |
| `IsNil(value)` | Check if value is nil | O(1) |
| `IsNotNil(value)` | Check if value is not nil | O(1) |
| `IsEmpty(value)` | Check if value is empty | O(1) |
| `IsNotEmpty(value)` | Check if value is not empty | O(1) |

### Safe Operations

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `Must(value, error)` | Return value or panic on error | O(1) |
| `SafeDeref(pointer)` | Safe pointer dereferencing | O(1) |
| `SafeDerefDef(pointer, default)` | Safe pointer dereferencing with default | O(1) |
| `Value(interface{})` | Extract value from interface | O(1) |
| `Def(value, default)` | Return value or default if empty | O(1) |

### Comparison Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `Equal(a, b)` | Check if values are equal | O(1) |
| `NotEqual(a, b)` | Check if values are not equal | O(1) |
| `DeepEqual(a, b)` | Deep comparison using reflection | O(n) |

## Use Cases

### 1. Configuration Management

```go
import "github.com/go4x/goal/value"

// Environment-based configuration
func LoadConfig() Config {
    return Config{
        Host:    value.OrElse("localhost", os.Getenv("HOST"), ""),
        Port:    value.OrElse(8080, parseInt(os.Getenv("PORT")), 0),
        Debug:   value.IfElse(os.Getenv("DEBUG") == "true", true, false),
        Timeout: value.OrElse(30, parseInt(os.Getenv("TIMEOUT")), 0),
    }
}
```

### 2. Data Validation

```go
import "github.com/go4x/goal/value"

// Validate input data
func ValidateInput(data InputData) error {
    if value.IsEmpty(data.Name) {
        return errors.New("name is required")
    }
    
    if value.IsZero(data.Age) || data.Age < 0 {
        return errors.New("age must be positive")
    }
    
    if value.IsNil(data.Email) || value.IsEmpty(*data.Email) {
        return errors.New("email is required")
    }
    
    return nil
}
```

### 3. API Response Processing

```go
import "github.com/go4x/goal/value"

// Process API responses with fallbacks
func ProcessResponse(response *APIResponse) string {
    message := value.Coalesce(
        response.Message,
        response.Error,
        "No message available",
    )
    
    status := value.SafeDerefDef(response.Status, "unknown")
    
    return value.IfElse(
        value.IsNotEmpty(message),
        fmt.Sprintf("[%s] %s", status, message),
        fmt.Sprintf("[%s] No message", status),
    )
}
```

### 4. Error Handling

```go
import "github.com/go4x/goal/value"

// Safe error handling
func ProcessData(data string) (string, error) {
    if value.IsEmpty(data) {
        return "", errors.New("data is empty")
    }
    
    result := value.Must(transformData(data))
    return result, nil
}
```

## License

This package is part of the goal project and follows the same license terms.
