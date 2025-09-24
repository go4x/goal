# Retry 重试包

一个灵活且强大的 Go 应用程序重试机制，支持可配置的重试策略、间隔和回调函数。

## 特性

- **灵活的重试逻辑**: 支持自定义重试条件和错误处理
- **可配置间隔**: 内置指数退避和固定间隔策略
- **可扩展设计**: 易于实现自定义间隔策略
- **回调支持**: 通过自定义回调函数监控重试尝试
- **函数式选项**: 使用函数式选项模式的简洁 API
- **类型安全**: 完全类型化，具有全面的错误处理

## 安装

```bash
go get github.com/go4x/goal/retry
```

## 快速开始

### 基本用法

```go
package main

import (
    "errors"
    "fmt"
    "time"
    
    "github.com/go4x/goal/retry"
)

func main() {
    // 使用指数退避的简单重试
    err := retry.Do(func() (bool, error) {
        // 你的操作在这里
        return false, errors.New("临时失败")
    }, retry.Times(3))
    
    if err != nil {
        fmt.Printf("操作在重试后失败: %v\n", err)
    }
}
```

### 高级用法

```go
// 使用自定义间隔和回调的重试
err := retry.Do(func() (bool, error) {
    // 你的操作在这里
    return false, someError
}, 
    retry.Times(5),
    retry.Interval(retry.ConstantInterval(100 * time.Millisecond)),
    retry.Callback(func(n uint, err error) {
        log.Printf("重试尝试 %d 失败: %v", n, err)
    }),
)
```

## API 参考

### 核心类型

#### `F` 函数类型
```go
type F func() (bool, error)
```

要重试的函数类型。返回：
- `bool`: `true` 停止重试，`false` 继续重试
- `error`: `nil` 表示成功，非 nil 表示失败

**重试行为:**
- 如果错误为 `nil`，无论 bool 值如何，重试都会立即停止
- 如果错误不为 `nil`，bool 值决定是否继续重试

#### `Intervaler` 接口
```go
type Intervaler interface {
    Interval(n uint)
}
```

定义重试间隔策略的接口。实现决定重试尝试之间的等待时间。

### 配置函数

#### `Times(n uint)`
设置最大重试尝试次数。

```go
retry.Times(3) // 最多重试 3 次
```

#### `Interval(s Intervaler)`
设置确定重试之间等待时间的间隔策略。

```go
// 使用固定间隔
retry.Interval(retry.ConstantInterval(time.Second))

// 使用默认指数退避
retry.Interval(retry.DefaultInterval())
```

#### `Callback(c func(n uint, err error))`
设置每次重试尝试后调用的回调函数。

```go
retry.Callback(func(n uint, err error) {
    log.Printf("重试尝试 %d 失败: %v", n, err)
})
```

### 主要函数

#### `Do(f F, pf ...settings) error`
使用重试逻辑执行给定函数。

**参数:**
- `f`: 要重试的函数
- `pf`: 配置重试行为的可选设置

**返回:**
- `error`: 如果函数成功则为 `nil`，如果所有重试都失败则返回最后一个错误

## 内置间隔策略

### 默认指数退避
```go
retry.DefaultInterval()
```

实现带抖动的指数退避：在重试尝试之间睡眠 2^n 秒加上随机抖动，防止雷群问题。

### 固定间隔
```go
retry.ConstantInterval(duration)
```

在所有重试尝试之间睡眠相同的时间。

```go
retry.ConstantInterval(100 * time.Millisecond)
```

### 带抖动的指数退避
```go
retry.ExponentialBackoffWithJitter(base, jitter)
```

实现带可配置抖动的指数退避，防止雷群问题。

**参数:**
- `base`: 第一次重试的基础时间
- `jitter`: 抖动因子（0.0 到 1.0），其中 0.0 = 无抖动，1.0 = 100% 抖动

```go
// 10% 抖动
retry.ExponentialBackoffWithJitter(time.Second, 0.1)

// 50% 抖动
retry.ExponentialBackoffWithJitter(time.Second, 0.5)
```

## 自定义间隔策略

您可以通过实现 `Intervaler` 接口来实现自己的间隔策略：

```go
type customInterval struct {
    base time.Duration
}

func (c *customInterval) Interval(n uint) {
    // 自定义逻辑
    time.Sleep(c.base * time.Duration(n))
}

// 使用
err := retry.Do(func() (bool, error) {
    return false, someError
}, 
    retry.Times(3),
    retry.Interval(&customInterval{base: 100 * time.Millisecond}),
)
```

## 示例

### HTTP 请求重试
```go
func makeHTTPRequest() (bool, error) {
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        return false, err // 继续重试
    }
    defer resp.Body.Close()
    
    if resp.StatusCode >= 500 {
        return false, fmt.Errorf("服务器错误: %d", resp.StatusCode)
    }
    
    if resp.StatusCode >= 400 {
        return true, fmt.Errorf("客户端错误: %d", resp.StatusCode) // 停止重试
    }
    
    return true, nil // 成功
}

err := retry.Do(makeHTTPRequest, 
    retry.Times(5),
    retry.Interval(retry.ConstantInterval(1 * time.Second)),
    retry.Callback(func(n uint, err error) {
        log.Printf("HTTP 请求尝试 %d 失败: %v", n, err)
    }),
)
```

### 数据库操作重试
```go
func saveToDatabase(data interface{}) (bool, error) {
    err := db.Save(data)
    if err != nil {
        // 检查是否为可重试错误
        if isRetryableError(err) {
            return false, err // 继续重试
        }
        return true, err // 停止重试
    }
    return true, nil // 成功
}

err := retry.Do(func() (bool, error) {
    return saveToDatabase(myData)
}, retry.Times(3))
```

### 条件重试逻辑
```go
func processWithCondition() (bool, error) {
    result, err := someOperation()
    if err != nil {
        // 只对特定错误重试
        if errors.Is(err, ErrTemporary) {
            return false, err
        }
        return true, err // 对永久错误不重试
    }
    
    // 检查结果并决定是否重试
    if result.NeedsRetry {
        return false, errors.New("结果需要重试")
    }
    
    return true, nil // 成功
}
```

## 最佳实践

1. **选择适当的重试次数**: 不要对可能永久失败的操作重试太多次
2. **使用指数退避**: 对于分布式系统，使用指数退避避免雷群问题
3. **处理不同类型的错误**: 区分可重试和不可重试的错误
4. **设置超时**: 考虑重试操作的整体超时
5. **记录重试尝试**: 使用回调监控和记录重试行为
6. **测试重试逻辑**: 确保重试逻辑在各种场景下正确工作

## 错误处理

如果所有重试都耗尽，重试包会返回遇到的最后一个错误。确保适当处理：

```go
err := retry.Do(myFunction, retry.Times(3))
if err != nil {
    // 处理最终错误
    log.Printf("操作在 3 次重试后失败: %v", err)
}
```

## 性能考虑

- **内存使用**: 该包轻量级，内存开销最小
- **Goroutine 安全**: 所有函数都是并发安全的
- **CPU 使用**: 指数退避有助于减少重试期间的 CPU 使用
- **网络效率**: 使用适当的间隔避免压倒远程服务

## 许可证

此包是 goal 项目的一部分。有关详细信息，请参阅主项目许可证。

## 贡献

欢迎贡献！请随时提交 Pull Request。
