# ReflectX

一个全面的Go反射工具包，提供广泛的基于反射的操作，包括类型检查、转换、结构体字段操作、值操作、方法调用以及接口/泛型类型检查。

## 功能特性

- **类型检查**: 检查值是否为nil、零值或特定类型
- **类型信息**: 获取详细的类型信息，包括名称、种类、大小、对齐方式
- **类型转换**: 安全的类型转换，带错误处理
- **结构体字段操作**: 获取/设置字段值、名称、标签和详细信息
- **值操作**: 通过索引操作操作切片、映射、数组
- **方法操作**: 调用方法、检查方法存在性、获取方法信息
- **接口工具**: 检查接口实现、获取接口方法
- **泛型工具**: 检查泛型类型和获取泛型类型信息
- **性能优化**: 高效实现，包含全面的基准测试

## 安装

```bash
go get github.com/go4x/goal/reflectx
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/reflectx"
)

func main() {
    // 检查值是否为nil
    var ptr *int
    fmt.Println("是否为nil:", reflectx.IsNil(ptr)) // true
    
    // 检查值是否为零值
    fmt.Println("是否为零值:", reflectx.IsZero(0)) // true
    
    // 获取类型信息
    fmt.Println("类型名称:", reflectx.GetTypeName("hello")) // string
    
    // 转换类型
    result, err := reflectx.Convert(42, reflect.TypeOf(0.0))
    if err == nil {
        fmt.Println("转换结果:", result) // 42.0
    }
}
```

## API 参考

### 类型检查

#### `IsNil(v interface{}) bool`
检查值是否为nil。

```go
var ptr *int
fmt.Println(reflectx.IsNil(ptr)) // true
fmt.Println(reflectx.IsNil(42))  // false
```

#### `IsZero(v interface{}) bool`
检查值是否为其类型的零值。

```go
fmt.Println(reflectx.IsZero(0))     // true
fmt.Println(reflectx.IsZero(""))     // true
fmt.Println(reflectx.IsZero(42))     // false
```

#### 类型检查器
- `IsPointer(v interface{}) bool`
- `IsSlice(v interface{}) bool`
- `IsMap(v interface{}) bool`
- `IsStruct(v interface{}) bool`
- `IsInterface(v interface{}) bool`
- `IsFunc(v interface{}) bool`
- `IsChannel(v interface{}) bool`
- `IsArray(v interface{}) bool`
- `IsString(v interface{}) bool`
- `IsInt(v interface{}) bool`
- `IsUint(v interface{}) bool`
- `IsFloat(v interface{}) bool`
- `IsBool(v interface{}) bool`
- `IsComplex(v interface{}) bool`

### 类型信息

#### `GetTypeName(v interface{}) string`
返回值的类型名称。

```go
type User struct{ Name string }
fmt.Println(reflectx.GetTypeName(User{})) // "main.User"
```

#### `GetKind(v interface{}) string`
返回值的种类。

```go
fmt.Println(reflectx.GetKind([]int{1, 2, 3})) // "slice"
```

#### `GetSize(v interface{}) uintptr`
返回值的字节大小。

```go
fmt.Println(reflectx.GetSize(int(42))) // 8
```

### 类型转换

#### `Convert(v interface{}, targetType reflect.Type) (interface{}, error)`
将值转换为目标类型。

```go
result, err := reflectx.Convert(42, reflect.TypeOf(0.0))
if err == nil {
    fmt.Println(result) // 42.0
}
```

#### `ConvertTo(v interface{}, target interface{}) (interface{}, error)`
将值转换为目标的类型。

```go
var target float64
result, err := reflectx.ConvertTo(42, target)
if err == nil {
    fmt.Println(result) // 42.0
}
```

### 结构体字段操作

#### `GetFieldNames(v interface{}) []string`
返回结构体中所有字段的名称。

```go
type Person struct {
    Name string
    Age  int
}
names := reflectx.GetFieldNames(Person{})
fmt.Println(names) // ["Name", "Age"]
```

#### `GetFieldValue(v interface{}, fieldName string) (interface{}, error)`
获取结构体字段的值。

```go
p := Person{Name: "Alice", Age: 30}
value, err := reflectx.GetFieldValue(p, "Name")
if err == nil {
    fmt.Println(value) // "Alice"
}
```

#### `SetFieldValue(v interface{}, fieldName string, value interface{}) error`
设置结构体字段的值。

```go
p := &Person{Name: "Alice", Age: 30}
err := reflectx.SetFieldValue(p, "Age", 31)
if err == nil {
    fmt.Println(p.Age) // 31
}
```

#### `GetFieldTags(v interface{}) map[string]string`
返回结构体中所有字段的标签。

```go
type User struct {
    Name string `json:"name" db:"user_name"`
    Age  int    `json:"age" db:"user_age"`
}
tags := reflectx.GetFieldTags(User{})
fmt.Println(tags) // map[Age:json:"age" db:"user_age" Name:json:"name" db:"user_name"]
```

### 值操作

#### `GetValue(v interface{}) interface{}`
返回底层值。

```go
slice := []int{1, 2, 3}
value := reflectx.GetValue(slice)
fmt.Println(value) // [1 2 3]
```

#### `GetLen(v interface{}) int`
返回切片、数组或字符串的长度。

```go
slice := []int{1, 2, 3}
fmt.Println(reflectx.GetLen(slice)) // 3
```

#### `GetIndex(v interface{}, index int) (interface{}, error)`
通过索引从切片、数组或字符串中获取元素。

```go
slice := []string{"apple", "banana", "cherry"}
element, err := reflectx.GetIndex(slice, 1)
if err == nil {
    fmt.Println(element) // "banana"
}
```

#### `SetIndex(v interface{}, index int, value interface{}) error`
通过索引在切片或数组中设置元素。

```go
slice := &[]string{"apple", "banana", "cherry"}
err := reflectx.SetIndex(slice, 1, "orange")
if err == nil {
    fmt.Println(*slice) // [apple orange cherry]
}
```

### 映射操作

#### `GetMapValue(m interface{}, key interface{}) (interface{}, error)`
通过键从映射中获取值。

```go
data := map[string]int{"apple": 5, "banana": 3}
value, err := reflectx.GetMapValue(data, "banana")
if err == nil {
    fmt.Println(value) // 3
}
```

#### `SetMapValue(m interface{}, key, value interface{}) error`
通过键在映射中设置值。

```go
data := map[string]int{"apple": 5}
err := reflectx.SetMapValue(data, "banana", 3)
if err == nil {
    fmt.Println(data["banana"]) // 3
}
```

#### `GetMapKeys(m interface{}) []interface{}`
返回映射中的所有键。

```go
data := map[string]int{"apple": 5, "banana": 3}
keys := reflectx.GetMapKeys(data)
fmt.Println(keys) // ["apple", "banana"]
```

### 方法操作

#### `CallMethod(v interface{}, methodName string, args ...interface{}) ([]interface{}, error)`
在值上调用方法。

```go
type Calculator struct{ Value int }
func (c Calculator) Add(x int) int { return c.Value + x }

calc := Calculator{Value: 10}
result, err := reflectx.CallMethod(calc, "Add", 5)
if err == nil {
    fmt.Println(result[0]) // 15
}
```

#### `HasMethod(v interface{}, methodName string) bool`
检查值是否有方法。

```go
type Service struct{}
func (s Service) Process() string { return "processed" }

service := Service{}
fmt.Println(reflectx.HasMethod(service, "Process")) // true
```

#### `GetMethodNames(v interface{}) []string`
返回值上所有方法的名称。

```go
type User struct{ Name string }
func (u User) GetName() string { return u.Name }
func (u User) SetName(name string) { u.Name = name }

user := User{}
names := reflectx.GetMethodNames(user)
fmt.Println(names) // ["GetName", "SetName"]
```

### 接口工具

#### `Implements(v interface{}, iface interface{}) bool`
检查值是否实现了接口。

```go
type Reader interface { Read() string }
type Book struct{ Title string }
func (b Book) Read() string { return b.Title }

book := Book{Title: "Go Programming"}
fmt.Println(reflectx.Implements(book, (*Reader)(nil))) // true
```

#### `GetInterfaceMethods(iface interface{}) []string`
返回接口中所有方法的名称。

```go
type Writer interface {
    Write(data string) error
    Close() error
}
methods := reflectx.GetInterfaceMethods((*Writer)(nil))
fmt.Println(methods) // ["Write", "Close"]
```

### 泛型工具

#### `IsGeneric(v interface{}) bool`
检查值是否为泛型类型。

```go
type GenericSlice[T any] []T
var slice GenericSlice[int]
fmt.Println(reflectx.IsGeneric(slice)) // false (Go反射的限制)
```

## 性能

该包针对性能进行了优化，包含全面的基准测试。主要性能特征：

- **类型检查操作**: 每次操作约10-50ns
- **字段操作**: 每次操作约100-500ns
- **方法操作**: 每次操作约200-1000ns
- **大型结构体操作**: 随字段数量线性扩展

运行基准测试：

```bash
go test -bench=. ./reflectx
```

## 测试

该包包含全面的测试，所有函数的覆盖率达到100%：

```bash
go test ./reflectx -v
```

## 示例

查看 `example_test.go` 文件了解所有函数的详细使用示例。

## 依赖

- Go 1.18+ (支持泛型)
- 仅使用标准库 (无外部依赖)

## 许可证

MIT许可证

## 贡献

欢迎贡献！请随时提交Pull Request。

## 相关库

- [reflect](https://pkg.go.dev/reflect) - Go的标准反射包
- [go4x/goal](https://github.com/go4x/goal) - 包含多个工具包的父项目
