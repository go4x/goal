# slicex

A comprehensive generic slice operations package for Go that provides enhanced slice functionality with immutability guarantees and rich operations.

## Features

- **Immutable Operations**: All methods return new slices without modifying originals
- **Generic Type Support**: Works with any comparable type
- **Rich Functionality**: Filter, map, sort, reverse, union, intersect, and more
- **Performance Optimized**: Uses hash maps for O(n+m) operations where possible
- **Functional Programming**: Chainable operations for fluent API
- **Type Safety**: Full compile-time type checking

## Installation

```bash
go get github.com/go4x/goal/col/slicex
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/col/slicex"
)

func main() {
    // Create a slice
    numbers := slicex.From([]int{3, 1, 4, 1, 5})
    
    // Filter and sort
    filtered := numbers.Filter(func(x int) bool { return x > 2 })
    sorted := numbers.Sort(func(a, b int) bool { return a < b })
    
    // Original slice remains unchanged
    fmt.Println(numbers.To()) // [3 1 4 1 5]
    fmt.Println(filtered.To()) // [3 4 5]
    fmt.Println(sorted.To()) // [1 1 3 4 5]
}
```

## Core Types

### S[T] - Generic Slice Type

The main generic slice type with enhanced methods:

```go
import "github.com/go4x/goal/col/slicex"

// Create from existing slice
numbers := slicex.From([]int{1, 2, 3, 4, 5})

// Create new empty slice
empty := slicex.New[int]()

// Convert back to Go slice
goSlice := numbers.To()
```

### SortableSlice[T] - Sorting Helper

A helper type for sorting operations:

```go
import "github.com/go4x/goal/col/slicex"

// Create sortable slice
sortable := slicex.NewSortableSlice([]int{3, 1, 4, 1, 5})

// Sort with custom comparator
sorted := sortable.Sort(func(a, b int) bool { return a < b })
```

## Basic Operations

### Creating Slices

```go
import "github.com/go4x/goal/col/slicex"

// From existing slice
original := []int{1, 2, 3, 4, 5}
slice := slicex.From(original)

// New empty slice
empty := slicex.New[int]()

// From variadic arguments
slice = slicex.Of(1, 2, 3, 4, 5)

// From function
slice = slicex.Generate(5, func(i int) int { return i * 2 }) // [0, 2, 4, 6, 8]
```

### Basic Properties

```go
import "github.com/go4x/goal/col/slicex"

slice := slicex.From([]int{1, 2, 3, 4, 5})

// Get length
fmt.Println(slice.Len()) // 5

// Check if empty
fmt.Println(slice.IsEmpty()) // false

// Get element at index
fmt.Println(slice.Get(2)) // 3 true

// Set element at index
newSlice := slice.Set(2, 10)
fmt.Println(newSlice.To()) // [1 2 10 4 5]
```

### Filtering and Mapping

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// Filter even numbers
evens := numbers.Filter(func(x int) bool { return x%2 == 0 })
fmt.Println(evens.To()) // [2 4 6 8 10]

// Map to squares
squares := numbers.Map(func(x int) int { return x * x })
fmt.Println(squares.To()) // [1 4 9 16 25 36 49 64 81 100]

// Filter and map in one operation
result := numbers.FilterMap(func(x int) (int, bool) {
    if x%2 == 0 {
        return x * x, true
    }
    return 0, false
})
fmt.Println(result.To()) // [4 16 36 64 100]
```

### Sorting

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{3, 1, 4, 1, 5, 9, 2, 6})

// Sort ascending
ascending := numbers.Sort(func(a, b int) bool { return a < b })
fmt.Println(ascending.To()) // [1 1 2 3 4 5 6 9]

// Sort descending
descending := numbers.Sort(func(a, b int) bool { return a > b })
fmt.Println(descending.To()) // [9 6 5 4 3 2 1 1]

// Sort by custom criteria
strings := slicex.From([]string{"apple", "banana", "cherry"})
byLength := strings.Sort(func(a, b string) bool { return len(a) < len(b) })
fmt.Println(byLength.To()) // [apple cherry banana]
```

### Searching

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5})

// Find first element
first := numbers.Find(func(x int) bool { return x > 2 })
fmt.Println(first) // 3 true

// Find last element
last := numbers.FindLast(func(x int) bool { return x > 2 })
fmt.Println(last) // 5 true

// Check if any element matches
hasEven := numbers.Any(func(x int) bool { return x%2 == 0 })
fmt.Println(hasEven) // true

// Check if all elements match
allPositive := numbers.All(func(x int) bool { return x > 0 })
fmt.Println(allPositive) // true
```

## Advanced Operations

### Set Operations

```go
import "github.com/go4x/goal/col/slicex"

slice1 := slicex.From([]int{1, 2, 3, 4, 5})
slice2 := slicex.From([]int{4, 5, 6, 7, 8})

// Union (all unique elements)
union := slice1.Union(slice2)
fmt.Println(union.To()) // [1 2 3 4 5 6 7 8]

// Intersection (common elements)
intersection := slice1.Intersect(slice2)
fmt.Println(intersection.To()) // [4 5]

// Difference (elements in slice1 but not in slice2)
difference := slice1.Difference(slice2)
fmt.Println(difference.To()) // [1 2 3]

// Symmetric difference (elements in either slice but not both)
symmetricDiff := slice1.SymmetricDifference(slice2)
fmt.Println(symmetricDiff.To()) // [1 2 3 6 7 8]
```

### Chunking and Grouping

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

// Chunk into groups of 3
chunks := numbers.Chunk(3)
for i, chunk := range chunks {
    fmt.Printf("Chunk %d: %v\n", i, chunk.To())
}
// Chunk 0: [1 2 3]
// Chunk 1: [4 5 6]
// Chunk 2: [7 8 9]

// Group by criteria
grouped := numbers.GroupBy(func(x int) int { return x % 3 })
for key, group := range grouped {
    fmt.Printf("Group %d: %v\n", key, group.To())
}
// Group 0: [3 6 9]
// Group 1: [1 4 7]
// Group 2: [2 5 8]
```

### Reducing and Aggregating

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5})

// Sum all elements
sum := numbers.Reduce(0, func(acc, x int) int { return acc + x })
fmt.Println(sum) // 15

// Find maximum
max := numbers.Reduce(numbers.Get(0).Value, func(acc, x int) int {
    if x > acc {
        return x
    }
    return acc
})
fmt.Println(max) // 5

// Count elements
count := numbers.Count(func(x int) bool { return x > 2 })
fmt.Println(count) // 3
```

## Utility Functions

### Comparison Functions

```go
import "github.com/go4x/goal/col/slicex"

slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{1, 2, 4}

// Check equality
fmt.Println(slicex.Equal(slice1, slice2)) // true
fmt.Println(slicex.Equal(slice1, slice3)) // false

// Check equality with custom function
fmt.Println(slicex.EqualFunc(slice1, slice2, func(a, b int) bool { return a == b })) // true
```

### Searching Functions

```go
import "github.com/go4x/goal/col/slicex"

slice := []int{1, 2, 3, 2, 4, 2}

// Find index of element
fmt.Println(slicex.IndexOf(slice, 2)) // 1
fmt.Println(slicex.LastIndexOf(slice, 2)) // 5

// Check if slice contains element
fmt.Println(slicex.Contains(slice, 3)) // true
fmt.Println(slicex.Contains(slice, 5)) // false
```

### Transformation Functions

```go
import "github.com/go4x/goal/col/slicex"

slice := []int{1, 2, 3, 4, 5}

// Take first 3 elements
first3 := slicex.Take(slice, 3)
fmt.Println(first3) // [1 2 3]

// Drop first 2 elements
after2 := slicex.Drop(slice, 2)
fmt.Println(after2) // [3 4 5]

// Take while condition is true
while := slicex.TakeWhile(slice, func(x int) bool { return x < 4 })
fmt.Println(while) // [1 2 3]

// Drop while condition is true
dropWhile := slicex.DropWhile(slice, func(x int) bool { return x < 3 })
fmt.Println(dropWhile) // [3 4 5]
```

## Performance Characteristics

### Time Complexity

| Operation | Time Complexity | Description |
|-----------|----------------|-------------|
| Filter | O(n) | Linear scan through all elements |
| Map | O(n) | Apply function to all elements |
| Sort | O(n log n) | Comparison-based sorting |
| Union | O(n + m) | Hash map for deduplication |
| Intersect | O(n + m) | Hash map for intersection |
| Find | O(n) | Linear search |
| Contains | O(n) | Linear search |
| Chunk | O(n) | Linear scan with grouping |

### Memory Usage

- **Immutable operations**: Each operation creates new slices
- **Hash map operations**: Union, intersect use O(n + m) memory
- **Sorting**: In-place sorting for efficiency
- **Chunking**: Creates new slices for each chunk

### Performance Tips

1. **Use hash map operations** for large datasets when possible
2. **Chain operations** to avoid intermediate allocations
3. **Use Take/Drop** instead of slicing when possible
4. **Consider sorting** for repeated search operations

## Thread Safety

⚠️ **Important**: All slice operations are **NOT thread-safe**. If you need concurrent access, you must use synchronization primitives:

```go
import (
    "sync"
    "github.com/go4x/goal/col/slicex"
)

type SafeSlice[T comparable] struct {
    mu sync.RWMutex
    s  slicex.S[T]
}

func (s *SafeSlice[T]) Filter(predicate func(T) bool) slicex.S[T] {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.s.Filter(predicate)
}
```

## API Reference

### S[T] Methods

| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `Len()` | Get length | O(1) |
| `IsEmpty()` | Check if empty | O(1) |
| `Get(index)` | Get element at index | O(1) |
| `Set(index, value)` | Set element at index | O(1) |
| `Filter(predicate)` | Filter elements | O(n) |
| `Map(transform)` | Transform elements | O(n) |
| `Sort(comparator)` | Sort elements | O(n log n) |
| `Find(predicate)` | Find first matching element | O(n) |
| `FindLast(predicate)` | Find last matching element | O(n) |
| `Any(predicate)` | Check if any element matches | O(n) |
| `All(predicate)` | Check if all elements match | O(n) |
| `Union(other)` | Union with another slice | O(n + m) |
| `Intersect(other)` | Intersection with another slice | O(n + m) |
| `Difference(other)` | Difference with another slice | O(n + m) |
| `SymmetricDifference(other)` | Symmetric difference | O(n + m) |
| `Chunk(size)` | Split into chunks | O(n) |
| `GroupBy(keyFunc)` | Group by key function | O(n) |
| `Reduce(initial, reducer)` | Reduce to single value | O(n) |
| `Count(predicate)` | Count matching elements | O(n) |
| `To()` | Convert to Go slice | O(n) |

### Utility Functions

| Function | Description | Time Complexity |
|----------|-------------|-----------------|
| `From(slice)` | Create from Go slice | O(n) |
| `New[T]()` | Create empty slice | O(1) |
| `Of(elements...)` | Create from elements | O(n) |
| `Generate(n, func)` | Generate from function | O(n) |
| `Equal(s1, s2)` | Check equality | O(n) |
| `EqualFunc(s1, s2, eq)` | Check equality with function | O(n) |
| `IndexOf(slice, element)` | Find index of element | O(n) |
| `LastIndexOf(slice, element)` | Find last index of element | O(n) |
| `Contains(slice, element)` | Check if contains element | O(n) |
| `Take(slice, n)` | Take first n elements | O(n) |
| `Drop(slice, n)` | Drop first n elements | O(n) |
| `TakeWhile(slice, predicate)` | Take while predicate is true | O(n) |
| `DropWhile(slice, predicate)` | Drop while predicate is true | O(n) |

## Use Cases

### 1. Data Processing

```go
import "github.com/go4x/goal/col/slicex"

// Process user data
type User struct {
    ID    int
    Name  string
    Age   int
    Email string
}

func ProcessUsers(users []User) []User {
    return slicex.From(users).
        Filter(func(u User) bool { return u.Age >= 18 }).
        Map(func(u User) User {
            u.Email = strings.ToLower(u.Email)
            return u
        }).
        Sort(func(a, b User) bool { return a.Age < b.Age }).
        To()
}
```

### 2. Data Analysis

```go
import "github.com/go4x/goal/col/slicex"

// Analyze sales data
type Sale struct {
    Product string
    Amount  float64
    Date    time.Time
}

func AnalyzeSales(sales []Sale) map[string]float64 {
    return slicex.From(sales).
        GroupBy(func(s Sale) string { return s.Product }).
        Map(func(product string, sales slicex.S[Sale]) (string, float64) {
            total := sales.Reduce(0.0, func(acc float64, s Sale) float64 {
                return acc + s.Amount
            })
            return product, total
        })
}
```

### 3. API Response Processing

```go
import "github.com/go4x/goal/col/slicex"

// Process API responses
func ProcessAPIResponse(data []map[string]interface{}) []string {
    return slicex.From(data).
        Filter(func(item map[string]interface{}) bool {
            return item["status"] == "active"
        }).
        Map(func(item map[string]interface{}) string {
            return item["name"].(string)
        }).
        Sort(func(a, b string) bool { return a < b }).
        To()
}
```

### 4. Configuration Management

```go
import "github.com/go4x/goal/col/slicex"

// Manage configuration entries
type ConfigEntry struct {
    Key   string
    Value string
    Env   string
}

func GetConfigForEnv(configs []ConfigEntry, env string) []ConfigEntry {
    return slicex.From(configs).
        Filter(func(c ConfigEntry) bool { return c.Env == env || c.Env == "default" }).
        Sort(func(a, b ConfigEntry) bool { return a.Key < b.Key }).
        To()
}
```

## License

This package is part of the goal project and follows the same license terms.
