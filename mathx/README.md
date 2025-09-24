# MathX - High Precision Decimal Math Library

A high-precision decimal math library for Go that provides accurate floating-point arithmetic operations using the `shopspring/decimal` library under the hood. MathX offers a clean, chainable API for mathematical operations with precise decimal calculations.

## Features

- **High Precision**: Uses `shopspring/decimal` for accurate decimal arithmetic
- **Chainable API**: Fluent interface for mathematical operations
- **Type Safety**: Compile-time type checking with generic support
- **Comprehensive**: Basic operations, statistical functions, formatting utilities
- **Zero Dependencies**: Only depends on `shopspring/decimal` and standard library

## Installation

```bash
go get github.com/go4x/goal/mathx
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/mathx"
)

func main() {
    // Basic arithmetic with precision
    result := mathx.Add(0.1, 0.2)
    fmt.Println(result) // 0.3 (not 0.30000000000000004)
    
    // Chainable operations
    formatted := mathx.Add(0.1, 0.2).
        Mul(10).
        Div(3, 2).
        Round(2).
        ToString()
    fmt.Println(formatted) // "1.00"
    
    // Money formatting
    price := mathx.Mul(99.99, 1.15).
        Round(2).
        FormatMoney(2)
    fmt.Printf("Price: $%s\n", price) // Price: $114.99
}
```

## Core Concepts

### Result Type

All mathematical operations return a `Result` type that can be used in two ways:

1. **Direct usage**: Returns `float64` when used directly
2. **Chainable methods**: Call methods for formatting and further operations

```go
// Direct usage
sum := mathx.Add(1, 2) // Returns Result, but can be used as float64
fmt.Println(sum) // 3

// Chainable usage
str := mathx.Add(1, 2).ToString() // "3"
```

## API Reference

### Basic Arithmetic

#### Add
```go
func Add(a, b float64) Result
```
Adds two numbers with decimal precision.

```go
result := mathx.Add(0.1, 0.2) // 0.3
```

#### Sub
```go
func Sub(a, b float64) Result
```
Subtracts two numbers with decimal precision.

```go
result := mathx.Sub(5.0, 2.0) // 3.0
```

#### Mul
```go
func Mul(a, b float64) Result
```
Multiplies two numbers with decimal precision.

```go
result := mathx.Mul(3.0, 4.0) // 12.0
```

#### Div
```go
func Div(a, b float64, precision int32) Result
```
Divides two numbers with specified precision.

```go
result := mathx.Div(10.0, 3.0, 2) // 3.33
```

#### DivTrunc
```go
func DivTrunc(a, b float64, precision int32) Result
```
Divides two numbers and truncates to specified precision.

```go
result := mathx.DivTrunc(10.0, 3.0, 2) // 3.33
```

### Mathematical Functions

#### Round
```go
func Round(value float64, precision int32) Result
```
Rounds a number to specified decimal places.

```go
result := mathx.Round(3.145, 2) // 3.15
```

#### Truncate
```go
func Truncate(value float64, precision int32) Result
```
Truncates a number to specified decimal places.

```go
result := mathx.Truncate(3.145, 2) // 3.14
```

#### Abs
```go
func Abs(value float64) float64
```
Returns the absolute value of a number.

```go
result := mathx.Abs(-3.14) // 3.14
```

#### Ceil
```go
func Ceil(value float64) float64
```
Returns the smallest integer greater than or equal to the value.

```go
result := mathx.Ceil(3.2) // 4.0
```

#### Floor
```go
func Floor(value float64) float64
```
Returns the largest integer less than or equal to the value.

```go
result := mathx.Floor(3.8) // 3.0
```

#### Pow
```go
func Pow(base, exponent float64) float64
```
Raises a number to the power of another.

```go
result := mathx.Pow(2.0, 3.0) // 8.0
```

#### Sqrt
```go
func Sqrt(value float64) float64
```
Returns the square root of a number.

```go
result := mathx.Sqrt(16.0) // 4.0
```

### Statistical Functions

#### Average
```go
func Average[T constraints.Integer | constraints.Float](ns ...T) float64
```
Calculates the average of a slice of numbers.

```go
avg := mathx.Average(1, 2, 3, 4, 5) // 3.0
```

#### Median
```go
func Median[T constraints.Integer | constraints.Float](ns ...T) float64
```
Calculates the median of a slice of numbers.

```go
median := mathx.Median(1, 2, 3, 4, 5) // 3.0
```

#### StandardDeviation
```go
func StandardDeviation[T constraints.Integer | constraints.Float](ns ...T) float64
```
Calculates the standard deviation of a slice of numbers.

```go
std := mathx.StandardDeviation(1, 2, 3, 4, 5) // 1.4142135623730951
```

#### Max
```go
func Max[T constraints.Ordered](ns ...T) T
```
Returns the maximum value from a slice of numbers.

```go
max := mathx.Max(1, 5, 3, 9, 2) // 9.0
```

#### Min
```go
func Min[T constraints.Ordered](ns ...T) T
```
Returns the minimum value from a slice of numbers.

```go
min := mathx.Min(1, 5, 3, 9, 2) // 1.0
```

#### Sum
```go
func Sum[T constraints.Integer | constraints.Float](ns ...T) T
```
Returns the sum of a slice of numbers.

```go
sum := mathx.Sum(1, 2, 3, 4, 5) // 15.0
```

### Utility Functions

#### Clamp
```go
func Clamp(value, min, max float64) float64
```
Clamps a value between min and max.

```go
result := mathx.Clamp(15.0, 0.0, 10.0) // 10.0
```

#### Lerp
```go
func Lerp(a, b, t float64) float64
```
Performs linear interpolation between two values.

```go
result := mathx.Lerp(0.0, 10.0, 0.5) // 5.0
```

#### Percentage
```go
func Percentage(value, percent float64) float64
```
Calculates percentage of a value.

```go
result := mathx.Percentage(100.0, 15.0) // 15.0
```

#### CompoundInterest
```go
func CompoundInterest(principal, rate float64, periods int) float64
```
Calculates compound interest.

```go
result := mathx.CompoundInterest(1000.0, 0.1, 2) // 1210.0
```

#### SafeDiv
```go
func SafeDiv(dividend, divisor float64, precision int32) float64
```
Safely divides two numbers, returns 0 if divisor is 0.

```go
result := mathx.SafeDiv(10.0, 0.0, 2) // 0.0
```

### String Conversion

#### ToString
```go
func ToString(value float64) string
```
Converts a float64 to string.

```go
str := mathx.ToString(3.14) // "3.14"
```

#### ToStringFixed
```go
func ToStringFixed(value float64, places int32) string
```
Converts a float64 to string with fixed decimal places.

```go
str := mathx.ToStringFixed(3.14159, 2) // "3.14"
```

#### ToStringBank
```go
func ToStringBank(value float64, places int32) string
```
Converts a float64 to string with banker's rounding.

```go
str := mathx.ToStringBank(3.145, 2) // "3.14"
```

#### FormatMoney
```go
func FormatMoney(amount float64, decimalPlaces int32) string
```
Formats a number as currency with thousands separator.

```go
str := mathx.FormatMoney(1234567.89, 2) // "1,234,567.89"
```

#### ParseFloat
```go
func ParseFloat(s string) (float64, error)
```
Safely parses a string to float64.

```go
value, err := mathx.ParseFloat("3.14")
if err == nil {
    fmt.Println(value) // 3.14
}
```

### Result Methods

The `Result` type provides chainable methods for further operations:

#### String Conversion
```go
result.ToString() string
result.ToStringFixed(places int32) string
result.ToStringBank(places int32) string
```

#### Formatting
```go
result.FormatMoney(decimalPlaces int32) string
result.Clean() Result // Removes trailing zeros
```

#### Mathematical Operations
```go
result.Add(other float64) Result
result.Sub(other float64) Result
result.Mul(other float64) Result
result.Div(other float64, precision int32) Result
result.Round(places int32) Result
result.Truncate(places int32) Result
result.Abs() Result
result.Neg() Result
```

#### Value Extraction
```go
result.Float64() float64
result.String() string
```

## Examples

### Financial Calculations

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/mathx"
)

func main() {
    // Calculate tax
    price := 99.99
    taxRate := 0.08
    tax := mathx.Mul(price, taxRate).Round(2)
    
    // Calculate total
    total := mathx.Add(price, tax.Float64())
    
    // Format as currency
    formatted := total.FormatMoney(2)
    fmt.Printf("Total: $%s\n", formatted) // Total: $107.99
}
```

### Statistical Analysis

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/mathx"
)

func main() {
    scores := []float64{85, 92, 78, 96, 88}
    
    avg := mathx.Average(scores...)
    median := mathx.Median(scores...)
    std := mathx.StandardDeviation(scores...)
    
    fmt.Printf("Average: %.2f\n", avg)
    fmt.Printf("Median: %.2f\n", median)
    fmt.Printf("Standard Deviation: %.2f\n", std)
}
```

### Chainable Operations

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/mathx"
)

func main() {
    // Complex calculation with formatting
    result := mathx.Add(0.1, 0.2).
        Mul(100).
        Div(3, 4).
        Round(2).
        Clean().
        ToString()
    
    fmt.Println(result) // "10.00"
}
```

### Percentage Calculations

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/mathx"
)

func main() {
    // Calculate tip
    bill := 50.00
    tipPercent := 18.0
    tip := mathx.Percentage(bill, tipPercent)
    
    // Calculate total
    total := mathx.Add(bill, tip)
    
    fmt.Printf("Bill: $%.2f\n", bill)
    fmt.Printf("Tip (%.0f%%): $%.2f\n", tipPercent, tip)
    fmt.Printf("Total: $%.2f\n", total.Float64())
}
```

## Performance Considerations

- MathX uses `shopspring/decimal` internally for high precision
- String conversions may have slight overhead compared to native float64 operations
- For performance-critical applications, consider using the direct `float64` return values
- Chainable operations create new `Result` instances, so use judiciously in tight loops

## Error Handling

Most functions return `float64` or `Result` types directly. Functions that can fail (like `ParseFloat`) return `(value, error)` tuples.

```go
value, err := mathx.ParseFloat("invalid")
if err != nil {
    // Handle error
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Built on top of the excellent `shopspring/decimal` library
- Inspired by the need for precise decimal arithmetic in Go applications
