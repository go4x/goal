# HTTPX - Go 语言高级 HTTP 客户端库

[![Go 版本](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org/)
[![许可证](https://img.shields.io/badge/license-apache2.0-green.svg)](LICENSE)
[![覆盖率](https://img.shields.io/badge/coverage-92.1%25-brightgreen.svg)](coverage.out)

HTTPX 是一个功能全面的 Go 语言 HTTP 客户端库，提供同步和异步 HTTP 请求功能。它提供了清晰、直观的 API，具有智能方法选择、全面的错误处理和广泛的定制选项。

## 特性

### 🚀 核心功能
- **同步 HTTP 客户端**: 功能完整的 REST 客户端，支持所有 HTTP 方法
- **异步 HTTP 客户端**: 基于通道的非阻塞 HTTP 请求
- **智能方法选择**: 根据内容类型自动选择合适的 HTTP 方法
- **请求接口**: 灵活的请求抽象，支持流式 API
- **响应包装器**: 增强的响应处理，提供便利方法
- **包级便利方法**: 通过默认客户端直接访问 HTTP 方法

### 🔧 HTTP 方法支持
- **GET**: 支持查询参数的标准 GET 请求
- **POST**: 支持 JSON、表单数据或自定义请求体的 POST 请求
- **PUT**: 支持 JSON、表单数据或自定义请求体的 PUT 请求
- **PATCH**: 支持 JSON、表单数据或自定义请求体的 PATCH 请求
- **DELETE**: 支持可选 JSON/表单请求体的 DELETE 请求
- **OPTIONS**: 用于 CORS 预检的 OPTIONS 请求

### 🎯 高级功能
- **上下文支持**: 完整的 context.Context 集成，支持超时和取消
- **请求选项**: 用于请求配置的函数选项模式
- **批量操作**: 并发请求执行和结果聚合
- **资源管理**: 自动响应体关闭和适当的资源清理
- **错误处理**: 全面的错误处理，提供详细的错误信息
- **竞态条件安全**: 通过 Go 的竞态检测器测试

## 安装

```bash
go get github.com/go4x/goal/httpx
```

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/httpx"
)

func main() {
    // 简单的 GET 请求
    resp, err := httpx.Get("https://httpbin.org/get")
    if err != nil {
        panic(err)
    }
    defer resp.Close()
    
    fmt.Printf("状态码: %d\n", resp.StatusCode)
}
```

### 使用 RestClient

```go
// 创建新的 REST 客户端
client := httpx.NewRestClient(nil)

// 带查询参数的 GET 请求
params := url.Values{"key": {"value"}}
resp, err := client.Get("https://httpbin.org/get", params)
if err != nil {
    panic(err)
}
defer resp.Close()

// POST JSON 请求
jsonBody := strings.NewReader(`{"name": "hank", "age": 30}`)
resp, err = client.PostJson("https://httpbin.org/post", jsonBody)
if err != nil {
    panic(err)
}
defer resp.Close()
```

### 使用 AsyncClient

```go
// 创建异步客户端
asyncClient := httpx.NewAsyncClient(nil)

// 发送异步请求
getResult := <-asyncClient.GetAsync("https://httpbin.org/get", nil)
postResult := <-asyncClient.PostJsonAsync("https://httpbin.org/post", jsonBody)

// 处理结果
if getResult.Err != nil {
    fmt.Printf("GET 请求失败: %v\n", getResult.Err)
} else {
    defer getResult.Resp.Close()
    fmt.Printf("GET 状态码: %d\n", getResult.Resp.StatusCode)
}

if postResult.Err != nil {
    fmt.Printf("POST 请求失败: %v\n", postResult.Err)
} else {
    defer postResult.Resp.Close()
    fmt.Printf("POST 状态码: %d\n", postResult.Resp.StatusCode)
}
```

### 批量请求

```go
// 创建多个请求
requests := []httpx.Request{
    httpx.MustNewRequest("GET", "https://httpbin.org/get", nil),
    httpx.MustNewRequest("POST", "https://httpbin.org/post", jsonBody, 
        httpx.WithContentType("application/json")),
}

// 执行批量请求
asyncClient := httpx.NewAsyncClient(nil)
results := <-asyncClient.BatchAsync(requests)

// 处理所有结果
for i, result := range results {
    if result.Err != nil {
        fmt.Printf("请求 %d 失败: %v\n", i, result.Err)
    } else {
        defer result.Resp.Close()
        fmt.Printf("请求 %d 成功: %d\n", i, result.Resp.StatusCode)
    }
}
```

## API 参考

### RestClient

#### 构造函数
```go
func NewRestClient(client *http.Client) *RestClient
```

#### HTTP 方法
```go
// GET 请求
func (rc *RestClient) Get(url string, params url.Values) (*Response, error)
func (rc *RestClient) GetWithBody(url string, body io.Reader) (*Response, error)
func (rc *RestClient) GetJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) GetForm(url string, body io.Reader) (*Response, error)

// POST 请求
func (rc *RestClient) Post(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PostJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PostForm(url string, body io.Reader) (*Response, error)

// PUT 请求
func (rc *RestClient) Put(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PutJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PutForm(url string, body io.Reader) (*Response, error)

// PATCH 请求
func (rc *RestClient) Patch(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PatchJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) PatchForm(url string, body io.Reader) (*Response, error)

// DELETE 请求
func (rc *RestClient) Delete(url string) (*Response, error)
func (rc *RestClient) DeleteJson(url string, body io.Reader) (*Response, error)
func (rc *RestClient) DeleteForm(url string, body io.Reader) (*Response, error)

// OPTIONS 请求
func (rc *RestClient) Options(url string, options ...RequestOption) (*Response, error)
```

### AsyncClient

#### 构造函数
```go
func NewAsyncClient(client *http.Client) *AsyncClient
```

#### 异步方法
所有 RestClient 方法都有对应的异步版本，返回 `<-chan AsyncResult`：

```go
func (ac *AsyncClient) GetAsync(url string, params url.Values) <-chan AsyncResult
func (ac *AsyncClient) PostJsonAsync(url string, body io.Reader) <-chan AsyncResult
// ... 所有 HTTP 方法都有对应的异步版本
```

#### 批量操作
```go
func (ac *AsyncClient) BatchAsync(requests []Request) <-chan []AsyncResult
func WaitForAll(resultChans []<-chan AsyncResult) []AsyncResult
func WaitForFirst(resultChans []<-chan AsyncResult) AsyncResult
```

### Request 接口

```go
type Request interface {
    GetURL() *url.URL
    GetMethod() string
    GetBody() (io.ReadCloser, error)
    GetHeader() http.Header
    GetContext() context.Context
    // ... 其他 getter 方法
    
    WithURL(url *url.URL) Request
    WithMethod(method string) Request
    WithBody(body io.ReadCloser) Request
    WithHeader(key string, value string) Request
    WithContext(ctx context.Context) Request
    // ... 其他 setter 方法
}
```

#### 请求创建
```go
func NewRequest(method string, url string, body io.Reader, options ...RequestOption) (Request, error)
func MustNewRequest(method string, url string, body io.Reader, options ...RequestOption) Request
func NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader, options ...RequestOption) (Request, error)
func MustNewRequestWithContext(ctx context.Context, method string, url string, body io.Reader, options ...RequestOption) Request
```

#### 请求选项
```go
func WithContentType(contentType string) RequestOption
func WithAuthorization(authorization string) RequestOption
func WithHeader(key string, value string) RequestOption
func WithContext(ctx context.Context) RequestOption
```

### Response 包装器

```go
type Response struct {
    *http.Response
}

// 状态检查
func (r *Response) IsSuccess() bool
func (r *Response) IsClientError() bool
func (r *Response) IsServerError() bool
func (r *Response) Status() int

// 响应体读取
func (r *Response) Bytes() ([]byte, error)
func (r *Response) String() (string, error)
func (r *Response) JSON(v interface{}) error

// 头部工具
func (r *Response) HeaderValue(key string) string
func (r *Response) ContentType() string
func (r *Response) ContentLength() int64

// 资源管理
func (r *Response) Close() error
```

### 包级便利方法

```go
// 所有 HTTP 方法都可以在包级别使用
func Get(url string, params url.Values) (*Response, error)
func PostJson(url string, body io.Reader) (*Response, error)
func PutForm(url string, body io.Reader) (*Response, error)
// ... 所有 HTTP 方法都有包级版本
```

## 示例

### 带超时的上下文

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 使用 RestClient
resp, err := client.GetWithContext(ctx, "https://httpbin.org/delay/2")
if err != nil {
    panic(err)
}
defer resp.Close()

// 使用 AsyncClient
result := <-asyncClient.WithContextAsync(ctx, "GET", "https://httpbin.org/delay/2", nil, nil)
if result.Err != nil {
    fmt.Printf("请求失败: %v\n", result.Err)
}
```

### 自定义头部

```go
// 使用请求选项
req, err := httpx.NewRequest("POST", "https://api.example.com/data", body,
    httpx.WithContentType("application/json"),
    httpx.WithAuthorization("Bearer token123"),
    httpx.WithHeader("X-Custom-Header", "value"),
)
if err != nil {
    panic(err)
}

// 使用 RestClient 带选项
resp, err := client.PostJson("https://api.example.com/data", body,
    httpx.WithAuthorization("Bearer token123"),
)
```

### 错误处理

```go
resp, err := httpx.Get("https://httpbin.org/status/404")
if err != nil {
    // 网络或请求创建错误
    fmt.Printf("请求失败: %v\n", err)
    return
}
defer resp.Close()

if !resp.IsSuccess() {
    // HTTP 错误响应
    body, _ := resp.String()
    fmt.Printf("HTTP 错误 %d: %s\n", resp.StatusCode, body)
    return
}

// 成功
fmt.Printf("成功: %d\n", resp.StatusCode)
```

### JSON 处理

```go
// 发送 JSON
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

user := User{Name: "hank", Age: 30}
jsonData, _ := json.Marshal(user)
resp, err := httpx.PostJson("https://api.example.com/users", strings.NewReader(string(jsonData)))

// 接收 JSON
type ResponseData struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var responseData ResponseData
err = resp.JSON(&responseData)
if err != nil {
    panic(err)
}
```

## 最佳实践

### 1. 始终关闭响应
```go
resp, err := httpx.Get("https://example.com")
if err != nil {
    panic(err)
}
defer resp.Close() // 始终关闭响应
```

### 2. 使用上下文进行超时控制
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

resp, err := httpx.GetWithContext(ctx, "https://example.com")
```

### 3. 适当处理错误
```go
resp, err := httpx.Get("https://example.com")
if err != nil {
    // 处理网络错误
    return fmt.Errorf("请求失败: %w", err)
}
defer resp.Close()

if !resp.IsSuccess() {
    // 处理 HTTP 错误
    return fmt.Errorf("HTTP 错误: %d", resp.StatusCode)
}
```

### 4. 使用批量操作处理多个请求
```go
requests := []httpx.Request{
    httpx.MustNewRequest("GET", "https://api1.example.com", nil),
    httpx.MustNewRequest("GET", "https://api2.example.com", nil),
}

asyncClient := httpx.NewAsyncClient(nil)
results := <-asyncClient.BatchAsync(requests)

for _, result := range results {
    if result.Err != nil {
        // 处理单个请求错误
        continue
    }
    defer result.Resp.Close()
    // 处理成功的响应
}
```

## 性能

HTTPX 专为高性能设计，开销最小：

- **零拷贝操作**，在可能的地方
- **高效的资源管理**，自动清理
- **并发请求执行**，适当的同步
- **最少的内存分配**，在热路径中

### 基准测试

```bash
go test -bench=. ./httpx
```

## 测试

运行测试套件：

```bash
# 运行所有测试
go test ./httpx

# 运行带覆盖率的测试
go test -cover ./httpx

# 运行带竞态检测的测试
go test -race ./httpx

# 运行基准测试
go test -bench=. ./httpx
```

## 贡献

1. Fork 仓库
2. 创建你的功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开 Pull Request

## 许可证

本项目采用 Apache License 2.0 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 致谢

- 基于 Go 标准 `net/http` 包构建
- 受到现代 HTTP 客户端库的启发
- 感谢 Go 社区的反馈和贡献
