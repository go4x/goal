# envx

`envx` 提供一组读取环境变量的小工具。

该包只使用 Go 标准库。它不负责加载 `.env` 文件、不做结构体绑定、不监听配置变化，也不引入外部依赖。

## 安装

```bash
go get github.com/go4x/goal/envx
```

## 用法

```go
package main

import (
	"fmt"
	"time"

	"github.com/go4x/goal/envx"
)

func main() {
	host := envx.GetDefault("APP_HOST", "127.0.0.1")
	port := envx.GetInt("APP_PORT", 8080)
	debug := envx.GetBool("APP_DEBUG", false)
	timeout := envx.GetDuration("APP_TIMEOUT", 5*time.Second)
	services := envx.GetSlice("APP_SERVICES", ",", []string{"api"})

	fmt.Println(host, port, debug, timeout, services)
}
```

## API

| 函数 | 说明 |
| --- | --- |
| `Get(key string) string` | 返回原始值；未设置时返回空字符串。 |
| `Exists(key string) bool` | 判断变量是否已设置，即使值为空也算已设置。 |
| `GetDefault(key, fallback string) string` | 仅在变量未设置时返回 fallback。 |
| `Require(key string) (string, error)` | 变量未设置时返回错误。 |
| `GetInt(key string, fallback int) int` | 解析 int，未设置或解析失败时返回 fallback。 |
| `GetBool(key string, fallback bool) bool` | 解析 bool，未设置或解析失败时返回 fallback。 |
| `GetDuration(key string, fallback time.Duration) time.Duration` | 解析 Go duration，未设置或解析失败时返回 fallback。 |
| `GetSlice(key, sep string, fallback []string) []string` | 分割字符串并去除空项。 |

## 说明

- `GetDefault` 会把“已设置但为空”的变量视为有效值。
- `Require` 只在变量未设置时失败。
- 类型解析函数在变量缺失或格式无效时都会返回 fallback。
- `GetSlice` 在 `sep` 为空时默认使用 `,`。
