# Limiter 包

一个基于令牌桶算法的高性能、线程安全的 Go 限流库。

## 特性

- **令牌桶算法**: 允许突发流量达到桶容量，同时保持稳定的令牌生成速率
- **线程安全**: 可在多个 goroutine 中安全并发使用
- **灵活的 API**: 支持阻塞、非阻塞和基于超时的令牌获取
- **统计跟踪**: 内置监控和可观测性功能
- **基于接口的设计**: 清晰的抽象，允许不同的限流实现
- **优雅关闭**: 适当的清理和信号处理

## 安装

```bash
go get github.com/go4x/goal/limiter
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"
    "github.com/go4x/goal/limiter"
)

func main() {
    // 创建令牌桶: 容量=10, 速率=5个令牌/秒
    limiter := limiter.NewTokenBucket(10, 5, time.Second)
    limiter.Start()
    defer limiter.Stop()

    // 发送请求
    for i := 0; i < 15; i++ {
        if limiter.TryTake() {
            fmt.Printf("请求 %d: 允许\n", i+1)
        } else {
            fmt.Printf("请求 %d: 限流\n", i+1)
        }
        time.Sleep(time.Millisecond * 100)
    }
}
```

## API 参考

### Limiter 接口

`Limiter` 接口为限流实现提供了通用契约：

```go
type Limiter interface {
    Start()                                                    // 启动限流器
    TryTake() bool                                           // 非阻塞令牌获取
    Take()                                                   // 阻塞令牌获取
    TakeWithTimeout(timeout time.Duration) bool              // 基于超时的获取
    Stat() (total, blocked int64, successRate float64)       // 获取统计信息
    Stop()                                                   // 停止限流器
}
```

### TokenBucket

`TokenBucket` 结构体使用令牌桶算法实现 `Limiter` 接口。

#### 构造函数

```go
func NewTokenBucket(capacity int, rate int, window time.Duration) *TokenBucket
```

**参数:**
- `capacity`: 桶可以容纳的最大令牌数（突发容量）
- `rate`: 每个时间窗口生成的令牌数
- `window`: 令牌生成的时间窗口（如 1 秒、1 分钟）

**示例:**
```go
// 容量100，每秒10个令牌
limiter := limiter.NewTokenBucket(100, 10, time.Second)
```

#### 方法

##### Start()
在后台 goroutine 中启动令牌生成过程。

```go
limiter.Start()
```

##### TryTake() bool
尝试获取令牌而不阻塞。成功返回 `true`，无令牌可用返回 `false`。

```go
if limiter.TryTake() {
    // 处理请求
} else {
    // 处理限流
}
```

##### Take()
获取令牌，阻塞直到有令牌可用。

```go
limiter.Take() // 会等待直到令牌可用
// 处理请求
```

##### TakeWithTimeout(timeout time.Duration) bool
在指定超时时间内尝试获取令牌。成功返回 `true`，超时返回 `false`。

```go
if limiter.TakeWithTimeout(time.Second) {
    // 处理请求
} else {
    // 处理超时
}
```

##### Stat() (total, blocked int64, successRate float64)
返回当前统计信息：
- `total`: 总请求数
- `blocked`: 被阻塞的请求数
- `successRate`: 成功率百分比（0.0 到 100.0）

```go
total, blocked, rate := limiter.Stat()
fmt.Printf("总数: %d, 阻塞: %d, 成功率: %.2f%%\n", total, blocked, rate)
```

##### Stop()
停止令牌生成过程。

```go
limiter.Stop()
```

##### ResetStat()
将所有统计计数器重置为零。

```go
limiter.ResetStat()
```

## 使用示例

### 基本限流

```go
limiter := limiter.NewTokenBucket(10, 5, time.Second)
limiter.Start()
defer limiter.Stop()

// 非阻塞方式
if limiter.TryTake() {
    processRequest()
} else {
    handleRateLimit()
}
```

### 阻塞限流

```go
limiter := limiter.NewTokenBucket(5, 2, time.Second)
limiter.Start()
defer limiter.Stop()

// 阻塞方式 - 会等待令牌
limiter.Take()
processRequest()
```

### 基于超时的限流

```go
limiter := limiter.NewTokenBucket(3, 1, time.Second)
limiter.Start()
defer limiter.Stop()

// 超时方式 - 最多等待5秒
if limiter.TakeWithTimeout(5 * time.Second) {
    processRequest()
} else {
    handleTimeout()
}
```

### 统计监控

```go
limiter := limiter.NewTokenBucket(100, 10, time.Second)
limiter.Start()
defer limiter.Stop()

// 发送一些请求...

// 检查统计信息
total, blocked, rate := limiter.Stat()
fmt.Printf("成功率: %.2f%%\n", rate)

// 重置统计信息用于新的测量周期
limiter.ResetStat()
```

### 接口使用

```go
// 使用接口实现多态
var rateLimiter limiter.Limiter = limiter.NewTokenBucket(10, 5, time.Second)
rateLimiter.Start()
defer rateLimiter.Stop()

// 所有接口方法工作方式相同
if rateLimiter.TryTake() {
    // 处理请求
}
```

### 并发使用

```go
limiter := limiter.NewTokenBucket(10, 5, time.Second)
limiter.Start()
defer limiter.Stop()

var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        if limiter.TryTake() {
            processRequest()
        }
    }()
}
wg.Wait()
```

## 性能

限流器专为高性能场景设计：

- **低延迟**: 非阻塞操作开销最小
- **高吞吐量**: 针对并发访问模式优化
- **内存高效**: 使用通道进行令牌管理
- **线程安全**: 所有操作都是并发安全的

### 基准测试结果

运行基准测试查看性能特征：

```bash
go test -bench=. ./limiter
```

## 错误处理

限流器优雅地处理边界情况：

- **零/负参数**: 自动修正为安全默认值
- **多次启动/停止调用**: 可安全多次调用
- **通道操作**: 防止 panic 条件

## 最佳实践

1. **始终调用 Start()**: 限流器在启动前不会生成令牌
2. **清理资源**: 完成后始终调用 Stop()
3. **监控统计信息**: 使用 Stat() 跟踪性能
4. **选择适当的容量**: 在突发允许和内存使用之间平衡
5. **使用超时**: 为了更好的控制，优先使用 TakeWithTimeout() 而不是阻塞的 Take()

## 线程安全

所有方法都是线程安全的，可以从多个 goroutine 并发调用。实现使用：

- 互斥锁保护统计信息
- 通道进行线程安全的令牌管理
- 适当的地方使用原子操作

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
