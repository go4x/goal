# mapx

A comprehensive generic map implementation package for Go that provides multiple map implementations optimized for different use cases.

## Features

- **Multiple Map Implementations**: Regular Map, ArrayMap, and LinkedMap
- **Generic Type Support**: Works with any comparable key type and any value type
- **Polymorphic Interface**: Unified `Map[K, V]` interface for all implementations
- **Performance Optimized**: Different implementations for different performance requirements
- **Order Preservation**: ArrayMap and LinkedMap maintain insertion order
- **Memory Efficient**: Optimized memory usage for different scenarios

## Installation

```bash
go get github.com/go4x/goal/col/mapx
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/col/mapx"
)

func main() {
    // Create a map (defaults to regular map)
    myMap := mapx.New[string, int]()
    myMap.Put("apple", 1).Put("banana", 2).Put("apple", 3) // Overwrites "apple"
    
    fmt.Println(myMap.Size()) // Output: 2
    fmt.Println(myMap.Get("apple")) // Output: 3 true
    
    // Get all entries
    entries := myMap.Entries()
    fmt.Println(entries) // Output: [apple:3 banana:2] (order may vary)
}
```

## Map Implementations

### 1. Regular Map (Default)

**Best for**: General-purpose map operations, performance-critical applications

```go
import "github.com/go4x/goal/col/mapx"

// Create a regular map
regularMap := mapx.New[string, int]()
regularMap.Put("first", 1).Put("second", 2).Put("first", 3) // Overwrites

// O(1) operations
fmt.Println(regularMap.Get("first")) // 3 true
regularMap.Remove("second")
fmt.Println(regularMap.Size()) // 1
```

**Characteristics:**
- ⚡ **Fastest performance**: O(1) average-case for all operations
- 🔀 **No order guarantee**: Entries may appear in any order
- 💾 **Memory efficient**: Uses Go's built-in map internally
- 🎯 **Best for**: Large datasets, performance-critical code

### 2. ArrayMap

**Best for**: Small datasets, when insertion order matters

```go
import "github.com/go4x/goal/col/mapx"

// Create an ArrayMap
arrayMap := mapx.NewArrayMap[string, int]()
arrayMap.Put("first", 1).Put("second", 2).Put("third", 3)

// Maintains insertion order
entries := arrayMap.Entries()
fmt.Println(entries) // Output: [first:1 second:2 third:3]
```

**Characteristics:**
- 📋 **Maintains order**: Entries appear in insertion order
- 🐌 **O(n) operations**: Linear time complexity
- 💾 **Memory efficient**: Good for small datasets
- 🎯 **Best for**: Small datasets (< 1000 entries), when order matters

### 3. LinkedMap

**Best for**: Large datasets requiring order, LRU cache implementations

```go
import "github.com/go4x/goal/col/mapx"

// Create a LinkedMap
linkedMap := mapx.NewLinkedMap[string, int]()
linkedMap.Put("first", 1).Put("second", 2).Put("third", 3)

// O(1) operations with order
fmt.Println(linkedMap.Get("first")) // 1 true
entries := linkedMap.Entries()
fmt.Println(entries) // Output: [first:1 second:2 third:3]

// LRU cache operations
linkedMapTyped := linkedMap.(*mapx.LinkedMap[string, int])
linkedMapTyped.MoveToEnd("first") // Move to end (most recently used)
linkedMapTyped.MoveToFront("second") // Move to front
```

**Characteristics:**
- ⚡ **O(1) performance**: Fast operations with order
- 📋 **Maintains order**: Entries appear in insertion order
- 🔄 **LRU support**: MoveToEnd/MoveToFront operations
- 🎯 **Best for**: Large datasets, LRU caches, when you need both speed and order

## Decision Guide

| Use Case | Recommended Implementation | Reason |
|----------|---------------------------|--------|
| General-purpose, don't care about order | `New[K, V]()` | Fastest O(1) operations |
| Small dataset (< 1000), need order | `NewArrayMap[K, V]()` | Simple, memory efficient |
| Large dataset, need order | `NewLinkedMap[K, V]()` | O(1) operations with order |
| Building LRU cache | `NewLinkedMap[K, V]()` | Built-in LRU operations |
| Default choice | `New[K, V]()` | Best general-purpose option |

## Common Operations

### Basic Operations

```go
import "github.com/go4x/goal/col/mapx"

// Create a map
myMap := mapx.New[string, int]()

// Put entries
myMap.Put("apple", 1).Put("banana", 2).Put("cherry", 3)

// Check if empty
fmt.Println(myMap.IsEmpty()) // false

// Get size
fmt.Println(myMap.Size()) // 3

// Get value
value, exists := myMap.Get("apple")
fmt.Println(value, exists) // 1 true

// Check if key exists
fmt.Println(myMap.Contains("banana")) // true
fmt.Println(myMap.Contains("grape")) // false

// Remove entry
myMap.Remove("banana")
fmt.Println(myMap.Contains("banana")) // false

// Get all entries
entries := myMap.Entries()
fmt.Println(entries) // [apple:1 cherry:3] (order may vary for regular map)

// Clear all entries
myMap.Clear()
fmt.Println(myMap.IsEmpty()) // true
```

### Chaining Operations

```go
import "github.com/go4x/goal/col/mapx"

// Method chaining for fluent API
myMap := mapx.New[string, int]().
    Put("apple", 1).
    Put("banana", 2).
    Put("cherry", 3).
    Remove("banana")

fmt.Println(myMap.Entries()) // [apple:1 cherry:3]
```

### Type Safety

```go
import "github.com/go4x/goal/col/mapx"

// Works with any comparable key type and any value type
stringIntMap := mapx.New[string, int]()
intStringMap := mapx.New[int, string]()
structMap := mapx.New[MyKey, MyValue]()

type MyKey struct {
    ID   int
    Name string
}

type MyValue struct {
    Data string
    Flag bool
}

// Custom types must be comparable for keys
structMap.Put(MyKey{ID: 1, Name: "test"}, MyValue{Data: "value", Flag: true})
```

## Performance Characteristics

### Time Complexity

| Operation | Regular Map | ArrayMap | LinkedMap |
|-----------|-------------|----------|-----------|
| Put | O(1) avg | O(n) if exists, O(1) if new | O(1) avg |
| Get | O(1) avg | O(n) | O(1) avg |
| Remove | O(1) avg | O(n) | O(1) avg |
| Contains | O(1) avg | O(n) | O(1) avg |
| Size/IsEmpty | O(1) | O(1) | O(1) |
| Entries | O(n) | O(n) | O(n) |
| MoveToEnd | N/A | N/A | O(1) |
| MoveToFront | N/A | N/A | O(1) |

### Memory Usage

- **Regular Map**: Most memory efficient for large datasets
- **ArrayMap**: Good for small datasets, linear memory growth
- **LinkedMap**: Slightly more memory overhead due to linked structure

### Performance Recommendations

1. **Use Regular Map** when:
   - Order doesn't matter
   - You need maximum performance
   - Working with large datasets

2. **Use ArrayMap** when:
   - Dataset is small (< 1000 entries)
   - Order is important
   - Memory usage is a concern

3. **Use LinkedMap** when:
   - You need both O(1) performance and order
   - Building LRU caches
   - Large datasets with order requirements

## Thread Safety

⚠️ **Important**: All map implementations are **NOT thread-safe**. If you need concurrent access, you must use synchronization primitives:

```go
import (
    "sync"
    "github.com/go4x/goal/col/mapx"
)

type SafeMap[K comparable, V any] struct {
    mu sync.RWMutex
    m  mapx.Map[K, V]
}

func (s *SafeMap[K, V]) Put(key K, value V) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m.Put(key, value)
}

func (s *SafeMap[K, V]) Get(key K) (V, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.m.Get(key)
}
```

## API Reference

### Map Interface

```go
type Map[K comparable, V any] interface {
    Put(key K, value V) Map[K, V]     // Put key-value pair
    Get(key K) (V, bool)              // Get value by key
    Remove(key K) Map[K, V]            // Remove key
    Size() int                         // Get number of entries
    IsEmpty() bool                     // Check if empty
    Contains(key K) bool              // Check if key exists
    Clear() Map[K, V]                  // Remove all entries
    Entries() []Entry[K, V]            // Get all entries
}
```

### Entry Type

```go
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}
```

### Constructors

| Function | Description | Use Case |
|----------|-------------|----------|
| `New[K, V]()` | Creates regular map (default) | General purpose |
| `NewArrayMap[K, V]()` | Creates ArrayMap | Small datasets, order needed |
| `NewLinkedMap[K, V]()` | Creates LinkedMap | Large datasets, order needed |

### LinkedMap Specific Methods

| Method | Description |
|--------|-------------|
| `MoveToEnd(key K)` | Move entry to end (most recently used) |
| `MoveToFront(key K)` | Move entry to front (least recently used) |

## Use Cases

### 1. Configuration Management

```go
import "github.com/go4x/goal/col/mapx"

// Configuration with order
type Config struct {
    settings *mapx.LinkedMap[string, interface{}]
}

func NewConfig() *Config {
    return &Config{
        settings: mapx.NewLinkedMap[string, interface{}]().(*mapx.LinkedMap[string, interface{}]),
    }
}

func (c *Config) Set(key string, value interface{}) {
    c.settings.Put(key, value)
}

func (c *Config) Get(key string) (interface{}, bool) {
    return c.settings.Get(key)
}

func (c *Config) GetAll() []mapx.Entry[string, interface{}] {
    return c.settings.Entries() // Returns in insertion order
}
```

### 2. Session Management

```go
import "github.com/go4x/goal/col/mapx"

// Session store with LRU eviction
type SessionStore struct {
    maxSize int
    sessions *mapx.LinkedMap[string, Session]
}

type Session struct {
    UserID string
    Data   map[string]interface{}
}

func NewSessionStore(maxSize int) *SessionStore {
    return &SessionStore{
        maxSize: maxSize,
        sessions: mapx.NewLinkedMap[string, Session]().(*mapx.LinkedMap[string, Session]),
    }
}

func (s *SessionStore) Get(sessionID string) (Session, bool) {
    if session, exists := s.sessions.Get(sessionID); exists {
        s.sessions.MoveToEnd(sessionID) // Mark as recently used
        return session, true
    }
    return Session{}, false
}

func (s *SessionStore) Put(sessionID string, session Session) {
    if s.sessions.Contains(sessionID) {
        s.sessions.Put(sessionID, session)
        s.sessions.MoveToEnd(sessionID)
    } else {
        if s.sessions.Size() >= s.maxSize {
            // Remove least recently used
            entries := s.sessions.Entries()
            if len(entries) > 0 {
                s.sessions.Remove(entries[0].Key)
            }
        }
        s.sessions.Put(sessionID, session)
    }
}
```

### 3. Cache Implementation

```go
import "github.com/go4x/goal/col/mapx"

// Simple cache with LRU eviction
type Cache struct {
    maxSize int
    items   *mapx.LinkedMap[string, interface{}]
}

func NewCache(maxSize int) *Cache {
    return &Cache{
        maxSize: maxSize,
        items:   mapx.NewLinkedMap[string, interface{}]().(*mapx.LinkedMap[string, interface{}]),
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    if value, exists := c.items.Get(key); exists {
        c.items.MoveToEnd(key) // Mark as recently used
        return value, true
    }
    return nil, false
}

func (c *Cache) Put(key string, value interface{}) {
    if c.items.Contains(key) {
        c.items.Put(key, value)
        c.items.MoveToEnd(key)
    } else {
        if c.items.Size() >= c.maxSize {
            // Remove least recently used
            entries := c.items.Entries()
            if len(entries) > 0 {
                c.items.Remove(entries[0].Key)
            }
        }
        c.items.Put(key, value)
    }
}
```

## License

This package is part of the goal project and follows the same license terms.