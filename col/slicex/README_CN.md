# slicex

一个功能全面的 Go 泛型切片操作包，提供增强的切片功能，具有不可变性保证和丰富的操作。

## 功能特性

- **不可变操作**：所有方法返回新切片而不修改原始切片
- **泛型类型支持**：支持任何可比较类型
- **丰富功能**：过滤、映射、排序、反转、并集、交集等
- **性能优化**：在可能的地方使用哈希映射进行 O(n+m) 操作
- **函数式编程**：可链式操作，流畅的 API
- **类型安全**：完整的编译时类型检查

## 安装

```bash
go get github.com/go4x/goal/col/slicex
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/col/slicex"
)

func main() {
    // 创建切片
    numbers := slicex.From([]int{3, 1, 4, 1, 5})
    
    // 过滤和排序
    filtered := numbers.Filter(func(x int) bool { return x > 2 })
    sorted := numbers.Sort(func(a, b int) bool { return a < b })
    
    // 原始切片保持不变
    fmt.Println(numbers.To()) // [3 1 4 1 5]
    fmt.Println(filtered.To()) // [3 4 5]
    fmt.Println(sorted.To()) // [1 1 3 4 5]
}
```

## 核心类型

### S[T] - 泛型切片类型

主要的泛型切片类型，具有增强的方法：

```go
import "github.com/go4x/goal/col/slicex"

// 从现有切片创建
numbers := slicex.From([]int{1, 2, 3, 4, 5})

// 创建新的空切片
empty := slicex.New[int]()

// 转换回 Go 切片
goSlice := numbers.To()
```

### SortableSlice[T] - 排序辅助类型

用于排序操作的辅助类型：

```go
import "github.com/go4x/goal/col/slicex"

// 创建可排序切片
sortable := slicex.NewSortableSlice([]int{3, 1, 4, 1, 5})

// 使用自定义比较器排序
sorted := sortable.Sort(func(a, b int) bool { return a < b })
```

## 基础操作

### 创建切片

```go
import "github.com/go4x/goal/col/slicex"

// 从现有切片创建
original := []int{1, 2, 3, 4, 5}
slice := slicex.From(original)

// 新的空切片
empty := slicex.New[int]()

// 从可变参数创建
slice = slicex.Of(1, 2, 3, 4, 5)

// 从函数创建
slice = slicex.Generate(5, func(i int) int { return i * 2 }) // [0, 2, 4, 6, 8]
```

### 基础属性

```go
import "github.com/go4x/goal/col/slicex"

slice := slicex.From([]int{1, 2, 3, 4, 5})

// 获取长度
fmt.Println(slice.Len()) // 5

// 检查是否为空
fmt.Println(slice.IsEmpty()) // false

// 获取索引处的元素
fmt.Println(slice.Get(2)) // 3 true

// 设置索引处的元素
newSlice := slice.Set(2, 10)
fmt.Println(newSlice.To()) // [1 2 10 4 5]
```

### 过滤和映射

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// 过滤偶数
evens := numbers.Filter(func(x int) bool { return x%2 == 0 })
fmt.Println(evens.To()) // [2 4 6 8 10]

// 映射为平方
squares := numbers.Map(func(x int) int { return x * x })
fmt.Println(squares.To()) // [1 4 9 16 25 36 49 64 81 100]

// 在一个操作中过滤和映射
result := numbers.FilterMap(func(x int) (int, bool) {
    if x%2 == 0 {
        return x * x, true
    }
    return 0, false
})
fmt.Println(result.To()) // [4 16 36 64 100]
```

### 排序

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{3, 1, 4, 1, 5, 9, 2, 6})

// 升序排序
ascending := numbers.Sort(func(a, b int) bool { return a < b })
fmt.Println(ascending.To()) // [1 1 2 3 4 5 6 9]

// 降序排序
descending := numbers.Sort(func(a, b int) bool { return a > b })
fmt.Println(descending.To()) // [9 6 5 4 3 2 1 1]

// 按自定义条件排序
strings := slicex.From([]string{"apple", "banana", "cherry"})
byLength := strings.Sort(func(a, b string) bool { return len(a) < len(b) })
fmt.Println(byLength.To()) // [apple cherry banana]
```

### 搜索

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5})

// 查找第一个元素
first := numbers.Find(func(x int) bool { return x > 2 })
fmt.Println(first) // 3 true

// 查找最后一个元素
last := numbers.FindLast(func(x int) bool { return x > 2 })
fmt.Println(last) // 5 true

// 检查是否有元素匹配
hasEven := numbers.Any(func(x int) bool { return x%2 == 0 })
fmt.Println(hasEven) // true

// 检查所有元素是否匹配
allPositive := numbers.All(func(x int) bool { return x > 0 })
fmt.Println(allPositive) // true
```

## 高级操作

### 集合操作

```go
import "github.com/go4x/goal/col/slicex"

slice1 := slicex.From([]int{1, 2, 3, 4, 5})
slice2 := slicex.From([]int{4, 5, 6, 7, 8})

// 并集（所有唯一元素）
union := slice1.Union(slice2)
fmt.Println(union.To()) // [1 2 3 4 5 6 7 8]

// 交集（公共元素）
intersection := slice1.Intersect(slice2)
fmt.Println(intersection.To()) // [4 5]

// 差集（在 slice1 中但不在 slice2 中的元素）
difference := slice1.Difference(slice2)
fmt.Println(difference.To()) // [1 2 3]

// 对称差集（在任一切片中但不在两者中的元素）
symmetricDiff := slice1.SymmetricDifference(slice2)
fmt.Println(symmetricDiff.To()) // [1 2 3 6 7 8]
```

### 分块和分组

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

// 分成大小为 3 的块
chunks := numbers.Chunk(3)
for i, chunk := range chunks {
    fmt.Printf("块 %d: %v\n", i, chunk.To())
}
// 块 0: [1 2 3]
// 块 1: [4 5 6]
// 块 2: [7 8 9]

// 按条件分组
grouped := numbers.GroupBy(func(x int) int { return x % 3 })
for key, group := range grouped {
    fmt.Printf("组 %d: %v\n", key, group.To())
}
// 组 0: [3 6 9]
// 组 1: [1 4 7]
// 组 2: [2 5 8]
```

### 归约和聚合

```go
import "github.com/go4x/goal/col/slicex"

numbers := slicex.From([]int{1, 2, 3, 4, 5})

// 求和所有元素
sum := numbers.Reduce(0, func(acc, x int) int { return acc + x })
fmt.Println(sum) // 15

// 查找最大值
max := numbers.Reduce(numbers.Get(0).Value, func(acc, x int) int {
    if x > acc {
        return x
    }
    return acc
})
fmt.Println(max) // 5

// 计算元素数量
count := numbers.Count(func(x int) bool { return x > 2 })
fmt.Println(count) // 3
```

## 工具函数

### 比较函数

```go
import "github.com/go4x/goal/col/slicex"

slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{1, 2, 4}

// 检查相等性
fmt.Println(slicex.Equal(slice1, slice2)) // true
fmt.Println(slicex.Equal(slice1, slice3)) // false

// 使用自定义函数检查相等性
fmt.Println(slicex.EqualFunc(slice1, slice2, func(a, b int) bool { return a == b })) // true
```

### 搜索函数

```go
import "github.com/go4x/goal/col/slicex"

slice := []int{1, 2, 3, 2, 4, 2}

// 查找元素索引
fmt.Println(slicex.IndexOf(slice, 2)) // 1
fmt.Println(slicex.LastIndexOf(slice, 2)) // 5

// 检查切片是否包含元素
fmt.Println(slicex.Contains(slice, 3)) // true
fmt.Println(slicex.Contains(slice, 5)) // false
```

### 转换函数

```go
import "github.com/go4x/goal/col/slicex"

slice := []int{1, 2, 3, 4, 5}

// 取前 3 个元素
first3 := slicex.Take(slice, 3)
fmt.Println(first3) // [1 2 3]

// 丢弃前 2 个元素
after2 := slicex.Drop(slice, 2)
fmt.Println(after2) // [3 4 5]

// 在条件为真时取元素
while := slicex.TakeWhile(slice, func(x int) bool { return x < 4 })
fmt.Println(while) // [1 2 3]

// 在条件为真时丢弃元素
dropWhile := slicex.DropWhile(slice, func(x int) bool { return x < 3 })
fmt.Println(dropWhile) // [3 4 5]
```

## 性能特征

### 时间复杂度

| 操作 | 时间复杂度 | 描述 |
|------|------------|------|
| Filter | O(n) | 线性扫描所有元素 |
| Map | O(n) | 对所有元素应用函数 |
| Sort | O(n log n) | 基于比较的排序 |
| Union | O(n + m) | 哈希映射去重 |
| Intersect | O(n + m) | 哈希映射交集 |
| Find | O(n) | 线性搜索 |
| Contains | O(n) | 线性搜索 |
| Chunk | O(n) | 线性扫描分组 |

### 内存使用

- **不可变操作**：每个操作创建新切片
- **哈希映射操作**：并集、交集使用 O(n + m) 内存
- **排序**：就地排序以提高效率
- **分块**：为每个块创建新切片

### 性能提示

1. **使用哈希映射操作** 处理大数据集
2. **链式操作** 避免中间分配
3. **使用 Take/Drop** 而不是切片操作
4. **考虑排序** 用于重复搜索操作

## 线程安全

⚠️ **重要**：所有切片操作都**不是线程安全的**。如果需要并发访问，必须使用同步原语：

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

## API 参考

### S[T] 方法

| 方法 | 描述 | 时间复杂度 |
|------|------|------------|
| `Len()` | 获取长度 | O(1) |
| `IsEmpty()` | 检查是否为空 | O(1) |
| `Get(index)` | 获取索引处的元素 | O(1) |
| `Set(index, value)` | 设置索引处的元素 | O(1) |
| `Filter(predicate)` | 过滤元素 | O(n) |
| `Map(transform)` | 转换元素 | O(n) |
| `Sort(comparator)` | 排序元素 | O(n log n) |
| `Find(predicate)` | 查找第一个匹配元素 | O(n) |
| `FindLast(predicate)` | 查找最后一个匹配元素 | O(n) |
| `Any(predicate)` | 检查是否有元素匹配 | O(n) |
| `All(predicate)` | 检查所有元素是否匹配 | O(n) |
| `Union(other)` | 与另一个切片的并集 | O(n + m) |
| `Intersect(other)` | 与另一个切片的交集 | O(n + m) |
| `Difference(other)` | 与另一个切片的差集 | O(n + m) |
| `SymmetricDifference(other)` | 对称差集 | O(n + m) |
| `Chunk(size)` | 分割成块 | O(n) |
| `GroupBy(keyFunc)` | 按键函数分组 | O(n) |
| `Reduce(initial, reducer)` | 归约为单个值 | O(n) |
| `Count(predicate)` | 计算匹配元素数量 | O(n) |
| `To()` | 转换为 Go 切片 | O(n) |

### 工具函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `From(slice)` | 从 Go 切片创建 | O(n) |
| `New[T]()` | 创建空切片 | O(1) |
| `Of(elements...)` | 从元素创建 | O(n) |
| `Generate(n, func)` | 从函数生成 | O(n) |
| `Equal(s1, s2)` | 检查相等性 | O(n) |
| `EqualFunc(s1, s2, eq)` | 使用函数检查相等性 | O(n) |
| `IndexOf(slice, element)` | 查找元素索引 | O(n) |
| `LastIndexOf(slice, element)` | 查找最后一个元素索引 | O(n) |
| `Contains(slice, element)` | 检查是否包含元素 | O(n) |
| `Take(slice, n)` | 取前 n 个元素 | O(n) |
| `Drop(slice, n)` | 丢弃前 n 个元素 | O(n) |
| `TakeWhile(slice, predicate)` | 在谓词为真时取元素 | O(n) |
| `DropWhile(slice, predicate)` | 在谓词为真时丢弃元素 | O(n) |

## 使用场景

### 1. 数据处理

```go
import "github.com/go4x/goal/col/slicex"

// 处理用户数据
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

### 2. 数据分析

```go
import "github.com/go4x/goal/col/slicex"

// 分析销售数据
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

### 3. API 响应处理

```go
import "github.com/go4x/goal/col/slicex"

// 处理 API 响应
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

### 4. 配置管理

```go
import "github.com/go4x/goal/col/slicex"

// 管理配置条目
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

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
