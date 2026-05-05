# Goal

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache--2.0-green.svg)](LICENSE)

[English](./README.md)

Goal 是一个 Go 工具库，提供集合、类型转换、错误处理、HTTP、I/O、JSON、随机数、
重试、时间、UUID、值处理等常用功能包。

## 安装

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
