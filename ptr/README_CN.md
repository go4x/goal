# ptr - Go 指针操作工具包

一个为 Go 提供安全高效指针操作的综合性工具包，具有完整的 nil 安全支持。

## 特性

- **类型安全**：使用 Go 泛型确保完整的类型安全
- **nil 安全**：所有函数都安全处理 nil 指针，不会引发 panic
- **功能全面**：为常见使用场景提供完整的指针操作集合
- **性能优化**：高效实现，最小化内存分配
- **易于使用**：简单直观的 API 设计

## 安装

```bash
go get github.com/go4x/goal/ptr
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/ptr"
)

func main() {
    // 基本指针操作
    value := 42
    p := ptr.To(value)        // 将值转换为指针
    result := ptr.From(p)    // 将指针转换为值
    
    // 安全解引用
    var nilPtr *int
    safeValue := ptr.Deref(nilPtr)  // 返回 0，不会 panic
    
    // 指针验证
    if ptr.IsNil(nilPtr) {
        fmt.Println("指针为 nil")
    }
    
    fmt.Printf("值: %d, 结果: %d, 安全值: %d\n", value, result, safeValue)
}
```

## API 参考

### 基本操作

#### `To[T any](v T) *T`
将值转换为指针。

```go
ptr := ptr.To(42)  // 值为 42 的 *int
```

#### `From[T any](v *T) T`
将指针转换为值。**注意**：如果指针为 nil 会引发 panic。

```go
value := ptr.From(&someInt)  // int 值
```

### 安全操作

#### `IsNil[T any](v *T) bool`
检查指针是否为 nil。

```go
var p *int
if ptr.IsNil(p) {
    fmt.Println("指针为 nil")
}
```

#### `IsNotNil[T any](v *T) bool`
检查指针是否不为 nil。

```go
p := ptr.To(42)
if ptr.IsNotNil(p) {
    fmt.Println("指针不为 nil")
}
```

#### `Deref[T any](v *T) T`
安全地解引用指针，如果为 nil 则返回零值。

```go
var nilPtr *int
value := ptr.Deref(nilPtr)  // 返回 0，不会 panic
```

#### `DerefOr[T any](v *T, defaultValue T) T`
安全地解引用指针，如果为 nil 则返回指定值。

```go
var nilPtr *int
value := ptr.DerefOr(nilPtr, 100)  // 返回 100
```

#### `ValueOr[T any](v *T, defaultValue T) T`
如果指针不为 nil 则返回其值，否则返回默认值。

```go
var nilPtr *int
value := ptr.ValueOr(nilPtr, 100)  // 返回 100
```

#### `ValueOrDefault[T any](v *T) T`
如果指针不为 nil 则返回其值，否则返回零值。

```go
var nilPtr *int
value := ptr.ValueOrDefault(nilPtr)  // 返回 0
```

### 比较操作

#### `Equal[T comparable](a, b *T) bool`
比较两个指针的值是否相等（处理 nil 情况）。

```go
p1 := ptr.To(42)
p2 := ptr.To(42)
p3 := ptr.To(100)
var nilPtr *int

fmt.Println(ptr.Equal(p1, p2))    // true
fmt.Println(ptr.Equal(p1, p3))    // false
fmt.Println(ptr.Equal(p1, nilPtr)) // false
```

#### `DeepEqual[T any](a, b *T) bool`
对两个指针的值进行深度比较。

```go
type Person struct {
    Name string
    Age  int
}

p1 := &Person{Name: "Alice", Age: 30}
p2 := &Person{Name: "Alice", Age: 30}
p3 := &Person{Name: "Bob", Age: 25}

fmt.Println(ptr.DeepEqual(p1, p2))  // true
fmt.Println(ptr.DeepEqual(p1, p3))  // false
```

### 克隆操作

#### `Clone[T any](v *T) *T`
克隆指针指向的值。

```go
original := ptr.To(42)
cloned := ptr.Clone(original)
// cloned 是一个具有相同值的新指针
```

#### `CloneSlice[T any](v []*T) []*T`
克隆指针切片。

```go
slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
cloned := ptr.CloneSlice(slice)
// cloned 是一个包含克隆指针的新切片
```

### 切片操作

#### `ToSlice[T any](v []T) []*T`
将值切片转换为指针切片。

```go
values := []int{1, 2, 3}
pointers := ptr.ToSlice(values)  // []*int
```

#### `FromSlice[T any](v []*T) []T`
将指针切片转换为值切片。

```go
pointers := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
values := ptr.FromSlice(pointers)  // []int{1, 2, 3}
```

#### `Filter[T any](v []*T) []*T`
过滤指针切片，返回非 nil 指针的切片。

```go
slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
filtered := ptr.Filter(slice)  // 包含非 nil 指针的 []*int
```

#### `FilterValues[T any](v []*T) []T`
过滤指针切片，返回非 nil 指针的值切片。

```go
slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
values := ptr.FilterValues(slice)  // []int{1, 3, 5}
```

#### `Map[T, U any](v []*T, fn func(*T) *U) []*U`
使用提供的函数映射指针切片。

```go
slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
strings := ptr.Map(slice, func(ptr *int) *string {
    if ptr == nil {
        return nil
    }
    return ptr.To(strconv.Itoa(*ptr))
})
```

#### `MapValues[T, U any](v []*T, fn func(T) U) []U`
使用提供的函数映射指针切片的值。

```go
slice := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
strings := ptr.MapValues(slice, func(val int) string {
    return strconv.Itoa(val)
})  // []string{"1", "2", "3"}
```

### 查询操作

#### `Any[T any](v []*T) bool`
检查切片中是否有任何非 nil 指针。

```go
slice := []*int{ptr.To(1), nil, ptr.To(3)}
hasAny := ptr.Any(slice)  // true
```

#### `All[T any](v []*T) bool`
检查切片中是否所有指针都不为 nil。

```go
slice1 := []*int{ptr.To(1), ptr.To(2), ptr.To(3)}
slice2 := []*int{ptr.To(1), nil, ptr.To(3)}

all1 := ptr.All(slice1)  // true
all2 := ptr.All(slice2)  // false
```

#### `Count[T any](v []*T) int`
计算切片中非 nil 指针的数量。

```go
slice := []*int{ptr.To(1), nil, ptr.To(3), nil, ptr.To(5)}
count := ptr.Count(slice)  // 3
```

#### `First[T any](v []*T) *T`
返回切片中第一个非 nil 指针。

```go
slice := []*int{nil, ptr.To(2), ptr.To(3)}
first := ptr.First(slice)  // 指向 2 的 *int
```

#### `Last[T any](v []*T) *T`
返回切片中最后一个非 nil 指针。

```go
slice := []*int{ptr.To(1), ptr.To(2), nil, ptr.To(4)}
last := ptr.Last(slice)  // 指向 4 的 *int
```

### 修改操作

#### `Set[T any](v *T, value T)`
设置指针指向的值。

```go
value := 42
ptr := &value
ptr.Set(ptr, 100)  // *ptr 现在是 100
```

#### `Zero[T any](v *T)`
将指针指向的值设置为其零值。

```go
value := 42
ptr := &value
ptr.Zero(ptr)  // *ptr 现在是 0
```

#### `Swap[T any](a, b *T)`
交换两个指针指向的值。

```go
a := ptr.To(1)
b := ptr.To(2)
ptr.Swap(a, b)  // *a 现在是 2，*b 现在是 1
```

## 使用示例

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/ptr"
)

func main() {
    // 在值和指针之间转换
    value := 42
    p := ptr.To(value)
    result := ptr.From(p)
    
    // 安全操作
    var nilPtr *int
    safeValue := ptr.Deref(nilPtr)  // 0，不会 panic
    
    fmt.Printf("原始值: %d, 结果: %d, 安全值: %d\n", value, result, safeValue)
}
```

### 处理切片

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/ptr"
)

func main() {
    // 创建指针切片
    values := []int{1, 2, 3, 4, 5}
    pointers := ptr.ToSlice(values)
    
    // 过滤掉 nil 指针（如果有的话）
    valid := ptr.Filter(pointers)
    
    // 计算非 nil 指针数量
    count := ptr.Count(pointers)
    
    // 获取第一个非 nil 指针
    first := ptr.First(pointers)
    
    fmt.Printf("有效指针: %d, 数量: %d, 第一个: %d\n", 
        len(valid), count, ptr.Deref(first))
}
```

### 复杂使用场景（结构体）

```go
package main

import (
    "fmt"
    "strings"
    "github.com/go4x/goal/ptr"
)

type User struct {
    ID    int
    Name  string
    Email *string // 可选字段
}

func main() {
    users := []*User{
        {ID: 1, Name: "Alice", Email: ptr.To("alice@example.com")},
        {ID: 2, Name: "Bob", Email: nil}, // 没有邮箱
        {ID: 3, Name: "Charlie", Email: ptr.To("charlie@example.com")},
    }
    
    // 过滤出非 nil 用户指针
    validUsers := ptr.Filter(users)
    
    // 获取所有邮箱地址（过滤掉 nil 邮箱）
    emails := []string{}
    for _, user := range validUsers {
        if ptr.IsNotNil(user.Email) {
            emails = append(emails, ptr.Deref(user.Email))
        }
    }
    
    // 检查是否所有用户指针都不为 nil
    allValid := ptr.All(users)
    
    fmt.Printf("有效用户: %d\n", len(validUsers))
    fmt.Printf("邮箱: %s\n", strings.Join(emails, ", "))
    fmt.Printf("所有用户有效: %t\n", allValid)
}
```

## 性能

`ptr` 包专为性能而设计，具有最小的内存分配：

- **零拷贝操作**：大多数操作不分配新内存
- **高效过滤**：针对切片操作优化的算法
- **类型安全**：使用泛型进行编译时类型检查
- **nil 安全**：不会因 nil 指针解引用而引发运行时 panic

### 性能基准

- **基本操作**（To, From, IsNil, Deref等）：~0.27ns/op，0 分配
- **切片操作**：高效的内存使用和分配
- **查询操作**（Any, All, Count）：高性能，特别是 Any 和 First 操作
- **深度比较**：合理的性能表现

## 测试

运行测试：

```bash
go test ./ptr
```

运行覆盖率测试：

```bash
go test ./ptr -cover
```

运行示例：

```bash
go test ./ptr -run Example
```

运行基准测试：

```bash
go test ./ptr -bench=.
```

## 许可证

此包是 `goal` 项目的一部分，遵循相同的许可证条款。

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 相关包

- [mathx](../mathx) - 具有小数精度的数学工具
- [stringx](../stringx) - 字符串操作工具
- [iox](../iox) - I/O 工具和文件操作
