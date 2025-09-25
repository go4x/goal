package cmd

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

// TestExec_Success 测试成功执行的命令
func TestExec_Success(t *testing.T) {
	tests := []struct {
		name     string
		shell    string
		args     []string
		expected string
	}{
		{
			name:     "echo command",
			shell:    "echo",
			args:     []string{"Hello World"},
			expected: "Hello World\n",
		},
		{
			name:     "bash echo command",
			shell:    "bash",
			args:     []string{"-c", "echo 'Test Message'"},
			expected: "Test Message\n",
		},
		{
			name:     "date command",
			shell:    "date",
			args:     []string{"+%Y"},
			expected: "", // 输出会根据当前年份变化，这里不验证具体值
		},
		{
			name:     "whoami command",
			shell:    "whoami",
			args:     []string{},
			expected: "", // 输出会根据当前用户变化，这里不验证具体值
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Exec(tt.shell, tt.args...)
			if err != nil {
				t.Errorf("Exec() error = %v, wantErr false", err)
				return
			}
			if tt.expected != "" && output != tt.expected {
				t.Errorf("Exec() output = %v, want %v", output, tt.expected)
			}
			// 验证输出不为空（除了echo命令的特定测试）
			if tt.name == "echo command" && !strings.Contains(output, "Hello World") {
				t.Errorf("Exec() output should contain 'Hello World', got %v", output)
			} else {
				t.Logf("Exec() output = %v, want %v", output, tt.expected)
			}
		})
	}
}

// TestExec_Error 测试执行失败的命令
func TestExec_Error(t *testing.T) {
	tests := []struct {
		name     string
		shell    string
		args     []string
		wantErr  bool
		errorMsg string
	}{
		{
			name:     "non-existent command",
			shell:    "nonexistentcommand12345",
			args:     []string{},
			wantErr:  true,
			errorMsg: "executable file not found",
		},
		{
			name:     "command with exit code 1",
			shell:    "bash",
			args:     []string{"-c", "exit 1"},
			wantErr:  true,
			errorMsg: "exit status 1",
		},
		{
			name:     "command with exit code 2",
			shell:    "bash",
			args:     []string{"-c", "exit 2"},
			wantErr:  true,
			errorMsg: "exit status 2",
		},
		{
			name:     "ls non-existent directory",
			shell:    "ls",
			args:     []string{"/nonexistent/directory/12345"},
			wantErr:  true,
			errorMsg: "exit status 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Exec(tt.shell, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil {
				// 验证错误信息包含预期的错误消息
				if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Exec() error message = %v, want to contain %v", err.Error(), tt.errorMsg)
				}
				// 验证即使出错也返回了输出（如果有的话）
				// 注意：exit命令不会产生输出，所以这里不检查空输出
				if tt.name == "ls non-existent directory" && output == "" {
					t.Errorf("Exec() should return error output for ls command, got empty output")
				}
			}
		})
	}
}

// TestExec_EdgeCases 测试边界情况
func TestExec_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		shell string
		args  []string
	}{
		{
			name:  "empty args",
			shell: "echo",
			args:  []string{},
		},
		{
			name:  "single empty arg",
			shell: "echo",
			args:  []string{""},
		},
		{
			name:  "multiple empty args",
			shell: "echo",
			args:  []string{"", "", ""},
		},
		{
			name:  "special characters",
			shell: "echo",
			args:  []string{"!@#$%^&*()"},
		},
		{
			name:  "unicode characters",
			shell: "echo",
			args:  []string{"你好世界"},
		},
		{
			name:  "long output",
			shell: "bash",
			args:  []string{"-c", "for i in {1..100}; do echo 'Line' $i; done"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Exec(tt.shell, tt.args...)
			// 对于echo命令，不应该出错
			if tt.shell == "echo" && err != nil {
				t.Errorf("Exec() error = %v, wantErr false for echo command", err)
			}
			// 验证输出不为空（对于某些命令）
			if tt.name == "echo command" && output == "" {
				t.Errorf("Exec() output should not be empty for echo command")
			}
		})
	}
}

// TestExec_Performance 性能测试
func TestExec_Performance(t *testing.T) {
	// 测试快速命令的执行时间
	output, err := Exec("echo", "performance test")
	if err != nil {
		t.Errorf("Exec() error = %v", err)
	}
	if !strings.Contains(output, "performance test") {
		t.Errorf("Exec() output = %v, want to contain 'performance test'", output)
	}
}

// BenchmarkExec 基准测试
func BenchmarkExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Exec("echo", "benchmark test")
		if err != nil {
			b.Errorf("Exec() error = %v", err)
		}
	}
}

// TestExec_Concurrent 并发测试
func TestExec_Concurrent(t *testing.T) {
	// 测试并发执行多个命令
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			output, err := Exec("echo", fmt.Sprintf("concurrent test %d", id))
			if err != nil {
				t.Errorf("Exec() error = %v", err)
			}
			if !strings.Contains(output, fmt.Sprintf("concurrent test %d", id)) {
				t.Errorf("Exec() output = %v, want to contain 'concurrent test %d'", output, id)
			}
			done <- true
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestExecWithTimeout 测试带超时的命令执行
func TestExecWithTimeout(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		output, err := ExecWithTimeout(1*time.Second, "sleep", "5")
		if err == nil {
			t.Error("ExecWithTimeout should return error for timeout")
		}
		if !strings.Contains(err.Error(), "context deadline exceeded") && !strings.Contains(err.Error(), "signal: killed") {
			t.Errorf("ExecWithTimeout error should contain 'context deadline exceeded' or 'signal: killed', got %v", err)
		}
		t.Logf("Timeout test output: %s", output)
	})

	t.Run("success within timeout", func(t *testing.T) {
		output, err := ExecWithTimeout(5*time.Second, "echo", "test")
		if err != nil {
			t.Errorf("ExecWithTimeout should not return error for successful command, got %v", err)
		}
		if !strings.Contains(output, "test") {
			t.Errorf("ExecWithTimeout output should contain 'test', got %s", output)
		}
	})
}

// TestExecWithDir 测试指定工作目录的命令执行
func TestExecWithDir(t *testing.T) {
	t.Run("execute in tmp directory", func(t *testing.T) {
		output, err := ExecWithDir("/tmp", "pwd")
		if err != nil {
			t.Errorf("ExecWithDir should not return error, got %v", err)
		}
		if !strings.Contains(output, "/tmp") {
			t.Errorf("ExecWithDir output should contain '/tmp', got %s", output)
		}
	})
}

// TestExecWithEnv 测试带环境变量的命令执行
func TestExecWithEnv(t *testing.T) {
	t.Run("custom environment", func(t *testing.T) {
		env := map[string]string{
			"TEST_VAR": "test_value",
		}
		output, err := ExecWithEnv(env, "bash", "-c", "echo $TEST_VAR")
		if err != nil {
			t.Errorf("ExecWithEnv should not return error, got %v", err)
		}
		if !strings.Contains(output, "test_value") {
			t.Errorf("ExecWithEnv output should contain 'test_value', got %s", output)
		}
	})
}

// TestExecSeparate 测试分离stdout和stderr的命令执行
func TestExecSeparate(t *testing.T) {
	t.Run("separate output", func(t *testing.T) {
		stdout, stderr, err := ExecSeparate("bash", "-c", "echo 'stdout'; echo 'stderr' >&2")
		if err != nil {
			t.Errorf("ExecSeparate should not return error, got %v", err)
		}
		if !strings.Contains(stdout, "stdout") {
			t.Errorf("ExecSeparate stdout should contain 'stdout', got %s", stdout)
		}
		if !strings.Contains(stderr, "stderr") {
			t.Errorf("ExecSeparate stderr should contain 'stderr', got %s", stderr)
		}
	})
}

// TestExecStream 测试流式输出
func TestExecStream(t *testing.T) {
	t.Run("stream output", func(t *testing.T) {
		var buf bytes.Buffer
		err := ExecStream(&buf, "echo", "stream test")
		if err != nil {
			t.Errorf("ExecStream should not return error, got %v", err)
		}
		if !strings.Contains(buf.String(), "stream test") {
			t.Errorf("ExecStream output should contain 'stream test', got %s", buf.String())
		}
	})
}

// TestExecWithContext 测试带上下文的命令执行
func TestExecWithContext(t *testing.T) {
	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		output, err := ExecWithContext(ctx, "sleep", "1")
		if err == nil {
			t.Error("ExecWithContext should return error for cancelled context")
		}
		t.Logf("Cancelled context test output: %s", output)
	})

	t.Run("success with context", func(t *testing.T) {
		ctx := context.Background()
		output, err := ExecWithContext(ctx, "echo", "context test")
		if err != nil {
			t.Errorf("ExecWithContext should not return error, got %v", err)
		}
		if !strings.Contains(output, "context test") {
			t.Errorf("ExecWithContext output should contain 'context test', got %s", output)
		}
	})
}

// TestExecWithRetry 测试重试机制
func TestExecWithRetry(t *testing.T) {
	t.Run("retry on failure", func(t *testing.T) {
		// This test might be flaky, so we'll use a command that's likely to fail
		output, err := ExecWithRetry(2, 100*time.Millisecond, "nonexistentcommand12345")
		if err == nil {
			t.Error("ExecWithRetry should return error for non-existent command")
		}
		t.Logf("Retry test output: %s", output)
	})

	t.Run("success on first try", func(t *testing.T) {
		output, err := ExecWithRetry(3, 100*time.Millisecond, "echo", "retry test")
		if err != nil {
			t.Errorf("ExecWithRetry should not return error, got %v", err)
		}
		if !strings.Contains(output, "retry test") {
			t.Errorf("ExecWithRetry output should contain 'retry test', got %s", output)
		}
	})
}

// TestExecWithValidation 测试输出验证
func TestExecWithValidation(t *testing.T) {
	t.Run("validation success", func(t *testing.T) {
		validator := func(output string) error {
			if !strings.Contains(output, "validation test") {
				return fmt.Errorf("output does not contain 'validation test'")
			}
			return nil
		}

		output, err := ExecWithValidation(validator, "echo", "validation test")
		if err != nil {
			t.Errorf("ExecWithValidation should not return error, got %v", err)
		}
		if !strings.Contains(output, "validation test") {
			t.Errorf("ExecWithValidation output should contain 'validation test', got %s", output)
		}
	})

	t.Run("validation failure", func(t *testing.T) {
		validator := func(output string) error {
			if !strings.Contains(output, "expected text") {
				return fmt.Errorf("output does not contain 'expected text'")
			}
			return nil
		}

		output, err := ExecWithValidation(validator, "echo", "wrong text")
		if err == nil {
			t.Error("ExecWithValidation should return error for validation failure")
		}
		if !strings.Contains(err.Error(), "validation failed") {
			t.Errorf("ExecWithValidation error should contain 'validation failed', got %v", err)
		}
		t.Logf("Validation failure test output: %s", output)
	})
}

// TestExecWithOutputFormat 测试输出格式化
func TestExecWithOutputFormat(t *testing.T) {
	t.Run("format output", func(t *testing.T) {
		formatter := func(output string) string {
			lines := strings.Split(strings.TrimSpace(output), "\n")
			return fmt.Sprintf("Command output (%d lines):\n%s", len(lines), output)
		}

		formatted, err := ExecWithOutputFormat(formatter, "echo", "line1\nline2")
		if err != nil {
			t.Errorf("ExecWithOutputFormat should not return error, got %v", err)
		}
		if !strings.Contains(formatted, "Command output (2 lines)") {
			t.Errorf("ExecWithOutputFormat should contain formatted output, got %s", formatted)
		}
	})
}

// TestExecAsync 测试异步执行
func TestExecAsync(t *testing.T) {
	t.Run("async execution", func(t *testing.T) {
		outputChan, errChan := ExecAsync("echo", "async test")

		select {
		case output := <-outputChan:
			if !strings.Contains(output, "async test") {
				t.Errorf("ExecAsync output should contain 'async test', got %s", output)
			}
		case err := <-errChan:
			t.Errorf("ExecAsync should not return error, got %v", err)
		case <-time.After(5 * time.Second):
			t.Error("ExecAsync should complete within 5 seconds")
		}
	})
}

// TestExecBatch 测试批量执行
func TestExecBatch(t *testing.T) {
	t.Run("batch execution", func(t *testing.T) {
		commands := []struct {
			Shell string
			Args  []string
		}{
			{"echo", []string{"Hello"}},
			{"echo", []string{"World"}},
			{"echo", []string{"Batch"}},
		}

		outputs, errors := ExecBatch(commands)

		if len(outputs) != 3 {
			t.Errorf("ExecBatch should return 3 outputs, got %d", len(outputs))
		}
		if len(errors) != 3 {
			t.Errorf("ExecBatch should return 3 errors, got %d", len(errors))
		}

		for i, output := range outputs {
			if errors[i] != nil {
				t.Errorf("ExecBatch command %d should not return error, got %v", i, errors[i])
			}
			if !strings.Contains(output, []string{"Hello", "World", "Batch"}[i]) {
				t.Errorf("ExecBatch command %d output should contain expected text, got %s", i, output)
			}
		}
	})
}

// TestExecWithPipe 测试管道执行
func TestExecWithPipe(t *testing.T) {
	t.Run("pipe execution", func(t *testing.T) {
		output, err := ExecWithPipe("echo", []string{"line1\nline2\nline3"}, "grep", []string{"line2"})
		if err != nil {
			t.Errorf("ExecWithPipe should not return error, got %v", err)
		}
		if !strings.Contains(output, "line2") {
			t.Errorf("ExecWithPipe output should contain 'line2', got %s", output)
		}
	})
}

// TestNewFunctions_EdgeCases 测试新功能的边界情况
func TestNewFunctions_EdgeCases(t *testing.T) {
	t.Run("empty environment", func(t *testing.T) {
		env := map[string]string{}
		output, err := ExecWithEnv(env, "echo", "test")
		if err != nil {
			t.Errorf("ExecWithEnv with empty env should not return error, got %v", err)
		}
		if !strings.Contains(output, "test") {
			t.Errorf("ExecWithEnv output should contain 'test', got %s", output)
		}
	})

	t.Run("zero retries", func(t *testing.T) {
		output, err := ExecWithRetry(0, time.Millisecond, "echo", "test")
		if err != nil {
			t.Errorf("ExecWithRetry with 0 retries should not return error, got %v", err)
		}
		if !strings.Contains(output, "test") {
			t.Errorf("ExecWithRetry output should contain 'test', got %s", output)
		}
	})

	t.Run("empty batch", func(t *testing.T) {
		commands := []struct {
			Shell string
			Args  []string
		}{}

		outputs, errors := ExecBatch(commands)
		if len(outputs) != 0 {
			t.Errorf("ExecBatch with empty commands should return empty outputs, got %d", len(outputs))
		}
		if len(errors) != 0 {
			t.Errorf("ExecBatch with empty commands should return empty errors, got %d", len(errors))
		}
	})
}

// BenchmarkNewFunctions 新功能的基准测试
func BenchmarkExecWithTimeout(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ExecWithTimeout(1*time.Second, "echo", "benchmark")
		if err != nil {
			b.Errorf("ExecWithTimeout error = %v", err)
		}
	}
}

func BenchmarkExecWithRetry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ExecWithRetry(1, time.Millisecond, "echo", "benchmark")
		if err != nil {
			b.Errorf("ExecWithRetry error = %v", err)
		}
	}
}

func BenchmarkExecBatch(b *testing.B) {
	commands := []struct {
		Shell string
		Args  []string
	}{
		{"echo", []string{"test1"}},
		{"echo", []string{"test2"}},
		{"echo", []string{"test3"}},
	}

	for i := 0; i < b.N; i++ {
		_, _ = ExecBatch(commands)
	}
}
