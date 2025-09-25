# cmd

一个功能丰富的 Go 命令执行包，提供了各种执行 shell 命令的工具，包括超时控制、重试机制、输出验证等高级功能。

## 功能特性

- **基础命令执行**：简单的命令执行，返回合并输出
- **超时控制**：可配置超时的命令执行
- **工作目录**：在指定目录中执行命令
- **环境变量**：使用自定义环境变量执行命令
- **分离输出**：分别捕获 stdout 和 stderr
- **流式输出**：将命令输出流式写入自定义写入器
- **上下文控制**：使用上下文取消机制执行命令
- **重试机制**：可配置重试次数和延迟的自动重试
- **输出验证**：使用自定义验证器验证命令输出
- **输出格式化**：使用自定义格式化器格式化命令输出
- **异步执行**：使用通道异步执行命令
- **批量执行**：顺序执行多个命令
- **管道操作**：执行管道命令（command1 | command2）

## 安装

```bash
go get github.com/go4x/goal/cmd
```

## 快速开始

```go
package main

import (
    "fmt"
    "log"
    "github.com/go4x/goal/cmd"
)

func main() {
    // 基础命令执行
    output, err := cmd.Exec("echo", "Hello World")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(output) // 输出: Hello World
}
```

## 基础用法

### 简单命令执行

```go
// 执行简单命令
output, err := cmd.Exec("ls", "-la")
if err != nil {
    log.Printf("命令执行失败: %v", err)
}
fmt.Println(output)
```

### 带超时的命令执行

```go
// 执行带超时的命令
output, err := cmd.ExecWithTimeout(5*time.Second, "sleep", "10")
if err != nil {
    log.Printf("命令超时: %v", err)
}
```

### 在指定目录中执行命令

```go
// 在指定目录中执行命令
output, err := cmd.ExecWithDir("/tmp", "pwd")
if err != nil {
    log.Printf("命令执行失败: %v", err)
}
fmt.Println(output) // 输出: /tmp
```

### 使用环境变量执行命令

```go
// 使用自定义环境变量执行命令
env := map[string]string{
    "DEBUG": "1",
    "PATH": "/usr/local/bin:/usr/bin",
}
output, err := cmd.ExecWithEnv(env, "echo", "$DEBUG")
if err != nil {
    log.Printf("命令执行失败: %v", err)
}
```

## 高级用法

### 分离 stdout 和 stderr

```go
// 分别捕获 stdout 和 stderr
stdout, stderr, err := cmd.ExecSeparate("bash", "-c", "echo 'stdout'; echo 'stderr' >&2")
if err != nil {
    log.Printf("命令执行失败: %v", err)
}
fmt.Printf("标准输出: %s", stdout)
fmt.Printf("错误输出: %s", stderr)
```

### 流式输出

```go
// 将输出流式写入自定义写入器
var buf bytes.Buffer
err := cmd.ExecStream(&buf, "echo", "流式输出")
if err != nil {
    log.Printf("命令执行失败: %v", err)
}
fmt.Println(buf.String())
```

### 上下文取消

```go
// 使用上下文取消机制执行命令
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// 1秒后取消
go func() {
    time.Sleep(1 * time.Second)
    cancel()
}()

output, err := cmd.ExecWithContext(ctx, "sleep", "10")
if err != nil {
    log.Printf("命令被取消: %v", err)
}
```

### 重试机制

```go
// 执行带重试的命令
output, err := cmd.ExecWithRetry(3, time.Second, "curl", "http://example.com")
if err != nil {
    log.Printf("重试后命令仍然失败: %v", err)
}
```

### 输出验证

```go
// 执行带输出验证的命令
validator := func(output string) error {
    if !strings.Contains(output, "success") {
        return fmt.Errorf("输出不包含 'success'")
    }
    return nil
}

output, err := cmd.ExecWithValidation(validator, "echo", "operation success")
if err != nil {
    log.Printf("验证失败: %v", err)
}
```

### 输出格式化

```go
// 执行带输出格式化的命令
formatter := func(output string) string {
    lines := strings.Split(strings.TrimSpace(output), "\n")
    return fmt.Sprintf("命令输出 (%d 行):\n%s", len(lines), output)
}

formatted, err := cmd.ExecWithOutputFormat(formatter, "ls", "-la")
if err != nil {
    log.Printf("命令执行失败: %v", err)
}
fmt.Println(formatted)
```

### 异步执行

```go
// 异步执行命令
outputChan, errChan := cmd.ExecAsync("sleep", "5")

select {
case output := <-outputChan:
    fmt.Printf("命令完成: %s", output)
case err := <-errChan:
    log.Printf("命令失败: %v", err)
case <-time.After(10 * time.Second):
    fmt.Println("超时")
}
```

### 批量执行

```go
// 执行多个命令
commands := []struct {
    Shell string
    Args  []string
}{
    {"echo", []string{"Hello"}},
    {"echo", []string{"World"}},
    {"ls", []string{"-la"}},
}

outputs, errors := cmd.ExecBatch(commands)
for i, output := range outputs {
    if errors[i] != nil {
        log.Printf("命令 %d 失败: %v", i, errors[i])
    } else {
        fmt.Printf("命令 %d 输出: %s", i, output)
    }
}
```

### 管道操作

```go
// 执行管道命令
output, err := cmd.ExecWithPipe("ls", []string{"-la"}, "grep", []string{"txt"})
if err != nil {
    log.Printf("管道命令失败: %v", err)
}
fmt.Println(output)
```

## API 参考

### 基础函数

| 函数 | 描述 |
|------|------|
| `Exec(shell, args...)` | 执行命令并返回合并输出 |
| `ExecWithTimeout(timeout, shell, args...)` | 执行带超时的命令 |
| `ExecWithDir(dir, shell, args...)` | 在指定目录中执行命令 |
| `ExecWithEnv(env, shell, args...)` | 使用环境变量执行命令 |

### 高级函数

| 函数 | 描述 |
|------|------|
| `ExecSeparate(shell, args...)` | 执行命令并分别返回 stdout/stderr |
| `ExecStream(writer, shell, args...)` | 将命令输出流式写入写入器 |
| `ExecWithContext(ctx, shell, args...)` | 使用上下文执行命令 |
| `ExecWithRetry(maxRetries, delay, shell, args...)` | 执行带重试的命令 |
| `ExecWithValidation(validator, shell, args...)` | 执行带输出验证的命令 |
| `ExecWithOutputFormat(formatter, shell, args...)` | 执行带输出格式化的命令 |
| `ExecAsync(shell, args...)` | 异步执行命令 |
| `ExecBatch(commands)` | 执行多个命令 |
| `ExecWithPipe(firstShell, firstArgs, secondShell, secondArgs)` | 执行管道命令 |

## 错误处理

所有函数都返回可以检查的错误：

```go
output, err := cmd.Exec("nonexistent-command")
if err != nil {
    switch {
    case strings.Contains(err.Error(), "executable file not found"):
        log.Println("命令未找到")
    case strings.Contains(err.Error(), "exit status"):
        log.Println("命令以非零退出码失败")
    default:
        log.Printf("未知错误: %v", err)
    }
}
```

## 性能考虑

- **同步执行**（`Exec`、`ExecWithTimeout` 等）会阻塞直到完成
- **异步执行**（`ExecAsync`）立即返回并通过通道提供结果
- **批量执行**（`ExecBatch`）顺序执行命令
- **流式处理**（`ExecStream`）对于大输出很有用，可以避免内存问题

## 线程安全

所有函数都是线程安全的，可以从多个 goroutine 并发调用。

## 示例

查看 `example_test.go` 文件了解所有功能的综合示例。

## 许可证

此包是 goal 项目的一部分，遵循相同的许可证条款。
