# cmd

A comprehensive command execution package for Go that provides various utilities for executing shell commands with advanced features like timeout, retry, validation, and more.

## Features

- **Basic Command Execution**: Simple command execution with combined output
- **Timeout Control**: Execute commands with configurable timeouts
- **Working Directory**: Execute commands in specific directories
- **Environment Variables**: Execute commands with custom environment variables
- **Separate Output**: Capture stdout and stderr separately
- **Streaming Output**: Stream command output to custom writers
- **Context Control**: Execute commands with context cancellation
- **Retry Mechanism**: Automatic retry with configurable attempts and delays
- **Output Validation**: Validate command output with custom validators
- **Output Formatting**: Format command output with custom formatters
- **Async Execution**: Execute commands asynchronously with channels
- **Batch Execution**: Execute multiple commands in sequence
- **Pipe Operations**: Execute piped commands (command1 | command2)

## Installation

```bash
go get github.com/go4x/goal/cmd
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/go4x/goal/cmd"
)

func main() {
    // Basic command execution
    output, err := cmd.Exec("echo", "Hello World")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(output) // Output: Hello World
}
```

## Basic Usage

### Simple Command Execution

```go
// Execute a simple command
output, err := cmd.Exec("ls", "-la")
if err != nil {
    log.Printf("Command failed: %v", err)
}
fmt.Println(output)
```

### Command with Timeout

```go
// Execute command with timeout
output, err := cmd.ExecWithTimeout(5*time.Second, "sleep", "10")
if err != nil {
    log.Printf("Command timed out: %v", err)
}
```

### Command in Specific Directory

```go
// Execute command in specific directory
output, err := cmd.ExecWithDir("/tmp", "pwd")
if err != nil {
    log.Printf("Command failed: %v", err)
}
fmt.Println(output) // Output: /tmp
```

### Command with Environment Variables

```go
// Execute command with custom environment
env := map[string]string{
    "DEBUG": "1",
    "PATH": "/usr/local/bin:/usr/bin",
}
output, err := cmd.ExecWithEnv(env, "echo", "$DEBUG")
if err != nil {
    log.Printf("Command failed: %v", err)
}
```

## Advanced Usage

### Separate stdout and stderr

```go
// Capture stdout and stderr separately
stdout, stderr, err := cmd.ExecSeparate("bash", "-c", "echo 'stdout'; echo 'stderr' >&2")
if err != nil {
    log.Printf("Command failed: %v", err)
}
fmt.Printf("Stdout: %s", stdout)
fmt.Printf("Stderr: %s", stderr)
```

### Streaming Output

```go
// Stream output to custom writer
var buf bytes.Buffer
err := cmd.ExecStream(&buf, "echo", "streamed output")
if err != nil {
    log.Printf("Command failed: %v", err)
}
fmt.Println(buf.String())
```

### Context Cancellation

```go
// Execute command with context cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Cancel after 1 second
go func() {
    time.Sleep(1 * time.Second)
    cancel()
}()

output, err := cmd.ExecWithContext(ctx, "sleep", "10")
if err != nil {
    log.Printf("Command cancelled: %v", err)
}
```

### Retry Mechanism

```go
// Execute command with retry
output, err := cmd.ExecWithRetry(3, time.Second, "curl", "http://example.com")
if err != nil {
    log.Printf("Command failed after retries: %v", err)
}
```

### Output Validation

```go
// Execute command with output validation
validator := func(output string) error {
    if !strings.Contains(output, "success") {
        return fmt.Errorf("output does not contain 'success'")
    }
    return nil
}

output, err := cmd.ExecWithValidation(validator, "echo", "operation success")
if err != nil {
    log.Printf("Validation failed: %v", err)
}
```

### Output Formatting

```go
// Execute command with output formatting
formatter := func(output string) string {
    lines := strings.Split(strings.TrimSpace(output), "\n")
    return fmt.Sprintf("Command output (%d lines):\n%s", len(lines), output)
}

formatted, err := cmd.ExecWithOutputFormat(formatter, "ls", "-la")
if err != nil {
    log.Printf("Command failed: %v", err)
}
fmt.Println(formatted)
```

### Async Execution

```go
// Execute command asynchronously
outputChan, errChan := cmd.ExecAsync("sleep", "5")

select {
case output := <-outputChan:
    fmt.Printf("Command completed: %s", output)
case err := <-errChan:
    log.Printf("Command failed: %v", err)
case <-time.After(10 * time.Second):
    fmt.Println("Timeout")
}
```

### Batch Execution

```go
// Execute multiple commands
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
        log.Printf("Command %d failed: %v", i, errors[i])
    } else {
        fmt.Printf("Command %d output: %s", i, output)
    }
}
```

### Pipe Operations

```go
// Execute piped commands
output, err := cmd.ExecWithPipe("ls", []string{"-la"}, "grep", []string{"txt"})
if err != nil {
    log.Printf("Piped command failed: %v", err)
}
fmt.Println(output)
```

## API Reference

### Basic Functions

| Function | Description |
|----------|-------------|
| `Exec(shell, args...)` | Execute command with combined output |
| `ExecWithTimeout(timeout, shell, args...)` | Execute command with timeout |
| `ExecWithDir(dir, shell, args...)` | Execute command in specific directory |
| `ExecWithEnv(env, shell, args...)` | Execute command with environment variables |

### Advanced Functions

| Function | Description |
|----------|-------------|
| `ExecSeparate(shell, args...)` | Execute command with separate stdout/stderr |
| `ExecStream(writer, shell, args...)` | Stream command output to writer |
| `ExecWithContext(ctx, shell, args...)` | Execute command with context |
| `ExecWithRetry(maxRetries, delay, shell, args...)` | Execute command with retry |
| `ExecWithValidation(validator, shell, args...)` | Execute command with output validation |
| `ExecWithOutputFormat(formatter, shell, args...)` | Execute command with output formatting |
| `ExecAsync(shell, args...)` | Execute command asynchronously |
| `ExecBatch(commands)` | Execute multiple commands |
| `ExecWithPipe(firstShell, firstArgs, secondShell, secondArgs)` | Execute piped commands |

## Error Handling

All functions return errors that can be checked:

```go
output, err := cmd.Exec("nonexistent-command")
if err != nil {
    switch {
    case strings.Contains(err.Error(), "executable file not found"):
        log.Println("Command not found")
    case strings.Contains(err.Error(), "exit status"):
        log.Println("Command failed with non-zero exit code")
    default:
        log.Printf("Unknown error: %v", err)
    }
}
```

## Performance Considerations

- **Synchronous execution** (`Exec`, `ExecWithTimeout`, etc.) blocks until completion
- **Asynchronous execution** (`ExecAsync`) returns immediately and provides results via channels
- **Batch execution** (`ExecBatch`) executes commands sequentially
- **Streaming** (`ExecStream`) is useful for large outputs to avoid memory issues

## Thread Safety

All functions are thread-safe and can be called concurrently from multiple goroutines.

## Examples

See the `example_test.go` file for comprehensive examples of all features.

## License

This package is part of the goal project and follows the same license terms.
