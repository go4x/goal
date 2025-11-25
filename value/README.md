# value

A comprehensive generic value handling package for Go that provides utilities for value manipulation, conditional logic, null/empty checks, and safe operations.

## Features

- **Generic Type Support**: Works with any comparable type
- **Null/Empty Handling**: Safe operations for nil and empty values
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
    if s == "" {
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
    if data == "" {
        return "", errors.New("data is empty")
    }
    
    // Process data safely
    result := value.Must(transformData(data))
    return result, nil
}

func transformData(data string) (string, error) {
    if data == "" {
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
    if user.Name == "" {
        return errors.New("name is required")
    }
    
    if user.Age <= 0 {
        return errors.New("age must be positive")
    }
    
    if user.Email == nil || *user.Email == "" {
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
        message != nil && *message != "",
        fmt.Sprintf("[%s] %s", status, *message),
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
            return item.Name != "" && item.Price != 0.0
        }).
        Map(func(item Item) Item {
            return Item{
                Name:  value.OrElse("Unknown", item.Name, ""),
                Price: value.OrElse(0.0, item.Price, 0.0),
                Category: value.IfElse(
                    item.Category != "",
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

### Coalescing Operations

- **Or/OrElse**: O(n) - Linear scan through values
- **Coalesce**: O(n) - Linear scan through values
- **SafeDeref**: O(1) - Direct pointer dereferencing

### Conditional Operations

- **IfElse/If**: O(1) - Direct conditional evaluation
- **When/WhenElse**: O(1) - Function call overhead
- **Must**: O(1) - Error check and panic if needed

### Best Practices

1. **Use Or/OrElse** for fallback value chains
2. **Use SafeDeref** for safe pointer operations
3. **Use Must** only when you're certain operations will succeed
4. **Use IfElse/If** for conditional value selection
5. **Use Coalesce** for pointer-based fallback chains

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

### Safe Operations

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `Must(value, error)` | Return value or panic on error | O(1) |
| `SafeDeref(pointer)` | Safe pointer dereferencing | O(1) |
| `SafeDerefDef(pointer, default)` | Safe pointer dereferencing with default | O(1) |
| `Value(interface{})` | Extract value from interface | O(1) |
| `Def(value, default)` | Return value or default if empty | O(1) |

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
    if data.Name == "" {
        return errors.New("name is required")
    }
    
    if data.Age <= 0 {
        return errors.New("age must be positive")
    }
    
    if data.Email == nil || *data.Email == "" {
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
        message != nil && *message != "",
        fmt.Sprintf("[%s] %s", status, *message),
        fmt.Sprintf("[%s] No message", status),
    )
}
```

### 4. Error Handling

```go
import "github.com/go4x/goal/value"

// Safe error handling
func ProcessData(data string) (string, error) {
    if data == "" {
        return "", errors.New("data is empty")
    }
    
    result := value.Must(transformData(data))
    return result, nil
}
```

## License

This package is part of the goal project and follows the same license terms.
