# value

一个功能全面的 Go 泛型值处理包，提供值操作、条件逻辑、空值检查和安全操作的实用工具。

## 功能特性

- **泛型类型支持**：支持任何可比较类型
- **空值处理**：nil 和空值的安全操作
- **条件逻辑**：函数式条件操作方法
- **安全操作**：具有适当错误处理的 panic 安全操作
- **指针操作**：安全的解引用和指针操作
- **值合并**：回退值操作
- **类型安全**：完整的编译时类型检查

## 安装

```bash
go get github.com/go4x/goal/value
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/value"
)

func main() {
    // 条件逻辑
    result := value.IfElse(age >= 18, "adult", "minor")
    
    // 安全操作
    safeValue := value.Must(strconv.Atoi("123"))
    
    // 值合并
    fallback := value.Or("", "", "default")
}
```

## 核心函数

### 条件逻辑

#### IfElse - 条件操作符

```go
import "github.com/go4x/goal/value"

// 基础条件逻辑
result := value.IfElse(age >= 18, "adult", "minor")
fmt.Println(result) // 如果 age >= 18 则为 "adult"，否则为 "minor"

// 不同类型
max := value.IfElse(a > b, a, b)
status := value.IfElse(isActive, "active", "inactive")
```

#### If - 单条件

```go
import "github.com/go4x/goal/value"

// 如果条件为真则返回值，否则返回零值
result := value.If(age >= 18, "adult")
fmt.Println(result) // 如果 age >= 18 则为 "adult"，否则为 ""

// 数字类型
positive := value.If(x > 0, x)
fmt.Println(positive) // 如果 x > 0 则为 x，否则为 0
```

#### When - 带函数的条件

```go
import "github.com/go4x/goal/value"

// 如果条件为真则执行函数
result := value.When(age >= 18, func() string {
    return "你是成年人"
})
fmt.Println(result) // 如果 age >= 18 则为函数结果，否则为 ""
```

#### WhenElse - 带两个函数的条件

```go
import "github.com/go4x/goal/value"

// 根据条件执行不同的函数
result := value.WhenElse(age >= 18, 
    func() string { return "adult" },
    func() string { return "minor" },
)
fmt.Println(result) // 根据条件返回 "adult" 或 "minor"
```

### 值合并

#### Or - 第一个非零值

```go
import "github.com/go4x/goal/value"

// 获取第一个非零值
result := value.Or("", "", "fallback", "ignored")
fmt.Println(result) // "fallback"

// 数字类型
number := value.Or(0, 0, 42, 100)
fmt.Println(number) // 42

// 所有零值
empty := value.Or("", "", "")
fmt.Println(empty) // "" (字符串的零值)
```

#### OrElse - 带默认值的第一个非零值

```go
import "github.com/go4x/goal/value"

// 获取第一个非零值，带默认回退
result := value.OrElse("default", "", "", "fallback")
fmt.Println(result) // "fallback"

// 所有零值返回默认值
defaultResult := value.OrElse("default", "", "", "")
fmt.Println(defaultResult) // "default"
```

#### Coalesce - 多值合并

```go
import "github.com/go4x/goal/value"

// 合并多个值
result := value.Coalesce("", "first", "second")
fmt.Println(result) // "first"

// 不同类型
var p1, p2 *int
val := 42
p3 := &val
number := value.Coalesce(p1, p2, p3)
fmt.Println(number) // 42
```

### 安全操作

#### Must - 错误时 Panic

```go
import (
    "strconv"
    "github.com/go4x/goal/value"
)

// 当你确定操作会成功时安全使用
result := value.Must(strconv.Atoi("123"))
fmt.Println(result) // 123

// 如果字符串不是有效整数，这会 panic
// result := value.Must(strconv.Atoi("invalid")) // panic!
```

#### SafeDeref - 安全指针解引用

```go
import "github.com/go4x/goal/value"

var ptr *int
var val int = 42
ptr = &val

// 安全解引用
result := value.SafeDeref(ptr)
fmt.Println(result) // 42 true

// 带默认值的安全解引用
result = value.SafeDerefDef(ptr, 0)
fmt.Println(result) // 42

// 处理 nil 指针
var nilPtr *int
result = value.SafeDeref(nilPtr)
fmt.Println(result) // 0 false

result = value.SafeDerefDef(nilPtr, -1)
fmt.Println(result) // -1
```

### 值操作

#### Value - 从接口提取值

```go
import "github.com/go4x/goal/value"

// 从 interface{} 提取值
var data interface{} = "hello"
result := value.Value(data)
fmt.Println(result) // "hello" true

// 处理 nil 接口
var nilData interface{}
result = value.Value(nilData)
fmt.Println(result) // nil false
```

#### Def - 默认值

```go
import "github.com/go4x/goal/value"

// 提供默认值
result := value.Def("", "default")
fmt.Println(result) // "default"

// 非空值
result = value.Def("hello", "default")
fmt.Println(result) // "hello"
```

## 高级用法

### 配置处理

```go
import "github.com/go4x/goal/value"

// 带回退的配置
type Config struct {
    Host     string
    Port     int
    Timeout  int
    Debug    bool
}

func LoadConfig() Config {
    return Config{
        Host:    value.OrElse("localhost", os.Getenv("HOST"), ""),
        Port:    value.OrElse(8080, parseInt(os.Getenv("PORT")), 0),
        Timeout: value.OrElse(30, parseInt(os.Getenv("TIMEOUT")), 0),
        Debug:   value.IfElse(os.Getenv("DEBUG") == "true", true, false),
    }
}

func parseInt(s string) int {
    if s == "" {
        return 0
    }
    return value.Must(strconv.Atoi(s))
}
```

### 错误处理

```go
import "github.com/go4x/goal/value"

// 安全错误处理
func ProcessData(data string) (string, error) {
    if data == "" {
        return "", errors.New("数据为空")
    }
    
    // 安全处理数据
    result := value.Must(transformData(data))
    return result, nil
}

func transformData(data string) (string, error) {
    if data == "" {
        return "", errors.New("空数据")
    }
    return strings.ToUpper(data), nil
}
```

### 数据验证

```go
import "github.com/go4x/goal/value"

// 验证用户输入
func ValidateUser(user User) error {
    if user.Name == "" {
        return errors.New("姓名是必需的")
    }
    
    if user.Age <= 0 {
        return errors.New("年龄必须为正数")
    }
    
    if user.Email == nil || *user.Email == "" {
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

### API 响应处理

```go
import "github.com/go4x/goal/value"

// 处理带回退的 API 响应
func ProcessAPIResponse(response *APIResponse) string {
    // 使用合并获取回退值
    message := value.Coalesce(
        response.Message,
        response.Error,
        "无可用消息",
    )
    
    // 安全解引用
    status := value.SafeDerefDef(response.Status, "unknown")
    
    // 条件格式化
    return value.IfElse(
        message != nil && *message != "",
        fmt.Sprintf("[%s] %s", status, *message),
        fmt.Sprintf("[%s] 无消息", status),
    )
}

type APIResponse struct {
    Message *string
    Error   *string
    Status  *string
}
```

### 函数式编程

```go
import "github.com/go4x/goal/value"

// 数据处理的函数式方法
func ProcessItems(items []Item) []Item {
    return slicex.From(items).
        Filter(func(item Item) bool {
            return item.Name != "" && item.Price != 0.0
        }).
        Map(func(item Item) Item {
            return Item{
                Name:  value.OrElse("Unknown", item.Name, ""),
                Price: value.OrElse(0.0, item.Price, 0.0),
                Category: value.IfElse(
                    item.Category != "",
                    item.Category,
                    "uncategorized",
                ),
            }
        }).
        To()
}

type Item struct {
    Name     string
    Price    float64
    Category string
}
```

## 性能考虑

### 合并操作

- **Or/OrElse**：O(n) - 线性扫描值
- **Coalesce**：O(n) - 线性扫描值
- **SafeDeref**：O(1) - 直接指针解引用

### 条件操作

- **IfElse/If**：O(1) - 直接条件评估
- **When/WhenElse**：O(1) - 函数调用开销
- **Must**：O(1) - 错误检查和必要时 panic

### 最佳实践

1. **使用 Or/OrElse** 进行回退值链
2. **使用 SafeDeref** 进行安全指针操作
3. **使用 Must** 仅在你确定操作会成功时
4. **使用 IfElse/If** 进行条件值选择
5. **使用 Coalesce** 进行基于指针的回退链

## 线程安全

⚠️ **重要**：此包中的所有函数都是**线程安全的**，可以从多个 goroutine 并发调用。但是，如果并发访问，被操作的基础数据必须是线程安全的。

## API 参考

### 条件函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `IfElse(condition, v1, v2)` | 如果条件为真返回 v1，否则返回 v2 | O(1) |
| `If(condition, value)` | 如果条件为真返回值，否则返回零值 | O(1) |
| `When(condition, func)` | 如果条件为真执行函数 | O(1) |
| `WhenElse(condition, func1, func2)` | 如果条件为真执行 func1，否则执行 func2 | O(1) |

### 合并函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `Or(values...)` | 返回第一个非零值 | O(n) |
| `OrElse(default, values...)` | 返回第一个非零值或默认值 | O(n) |
| `Coalesce(values...)` | 返回第一个非零值 | O(n) |
| `CoalesceValue(values...)` | 返回第一个非零值 | O(n) |
| `CoalesceValueDef(default, values...)` | 返回第一个非零值或默认值 | O(n) |

### 安全操作

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `Must(value, error)` | 返回值或在错误时 panic | O(1) |
| `SafeDeref(pointer)` | 安全指针解引用 | O(1) |
| `SafeDerefDef(pointer, default)` | 带默认值的安全指针解引用 | O(1) |
| `Value(interface{})` | 从接口提取值 | O(1) |
| `Def(value, default)` | 返回值或空时返回默认值 | O(1) |

## 使用场景

### 1. 配置管理

```go
import "github.com/go4x/goal/value"

// 基于环境的配置
func LoadConfig() Config {
    return Config{
        Host:    value.OrElse("localhost", os.Getenv("HOST"), ""),
        Port:    value.OrElse(8080, parseInt(os.Getenv("PORT")), 0),
        Debug:   value.IfElse(os.Getenv("DEBUG") == "true", true, false),
        Timeout: value.OrElse(30, parseInt(os.Getenv("TIMEOUT")), 0),
    }
}
```

### 2. 数据验证

```go
import "github.com/go4x/goal/value"

// 验证输入数据
func ValidateInput(data InputData) error {
    if data.Name == "" {
        return errors.New("姓名是必需的")
    }
    
    if data.Age <= 0 {
        return errors.New("年龄必须为正数")
    }
    
    if data.Email == nil || *data.Email == "" {
        return errors.New("邮箱是必需的")
    }
    
    return nil
}
```

### 3. API 响应处理

```go
import "github.com/go4x/goal/value"

// 处理带回退的 API 响应
func ProcessResponse(response *APIResponse) string {
    message := value.Coalesce(
        response.Message,
        response.Error,
        "无可用消息",
    )
    
    status := value.SafeDerefDef(response.Status, "unknown")
    
    return value.IfElse(
        message != nil && *message != "",
        fmt.Sprintf("[%s] %s", status, *message),
        fmt.Sprintf("[%s] 无消息", status),
    )
}
```

### 4. 错误处理

```go
import "github.com/go4x/goal/value"

// 安全错误处理
func ProcessData(data string) (string, error) {
    if data == "" {
        return "", errors.New("数据为空")
    }
    
    result := value.Must(transformData(data))
    return result, nil
}
```

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
