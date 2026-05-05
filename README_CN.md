# Goal

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache--2.0-green.svg)](LICENSE)

[English](./README.md)

Goal 是一个 Go 工具库，提供集合、类型转换、错误处理、HTTP、I/O、JSON、随机数、
重试、时间、UUID、值处理等常用功能包。

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

| 包                     | 描述         | 特性                                 |
| ---------------------- | ------------ | ------------------------------------ |
| [`assert`](./assert/)  | 测试断言     | 类型安全断言，自定义消息             |
| [`value`](./value/)    | 值处理工具   | 空值检查，条件逻辑，安全操作         |
| [`ptr`](./ptr/)        | 指针工具     | 安全解引用，指针操作                 |
| [`conv`](./conv/)      | 类型转换     | 安全转换，格式验证                   |
| [`envx`](./envx/)      | 环境变量工具 | 默认值、必填值、类型解析             |
| [`is`](./is/) (v1.1.0) | 值检查和比较 | 布尔操作，零值/空值/nil 检查，值比较 |

### 集合

| 包                            | 描述         | 特性                                    |
| ----------------------------- | ------------ | --------------------------------------- |
| [`col/set`](./col/set/)       | 集合实现     | HashSet、ArraySet、LinkedSet，O(1) 操作 |
| [`col/mapx`](./col/mapx/)     | 映射实现     | 常规映射、ArrayMap、LinkedMap，保持顺序 |
| [`col/slicex`](./col/slicex/) | 增强切片操作 | 不可变操作，函数式编程                  |

### 字符串和文本

| 包                      | 描述       | 特性                               |
| ----------------------- | ---------- | ---------------------------------- |
| [`stringx`](./stringx/) | 字符串工具 | 大小写转换，模糊处理，常量，构建器 |
| [`color`](./color/)     | 颜色操作   | RGB 操作，颜色转换                 |
| [`jsonx`](./jsonx/)     | JSON 工具  | 增强 JSON 操作，验证               |

### 系统和 I/O

| 包                  | 描述        | 特性                         |
| ------------------- | ----------- | ---------------------------- |
| [`cmd`](./cmd/)     | 命令执行    | 异步执行，超时处理，流式处理 |
| [`iox`](./iox/)     | I/O 工具    | 文件操作，目录遍历，路径处理 |
| [`httpx`](./httpx/) | HTTP 客户端 | 异步客户端，请求/响应处理    |

### 加密和安全

| 包                      | 描述      | 特性                              |
| ----------------------- | --------- | --------------------------------- |
| [`ciphers`](./ciphers/) | 加密函数  | AES 加密，哈希，数据压缩          |
| [`uuid`](./uuid/)       | UUID 生成 | 标准 UUID，分布式 ID（Sonyflake） |

### 错误处理

| 包                    | 描述     | 特性               |
| --------------------- | -------- | ------------------ |
| [`errorx`](./errorx/) | 错误工具 | 错误链，包装，恢复 |

### 数学和统计

| 包                    | 描述     | 特性                 |
| --------------------- | -------- | -------------------- |
| [`prob`](./prob/)     | 概率函数 | 统计操作，概率计算   |
| [`random`](./random/) | 随机生成 | 数字生成，字符串生成 |

### 工具

| 包                        | 描述     | 特性                   |
| ------------------------- | -------- | ---------------------- |
| [`timex`](./timex/)       | 时间工具 | 时间格式化，解析，操作 |
| [`limiter`](./limiter/)   | 限流     | 令牌桶，限流算法       |
| [`retry`](./retry/)       | 重试机制 | 指数退避，重试策略     |
| [`reflectx`](./reflectx/) | 反射工具 | 类型检查，反射辅助     |
| [`printer`](./printer/)   | 打印工具 | 格式化输出，美化打印   |

## 🚀 快速开始

### 安装

```bash
go get github.com/go4x/goal
```

## 示例

```go
package main

import (
	"fmt"

	"github.com/go4x/goal/col/set"
	"github.com/go4x/goal/is"
	"github.com/go4x/goal/value"
)

func main() {
	names := set.New[string]().Add("alice").Add("bob")

	fmt.Println(names.Contains("alice"))
	fmt.Println(value.Or("", "fallback"))
	fmt.Println(is.Empty([]string{}))
}
```

## 包列表

| 包 | 用途 |
| --- | --- |
| [`assert`](./assert/) | 测试断言 |
| [`ciphers`](./ciphers/) | AES、哈希、Base64 和编码工具 |
| [`cmd`](./cmd/) | 命令执行 |
| [`col/mapx`](./col/mapx/) | Map 实现 |
| [`col/set`](./col/set/) | Set 实现 |
| [`col/slicex`](./col/slicex/) | Slice 工具 |
| [`conv`](./conv/) | 类型转换 |
| [`errorx`](./errorx/) | 错误处理、包装和 recover 工具 |
| [`httpx`](./httpx/) | HTTP 请求和异步客户端 |
| [`iox`](./iox/) | 文件、目录、路径和遍历工具 |
| [`is`](./is/) | 值检查和比较 |
| [`jsonx`](./jsonx/) | JSON 工具 |
| [`limiter`](./limiter/) | 令牌桶限流器 |
| [`ptr`](./ptr/) | 指针工具 |
| [`random`](./random/) | 随机数字符串 |
| [`reflectx`](./reflectx/) | 反射工具 |
| [`retry`](./retry/) | 重试工具 |
| [`stringx`](./stringx/) | 字符串工具 |
| [`timex`](./timex/) | 时间工具 |
| [`uuid`](./uuid/) | UUID 和分布式 ID |
| [`value`](./value/) | 值选择和 Must 风格工具 |

更详细的 API 示例见各包目录。

## 开发

```bash
go test ./...
go test -race ./...
go test -bench=. ./...
```

部分 API 是明确的 `Must` 或 `Force` 风格，输入非法时会 panic。业务库代码中如果输入
不完全可信，优先使用返回 `error` 的 API。

## 文档

- [更新日志](./CHANGELOG.md)
- 各包目录中的包级 README
- 当本模块位于 go4x 工作区中时，工作区级文档位于 `../docs/`

## 许可证

Apache License 2.0。见 [LICENSE](./LICENSE)。
