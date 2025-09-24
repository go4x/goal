# random - Go Random Number Generation Package

A comprehensive Go package for random number generation, sampling, and probability operations with high performance and type safety.

## Features

- **High Performance**: Global random number generator for optimal performance
- **Type Safety**: Generic functions for any data type
- **Comprehensive**: Integer, float, boolean, and distribution random numbers
- **String Generation**: Extensive random string generation with multiple character sets
- **Sampling**: Shuffle, sample, and weighted selection operations
- **Probability**: Percentage-based probability functions
- **Statistical**: Normal and exponential distribution generation
- **Security**: Cryptographically secure random number generation

## Installation

```bash
go get github.com/go4x/goal/random
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Basic random numbers
    intValue := random.Int(100)           // Random integer 0-99
    floatValue := random.Float64()        // Random float 0.0-1.0
    boolValue := random.Bool()           // Random boolean
    
    // Range-based generation
    rangeInt := random.Between(10, 50)    // Random integer 10-49
    rangeFloat := random.Float64Between(1.5, 3.7) // Random float 1.5-3.7
    
    // Probability
    if random.Percent(30) {              // 30% chance
        fmt.Println("Success!")
    }
    
    // Sampling
    items := []string{"A", "B", "C", "D", "E"}
    choice := random.Choice(items)        // Random choice
    sampled := random.Sample(items, 3)   // Sample 3 items
    
    fmt.Printf("Random values: %d, %.4f, %t\n", intValue, floatValue, boolValue)
    fmt.Printf("Range values: %d, %.4f\n", rangeInt, rangeFloat)
    fmt.Printf("Choice: %s, Sampled: %v\n", choice, sampled)
}
```

## API Reference

### Basic Random Numbers

#### `Int(max int) int`
Generates a random integer in the range [0, max).

```go
value := random.Int(100)  // Returns 0-99
```

#### `Between(min, max int) int`
Generates a random integer in the range [min, max).

```go
value := random.Between(10, 50)  // Returns 10-49
```

#### `Float64() float64`
Generates a random float64 in the range [0.0, 1.0).

```go
value := random.Float64()  // Returns 0.0-0.999...
```

#### `Float64Between(min, max float64) float64`
Generates a random float64 in the range [min, max).

```go
value := random.Float64Between(1.5, 3.7)  // Returns 1.5-3.699...
```

#### `Bool() bool`
Generates a random boolean value.

```go
value := random.Bool()  // Returns true or false
```

### Selection and Sampling

#### `Choice[T any](slice []T) T`
Selects a random element from a slice.

```go
options := []string{"A", "B", "C"}
choice := random.Choice(options)  // Returns "A", "B", or "C"
```

#### `Shuffle[T any](slice []T)`
Shuffles a slice in place using Fisher-Yates algorithm.

```go
numbers := []int{1, 2, 3, 4, 5}
random.Shuffle(numbers)  // Shuffles in place
```

#### `Sample[T any](slice []T, k int) []T`
Samples k elements from a slice without replacement.

```go
items := []string{"A", "B", "C", "D", "E"}
sampled := random.Sample(items, 3)  // Returns 3 random items
```

#### `SelectWeighted[T any](choices []WeightedChoice[T]) T`
Selects a random element based on weights.

```go
choices := []random.WeightedChoice[string]{
    {Weight: 1, Value: "Common"},
    {Weight: 2, Value: "Uncommon"},
    {Weight: 3, Value: "Rare"},
}
choice := random.SelectWeighted(choices)
```

### Probability Functions

#### `Percent(percentage int) bool`
Returns true with the given percentage probability.

```go
if random.Percent(30) {  // 30% chance
    fmt.Println("Success!")
}
```

#### `PercentFloat(probability float64) bool`
Returns true with the given probability.

```go
if random.PercentFloat(0.3) {  // 30% chance
    fmt.Println("Success!")
}
```

### Statistical Distributions

#### `Normal(mean, stdDev float64) float64`
Generates a random number from a normal distribution.

```go
value := random.Normal(100, 15)  // Mean=100, StdDev=15
```

#### `Exponential(lambda float64) float64`
Generates a random number from an exponential distribution.

```go
value := random.Exponential(2.0)  // Rate=2.0
```

### String Generation

#### Basic Character Sets

#### `Lowercase(length int) string`
Generates a random lowercase string.

```go
str := random.Lowercase(8)  // Returns "abcdefgh"
```

#### `Uppercase(length int) string`
Generates a random uppercase string.

```go
str := random.Uppercase(8)  // Returns "ABCDEFGH"
```

#### `Digits(length int) string`
Generates a random digit-only string.

```go
str := random.Digits(6)  // Returns "123456"
```

#### `Symbols(length int) string`
Generates a random symbol-only string.

```go
str := random.Symbols(5)  // Returns "!@#$%"
```

#### `HexUpper(length int) string`
Generates a random uppercase hexadecimal string.

```go
str := random.HexUpper(8)  // Returns "12345678"
```

#### Advanced String Generation

#### `AlphanumericSymbols(length int) string`
Generates a random string with letters, numbers, and symbols.

```go
str := random.AlphanumericSymbols(10)  // Returns "aB3dEfGh!@"
```

#### `StrongPassword(length int) string`
Generates a strong password without confusing characters.

```go
str := random.StrongPassword(12)  // Returns "Kj9mN2pQ7&xY"
```

#### `Readable(length int) string`
Generates a readable string without confusing characters.

```go
str := random.Readable(10)  // Returns "abcdefghjk"
```

#### `ShortID(length int) string`
Generates a short ID suitable for URLs.

```go
str := random.ShortID(8)  // Returns "aB3dEfGh"
```

#### `Password(length int, includeSymbols bool) string`
Generates a password with optional symbols.

```go
str := random.Password(12, true)   // With symbols: "Kj9#mN2$pQ7&"
str := random.Password(12, false)  // Without symbols: "Kj9mN2pQ7xYz"
```

#### `Username(length int) string`
Generates a username with letters, numbers, and underscores.

```go
str := random.Username(8)  // Returns "user1234"
```

#### `Email(length int) string`
Generates an email prefix.

```go
str := random.Email(8)  // Returns "user1234"
```

#### `Token(length int) string`
Generates a security token.

```go
str := random.Token(32)  // Returns "aB3dEfGhJkLmN2pQ7rStUvWxYz123456"
```

#### Color Generation

#### `ColorHex() string`
Generates a random color hex code.

```go
color := random.ColorHex()  // Returns "#FF5733"
```

#### `ColorRGB() string`
Generates a random RGB color string.

```go
color := random.ColorRGB()  // Returns "rgb(255, 87, 51)"
```

#### Network Address Generation

#### `MACAddress() string`
Generates a random MAC address.

```go
mac := random.MACAddress()  // Returns "00:1B:44:11:3A:B7"
```

#### `IPAddress() string`
Generates a random IP address.

```go
ip := random.IPAddress()  // Returns "192.168.1.100"
```

#### Advanced String Generation

#### `WeightedString(chars []WeightedChar, length int) string`
Generates a string based on character weights.

```go
chars := []random.WeightedChar{
    {Char: 'a', Weight: 1},
    {Char: 'b', Weight: 2},
    {Char: 'c', Weight: 3},
}
str := random.WeightedString(chars, 10)  // Returns "ccbacccbac"
```

#### `PatternString(pattern string) string`
Generates a string based on a pattern.

```go
// Pattern characters:
// a = lowercase letter
// A = uppercase letter
// n = number
// s = symbol
// x = any alphanumeric
// ? = any character

str := random.PatternString("aAn")  // Returns "aB3"
```

## Usage Examples

### Basic Random Generation

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Generate random numbers
    fmt.Printf("Random int: %d\n", random.Int(100))
    fmt.Printf("Random float: %.4f\n", random.Float64())
    fmt.Printf("Random bool: %t\n", random.Bool())
    
    // Generate within ranges
    fmt.Printf("Random between 10-50: %d\n", random.Between(10, 50))
    fmt.Printf("Random float 1.5-3.7: %.4f\n", random.Float64Between(1.5, 3.7))
}
```

### Sampling and Selection

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Random selection
    colors := []string{"red", "green", "blue", "yellow", "purple"}
    color := random.Choice(colors)
    fmt.Printf("Random color: %s\n", color)
    
    // Shuffle a deck of cards
    cards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    random.Shuffle(cards)
    fmt.Printf("Shuffled cards: %v\n", cards)
    
    // Sample without replacement
    items := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    sampled := random.Sample(items, 3)
    fmt.Printf("Sampled items: %v\n", sampled)
}
```

### Weighted Selection

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Weighted loot system
    lootChoices := []random.WeightedChoice[string]{
        {Weight: 50, Value: "Common Item"},
        {Weight: 30, Value: "Uncommon Item"},
        {Weight: 15, Value: "Rare Item"},
        {Weight: 4, Value: "Epic Item"},
        {Weight: 1, Value: "Legendary Item"},
    }
    
    // Simulate 10 loot drops
    for i := 0; i < 10; i++ {
        loot := random.SelectWeighted(lootChoices)
        fmt.Printf("Drop %d: %s\n", i+1, loot)
    }
}
```

### Probability Systems

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Critical hit system
    if random.Percent(15) {  // 15% crit chance
        fmt.Println("Critical hit!")
    } else {
        fmt.Println("Normal hit")
    }
    
    // Skill check system
    if random.PercentFloat(0.75) {  // 75% success rate
        fmt.Println("Skill check passed!")
    } else {
        fmt.Println("Skill check failed!")
    }
}
```

### Statistical Analysis

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Generate sample data
    samples := make([]float64, 1000)
    for i := range samples {
        samples[i] = random.Normal(100, 15)  // Mean=100, StdDev=15
    }
    
    // Calculate statistics
    sum := 0.0
    for _, sample := range samples {
        sum += sample
    }
    mean := sum / float64(len(samples))
    
    fmt.Printf("Sample mean: %.2f\n", mean)
    fmt.Printf("Expected mean: 100.00\n")
}
```

### String Generation Examples

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Basic string generation
    fmt.Printf("Lowercase: %s\n", random.Lowercase(8))
    fmt.Printf("Uppercase: %s\n", random.Uppercase(8))
    fmt.Printf("Digits: %s\n", random.Digits(6))
    fmt.Printf("Symbols: %s\n", random.Symbols(5))
    
    // Advanced string generation
    fmt.Printf("Strong Password: %s\n", random.StrongPassword(12))
    fmt.Printf("Username: %s\n", random.Username(8))
    fmt.Printf("Email: %s@example.com\n", random.Email(8))
    fmt.Printf("Token: %s\n", random.Token(32))
    
    // Color generation
    fmt.Printf("Color Hex: %s\n", random.ColorHex())
    fmt.Printf("Color RGB: %s\n", random.ColorRGB())
    
    // Network addresses
    fmt.Printf("MAC Address: %s\n", random.MACAddress())
    fmt.Printf("IP Address: %s\n", random.IPAddress())
}
```

### Pattern-Based String Generation

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Pattern-based generation
    patterns := []string{
        "aaa",     // lowercase letters
        "AAA",     // uppercase letters
        "nnn",     // numbers
        "sss",     // symbols
        "xxx",     // alphanumeric
        "aAn",     // mixed pattern
    }
    
    for _, pattern := range patterns {
        fmt.Printf("Pattern %s: %s\n", pattern, random.PatternString(pattern))
    }
}
```

### Weighted String Generation

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/random"
)

func main() {
    // Weighted character selection
    chars := []random.WeightedChar{
        {Char: 'a', Weight: 1},  // 16.7% chance
        {Char: 'b', Weight: 2},  // 33.3% chance
        {Char: 'c', Weight: 3},  // 50.0% chance
    }
    
    // Generate weighted strings
    for i := 0; i < 5; i++ {
        str := random.WeightedString(chars, 10)
        fmt.Printf("Weighted string %d: %s\n", i+1, str)
    }
}
```

## Performance

The package is optimized for performance:

- **Global Random Generator**: Single shared random number generator
- **Efficient Algorithms**: Fisher-Yates shuffle, reservoir sampling
- **Minimal Allocations**: Reduced memory allocations where possible
- **Type Safety**: Generic functions for compile-time type checking

### Benchmark Results

- **Basic Operations**: ~2-5ns/op, 0 allocations
- **Selection Operations**: ~10-50ns/op, 0-1 allocations
- **Sampling Operations**: ~100ns-1μs/op, 1-3 allocations
- **Distribution Generation**: ~20-100ns/op, 0 allocations
- **String Generation**: ~50-200ns/op, 1-2 allocations
- **Large String Generation**: ~1-10μs/op, 1-3 allocations

## Testing

Run tests:

```bash
go test ./random
```

Run with coverage:

```bash
go test ./random -cover
```

Run examples:

```bash
go test ./random -run Example
```

Run benchmarks:

```bash
go test ./random -bench=.
```

Run fuzz tests:

```bash
go test ./random -fuzz=.
```

## License

This package is part of the `goal` project and follows the same license terms.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
