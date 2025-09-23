# AsyncClient - 异步 HTTP 客户端

`AsyncClient` 提供了基于 Go channel 的异步 HTTP 请求功能，允许您以非阻塞的方式发送 HTTP 请求并处理响应。

## 特性

- **非阻塞请求**: 所有 HTTP 方法都返回 channel，不会阻塞调用线程
- **完整 HTTP 方法支持**: GET, POST, PUT, PATCH, DELETE 及其变体
- **批量请求**: 支持并发发送多个请求
- **上下文支持**: 支持超时和取消控制
- **结果聚合**: 提供 `WaitForAll` 和 `WaitForFirst` 工具函数
- **类型安全**: 基于 Go 的 channel 机制，编译时类型检查

## 基本用法

### 创建异步客户端

```go
import "github.com/go4x/goal/httpx/v2"

// 使用默认客户端
client := httpx.NewAsyncClient(nil)

// 使用自定义客户端
customClient := &http.Client{Timeout: 5 * time.Second}
asyncClient := httpx.NewAsyncClient(customClient)
```

### 发送异步请求

```go
// GET 请求
resultChan := client.GetAsync("https://api.example.com/users", nil)

// 做其他工作...
fmt.Println("请求已发送，继续处理其他任务...")

// 等待结果
result := <-resultChan
if result.Err != nil {
    log.Fatal(result.Err)
}
defer result.Resp.Close()

body, _ := result.Resp.String()
fmt.Printf("响应: %s\n", body)
```

### POST JSON 请求

```go
jsonData := `{"name": "张三", "age": 25}`
body := strings.NewReader(jsonData)

resultChan := client.PostJsonAsync("https://api.example.com/users", body)
result := <-resultChan

if result.Err != nil {
    log.Fatal(result.Err)
}
defer result.Resp.Close()

fmt.Printf("状态码: %d\n", result.Resp.StatusCode)
```

## 批量请求

### 使用 BatchAsync 和智能方法选择

```go
// 创建不同类型的请求，客户端会根据 Content-Type 自动选择合适的方法
requests := []httpx.AsyncRequest{
    // GET 请求
    httpx.NewAsyncRequest("GET", "https://api1.example.com/data", nil, nil),
    
    // POST JSON 请求 - 自动使用 PostJson
    httpx.NewAsyncRequestWithContentType("POST", "https://api2.example.com/users", nil, 
        strings.NewReader(`{"name": "张三", "age": 25}`), "application/json"),
    
    // PUT 表单请求 - 自动使用 PutForm
    httpx.NewAsyncRequestWithContentType("PUT", "https://api3.example.com/users/1", nil,
        strings.NewReader("name=李四&age=30"), "application/x-www-form-urlencoded"),
}

resultChan := client.BatchAsync(requests)
results := <-resultChan

for i, result := range results {
    if result.Err != nil {
        fmt.Printf("请求 %d 失败: %v\n", i, result.Err)
        continue
    }
    defer result.Resp.Close()
    fmt.Printf("请求 %d 成功: %d\n", i, result.Resp.StatusCode)
}
```

### 使用 WaitForAll

```go
resultChans := []<-chan httpx.AsyncResult{
    client.GetAsync("https://api1.example.com/data", nil),
    client.GetAsync("https://api2.example.com/data", nil),
    client.GetAsync("https://api3.example.com/data", nil),
}

results := httpx.WaitForAll(resultChans)
fmt.Printf("所有 %d 个请求已完成\n", len(results))
```

### 使用 WaitForFirst

```go
resultChans := []<-chan httpx.AsyncResult{
    client.GetAsync("https://slow-api.example.com/data", nil), // 慢
    client.GetAsync("https://fast-api.example.com/data", nil), // 快
}

result := httpx.WaitForFirst(resultChans)
fmt.Printf("第一个响应: %s\n", result.Resp.StatusCode)
```

## 上下文控制

### 超时控制

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

resultChan := client.WithContextAsync(ctx, "GET", "https://api.example.com/data", nil, nil)
result := <-resultChan

if result.Err != nil {
    if errors.Is(result.Err, context.DeadlineExceeded) {
        fmt.Println("请求超时")
    } else {
        fmt.Printf("请求失败: %v\n", result.Err)
    }
    return
}
defer result.Resp.Close()
```

### 取消控制

```go
ctx, cancel := context.WithCancel(context.Background())

// 在另一个 goroutine 中取消请求
go func() {
    time.Sleep(2 * time.Second)
    cancel()
}()

resultChan := client.WithContextAsync(ctx, "GET", "https://api.example.com/data", nil, nil)
result := <-resultChan

if result.Err != nil {
    if errors.Is(result.Err, context.Canceled) {
        fmt.Println("请求被取消")
    } else {
        fmt.Printf("请求失败: %v\n", result.Err)
    }
}
```

## 智能方法选择

`AsyncClient` 支持智能方法选择，根据 HTTP 方法和 Content-Type 自动选择最合适的请求方法：

### 自动方法映射

| HTTP 方法 | Content-Type | 自动选择的方法 |
|-----------|--------------|----------------|
| GET | - | `Get()` |
| GET | application/json | `GetJson()` |
| GET | application/x-www-form-urlencoded | `GetForm()` |
| GET | 其他 | `GetWithBody()` |
| POST | - | `Post()` |
| POST | application/json | `PostJson()` |
| POST | application/x-www-form-urlencoded | `PostForm()` |
| PUT | - | `Put()` |
| PUT | application/json | `PutJson()` |
| PUT | application/x-www-form-urlencoded | `PutForm()` |
| PATCH | - | `Patch()` |
| PATCH | application/json | `PatchJson()` |
| PATCH | application/x-www-form-urlencoded | `PatchForm()` |
| DELETE | - | `Delete()` |
| DELETE | application/json | `DeleteJson()` |
| DELETE | application/x-www-form-urlencoded | `DeleteForm()` |
| OPTIONS | - | `Options()` |

### 使用方式

```go
// 方式1: 使用 NewAsyncRequestWithContentType
req := httpx.NewAsyncRequestWithContentType("POST", url, nil, body, "application/json")

// 方式2: 使用 WithContentType 链式调用
req := httpx.NewAsyncRequest("POST", url, nil, body).WithContentType("application/json")

// 方式3: 在 Headers 中设置 Content-Type
req := httpx.NewAsyncRequest("POST", url, nil, body)
req.Headers["Content-Type"] = "application/json"
```

### 智能选择示例

```go
// 这些请求会自动选择合适的方法
requests := []httpx.AsyncRequest{
    // 自动使用 PostJson
    httpx.NewAsyncRequestWithContentType("POST", url, nil, 
        strings.NewReader(`{"name": "test"}`), "application/json"),
    
    // 自动使用 PostForm
    httpx.NewAsyncRequestWithContentType("POST", url, nil,
        strings.NewReader("name=test&value=123"), "application/x-www-form-urlencoded"),
    
    // 自动使用 PutJson
    httpx.NewAsyncRequestWithContentType("PUT", url, nil,
        strings.NewReader(`{"id": 1}`), "application/json"),
    
    // 自动使用 DeleteJson
    httpx.NewAsyncRequestWithContentType("DELETE", url, nil,
        strings.NewReader(`{"id": 1}`), "application/json"),
    
    // 自动使用 Options
    httpx.NewAsyncRequest("OPTIONS", url, nil, nil),
}
```

## 所有支持的方法

### GET 方法
- `GetAsync(url, params)` - 基本 GET 请求
- `GetWithBodyAsync(url, body)` - GET 请求带 body
- `GetJsonAsync(url, body)` - GET 请求 JSON body
- `GetFormAsync(url, body)` - GET 请求表单 body

### POST 方法
- `PostAsync(url, body)` - 基本 POST 请求
- `PostJsonAsync(url, body)` - POST JSON 请求
- `PostFormAsync(url, body)` - POST 表单请求

### PUT 方法
- `PutAsync(url, body)` - 基本 PUT 请求
- `PutJsonAsync(url, body)` - PUT JSON 请求
- `PutFormAsync(url, body)` - PUT 表单请求

### PATCH 方法
- `PatchAsync(url, body)` - 基本 PATCH 请求
- `PatchJsonAsync(url, body)` - PATCH JSON 请求
- `PatchFormAsync(url, body)` - PATCH 表单请求

### DELETE 方法
- `DeleteAsync(url)` - 基本 DELETE 请求
- `DeleteJsonAsync(url, body)` - DELETE JSON 请求
- `DeleteFormAsync(url, body)` - DELETE 表单请求

### OPTIONS 方法
- `OptionsAsync(url, options...)` - OPTIONS 请求，常用于 CORS 预检请求

## 错误处理

```go
resultChan := client.GetAsync("https://invalid-url", nil)
result := <-resultChan

if result.Err != nil {
    // 处理网络错误
    if netErr, ok := result.Err.(net.Error); ok {
        if netErr.Timeout() {
            fmt.Println("请求超时")
        } else {
            fmt.Printf("网络错误: %v\n", netErr)
        }
    } else {
        fmt.Printf("其他错误: %v\n", result.Err)
    }
    return
}

// 处理 HTTP 错误状态码
if !result.Resp.IsSuccess() {
    fmt.Printf("HTTP 错误: %d %s\n", result.Resp.StatusCode, result.Resp.Status())
    return
}

defer result.Resp.Close()
// 处理成功响应...
```

## 性能考虑

1. **内存使用**: 每个异步请求都会创建一个 goroutine 和 channel
2. **并发控制**: 大量并发请求时考虑使用 worker pool 模式
3. **资源清理**: 始终记得关闭响应体 `result.Resp.Close()`

## 最佳实践

1. **超时设置**: 为长时间运行的请求设置合理的超时时间
2. **错误处理**: 总是检查 `result.Err` 并适当处理
3. **资源管理**: 使用 `defer result.Resp.Close()` 确保响应体被关闭
4. **并发限制**: 避免无限制地创建异步请求，考虑使用信号量或 worker pool

## 与同步客户端的对比

| 特性 | 同步客户端 | 异步客户端 |
|------|------------|------------|
| **阻塞性** | 阻塞调用线程 | 非阻塞，返回 channel |
| **并发** | 需要手动使用 goroutine | 内置并发支持 |
| **错误处理** | 直接返回 error | 通过 channel 传递 error |
| **性能** | 简单直接 | 更好的并发性能 |
| **复杂度** | 简单 | 中等复杂度 |

异步客户端特别适合需要高并发请求的场景，如批量 API 调用、数据抓取、微服务间通信等。
