# value

一个功能全面的 Go 泛型值处理包，提供值操作、条件逻辑、空值检查和安全操作的实用工具。

## 功能特性

- **泛型类型支持**：支持任何可比较类型
- **空值处理**：全面的 nil、空值和零值检查
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
    
    // 空值检查
    if value.IsNotEmpty(data) {
        fmt.Println("数据不为空")
    }
    
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

### 空值检查

#### IsZero/IsNotZero - 零值检查

```go
import "github.com/go4x/goal/value"

// 检查零值
fmt.Println(value.IsZero(0))        // true
fmt.Println(value.IsZero(""))      // true
fmt.Println(value.IsZero(false))  // true
fmt.Println(value.IsZero(42))     // false
fmt.Println(value.IsZero("hello")) // false

// 检查非零值
fmt.Println(value.IsNotZero(42))     // true
fmt.Println(value.IsNotZero("hello")) // true
fmt.Println(value.IsNotZero(0))      // false
```

#### IsNil/IsNotNil - Nil 检查

```go
import "github.com/go4x/goal/value"

var ptr *int
fmt.Println(value.IsNil(ptr))        // true
fmt.Println(value.IsNil(nil))       // true
fmt.Println(value.IsNil([]int{}))   // false (空切片不是 nil)
fmt.Println(value.IsNil((*int)(nil))) // true

// 检查非 nil 值
ptr = &42
fmt.Println(value.IsNotNil(ptr))     // true
fmt.Println(value.IsNotNil([]int{})) // true (空切片不是 nil)
fmt.Println(value.IsNotNil(nil))     // false
```

#### IsEmpty/IsNotEmpty - 空值检查

```go
import "github.com/go4x/goal/value"

// 检查空值
fmt.Println(value.IsEmpty(""))                    // true
fmt.Println(value.IsEmpty([]int{}))              // true
fmt.Println(value.IsEmpty(map[string]int{}))     // true
fmt.Println(value.IsEmpty(0))                   // true
fmt.Println(value.IsEmpty("hello"))              // false
fmt.Println(value.IsEmpty([]int{1, 2}))         // false

// 检查非空值
fmt.Println(value.IsNotEmpty("hello"))           // true
fmt.Println(value.IsNotEmpty([]int{1, 2}))      // true
fmt.Println(value.IsNotEmpty(""))                // false
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

### 比较操作

#### Equal/NotEqual - 值比较

```go
import "github.com/go4x/goal/value"

// 比较值
fmt.Println(value.Equal(42, 42))     // true
fmt.Println(value.Equal("hello", "hello")) // true
fmt.Println(value.Equal(42, 43))    // false

// 检查不等
fmt.Println(value.NotEqual(42, 43))  // true
fmt.Println(value.NotEqual(42, 42))  // false
```

#### DeepEqual - 深度值比较

```go
import "github.com/go4x/goal/value"

// 使用反射进行深度比较
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := []int{1, 2, 4}

fmt.Println(value.DeepEqual(slice1, slice2)) // true
fmt.Println(value.DeepEqual(slice1, slice3)) // false

// 比较结构体
type Person struct {
    Name string
    Age  int
}

p1 := Person{Name: "John", Age: 30}
p2 := Person{Name: "John", Age: 30}
p3 := Person{Name: "Jane", Age: 30}

fmt.Println(value.DeepEqual(p1, p2)) // true
fmt.Println(value.DeepEqual(p1, p3)) // false
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
    if value.IsEmpty(s) {
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
    if value.IsEmpty(data) {
        return "", errors.New("数据为空")
    }
    
    // 安全处理数据
    result := value.Must(transformData(data))
    return result, nil
}

func transformData(data string) (string, error) {
    if value.IsEmpty(data) {
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
    if value.IsEmpty(user.Name) {
        return errors.New("姓名是必需的")
    }
    
    if value.IsZero(user.Age) || user.Age < 0 {
        return errors.New("年龄必须为正数")
    }
    
    if value.IsNil(user.Email) || value.IsEmpty(*user.Email) {
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
        value.IsNotEmpty(message),
        fmt.Sprintf("[%s] %s", status, message),
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
            return value.IsNotEmpty(item.Name) && value.IsNotZero(item.Price)
        }).
        Map(func(item Item) Item {
            return Item{
                Name:  value.OrElse("Unknown", item.Name, ""),
                Price: value.OrElse(0.0, item.Price, 0.0),
                Category: value.IfElse(
                    value.IsNotEmpty(item.Category),
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

### 零值检查

- **IsZero/IsNotZero**：O(1) - 与零值直接比较
- **IsEmpty/IsNotEmpty**：O(1) 对于大多数类型，O(n) 对于切片/映射（长度检查）
- **IsNil/IsNotNil**：O(1) - 基于反射的 nil 检查

### 合并操作

- **Or/OrElse**：O(n) - 线性扫描值
- **Coalesce**：O(n) - 线性扫描值
- **SafeDeref**：O(1) - 直接指针解引用

### 条件操作

- **IfElse/If**：O(1) - 直接条件评估
- **When/WhenElse**：O(1) - 函数调用开销
- **Must**：O(1) - 错误检查和必要时 panic

### 最佳实践

1. **使用 IsZero/IsNotZero** 进行简单的零值检查
2. **使用 IsEmpty/IsNotEmpty** 进行全面的空值检查
3. **使用 Or/OrElse** 进行回退值链
4. **使用 SafeDeref** 进行安全指针操作
5. **使用 Must** 仅在你确定操作会成功时

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

### 检查函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `IsZero(value)` | 检查值是否为零 | O(1) |
| `IsNotZero(value)` | 检查值是否不为零 | O(1) |
| `IsNil(value)` | 检查值是否为 nil | O(1) |
| `IsNotNil(value)` | 检查值是否不为 nil | O(1) |
| `IsEmpty(value)` | 检查值是否为空 | O(1) |
| `IsNotEmpty(value)` | 检查值是否不为空 | O(1) |

### 安全操作

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `Must(value, error)` | 返回值或在错误时 panic | O(1) |
| `SafeDeref(pointer)` | 安全指针解引用 | O(1) |
| `SafeDerefDef(pointer, default)` | 带默认值的安全指针解引用 | O(1) |
| `Value(interface{})` | 从接口提取值 | O(1) |
| `Def(value, default)` | 返回值或空时返回默认值 | O(1) |

### 比较函数

| 函数 | 描述 | 时间复杂度 |
|------|------|------------|
| `Equal(a, b)` | 检查值是否相等 | O(1) |
| `NotEqual(a, b)` | 检查值是否不相等 | O(1) |
| `DeepEqual(a, b)` | 使用反射进行深度比较 | O(n) |

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
    if value.IsEmpty(data.Name) {
        return errors.New("姓名是必需的")
    }
    
    if value.IsZero(data.Age) || data.Age < 0 {
        return errors.New("年龄必须为正数")
    }
    
    if value.IsNil(data.Email) || value.IsEmpty(*data.Email) {
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
        value.IsNotEmpty(message),
        fmt.Sprintf("[%s] %s", status, message),
        fmt.Sprintf("[%s] 无消息", status),
    )
}
```

### 4. 错误处理

```go
import "github.com/go4x/goal/value"

// 安全错误处理
func ProcessData(data string) (string, error) {
    if value.IsEmpty(data) {
        return "", errors.New("数据为空")
    }
    
    result := value.Must(transformData(data))
    return result, nil
}
```

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
