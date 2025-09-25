# mapx

一个功能全面的 Go 泛型映射实现包，提供多种针对不同用例优化的映射实现。

## 功能特性

- **多种映射实现**：常规映射、ArrayMap 和 LinkedMap
- **泛型类型支持**：支持任何可比较键类型和任何值类型
- **多态接口**：统一的 `Map[K, V]` 接口适用于所有实现
- **性能优化**：针对不同性能需求的不同实现
- **顺序保持**：ArrayMap 和 LinkedMap 保持插入顺序
- **内存高效**：针对不同场景优化的内存使用

## 安装

```bash
go get github.com/go4x/goal/col/mapx
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/col/mapx"
)

func main() {
    // 创建映射（默认为常规映射）
    myMap := mapx.New[string, int]()
    myMap.Put("apple", 1).Put("banana", 2).Put("apple", 3) // 覆盖 "apple"
    
    fmt.Println(myMap.Size()) // 输出: 2
    fmt.Println(myMap.Get("apple")) // 输出: 3 true
    
    // 获取所有条目
    entries := myMap.Entries()
    fmt.Println(entries) // 输出: [apple:3 banana:2] (顺序可能不同)
}
```

## 映射实现

### 1. 常规映射（默认）

**适用于**：通用映射操作、性能关键应用

```go
import "github.com/go4x/goal/col/mapx"

// 创建常规映射
regularMap := mapx.New[string, int]()
regularMap.Put("first", 1).Put("second", 2).Put("first", 3) // 覆盖

// O(1) 操作
fmt.Println(regularMap.Get("first")) // 3 true
regularMap.Remove("second")
fmt.Println(regularMap.Size()) // 1
```

**特点：**
- ⚡ **最快性能**：所有操作平均 O(1) 时间复杂度
- 🔀 **无顺序保证**：条目可能以任意顺序出现
- 💾 **内存高效**：内部使用 Go 内置映射
- 🎯 **适用于**：大数据集、性能关键代码

### 2. ArrayMap

**适用于**：小数据集、需要保持插入顺序

```go
import "github.com/go4x/goal/col/mapx"

// 创建 ArrayMap
arrayMap := mapx.NewArrayMap[string, int]()
arrayMap.Put("first", 1).Put("second", 2).Put("third", 3)

// 保持插入顺序
entries := arrayMap.Entries()
fmt.Println(entries) // 输出: [first:1 second:2 third:3]
```

**特点：**
- 📋 **保持顺序**：条目按插入顺序出现
- 🐌 **O(n) 操作**：线性时间复杂度
- 💾 **内存高效**：适合小数据集
- 🎯 **适用于**：小数据集（< 1000 条目）、需要顺序的场景

### 3. LinkedMap

**适用于**：需要顺序的大数据集、LRU 缓存实现

```go
import "github.com/go4x/goal/col/mapx"

// 创建 LinkedMap
linkedMap := mapx.NewLinkedMap[string, int]()
linkedMap.Put("first", 1).Put("second", 2).Put("third", 3)

// 带顺序的 O(1) 操作
fmt.Println(linkedMap.Get("first")) // 1 true
entries := linkedMap.Entries()
fmt.Println(entries) // 输出: [first:1 second:2 third:3]

// LRU 缓存操作
linkedMapTyped := linkedMap.(*mapx.LinkedMap[string, int])
linkedMapTyped.MoveToEnd("first") // 移动到末尾（最近使用）
linkedMapTyped.MoveToFront("second") // 移动到开头
```

**特点：**
- ⚡ **O(1) 性能**：快速操作且保持顺序
- 📋 **保持顺序**：条目按插入顺序出现
- 🔄 **LRU 支持**：MoveToEnd/MoveToFront 操作
- 🎯 **适用于**：大数据集、LRU 缓存、需要速度和顺序的场景

## 选择指南

| 使用场景 | 推荐实现 | 原因 |
|----------|----------|------|
| 通用用途，不关心顺序 | `New[K, V]()` | 最快的 O(1) 操作 |
| 小数据集（< 1000），需要顺序 | `NewArrayMap[K, V]()` | 简单、内存高效 |
| 大数据集，需要顺序 | `NewLinkedMap[K, V]()` | 带顺序的 O(1) 操作 |
| 构建 LRU 缓存 | `NewLinkedMap[K, V]()` | 内置 LRU 操作 |
| 默认选择 | `New[K, V]()` | 最佳通用选择 |

## 常用操作

### 基础操作

```go
import "github.com/go4x/goal/col/mapx"

// 创建映射
myMap := mapx.New[string, int]()

// 添加条目
myMap.Put("apple", 1).Put("banana", 2).Put("cherry", 3)

// 检查是否为空
fmt.Println(myMap.IsEmpty()) // false

// 获取大小
fmt.Println(myMap.Size()) // 3

// 获取值
value, exists := myMap.Get("apple")
fmt.Println(value, exists) // 1 true

// 检查键是否存在
fmt.Println(myMap.Contains("banana")) // true
fmt.Println(myMap.Contains("grape")) // false

// 移除条目
myMap.Remove("banana")
fmt.Println(myMap.Contains("banana")) // false

// 获取所有条目
entries := myMap.Entries()
fmt.Println(entries) // [apple:1 cherry:3] (常规映射顺序可能不同)

// 清空所有条目
myMap.Clear()
fmt.Println(myMap.IsEmpty()) // true
```

### 链式操作

```go
import "github.com/go4x/goal/col/mapx"

// 方法链式调用，流畅的 API
myMap := mapx.New[string, int]().
    Put("apple", 1).
    Put("banana", 2).
    Put("cherry", 3).
    Remove("banana")

fmt.Println(myMap.Entries()) // [apple:1 cherry:3]
```

### 类型安全

```go
import "github.com/go4x/goal/col/mapx"

// 支持任何可比较键类型和任何值类型
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

// 自定义类型必须是可比较的（用于键）
structMap.Put(MyKey{ID: 1, Name: "test"}, MyValue{Data: "value", Flag: true})
```

## 性能特征

### 时间复杂度

| 操作 | 常规映射 | ArrayMap | LinkedMap |
|------|----------|----------|-----------|
| Put | O(1) 平均 | 存在时 O(n)，新条目时 O(1) | O(1) 平均 |
| Get | O(1) 平均 | O(n) | O(1) 平均 |
| Remove | O(1) 平均 | O(n) | O(1) 平均 |
| Contains | O(1) 平均 | O(n) | O(1) 平均 |
| Size/IsEmpty | O(1) | O(1) | O(1) |
| Entries | O(n) | O(n) | O(n) |
| MoveToEnd | 不适用 | 不适用 | O(1) |
| MoveToFront | 不适用 | 不适用 | O(1) |

### 内存使用

- **常规映射**：大数据集最内存高效
- **ArrayMap**：小数据集良好，线性内存增长
- **LinkedMap**：由于链表结构，内存开销稍大

### 性能建议

1. **使用常规映射** 当：
   - 顺序不重要
   - 需要最大性能
   - 处理大数据集

2. **使用 ArrayMap** 当：
   - 数据集较小（< 1000 条目）
   - 顺序很重要
   - 内存使用是考虑因素

3. **使用 LinkedMap** 当：
   - 需要 O(1) 性能和顺序
   - 构建 LRU 缓存
   - 大数据集且需要顺序

## 线程安全

⚠️ **重要**：所有映射实现都**不是线程安全的**。如果需要并发访问，必须使用同步原语：

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

## API 参考

### Map 接口

```go
type Map[K comparable, V any] interface {
    Put(key K, value V) Map[K, V]     // 添加键值对
    Get(key K) (V, bool)              // 根据键获取值
    Remove(key K) Map[K, V]            // 移除键
    Size() int                         // 获取条目数量
    IsEmpty() bool                     // 检查是否为空
    Contains(key K) bool              // 检查键是否存在
    Clear() Map[K, V]                  // 清空所有条目
    Entries() []Entry[K, V]            // 获取所有条目
}
```

### Entry 类型

```go
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}
```

### 构造函数

| 函数 | 描述 | 使用场景 |
|------|------|----------|
| `New[K, V]()` | 创建常规映射（默认） | 通用用途 |
| `NewArrayMap[K, V]()` | 创建 ArrayMap | 小数据集，需要顺序 |
| `NewLinkedMap[K, V]()` | 创建 LinkedMap | 大数据集，需要顺序 |

### LinkedMap 特定方法

| 方法 | 描述 |
|------|------|
| `MoveToEnd(key K)` | 将条目移动到末尾（最近使用） |
| `MoveToFront(key K)` | 将条目移动到开头（最少使用） |

## 使用场景

### 1. 配置管理

```go
import "github.com/go4x/goal/col/mapx"

// 带顺序的配置
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
    return c.settings.Entries() // 按插入顺序返回
}
```

### 2. 会话管理

```go
import "github.com/go4x/goal/col/mapx"

// 带 LRU 淘汰的会话存储
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
        s.sessions.MoveToEnd(sessionID) // 标记为最近使用
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
            // 移除最近最少使用的
            entries := s.sessions.Entries()
            if len(entries) > 0 {
                s.sessions.Remove(entries[0].Key)
            }
        }
        s.sessions.Put(sessionID, session)
    }
}
```

### 3. 缓存实现

```go
import "github.com/go4x/goal/col/mapx"

// 带 LRU 淘汰的简单缓存
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
        c.items.MoveToEnd(key) // 标记为最近使用
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
            // 移除最近最少使用的
            entries := c.items.Entries()
            if len(entries) > 0 {
                c.items.Remove(entries[0].Key)
            }
        }
        c.items.Put(key, value)
    }
}
```

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
