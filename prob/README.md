# prob - Go Probability and Statistics Package

A comprehensive Go package for probability calculations, statistical distributions, and random sampling operations.

## Features

- **Basic Probability**: Percentage-based probability calculations
- **Weighted Selection**: Select items based on weights with type safety
- **Statistical Distributions**: Binomial, Poisson, Normal, Geometric, Hypergeometric
- **Random Generation**: Uniform, Normal, Exponential distributions
- **Sampling**: Shuffle, sample, and weighted sampling operations
- **Performance Optimized**: Efficient algorithms with minimal allocations
- **Type Safe**: Generic functions for any data type

## Installation

```bash
go get github.com/go4x/goal/prob
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // Basic probability
    if prob.Percent(30) {
        fmt.Println("30% chance hit!")
    }
    
    // Weighted selection
    weights := []int{1, 2, 3, 4}
    index, _ := prob.Select(weights)
    fmt.Printf("Selected index: %d\n", index)
    
    // Statistical distributions
    probability := prob.Binomial(10, 3, 0.5)
    fmt.Printf("Binomial probability: %.4f\n", probability)
}
```

## API Reference

### Basic Probability

#### `Percent(percentage int) bool`
Calculates percentage-based probability (0-100).

```go
if prob.Percent(30) {
    // 30% chance of success
}
```

#### `PercentFloat(probability float64) bool`
Calculates probability using float64 (0.0-1.0).

```go
if prob.PercentFloat(0.3) {
    // 30% chance of success
}
```

#### `Half() bool`
Returns true with 50% probability.

```go
if prob.Half() {
    fmt.Println("Heads")
} else {
    fmt.Println("Tails")
}
```

### Weighted Selection

#### `Select(weights []int) (int, error)`
Selects an index based on integer weights.

```go
weights := []int{1, 2, 3, 4}
index, err := prob.Select(weights)
if err != nil {
    // Handle error
}
```

#### `SelectFloat(weights []float64) (int, error)`
Selects an index based on float64 weights.

```go
weights := []float64{0.1, 0.2, 0.3, 0.4}
index, err := prob.SelectFloat(weights)
```

#### `SelectSafe(weights []int) int`
Safe version that returns -1 on error instead of panicking.

```go
index := prob.SelectSafe(weights)
if index == -1 {
    // Handle error
}
```

#### `SelectWeighted[T any](choices []WeightedChoice[T]) (T, error)`
Selects a value from weighted choices with type safety.

```go
choices := []prob.WeightedChoice[string]{
    {Weight: 1, Value: "Low"},
    {Weight: 2, Value: "Medium"},
    {Weight: 3, Value: "High"},
}

value, err := prob.SelectWeighted(choices)
```

#### `SelectWeightedFloat[T any](choices []WeightedChoiceFloat[T]) (T, error)`
Selects a value from float64 weighted choices.

```go
choices := []prob.WeightedChoiceFloat[string]{
    {Weight: 0.1, Value: "Rare"},
    {Weight: 0.3, Value: "Common"},
    {Weight: 0.6, Value: "Very Common"},
}

value, err := prob.SelectWeightedFloat(choices)
```

### Statistical Distributions

#### `Binomial(n, k int, p float64) float64`
Calculates binomial probability.

```go
// Probability of exactly 3 successes in 10 trials with 0.5 probability
prob := prob.Binomial(10, 3, 0.5)
```

#### `Poisson(k int, lambda float64) float64`
Calculates Poisson probability.

```go
// Poisson probability for k=3, λ=2.0
prob := prob.Poisson(3, 2.0)
```

#### `Normal(x, mean, stdDev float64) float64`
Calculates normal distribution probability density.

```go
// Normal density at x=0, mean=0, stdDev=1
density := prob.Normal(0, 0, 1)
```

#### `Geometric(k int, p float64) float64`
Calculates geometric probability.

```go
// Probability of first success on trial 3
prob := prob.Geometric(3, 0.3)
```

#### `Hypergeometric(n, k, K, N int) float64`
Calculates hypergeometric probability.

```go
// Hypergeometric probability
prob := prob.Hypergeometric(10, 3, 20, 100)
```

### Random Generation

#### `Uniform(min, max float64) float64`
Generates random number from uniform distribution.

```go
// Random number between 0 and 10
value := prob.Uniform(0, 10)
```

#### `NormalRandom(mean, stdDev float64) float64`
Generates random number from normal distribution.

```go
// Normal random number with mean=0, stdDev=1
value := prob.NormalRandom(0, 1)
```

#### `Exponential(lambda float64) float64`
Generates random number from exponential distribution.

```go
// Exponential random number with λ=1.0
value := prob.Exponential(1.0)
```

### Sampling Operations

#### `Shuffle[T any](slice []T)`
Shuffles a slice in place using Fisher-Yates algorithm.

```go
slice := []int{1, 2, 3, 4, 5}
prob.Shuffle(slice)
```

#### `Sample[T any](slice []T, k int) []T`
Samples k elements from a slice without replacement.

```go
slice := []string{"A", "B", "C", "D", "E"}
sampled := prob.Sample(slice, 3)
```

#### `WeightedSample[T any](slice []T, weights []int, k int) ([]T, error)`
Samples k elements based on weights.

```go
slice := []string{"A", "B", "C", "D"}
weights := []int{1, 2, 3, 4}
sampled, err := prob.WeightedSample(slice, weights, 2)
```

## Usage Examples

### Basic Probability

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // 30% chance of success
    if prob.Percent(30) {
        fmt.Println("Success!")
    }
    
    // 50% chance
    if prob.Half() {
        fmt.Println("Heads")
    } else {
        fmt.Println("Tails")
    }
}
```

### Weighted Selection

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // Select from weighted choices
    choices := []prob.WeightedChoice[string]{
        {Weight: 1, Value: "Common"},
        {Weight: 2, Value: "Uncommon"},
        {Weight: 3, Value: "Rare"},
        {Weight: 4, Value: "Epic"},
        {Weight: 5, Value: "Legendary"},
    }
    
    value, err := prob.SelectWeighted(choices)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Selected item: %s\n", value)
}
```

### Statistical Analysis

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // Binomial distribution
    n, k, p := 20, 5, 0.3
    binomialProb := prob.Binomial(n, k, p)
    fmt.Printf("Binomial: P(X=%d) in %d trials = %.4f\n", k, n, binomialProb)
    
    // Poisson distribution
    lambda := 3.0
    poissonProb := prob.Poisson(2, lambda)
    fmt.Printf("Poisson: P(X=2) with λ=%.1f = %.4f\n", lambda, poissonProb)
    
    // Normal distribution
    normalDensity := prob.Normal(0, 0, 1)
    fmt.Printf("Normal density at 0: %.4f\n", normalDensity)
}
```

### Random Sampling

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // Shuffle a slice
    slice := []int{1, 2, 3, 4, 5}
    fmt.Printf("Original: %v\n", slice)
    
    prob.Shuffle(slice)
    fmt.Printf("Shuffled: %v\n", slice)
    
    // Sample elements
    sampled := prob.Sample(slice, 3)
    fmt.Printf("Sampled: %v\n", sampled)
}
```

### Gaming Applications

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/prob"
)

func main() {
    // Loot drop system
    lootChoices := []prob.WeightedChoice[string]{
        {Weight: 1, Value: "Legendary"},
        {Weight: 5, Value: "Epic"},
        {Weight: 15, Value: "Rare"},
        {Weight: 30, Value: "Common"},
        {Weight: 49, Value: "Trash"},
    }
    
    // Simulate loot drops
    for i := 0; i < 5; i++ {
        loot, _ := prob.SelectWeighted(lootChoices)
        fmt.Printf("Drop %d: %s\n", i+1, loot)
    }
    
    // Critical hit system
    critChance := 0.15 // 15% crit chance
    if prob.PercentFloat(critChance) {
        fmt.Println("Critical hit!")
    } else {
        fmt.Println("Normal hit")
    }
}
```

## Performance

The package is optimized for performance:

- **Global Random Generator**: Single shared random number generator
- **Efficient Algorithms**: Fisher-Yates shuffle, optimized sampling
- **Minimal Allocations**: Reduced memory allocations where possible
- **Type Safety**: Generic functions for compile-time type checking

### Benchmark Results

- **Basic Operations**: ~3-4ns/op, 0 allocations
- **Weighted Selection**: ~17-28ns/op, 0-1 allocations
- **Statistical Calculations**: ~8-60ns/op, 0 allocations
- **Sampling Operations**: ~400ns-4μs/op, 1-3 allocations

## Testing

Run tests:

```bash
go test ./prob
```

Run with coverage:

```bash
go test ./prob -cover
```

Run examples:

```bash
go test ./prob -run Example
```

Run benchmarks:

```bash
go test ./prob -bench=.
```

## License

This package is part of the `goal` project and follows the same license terms.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
