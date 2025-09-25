# Goal

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-apache2.0-green.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-v1.0.0-brightgreen.svg)](https://github.com/go4x/goal/releases/tag/v1.0.0)

[英文](./README.md)

一个功能全面的 Go 实用工具库，提供丰富的包集合，用于常见编程任务、数据结构和系统操作。

## 🚀 功能特性

- **全面覆盖**: 20+ 个专业包，覆盖广泛领域
- **泛型支持**: 现代 Go 泛型贯穿整个代码库
- **极少依赖**: 极少的外部依赖，易于维护
- **性能优化**: 精心设计，追求效率和速度
- **文档完善**: 详细的文档和示例
- **测试覆盖**: 全面的测试覆盖
- **生产就绪**: 在真实应用中经过实战测试

## 📦 包概览

### 核心工具

| 包 | 描述 | 特性 |
|----|------|------|
| [`assert`](./assert/) | 测试断言 | 类型安全断言，自定义消息 |
| [`value`](./value/) | 值处理工具 | 空值检查，条件逻辑，安全操作 |
| [`ptr`](./ptr/) | 指针工具 | 安全解引用，指针操作 |
| [`conv`](./conv/) | 类型转换 | 安全转换，格式验证 |

### 集合

| 包 | 描述 | 特性 |
|----|------|------|
| [`col/set`](./col/set/) | 集合实现 | HashSet、ArraySet、LinkedSet，O(1) 操作 |
| [`col/mapx`](./col/mapx/) | 映射实现 | 常规映射、ArrayMap、LinkedMap，保持顺序 |
| [`col/slicex`](./col/slicex/) | 增强切片操作 | 不可变操作，函数式编程 |

### 字符串和文本

| 包 | 描述 | 特性 |
|----|------|------|
| [`stringx`](./stringx/) | 字符串工具 | 大小写转换，模糊处理，常量，构建器 |
| [`color`](./color/) | 颜色操作 | RGB 操作，颜色转换 |
| [`jsonx`](./jsonx/) | JSON 工具 | 增强 JSON 操作，验证 |

### 系统和 I/O

| 包 | 描述 | 特性 |
|----|------|------|
| [`cmd`](./cmd/) | 命令执行 | 异步执行，超时处理，流式处理 |
| [`iox`](./iox/) | I/O 工具 | 文件操作，目录遍历，路径处理 |
| [`httpx`](./httpx/) | HTTP 客户端 | 异步客户端，请求/响应处理 |

### 加密和安全

| 包 | 描述 | 特性 |
|----|------|------|
| [`ciphers`](./ciphers/) | 加密函数 | AES 加密，哈希，数据压缩 |
| [`uuid`](./uuid/) | UUID 生成 | 标准 UUID，分布式 ID（Sonyflake） |

### 错误处理

| 包 | 描述 | 特性 |
|----|------|------|
| [`errorx`](./errorx/) | 错误工具 | 错误链，包装，恢复 |

### 数学和统计

| 包 | 描述 | 特性 |
|----|------|------|
| [`mathx`](./mathx/) | 数学工具 | 高级数学操作，计算 |
| [`prob`](./prob/) | 概率函数 | 统计操作，概率计算 |
| [`random`](./random/) | 随机生成 | 数字生成，字符串生成 |

### 工具

| 包 | 描述 | 特性 |
|----|------|------|
| [`timex`](./timex/) | 时间工具 | 时间格式化，解析，操作 |
| [`limiter`](./limiter/) | 限流 | 令牌桶，限流算法 |
| [`retry`](./retry/) | 重试机制 | 指数退避，重试策略 |
| [`reflectx`](./reflectx/) | 反射工具 | 类型检查，反射辅助 |
| [`printer`](./printer/) | 打印工具 | 格式化输出，美化打印 |

## 🚀 快速开始

### 安装

```bash
go get github.com/go4x/goal
```

### 基础用法

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/stringx"
    "github.com/go4x/goal/value"
    "github.com/go4x/goal/col/set"
)

func main() {
    // 字符串操作
    result := stringx.ToCamelCase("hello_world")
    fmt.Println(result) // "helloWorld"
    
    // 值处理
    safeValue := value.Or("", "", "fallback")
    fmt.Println(safeValue) // "fallback"
    
    // 集合操作
    mySet := set.New[string]()
    mySet.Add("apple").Add("banana")
    fmt.Println(mySet.Contains("apple")) // true
}
```

## 📚 包示例

### 字符串操作

```go
import "github.com/go4x/goal/stringx"

// 大小写转换
camel := stringx.ToCamelCase("hello_world")     // "helloWorld"
pascal := stringx.ToPascalCase("hello_world")   // "HelloWorld"
kebab := stringx.ToKebabCase("HelloWorld")      // "hello-world"

// 字符串模糊处理
blurred := stringx.BlurEmail("user@example.com") // "u****@example.com"

// 字符串构建
builder := stringx.NewBuilder()
builder.WriteString("Hello ").WriteString("World")
result := builder.String() // "Hello World"
```

### 值处理

```go
import "github.com/go4x/goal/value"

// 条件逻辑
result := value.IfElse(age >= 18, "adult", "minor")

// 空值检查
if value.IsNotEmpty(data) {
    // 处理数据
}

// 值合并
fallback := value.Or("", "", "default")

// 安全操作
safeValue := value.Must(strconv.Atoi("123"))
```

### 集合

```go
import "github.com/go4x/goal/col/set"
import "github.com/go4x/goal/col/mapx"

// 集合操作
mySet := set.New[string]()
mySet.Add("apple").Add("banana").Add("apple") // 重复项被忽略
fmt.Println(mySet.Size()) // 2

// 映射操作
myMap := mapx.New[string, int]()
myMap.Put("apple", 1).Put("banana", 2)
value, exists := myMap.Get("apple")
fmt.Println(value, exists) // 1 true
```

### HTTP 客户端

```go
import "github.com/go4x/goal/httpx"

// 简单 HTTP 请求
response, err := httpx.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer response.Close()

// 异步 HTTP 请求
client := httpx.NewAsyncClient()
future := client.GetAsync("https://api.example.com/data")
response, err := future.Get()
```

### 加密

```go
import "github.com/go4x/goal/ciphers"

// AES 加密
data := []byte("敏感数据")
key := []byte("your-32-byte-key-here-123456789012")
iv := []byte("random-16-byte-iv")

encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}

decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
if err != nil {
    log.Fatal(err)
}
```

### 命令执行

```go
import "github.com/go4x/goal/cmd"

// 带超时的命令执行
result, err := cmd.ExecWithTimeout("ls -la", 5*time.Second)
if err != nil {
    log.Fatal(err)
}

// 异步命令执行
future := cmd.ExecAsync("long-running-command")
result, err := future.Get()
```

## 🏗️ 架构

### 设计原则

1. **模块化**：每个包都是自包含的，专注于特定领域
2. **简洁易用**：每个包都尽可能简洁，避免不必要的复杂性，确保易用性
3. **极少依赖**：尽可能原生实现，减少依赖外部包，避免不必要的复杂性
4. **泛型优先**：现代 Go 泛型贯穿始终，确保类型安全
5. **追求性能**：为速度和内存效率而优化
6. **线程安全**：所有包都设计为并发使用
7. **文档完善**：详细的文档和示例

### 包结构

```
goal/
├── assert/          # 测试断言
├── ciphers/         # 加密函数
├── cmd/             # 命令执行
├── col/             # 集合
│   ├── mapx/        # 映射实现
│   ├── set/         # 集合实现
│   └── slicex/      # 增强切片操作
├── color/           # 颜色操作
├── conv/            # 类型转换
├── errorx/          # 错误处理
├── httpx/           # HTTP 客户端
├── iox/             # I/O 工具
├── jsonx/           # JSON 工具
├── limiter/         # 限流
├── mathx/           # 数学工具
├── printer/         # 打印工具
├── prob/            # 概率函数
├── ptr/             # 指针工具
├── random/          # 随机生成
├── reflectx/        # 反射工具
├── retry/           # 重试机制
├── stringx/         # 字符串工具
├── timex/           # 时间工具
├── uuid/            # UUID 生成
└── value/           # 值处理
```

## 🔧 开发

### 要求

- Go 1.24.0 或更高版本
- Git

### 构建

```bash
# 克隆仓库
git clone https://github.com/go4x/goal.git
cd goal

# 运行测试
go test ./...

# 运行基准测试
go test -bench=. ./...

# 检查覆盖率
go test -cover ./...
```

### 贡献

1. Fork 仓库
2. 创建功能分支
3. 进行更改
4. 为新功能添加测试
5. 确保所有测试通过
6. 提交拉取请求

## 📊 性能

### 基准测试

该库为高性能而设计：

- **集合**：大多数集合/映射操作的 O(1) 操作
- **字符串操作**：优化的字符串操作
- **HTTP 客户端**：带连接池的异步操作
- **加密**：硬件加速的加密操作

### 内存使用

- **高效**：最少的内存分配
- **池化**：在适当的地方使用连接和缓冲区池
- **不可变**：不可变操作以防止副作用

## 📖 文档

### 文档结构

每个包包括：

- **README.md**：英文文档
- **README_CN.md**：中文文档
- **示例**：全面的使用示例
- **API 参考**：完整的 API 文档
- **性能说明**：性能特征和提示

### 获取帮助

- **GitHub Issues**：报告错误和请求功能
- **文档**：全面的包文档
- **示例**：大量的代码示例
- **社区**：加入讨论

## 🤝 贡献

我们欢迎贡献！请查看我们的[贡献指南](CONTRIBUTING.md)了解详情。

### 贡献领域

- **新包**：建议新的工具包
- **性能**：优化现有代码
- **文档**：改进文档
- **示例**：添加更多使用示例
- **测试**：提高测试覆盖率

## 📄 许可证

本项目采用 Apache License 2.0 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- **Go 团队**：感谢优秀的 Go 编程语言
- **社区**：感谢反馈和贡献
- **依赖项**：感谢我们使用的优秀第三方包

### 版本历史

- **v1.0.0**：核心包的初始发布

## 📞 支持

- **GitHub Issues**：[报告问题](https://github.com/go4x/goal/issues)

---

**Goal** - 让 Go 开发更高效、更愉快！🎯
