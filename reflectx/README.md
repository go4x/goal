# ReflectX

A comprehensive Go reflection utility package that provides a wide range of reflection-based operations for type checking, conversion, struct field manipulation, value operations, method calling, and interface/generic type inspection.

## Features

- **Type Checking**: Check if values are nil, zero, or specific types
- **Type Information**: Get detailed type information including name, kind, size, alignment
- **Type Conversion**: Safe type conversion with error handling
- **Struct Field Operations**: Get/set field values, names, tags, and detailed information
- **Value Operations**: Manipulate slices, maps, arrays with index-based operations
- **Method Operations**: Call methods, check method existence, get method information
- **Interface Utilities**: Check interface implementation, get interface methods
- **Generic Utilities**: Check generic types and get generic type information
- **Performance Optimized**: Efficient implementations with comprehensive benchmarks

## Installation

```bash
go get github.com/go4x/goal/reflectx
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/reflectx"
)

func main() {
    // Check if value is nil
    var ptr *int
    fmt.Println("Is nil:", reflectx.IsNil(ptr)) // true
    
    // Check if value is zero
    fmt.Println("Is zero:", reflectx.IsZero(0)) // true
    
    // Get type information
    fmt.Println("Type name:", reflectx.GetTypeName("hello")) // string
    
    // Convert types
    result, err := reflectx.Convert(42, reflect.TypeOf(0.0))
    if err == nil {
        fmt.Println("Converted:", result) // 42
    }
}
```

## API Reference

### Type Checking

#### `IsNil(v interface{}) bool`
Checks if a value is nil.

```go
var ptr *int
fmt.Println(reflectx.IsNil(ptr)) // true
fmt.Println(reflectx.IsNil(42))  // false
```

#### `IsZero(v interface{}) bool`
Checks if a value is the zero value for its type.

```go
fmt.Println(reflectx.IsZero(0))     // true
fmt.Println(reflectx.IsZero(""))    // true
fmt.Println(reflectx.IsZero(42))    // false
```

#### Type Checkers
- `IsPointer(v interface{}) bool`
- `IsSlice(v interface{}) bool`
- `IsMap(v interface{}) bool`
- `IsStruct(v interface{}) bool`
- `IsInterface(v interface{}) bool`
- `IsFunc(v interface{}) bool`
- `IsChannel(v interface{}) bool`
- `IsArray(v interface{}) bool`
- `IsString(v interface{}) bool`
- `IsInt(v interface{}) bool`
- `IsUint(v interface{}) bool`
- `IsFloat(v interface{}) bool`
- `IsBool(v interface{}) bool`
- `IsComplex(v interface{}) bool`

### Type Information

#### `GetTypeName(v interface{}) string`
Returns the type name of a value.

```go
type User struct{ Name string }
fmt.Println(reflectx.GetTypeName(User{})) // "main.User"
```

#### `GetKind(v interface{}) string`
Returns the kind of a value.

```go
fmt.Println(reflectx.GetKind([]int{1, 2, 3})) // "slice"
```

#### `GetSize(v interface{}) uintptr`
Returns the size of a value in bytes.

```go
fmt.Println(reflectx.GetSize(int(42))) // 8
```

### Type Conversion

#### `Convert(v interface{}, targetType reflect.Type) (interface{}, error)`
Converts a value to the target type.

```go
result, err := reflectx.Convert(42, reflect.TypeOf(0.0))
if err == nil {
    fmt.Println(result) // 42.0
}
```

#### `ConvertTo(v interface{}, target interface{}) (interface{}, error)`
Converts a value to the type of the target.

```go
var target float64
result, err := reflectx.ConvertTo(42, target)
if err == nil {
    fmt.Println(result) // 42.0
}
```

### Struct Field Operations

#### `GetFieldNames(v interface{}) []string`
Returns the names of all fields in a struct.

```go
type Person struct {
    Name string
    Age  int
}
names := reflectx.GetFieldNames(Person{})
fmt.Println(names) // ["Name", "Age"]
```

#### `GetFieldValue(v interface{}, fieldName string) (interface{}, error)`
Gets the value of a struct field.

```go
p := Person{Name: "Alice", Age: 30}
value, err := reflectx.GetFieldValue(p, "Name")
if err == nil {
    fmt.Println(value) // "Alice"
}
```

#### `SetFieldValue(v interface{}, fieldName string, value interface{}) error`
Sets the value of a struct field.

```go
p := &Person{Name: "Alice", Age: 30}
err := reflectx.SetFieldValue(p, "Age", 31)
if err == nil {
    fmt.Println(p.Age) // 31
}
```

#### `GetFieldTags(v interface{}) map[string]string`
Returns all field tags in a struct.

```go
type User struct {
    Name string `json:"name" db:"user_name"`
    Age  int    `json:"age" db:"user_age"`
}
tags := reflectx.GetFieldTags(User{})
fmt.Println(tags) // map[Age:json:"age" db:"user_age" Name:json:"name" db:"user_name"]
```

### Value Operations

#### `GetValue(v interface{}) interface{}`
Returns the underlying value.

```go
slice := []int{1, 2, 3}
value := reflectx.GetValue(slice)
fmt.Println(value) // [1 2 3]
```

#### `GetLen(v interface{}) int`
Returns the length of a slice, array, or string.

```go
slice := []int{1, 2, 3}
fmt.Println(reflectx.GetLen(slice)) // 3
```

#### `GetIndex(v interface{}, index int) (interface{}, error)`
Gets an element by index from a slice, array, or string.

```go
slice := []string{"apple", "banana", "cherry"}
element, err := reflectx.GetIndex(slice, 1)
if err == nil {
    fmt.Println(element) // "banana"
}
```

#### `SetIndex(v interface{}, index int, value interface{}) error`
Sets an element by index in a slice or array.

```go
slice := &[]string{"apple", "banana", "cherry"}
err := reflectx.SetIndex(slice, 1, "orange")
if err == nil {
    fmt.Println(*slice) // [apple orange cherry]
}
```

### Map Operations

#### `GetMapValue(m interface{}, key interface{}) (interface{}, error)`
Gets a value from a map by key.

```go
data := map[string]int{"apple": 5, "banana": 3}
value, err := reflectx.GetMapValue(data, "banana")
if err == nil {
    fmt.Println(value) // 3
}
```

#### `SetMapValue(m interface{}, key, value interface{}) error`
Sets a value in a map by key.

```go
data := map[string]int{"apple": 5}
err := reflectx.SetMapValue(data, "banana", 3)
if err == nil {
    fmt.Println(data["banana"]) // 3
}
```

#### `GetMapKeys(m interface{}) []interface{}`
Returns all keys from a map.

```go
data := map[string]int{"apple": 5, "banana": 3}
keys := reflectx.GetMapKeys(data)
fmt.Println(keys) // ["apple", "banana"]
```

### Method Operations

#### `CallMethod(v interface{}, methodName string, args ...interface{}) ([]interface{}, error)`
Calls a method on a value.

```go
type Calculator struct{ Value int }
func (c Calculator) Add(x int) int { return c.Value + x }

calc := Calculator{Value: 10}
result, err := reflectx.CallMethod(calc, "Add", 5)
if err == nil {
    fmt.Println(result[0]) // 15
}
```

#### `HasMethod(v interface{}, methodName string) bool`
Checks if a value has a method.

```go
type Service struct{}
func (s Service) Process() string { return "processed" }

service := Service{}
fmt.Println(reflectx.HasMethod(service, "Process")) // true
```

#### `GetMethodNames(v interface{}) []string`
Returns the names of all methods on a value.

```go
type User struct{ Name string }
func (u User) GetName() string { return u.Name }
func (u User) SetName(name string) { u.Name = name }

user := User{}
names := reflectx.GetMethodNames(user)
fmt.Println(names) // ["GetName", "SetName"]
```

### Interface Utilities

#### `Implements(v interface{}, iface interface{}) bool`
Checks if a value implements an interface.

```go
type Reader interface { Read() string }
type Book struct{ Title string }
func (b Book) Read() string { return b.Title }

book := Book{Title: "Go Programming"}
fmt.Println(reflectx.Implements(book, (*Reader)(nil))) // true
```

#### `GetInterfaceMethods(iface interface{}) []string`
Returns the names of all methods in an interface.

```go
type Writer interface {
    Write(data string) error
    Close() error
}
methods := reflectx.GetInterfaceMethods((*Writer)(nil))
fmt.Println(methods) // ["Write", "Close"]
```

### Generic Utilities

#### `IsGeneric(v interface{}) bool`
Checks if a value is of a generic type.

```go
type GenericSlice[T any] []T
var slice GenericSlice[int]
fmt.Println(reflectx.IsGeneric(slice)) // false (limitation of Go reflection)
```

## Performance

The package is optimized for performance with comprehensive benchmarks. Key performance characteristics:

- **Type checking operations**: ~10-50ns per operation
- **Field operations**: ~100-500ns per operation
- **Method operations**: ~200-1000ns per operation
- **Large struct operations**: Scales linearly with field count

Run benchmarks:

```bash
go test -bench=. ./reflectx
```

## Testing

The package includes comprehensive tests with 100% coverage of all functions:

```bash
go test ./reflectx -v
```

## Examples

See the `example_test.go` file for detailed usage examples of all functions.

## Dependencies

- Go 1.18+ (for generics support)
- Standard library only (no external dependencies)

## License

This package is part of the goal project and follows the same license terms.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Related Packages

- [reflect](https://pkg.go.dev/reflect) - Go's standard reflection package
- [go4x/goal](https://github.com/go4x/goal) - The parent project containing multiple utility packages
