# printer - Go 增强打印工具包

一个为 Go 提供增强打印功能的综合性工具包，具有格式化、表格支持和结构化输出功能。

## 特性

- **增强格式化**：具有宽度控制的高级列格式化
- **表格支持**：内置表头和行数据的表格打印
- **结构化输出**：JSON 格式和结构体格式的输出
- **错误处理**：所有操作都有适当的错误处理
- **类型安全**：支持各种数据类型并自动格式化
- **性能优化**：高效的字符串构建和格式化

## 安装

```bash
go get github.com/go4x/goal/printer
```

## 快速开始

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // 基本打印
    printer.Println("你好 %s，你今年 %d 岁", "Alice", 25)
    
    // 列格式化
    printer.Printwln(10, "姓名", "年龄", "城市")
    printer.Printwln(10, "Alice", 25, "New York")
    
    // 表格打印
    headers := []string{"姓名", "年龄", "分数"}
    rows := [][]any{
        {"Alice", 25, 95.5},
        {"Bob", 30, 87.2},
    }
    printer.PrintTable(headers, rows, 12)
}
```

## API 参考

### 基本打印

#### `NewLine()`
打印一个换行符。

```go
printer.NewLine()  // 打印 \n
```

#### `NewSepLine()`
打印一个包含 80 个等号的分隔线。

```go
printer.NewSepLine()  // 打印 ========================================================================================
```

#### `Printf(format string, args ...any)`
fmt.Printf 的包装器，保持一致性。

```go
printer.Printf("你好 %s，你今年 %d 岁\n", "Alice", 25)
```

#### `Println(format string, args ...any)`
打印格式化文本并自动添加换行符。

```go
printer.Println("你好 %s，你今年 %d 岁", "Alice", 25)
// 输出: 你好 Alice，你今年 25 岁
```

### 列格式化

#### `Printw(width int, cols ...any) error`
使用指定宽度打印格式化列。

```go
err := printer.Printw(10, "姓名", "年龄", "城市")
if err != nil {
    // 处理错误
}
```

#### `Printwln(width int, cols ...any) error`
使用指定宽度打印格式化列并添加换行符。

```go
err := printer.Printwln(10, "姓名", "年龄", "城市")
if err != nil {
    // 处理错误
}
```

### 表格打印

#### `PrintTable(headers []string, rows [][]any, colWidth int) error`
打印带有表头和行的格式化表格。

```go
headers := []string{"姓名", "年龄", "分数"}
rows := [][]any{
    {"Alice", 25, 95.5},
    {"Bob", 30, 87.2},
}

err := printer.PrintTable(headers, rows, 12)
if err != nil {
    // 处理错误
}
```

### 结构化输出

#### `PrintJSON(data map[string]any)`
以 JSON 格式打印数据。

```go
data := map[string]any{
    "name": "Alice",
    "age":  25,
    "city": "New York",
}
printer.PrintJSON(data)
```

#### `PrintStruct(name string, fields map[string]any)`
以可读格式打印结构体。

```go
fields := map[string]any{
    "Name": "Alice",
    "Age":  25,
    "City": "New York",
}
printer.PrintStruct("Person", fields)
```

## 使用示例

### 基本使用

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // 基本打印
    printer.Println("欢迎使用 printer 包！")
    printer.NewSepLine()
    
    // 列格式化
    printer.Printwln(15, "产品", "价格", "库存")
    printer.Printwln(15, "笔记本", 999.99, 10)
    printer.Printwln(15, "鼠标", 29.99, 50)
}
```

### 表格打印

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // 创建销售报告
    headers := []string{"产品", "数量", "价格", "总计"}
    rows := [][]any{
        {"笔记本", 5, 999.99, 4999.95},
        {"鼠标", 20, 29.99, 599.80},
        {"键盘", 15, 79.99, 1199.85},
    }
    
    printer.Println("=== 销售报告 ===")
    printer.PrintTable(headers, rows, 15)
}
```

### 结构化数据

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // 以 JSON 格式打印用户信息
    user := map[string]any{
        "name": "Alice",
        "age":  25,
        "email": "alice@example.com",
        "active": true,
    }
    
    printer.Println("用户信息:")
    printer.PrintJSON(user)
    
    // 以结构体格式打印
    printer.Println("\n用户详情:")
    printer.PrintStruct("User", user)
}
```

### 错误处理

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/printer"
)

func main() {
    headers := []string{"姓名", "年龄"}
    rows := [][]any{
        {"Alice", 25, "额外"}, // 这会导致错误
    }
    
    if err := printer.PrintTable(headers, rows, 10); err != nil {
        fmt.Printf("错误: %v\n", err)
        // 适当处理错误
    }
}
```

## 类型支持

printer 包支持各种数据类型并自动格式化：

- **字符串**：左对齐，指定宽度
- **整数**：右对齐，指定宽度
- **浮点数**：右对齐，保留 2 位小数
- **其他类型**：使用 %v 格式化，指定宽度

## 性能

printer 包针对性能进行了优化：

- **高效字符串构建**：使用 strings.Builder
- **最小内存分配**：常见操作的内存分配最少
- **快速列格式化**：预计算格式字符串
- **优化表格打印**：批量操作

### 性能基准

- **基本操作**：~200-300ns/op，1-2 分配
- **列格式化**：~350ns/op，3-4 分配
- **表格打印**：~2μs/op，24 分配
- **大表格**：~560μs/op，5000+ 分配

## 测试

运行测试：

```bash
go test ./printer
```

运行覆盖率测试：

```bash
go test ./printer -cover
```

运行示例：

```bash
go test ./printer -run Example
```

运行基准测试：

```bash
go test ./printer -bench=.
```

## 许可证

此包是 `goal` 项目的一部分，遵循相同的许可证条款。

## 贡献

欢迎贡献！请随时提交 Pull Request。
