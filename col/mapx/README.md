# MapX Package - Map Implementations Comparison

This package provides three different map implementations, each optimized for different use cases.

## 📊 Quick Comparison Table

| Feature | **M** (Regular) | **ArrayMap** | **LinkedMap** |
|---------|----------------|--------------|---------------|
| **Implementation** | Go built-in map | Two parallel slices | Doubly linked list + hash map |
| **Insertion Order** | ❌ Not guaranteed | ✅ Guaranteed | ✅ Guaranteed |
| **Get Operation** | O(1) | O(n) | O(1) |
| **Put Operation** | O(1) | O(1) new / O(n) existing | O(1) |
| **Delete Operation** | O(1) | O(n) | O(1) |
| **Memory Overhead** | Low | Very Low | Medium |
| **Cache Performance** | Good | Excellent (small data) | Good |
| **Special Features** | None | First/Last | MoveToEnd/MoveToFront |
| **Best For** | General use | Small datasets | Large datasets |

## 🎯 Detailed Analysis

### M (Regular Map)
**Based on Go's built-in map**

#### ✅ Advantages:
- **Fastest performance**: O(1) for all operations
- **Lowest memory overhead**: Uses Go's optimized built-in map
- **Battle-tested**: Based on Go's standard library
- **No pointer chasing**: Excellent cache performance
- **Simple**: Easy to understand and debug

#### ❌ Disadvantages:
- **No order guarantee**: Iteration order is random
- **No special features**: Basic map operations only

#### 🎯 Use Cases:
- General-purpose mappings where order doesn't matter
- Performance-critical applications
- Large datasets where you don't need order
- When you need the fastest possible performance

#### 📈 Performance Characteristics:
```
Dataset Size: Any
Get: O(1)
Put: O(1)
Delete: O(1)
Memory: O(n)
```

### ArrayMap
**Based on two parallel slices**

#### ✅ Advantages:
- **Simple implementation**: Easy to understand and debug
- **Minimal memory overhead**: Just two slices, no pointers
- **Cache-friendly**: Sequential memory access for small datasets
- **Deterministic order**: Always maintains insertion order
- **Memory efficient**: No extra metadata or pointers

#### ❌ Disadvantages:
- **O(n) operations**: Linear search for Get/Delete
- **Poor scaling**: Performance degrades with size
- **No O(1) operations**: Except appending new keys

#### 🎯 Use Cases:
- Small to medium datasets (< 1000 elements)
- Configuration maps where order matters
- Prototyping or simple applications
- Memory-constrained environments
- When you rarely access elements (mostly iteration)

#### 📈 Performance Characteristics:
```
Dataset Size: < 1000 elements
Get: O(n) - linear search
Put: O(1) new / O(n) existing
Delete: O(n)
Memory: O(n) - minimal overhead
```

### LinkedMap
**Based on doubly linked list + hash map**

#### ✅ Advantages:
- **O(1) operations**: Fast Get, Put, Delete operations
- **Order guaranteed**: Maintains insertion order
- **Scalable**: Good performance for large datasets
- **Advanced features**: MoveToEnd, MoveToFront for LRU caches
- **Efficient deletions**: O(1) node removal

#### ❌ Disadvantages:
- **Higher memory overhead**: Pointers + hash map overhead
- **Complex implementation**: More code to maintain
- **Pointer chasing**: Potential cache misses
- **Overhead for small datasets**: Hash map overhead

#### 🎯 Use Cases:
- Large datasets (> 1000 elements)
- Frequent lookups, insertions, or deletions
- LRU cache implementations
- Performance-critical applications requiring order
- When you need both O(1) operations and order

#### 📈 Performance Characteristics:
```
Dataset Size: > 1000 elements
Get: O(1) - hash map lookup
Put: O(1) - hash map + linked list
Delete: O(1) - hash map + linked list
Memory: O(n) - hash map + node pointers
```

## 🤔 Decision Guide

### Choose **M** when:
- You need the fastest possible performance
- Order doesn't matter
- You're working with large datasets
- You want the simplest solution

### Choose **ArrayMap** when:
- You need insertion order
- Dataset is small (< 1000 elements)
- You have memory constraints
- You rarely access elements (mostly iteration)
- You want a simple, predictable implementation

### Choose **LinkedMap** when:
- You need insertion order AND performance
- Dataset is large (> 1000 elements)
- You need frequent lookups/deletions
- You're building an LRU cache
- You need advanced features like MoveToEnd

## 🔄 Polymorphic Usage

All three implementations implement the `Map[K, V]` interface, allowing for polymorphic usage:

```go
// Factory function
func createMap(mapType string) mapx.Map[string, int] {
    switch mapType {
    case "regular":
        return mapx.New[string, int]()
    case "array":
        return mapx.NewArrayMap[string, int]()
    case "linked":
        return mapx.NewLinkedMap[string, int]()
    default:
        return mapx.New[string, int]()
    }
}

// Use polymorphically
var maps []mapx.Map[string, int]
maps = append(maps, mapx.New[string, int]())
maps = append(maps, mapx.NewArrayMap[string, int]())
maps = append(maps, mapx.NewLinkedMap[string, int]())

for _, m := range maps {
    m.Put("key", 42)    // Same interface
    fmt.Println(m.Size()) // Same behavior
}
```

## 📝 Examples

### Basic Usage
```go
// Regular map
regular := mapx.New[string, int]()
regular.Put("a", 1).Put("b", 2)

// Array map (ordered)
array := mapx.NewArrayMap[string, int]()
array.Put("a", 1).Put("b", 2)
// Keys() returns ["a", "b"] in insertion order

// Linked map (ordered + fast)
linked := mapx.NewLinkedMap[string, int]()
linked.Put("a", 1).Put("b", 2)
// Keys() returns ["a", "b"] in insertion order
```

### LRU Cache Implementation
```go
type LRUCache struct {
    capacity int
    cache    *mapx.LinkedMap[string, int]
}

func (lru *LRUCache) Get(key string) (int, bool) {
    if val, ok := lru.cache.Get(key); ok {
        // Move to end (most recently used)
        lru.cache.MoveToEnd(key)
        return val, true
    }
    return 0, false
}

func (lru *LRUCache) Put(key string, value int) {
    if lru.cache.Contains(key) {
        lru.cache.Put(key, value)
        lru.cache.MoveToEnd(key)
    } else {
        if lru.cache.Size() >= lru.capacity {
            // Remove least recently used (first element)
            if k, _, ok := lru.cache.First(); ok {
                lru.cache.Del(k)
            }
        }
        lru.cache.Put(key, value)
    }
}
```

## ⚡ Performance Tips

1. **For small datasets (< 100 elements)**: ArrayMap might actually be faster than LinkedMap due to cache locality
2. **For large datasets (> 1000 elements)**: LinkedMap is significantly faster than ArrayMap
3. **For memory-constrained environments**: Use ArrayMap
4. **For maximum performance**: Use M (regular map) if order doesn't matter
5. **For LRU caches**: LinkedMap is the clear winner with MoveToEnd/MoveToFront

## 🧪 Testing

All implementations are thoroughly tested with 99.4% coverage. Run tests with:

```bash
go test ./col/mapx -v -cover
```

## 📚 See Also

- `array_map_test.go` - ArrayMap tests and examples
- `linked_map_test.go` - LinkedMap tests and examples  
- `interface_example_test.go` - Polymorphic usage examples
