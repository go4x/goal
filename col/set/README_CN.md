# set

一个功能全面的 Go 泛型集合实现包，提供多种针对不同用例优化的集合实现。

## 功能特性

- **多种集合实现**：HashSet、ArraySet 和 LinkedSet
- **泛型类型支持**：支持任何可比较类型
- **多态接口**：统一的 `Set[T]` 接口适用于所有实现
- **性能优化**：针对不同性能需求的不同实现
- **顺序保持**：ArraySet 和 LinkedSet 保持插入顺序
- **内存高效**：针对不同场景优化的内存使用

## 安装

```bash
go get github.com/go4x/goal/col/set
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/col/set"
)

func main() {
    // 创建集合（默认为 HashSet）
    mySet := set.New[string]()
    mySet.Add("apple").Add("banana").Add("apple") // "apple" 只添加一次
    
    fmt.Println(mySet.Size()) // 输出: 2
    fmt.Println(mySet.Contains("apple")) // 输出: true
    
    // 获取所有元素
    elements := mySet.Elems()
    fmt.Println(elements) // 输出: [apple banana] (顺序可能不同)
}
```

## 集合实现

### 1. HashSet（默认）

**适用于**：通用集合操作、性能关键应用

```go
import "github.com/go4x/goal/col/set"

// 创建 HashSet
hashSet := set.NewHashSet[string]()
hashSet.Add("first").Add("second").Add("first") // 重复项被忽略

// O(1) 操作
fmt.Println(hashSet.Contains("first")) // true
hashSet.Remove("second")
fmt.Println(hashSet.Size()) // 1
```

**特点：**
- ⚡ **最快性能**：所有操作平均 O(1) 时间复杂度
- 🔀 **无顺序保证**：元素可能以任意顺序出现
- 💾 **内存高效**：内部使用哈希映射
- 🎯 **适用于**：大数据集、性能关键代码

### 2. ArraySet

**适用于**：小数据集、需要保持插入顺序

```go
import "github.com/go4x/goal/col/set"

// 创建 ArraySet
arraySet := set.NewArraySet[string]()
arraySet.Add("first").Add("second").Add("third")

// 保持插入顺序
elements := arraySet.Elems()
fmt.Println(elements) // 输出: [first second third]
```

**特点：**
- 📋 **保持顺序**：元素按插入顺序出现
- 🐌 **O(n) 操作**：线性时间复杂度
- 💾 **内存高效**：适合小数据集
- 🎯 **适用于**：小数据集（< 1000 元素）、需要顺序的场景

### 3. LinkedSet

**适用于**：需要顺序的大数据集、LRU 缓存实现

```go
import "github.com/go4x/goal/col/set"

// 创建 LinkedSet
linkedSet := set.NewLinkedSet[string]()
linkedSet.Add("first").Add("second").Add("third")

// 带顺序的 O(1) 操作
fmt.Println(linkedSet.Contains("first")) // true
elements := linkedSet.Elems()
fmt.Println(elements) // 输出: [first second third]

// LRU 缓存操作
linkedSetTyped := linkedSet.(*set.LinkedSet[string])
linkedSetTyped.MoveToEnd("first") // 移动到末尾（最近使用）
linkedSetTyped.MoveToFront("second") // 移动到开头
```

**特点：**
- ⚡ **O(1) 性能**：快速操作且保持顺序
- 📋 **保持顺序**：元素按插入顺序出现
- 🔄 **LRU 支持**：MoveToEnd/MoveToFront 操作
- 🎯 **适用于**：大数据集、LRU 缓存、需要速度和顺序的场景

## 选择指南

| 使用场景 | 推荐实现 | 原因 |
|----------|----------|------|
| 通用用途，不关心顺序 | `NewHashSet()` | 最快的 O(1) 操作 |
| 小数据集（< 1000），需要顺序 | `NewArraySet()` | 简单、内存高效 |
| 大数据集，需要顺序 | `NewLinkedSet()` | 带顺序的 O(1) 操作 |
| 构建 LRU 缓存 | `NewLinkedSet()` | 内置 LRU 操作 |
| 默认选择 | `New()` (HashSet) | 最佳通用选择 |

## 常用操作

### 基础操作

```go
import "github.com/go4x/goal/col/set"

// 创建集合
mySet := set.New[int]()

// 添加元素
mySet.Add(1).Add(2).Add(3).Add(1) // 重复项被忽略

// 检查是否为空
fmt.Println(mySet.IsEmpty()) // false

// 获取大小
fmt.Println(mySet.Size()) // 3

// 检查包含关系
fmt.Println(mySet.Contains(2)) // true
fmt.Println(mySet.Contains(4)) // false

// 移除元素
mySet.Remove(2)
fmt.Println(mySet.Contains(2)) // false

// 获取所有元素
elements := mySet.Elems()
fmt.Println(elements) // [1 3] (HashSet 顺序可能不同)

// 清空所有元素
mySet.Clear()
fmt.Println(mySet.IsEmpty()) // true
```

### 链式操作

```go
import "github.com/go4x/goal/col/set"

// 方法链式调用，流畅的 API
mySet := set.New[string]().
    Add("apple").
    Add("banana").
    Add("cherry").
    Remove("banana")

fmt.Println(mySet.Elems()) // [apple cherry]
```

### 类型安全

```go
import "github.com/go4x/goal/col/set"

// 支持任何可比较类型
stringSet := set.New[string]()
intSet := set.New[int]()
structSet := set.New[MyStruct]()

type MyStruct struct {
    ID   int
    Name string
}

// 自定义类型必须是可比较的
structSet.Add(MyStruct{ID: 1, Name: "test"})
```

## 高级用法

### LRU 缓存实现

```go
import "github.com/go4x/goal/col/set"

// 使用 LinkedSet 实现 LRU 缓存
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
        // 移动到末尾（最近使用）
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
            // 移除最近最少使用的元素（第一个元素）
            elements := c.items.Elems()
            if len(elements) > 0 {
                c.items.Remove(elements[0])
            }
        }
        c.items.Add(key)
    }
}
```

### 集合运算

```go
import "github.com/go4x/goal/col/set"

// 两个集合的并集
func Union[T comparable](set1, set2 set.Set[T]) set.Set[T] {
    result := set.New[T]()
    
    // 添加 set1 的所有元素
    for _, elem := range set1.Elems() {
        result.Add(elem)
    }
    
    // 添加 set2 的所有元素
    for _, elem := range set2.Elems() {
        result.Add(elem)
    }
    
    return result
}

// 两个集合的交集
func Intersection[T comparable](set1, set2 set.Set[T]) set.Set[T] {
    result := set.New[T]()
    
    for _, elem := range set1.Elems() {
        if set2.Contains(elem) {
            result.Add(elem)
        }
    }
    
    return result
}

// 两个集合的差集
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

### 多态使用

```go
import "github.com/go4x/goal/col/set"

// 适用于任何集合实现的函数
func ProcessSet(s set.Set[string]) {
    s.Add("processed")
    fmt.Println("集合大小:", s.Size())
    fmt.Println("元素:", s.Elems())
}

func main() {
    // 适用于任何集合类型
    hashSet := set.NewHashSet[string]()
    arraySet := set.NewArraySet[string]()
    linkedSet := set.NewLinkedSet[string]()
    
    ProcessSet(hashSet)
    ProcessSet(arraySet)
    ProcessSet(linkedSet)
}
```

## 性能特征

### 时间复杂度

| 操作 | HashSet | ArraySet | LinkedSet |
|------|---------|----------|-----------|
| Add | O(1) 平均 | 存在时 O(n)，新元素时 O(1) | O(1) 平均 |
| Remove | O(1) 平均 | O(n) | O(1) 平均 |
| Contains | O(1) 平均 | O(n) | O(1) 平均 |
| Size/IsEmpty | O(1) | O(1) | O(1) |
| Elems | O(n) | O(n) | O(n) |
| MoveToEnd | 不适用 | 不适用 | O(1) |
| MoveToFront | 不适用 | 不适用 | O(1) |

### 内存使用

- **HashSet**：大数据集最内存高效
- **ArraySet**：小数据集良好，线性内存增长
- **LinkedSet**：由于链表结构，内存开销稍大

### 性能建议

1. **使用 HashSet** 当：
   - 顺序不重要
   - 需要最大性能
   - 处理大数据集

2. **使用 ArraySet** 当：
   - 数据集较小（< 1000 元素）
   - 顺序很重要
   - 内存使用是考虑因素

3. **使用 LinkedSet** 当：
   - 需要 O(1) 性能和顺序
   - 构建 LRU 缓存
   - 大数据集且需要顺序

## 线程安全

⚠️ **重要**：所有集合实现都**不是线程安全的**。如果需要并发访问，必须使用同步原语：

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

## API 参考

### Set 接口

```go
type Set[T any] interface {
    Add(t T) Set[T]           // 添加元素（无重复）
    Remove(t T) Set[T]        // 移除元素
    Size() int                // 获取元素数量
    IsEmpty() bool            // 检查是否为空
    Contains(t T) bool        // 检查是否包含元素
    Clear() Set[T]            // 清空所有元素
    Elems() []T               // 获取所有元素的切片
}
```

### 构造函数

| 函数 | 描述 | 使用场景 |
|------|------|----------|
| `New[T]()` | 创建 HashSet（默认） | 通用用途 |
| `NewHashSet[T]()` | 创建 HashSet | 性能关键 |
| `NewArraySet[T]()` | 创建 ArraySet | 小数据集，需要顺序 |
| `NewLinkedSet[T]()` | 创建 LinkedSet | 大数据集，需要顺序 |

### LinkedSet 特定方法

| 方法 | 描述 |
|------|------|
| `MoveToEnd(elem T)` | 将元素移动到末尾（最近使用） |
| `MoveToFront(elem T)` | 将元素移动到开头（最少使用） |

## 使用场景

### 1. 去重

```go
import "github.com/go4x/goal/col/set"

// 从切片中移除重复项
func RemoveDuplicates[T comparable](slice []T) []T {
    set := set.New[T]()
    for _, item := range slice {
        set.Add(item)
    }
    return set.Elems()
}

// 使用
numbers := []int{1, 2, 2, 3, 3, 3, 4}
unique := RemoveDuplicates(numbers)
fmt.Println(unique) // [1 2 3 4] (顺序可能不同)
```

### 2. 成员资格测试

```go
import "github.com/go4x/goal/col/set"

// 快速成员资格测试
allowedUsers := set.New[string]()
allowedUsers.Add("admin").Add("user").Add("guest")

func IsUserAllowed(username string) bool {
    return allowedUsers.Contains(username)
}
```

### 3. 标签管理

```go
import "github.com/go4x/goal/col/set"

// 带顺序的标签系统
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
    return a.Tags.Elems() // 按插入顺序返回
}
```

### 4. 缓存实现

```go
import "github.com/go4x/goal/col/set"

// 带 LRU 淘汰的简单缓存
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
        c.items.MoveToEnd(key) // 标记为最近使用
    } else {
        if c.items.Size() >= c.maxSize {
            // 移除最近最少使用的元素
            elements := c.items.Elems()
            if len(elements) > 0 {
                c.items.Remove(elements[0])
            }
        }
        c.items.Add(key)
    }
}
```

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
