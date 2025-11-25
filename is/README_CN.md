# is

一个功能全面的 Go 值检查和比较包，提供布尔操作、零值/nil/空值检查和值比较的实用工具。

## 功能特性

- **布尔操作**：逻辑取反和布尔值检查
- **零值检查**：检查值是否为其类型的零值
- **Nil 检查**：对引用类型进行全面的 nil 检查
- **空值检查**：检查不同类型中的空值
- **值比较**：相等性和深度相等性检查
- **类型安全**：使用泛型进行完整的编译时类型检查

## 安装

```bash
go get github.com/go4x/goal/is
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/is"
)

func main() {
    // 零值检查
    if is.Zero(0) {
        fmt.Println("值为零")
    }
    
    // Nil 检查
    var ptr *int
    if is.Nil(ptr) {
        fmt.Println("指针为 nil")
    }
    
    // 空值检查
    if is.Empty("") {
        fmt.Println("字符串为空")
    }
    
    // 值比较
    if is.Equal(42, 42) {
        fmt.Println("值相等")
    }
}
```

## 核心函数

### 布尔操作

#### Not - 逻辑取反

```go
import "github.com/go4x/goal/is"

// 对布尔值取反
fmt.Println(is.Not(true))   // false
fmt.Println(is.Not(false))  // true

// 在条件中使用
supported := true
if is.Not(supported) {
    fmt.Println("不支持")
} else {
    fmt.Println("支持")
}
```

#### True - 检查是否为真

```go
import "github.com/go4x/goal/is"

// 检查值是否为 true
fmt.Println(is.True(true))   // true
fmt.Println(is.True(false)) // false
```

#### False - 检查是否为假

```go
import "github.com/go4x/goal/is"

// 检查值是否为 false
fmt.Println(is.False(true))   // false
fmt.Println(is.False(false)) // true
```

### 零值检查

#### Zero - 零值检查

```go
import "github.com/go4x/goal/is"

// 检查零值
fmt.Println(is.Zero(0))        // true
fmt.Println(is.Zero(""))      // true
fmt.Println(is.Zero(false))   // true
fmt.Println(is.Zero(42))      // false
fmt.Println(is.Zero("hello")) // false

// 不同类型
fmt.Println(is.Zero(0.0))      // true
fmt.Println(is.Zero(0.0))     // true
fmt.Println(is.Zero([]int{})) // false (空切片不是零值)
```

#### NotZero - 非零值检查

```go
import "github.com/go4x/goal/is"

// 检查非零值
fmt.Println(is.NotZero(42))     // true
fmt.Println(is.NotZero("hello")) // true
fmt.Println(is.NotZero(0))      // false
fmt.Println(is.NotZero(""))     // false
```

### Nil 检查

#### Nil - Nil 检查

```go
import "github.com/go4x/goal/is"

var ptr *int
fmt.Println(is.Nil(ptr))        // true
fmt.Println(is.Nil(nil))       // true
fmt.Println(is.Nil([]int{}))  // false (空切片不是 nil)
fmt.Println(is.Nil((*int)(nil))) // true

// 不同的引用类型
var m map[string]int
fmt.Println(is.Nil(m))          // true

var ch chan int
fmt.Println(is.Nil(ch))         // true
```

#### NotNil - 非 Nil 检查

```go
import "github.com/go4x/goal/is"

// 检查非 nil 值
ptr := &42
fmt.Println(is.NotNil(ptr))     // true
fmt.Println(is.NotNil([]int{})) // true (空切片不是 nil)
fmt.Println(is.NotNil(nil))     // false

// 映射和通道
m := make(map[string]int)
fmt.Println(is.NotNil(m))       // true

ch := make(chan int)
fmt.Println(is.NotNil(ch))      // true
```

### 空值检查

#### Empty - 空值检查

```go
import "github.com/go4x/goal/is"

// 检查空值
fmt.Println(is.Empty(""))                    // true
fmt.Println(is.Empty([]int{}))              // true
fmt.Println(is.Empty(map[string]int{}))      // true
fmt.Println(is.Empty(0))                    // true
fmt.Println(is.Empty("hello"))              // false
fmt.Println(is.Empty([]int{1, 2}))          // false

// 不同类型
fmt.Println(is.Empty(nil))                  // true
var ptr *int
fmt.Println(is.Empty(ptr))                  // true
```

#### NotEmpty - 非空值检查

```go
import "github.com/go4x/goal/is"

// 检查非空值
fmt.Println(is.NotEmpty("hello"))           // true
fmt.Println(is.NotEmpty([]int{1, 2}))       // true
fmt.Println(is.NotEmpty(42))                // true
fmt.Println(is.NotEmpty(""))                // false
fmt.Println(is.NotEmpty([]int{}))           // false
```

### 值比较

#### Equal - 相等性检查

```go
import "github.com/go4x/goal/is"

// 比较值
fmt.Println(is.Equal(42, 42))              // true
fmt.Println(is.Equal("hello", "hello"))    // true
fmt.Println(is.Equal(42, 43))              // false
fmt.Println(is.Equal("hello", "hi"))       // false

// 不同类型
fmt.Println(is.Equal(true, true))           // true
fmt.Println(is.Equal(3.14, 3.14))          // true
```

#### NotEqual - 不等性检查

```go
import "github.com/go4x/goal/is"

// 检查不等
fmt.Println(is.NotEqual(42, 43))            // true
fmt.Println(is.NotEqual("hello", "hi"))    // true
fmt.Println(is.NotEqual(42, 42))            // false
fmt.Println(is.NotEqual("hello", "hello"))  // false
```

#### DeepEqual - 深度相等性检查

```go
import "github.com/go4x/goal/is"

// 使用反射进行深度比较
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{1, 2, 4}

fmt.Println(is.DeepEqual(slice1, slice2)) // true
fmt.Println(is.DeepEqual(slice1, slice3)) // false

// 比较映射
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"a": 1, "b": 2}
map3 := map[string]int{"a": 1, "b": 3}

fmt.Println(is.DeepEqual(map1, map2)) // true
fmt.Println(is.DeepEqual(map1, map3)) // false

// 比较结构体
type Person struct {
    Name string
    Age  int
}

p1 := Person{Name: "John", Age: 30}
p2 := Person{Name: "John", Age: 30}
p3 := Person{Name: "Jane", Age: 30}

fmt.Println(is.DeepEqual(p1, p2)) // true
fmt.Println(is.DeepEqual(p1, p3)) // false
```

## 高级用法

### 数据验证

```go
import "github.com/go4x/goal/is"

// 验证用户输入
func ValidateUser(user User) error {
    if is.Empty(user.Name) {
        return errors.New("姓名是必需的")
    }
    
    if is.Zero(user.Age) || user.Age < 0 {
        return errors.New("年龄必须为正数")
    }
    
    if is.Nil(user.Email) || is.Empty(*user.Email) {
        return errors.New("邮箱是必需的")
    }
    
    return nil
}

type User struct {
    Name  string
    Age   int
    Email *string
}
```

### 配置验证

```go
import "github.com/go4x/goal/is"

// 验证配置
func ValidateConfig(config Config) error {
    if is.Empty(config.Host) {
        return errors.New("主机不能为空")
    }
    
    if is.Zero(config.Port) || config.Port <= 0 {
        return errors.New("端口必须为正数")
    }
    
    if is.Nil(config.Database) {
        return errors.New("数据库配置是必需的")
    }
    
    return nil
}
```

### API 响应处理

```go
import "github.com/go4x/goal/is"

// 处理 API 响应
func ProcessResponse(response *APIResponse) error {
    if is.Nil(response) {
        return errors.New("响应为 nil")
    }
    
    if is.Empty(response.Data) {
        return errors.New("响应数据为空")
    }
    
    // 检查响应是否成功
    if is.NotEqual(response.Status, "success") {
        return fmt.Errorf("意外状态: %s", response.Status)
    }
    
    return nil
}
```

### 函数式编程

```go
import "github.com/go4x/goal/is"

// 基于检查过滤项目
func FilterValidItems(items []Item) []Item {
    var valid []Item
    for _, item := range items {
        if is.NotEmpty(item.Name) && is.NotZero(item.Price) {
            valid = append(valid, item)
        }
    }
    return valid
}

// 比较集合
func AreCollectionsEqual(a, b []int) bool {
    return is.DeepEqual(a, b)
}
```

## 性能考虑

### 零值检查

- **Zero/NotZero**：O(1) - 与零值直接比较
- **Empty/NotEmpty**：O(1) 对于大多数类型，O(n) 对于切片/映射（长度检查）
- **Nil/NotNil**：O(1) - 基于反射的 nil 检查

### 比较操作

- **Equal/NotEqual**：O(1) - 直接相等性检查
- **DeepEqual**：O(n) - 基于反射的深度比较

### 最佳实践

1. **使用 Zero/NotZero** 对可比较类型进行简单的零值检查
2. **使用 Empty/NotEmpty** 对所有类型进行全面的空值检查
3. **使用 Nil/NotNil** 进行引用类型的 nil 检查
4. **使用 Equal/NotEqual** 进行简单的值比较
5. **使用 DeepEqual** 进行复杂的嵌套结构比较

## 线程安全

⚠️ **重要**：此包中的所有函数都是**线程安全的**，可以从多个 goroutine 并发调用。但是，如果并发访问，被检查的基础数据必须是线程安全的。

## API 参考

### 布尔函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `Not(v)` | 返回布尔值的逻辑取反 | O(1) |
| `True(v)` | 如果值为 true 则返回 true | O(1) |
| `False(v)` | 如果值为 false 则返回 true | O(1) |

### 检查函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `Zero(value)` | 检查值是否为零 | O(1) |
| `NotZero(value)` | 检查值是否不为零 | O(1) |
| `Nil(value)` | 检查值是否为 nil | O(1) |
| `NotNil(value)` | 检查值是否不为 nil | O(1) |
| `Empty(value)` | 检查值是否为空 | O(1) |
| `NotEmpty(value)` | 检查值是否不为空 | O(1) |

### 比较函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `Equal(a, b)` | 检查值是否相等 | O(1) |
| `NotEqual(a, b)` | 检查值是否不相等 | O(1) |
| `DeepEqual(a, b)` | 使用反射进行深度比较 | O(n) |

## 使用场景

### 1. 输入验证

```go
import "github.com/go4x/goal/is"

// 验证输入数据
func ValidateInput(data InputData) error {
    if is.Empty(data.Name) {
        return errors.New("姓名是必需的")
    }
    
    if is.Zero(data.Age) || data.Age < 0 {
        return errors.New("年龄必须为正数")
    }
    
    if is.Nil(data.Email) || is.Empty(*data.Email) {
        return errors.New("邮箱是必需的")
    }
    
    return nil
}
```

### 2. 配置检查

```go
import "github.com/go4x/goal/is"

// 检查配置值
func CheckConfig(config Config) error {
    if is.Empty(config.Host) {
        return errors.New("主机不能为空")
    }
    
    if is.Zero(config.Port) {
        return errors.New("必须指定端口")
    }
    
    return nil
}
```

### 3. 数据比较

```go
import "github.com/go4x/goal/is"

// 比较数据结构
func CompareData(old, new Data) bool {
    return is.DeepEqual(old, new)
}

// 检查值是否改变
func HasChanged(old, new string) bool {
    return is.NotEqual(old, new)
}
```

### 4. 条件逻辑

```go
import "github.com/go4x/goal/is"

// 在条件逻辑中使用
func ProcessData(data string) error {
    if is.Empty(data) {
        return errors.New("数据为空")
    }
    
    if is.NotEmpty(data) {
        // 处理非空数据
        return process(data)
    }
    
    return nil
}
```

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。

