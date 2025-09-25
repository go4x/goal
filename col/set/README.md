# set

A comprehensive generic set implementation package for Go that provides multiple set implementations optimized for different use cases.

## Features

- **Multiple Set Implementations**: HashSet, ArraySet, and LinkedSet
- **Generic Type Support**: Works with any comparable type
- **Polymorphic Interface**: Unified `Set[T]` interface for all implementations
- **Performance Optimized**: Different implementations for different performance requirements
- **Order Preservation**: ArraySet and LinkedSet maintain insertion order
- **Memory Efficient**: Optimized memory usage for different scenarios

## Installation

```bash
go get github.com/go4x/goal/col/set
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/col/set"
)

func main() {
    // Create a set (defaults to HashSet)
    mySet := set.New[string]()
    mySet.Add("apple").Add("banana").Add("apple") // "apple" is only added once
    
    fmt.Println(mySet.Size()) // Output: 2
    fmt.Println(mySet.Contains("apple")) // Output: true
    
    // Get all elements
    elements := mySet.Elems()
    fmt.Println(elements) // Output: [apple banana] (order may vary)
}
```

## Set Implementations

### 1. HashSet (Default)

**Best for**: General-purpose set operations, performance-critical applications

```go
import "github.com/go4x/goal/col/set"

// Create a HashSet
hashSet := set.NewHashSet[string]()
hashSet.Add("first").Add("second").Add("first") // Duplicates ignored

// O(1) operations
fmt.Println(hashSet.Contains("first")) // true
hashSet.Remove("second")
fmt.Println(hashSet.Size()) // 1
```

**Characteristics:**
- ⚡ **Fastest performance**: O(1) average-case for all operations
- 🔀 **No order guarantee**: Elements may appear in any order
- 💾 **Memory efficient**: Uses hash map internally
- 🎯 **Best for**: Large datasets, performance-critical code

### 2. ArraySet

**Best for**: Small datasets, when insertion order matters

```go
import "github.com/go4x/goal/col/set"

// Create an ArraySet
arraySet := set.NewArraySet[string]()
arraySet.Add("first").Add("second").Add("third")

// Maintains insertion order
elements := arraySet.Elems()
fmt.Println(elements) // Output: [first second third]
```

**Characteristics:**
- 📋 **Maintains order**: Elements appear in insertion order
- 🐌 **O(n) operations**: Linear time complexity
- 💾 **Memory efficient**: Good for small datasets
- 🎯 **Best for**: Small datasets (< 1000 elements), when order matters

### 3. LinkedSet

**Best for**: Large datasets requiring order, LRU cache implementations

```go
import "github.com/go4x/goal/col/set"

// Create a LinkedSet
linkedSet := set.NewLinkedSet[string]()
linkedSet.Add("first").Add("second").Add("third")

// O(1) operations with order
fmt.Println(linkedSet.Contains("first")) // true
elements := linkedSet.Elems()
fmt.Println(elements) // Output: [first second third]

// LRU cache operations
linkedSetTyped := linkedSet.(*set.LinkedSet[string])
linkedSetTyped.MoveToEnd("first") // Move to end (most recently used)
linkedSetTyped.MoveToFront("second") // Move to front
```

**Characteristics:**
- ⚡ **O(1) performance**: Fast operations with order
- 📋 **Maintains order**: Elements appear in insertion order
- 🔄 **LRU support**: MoveToEnd/MoveToFront operations
- 🎯 **Best for**: Large datasets, LRU caches, when you need both speed and order

## Decision Guide

| Use Case | Recommended Implementation | Reason |
|----------|---------------------------|--------|
| General-purpose, don't care about order | `NewHashSet()` | Fastest O(1) operations |
| Small dataset (< 1000), need order | `NewArraySet()` | Simple, memory efficient |
| Large dataset, need order | `NewLinkedSet()` | O(1) operations with order |
| Building LRU cache | `NewLinkedSet()` | Built-in LRU operations |
| Default choice | `New()` (HashSet) | Best general-purpose option |

## Common Operations

### Basic Operations

```go
import "github.com/go4x/goal/col/set"

// Create a set
mySet := set.New[int]()

// Add elements
mySet.Add(1).Add(2).Add(3).Add(1) // Duplicates ignored

// Check if empty
fmt.Println(mySet.IsEmpty()) // false

// Get size
fmt.Println(mySet.Size()) // 3

// Check containment
fmt.Println(mySet.Contains(2)) // true
fmt.Println(mySet.Contains(4)) // false

// Remove elements
mySet.Remove(2)
fmt.Println(mySet.Contains(2)) // false

// Get all elements
elements := mySet.Elems()
fmt.Println(elements) // [1 3] (order may vary for HashSet)

// Clear all elements
mySet.Clear()
fmt.Println(mySet.IsEmpty()) // true
```

### Chaining Operations

```go
import "github.com/go4x/goal/col/set"

// Method chaining for fluent API
mySet := set.New[string]().
    Add("apple").
    Add("banana").
    Add("cherry").
    Remove("banana")

fmt.Println(mySet.Elems()) // [apple cherry]
```

### Type Safety

```go
import "github.com/go4x/goal/col/set"

// Works with any comparable type
stringSet := set.New[string]()
intSet := set.New[int]()
structSet := set.New[MyStruct]()

type MyStruct struct {
    ID   int
    Name string
}

// Custom types must be comparable
structSet.Add(MyStruct{ID: 1, Name: "test"})
```

## Advanced Usage

### LRU Cache Implementation

```go
import "github.com/go4x/goal/col/set"

// Use LinkedSet for LRU cache
type LRUCache struct {
    capacity int
    items    *set.LinkedSet[string]
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        items:    set.NewLinkedSet[string]().(*set.LinkedSet[string]),
    }
}

func (c *LRUCache) Get(key string) bool {
    if c.items.Contains(key) {
        // Move to end (most recently used)
        c.items.MoveToEnd(key)
        return true
    }
    return false
}

func (c *LRUCache) Put(key string) {
    if c.items.Contains(key) {
        c.items.MoveToEnd(key)
    } else {
        if c.items.Size() >= c.capacity {
            // Remove least recently used (first element)
            elements := c.items.Elems()
            if len(elements) > 0 {
                c.items.Remove(elements[0])
            }
        }
        c.items.Add(key)
    }
}
```

### Set Operations

```go
import "github.com/go4x/goal/col/set"

// Union of two sets
func Union[T comparable](set1, set2 set.Set[T]) set.Set[T] {
    result := set.New[T]()
    
    // Add all elements from set1
    for _, elem := range set1.Elems() {
        result.Add(elem)
    }
    
    // Add all elements from set2
    for _, elem := range set2.Elems() {
        result.Add(elem)
    }
    
    return result
}

// Intersection of two sets
func Intersection[T comparable](set1, set2 set.Set[T]) set.Set[T] {
    result := set.New[T]()
    
    for _, elem := range set1.Elems() {
        if set2.Contains(elem) {
            result.Add(elem)
        }
    }
    
    return result
}

// Difference of two sets
func Difference[T comparable](set1, set2 set.Set[T]) set.Set[T] {
    result := set.New[T]()
    
    for _, elem := range set1.Elems() {
        if !set2.Contains(elem) {
            result.Add(elem)
        }
    }
    
    return result
}
```

### Polymorphic Usage

```go
import "github.com/go4x/goal/col/set"

// Function that works with any set implementation
func ProcessSet(s set.Set[string]) {
    s.Add("processed")
    fmt.Println("Set size:", s.Size())
    fmt.Println("Elements:", s.Elems())
}

func main() {
    // Works with any set type
    hashSet := set.NewHashSet[string]()
    arraySet := set.NewArraySet[string]()
    linkedSet := set.NewLinkedSet[string]()
    
    ProcessSet(hashSet)
    ProcessSet(arraySet)
    ProcessSet(linkedSet)
}
```

## Performance Characteristics

### Time Complexity

| Operation | HashSet | ArraySet | LinkedSet |
|-----------|---------|----------|-----------|
| Add | O(1) avg | O(n) if exists, O(1) if new | O(1) avg |
| Remove | O(1) avg | O(n) | O(1) avg |
| Contains | O(1) avg | O(n) | O(1) avg |
| Size/IsEmpty | O(1) | O(1) | O(1) |
| Elems | O(n) | O(n) | O(n) |
| MoveToEnd | N/A | N/A | O(1) |
| MoveToFront | N/A | N/A | O(1) |

### Memory Usage

- **HashSet**: Most memory efficient for large datasets
- **ArraySet**: Good for small datasets, linear memory growth
- **LinkedSet**: Slightly more memory overhead due to linked structure

### Performance Recommendations

1. **Use HashSet** when:
   - Order doesn't matter
   - You need maximum performance
   - Working with large datasets

2. **Use ArraySet** when:
   - Dataset is small (< 1000 elements)
   - Order is important
   - Memory usage is a concern

3. **Use LinkedSet** when:
   - You need both O(1) performance and order
   - Building LRU caches
   - Large datasets with order requirements

## Thread Safety

⚠️ **Important**: All set implementations are **NOT thread-safe**. If you need concurrent access, you must use synchronization primitives:

```go
import (
    "sync"
    "github.com/go4x/goal/col/set"
)

type SafeSet[T comparable] struct {
    mu  sync.RWMutex
    set set.Set[T]
}

func (s *SafeSet[T]) Add(elem T) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.set.Add(elem)
}

func (s *SafeSet[T]) Contains(elem T) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.set.Contains(elem)
}
```

## API Reference

### Set Interface

```go
type Set[T any] interface {
    Add(t T) Set[T]           // Add element (no duplicates)
    Remove(t T) Set[T]       // Remove element
    Size() int               // Get number of elements
    IsEmpty() bool           // Check if empty
    Contains(t T) bool       // Check if contains element
    Clear() Set[T]           // Remove all elements
    Elems() []T              // Get all elements as slice
}
```

### Constructors

| Function | Description | Use Case |
|----------|-------------|----------|
| `New[T]()` | Creates HashSet (default) | General purpose |
| `NewHashSet[T]()` | Creates HashSet | Performance critical |
| `NewArraySet[T]()` | Creates ArraySet | Small datasets, order needed |
| `NewLinkedSet[T]()` | Creates LinkedSet | Large datasets, order needed |

### LinkedSet Specific Methods

| Method | Description |
|--------|-------------|
| `MoveToEnd(elem T)` | Move element to end (most recently used) |
| `MoveToFront(elem T)` | Move element to front (least recently used) |

## Use Cases

### 1. Deduplication

```go
import "github.com/go4x/goal/col/set"

// Remove duplicates from slice
func RemoveDuplicates[T comparable](slice []T) []T {
    set := set.New[T]()
    for _, item := range slice {
        set.Add(item)
    }
    return set.Elems()
}

// Usage
numbers := []int{1, 2, 2, 3, 3, 3, 4}
unique := RemoveDuplicates(numbers)
fmt.Println(unique) // [1 2 3 4] (order may vary)
```

### 2. Membership Testing

```go
import "github.com/go4x/goal/col/set"

// Fast membership testing
allowedUsers := set.New[string]()
allowedUsers.Add("admin").Add("user").Add("guest")

func IsUserAllowed(username string) bool {
    return allowedUsers.Contains(username)
}
```

### 3. Tag Management

```go
import "github.com/go4x/goal/col/set"

// Tag system with order
type Article struct {
    ID   int
    Tags *set.LinkedSet[string]
}

func NewArticle(id int) *Article {
    return &Article{
        ID:   id,
        Tags: set.NewLinkedSet[string]().(*set.LinkedSet[string]),
    }
}

func (a *Article) AddTag(tag string) {
    a.Tags.Add(tag)
}

func (a *Article) GetTags() []string {
    return a.Tags.Elems() // Returns in insertion order
}
```

### 4. Cache Implementation

```go
import "github.com/go4x/goal/col/set"

// Simple cache with LRU eviction
type Cache struct {
    maxSize int
    items   *set.LinkedSet[string]
}

func NewCache(maxSize int) *Cache {
    return &Cache{
        maxSize: maxSize,
        items:   set.NewLinkedSet[string]().(*set.LinkedSet[string]),
    }
}

func (c *Cache) Access(key string) {
    if c.items.Contains(key) {
        c.items.MoveToEnd(key) // Mark as recently used
    } else {
        if c.items.Size() >= c.maxSize {
            // Remove least recently used
            elements := c.items.Elems()
            if len(elements) > 0 {
                c.items.Remove(elements[0])
            }
        }
        c.items.Add(key)
    }
}
```

## License

This package is part of the goal project and follows the same license terms.
