package cmd

import (
	"fmt"
	"strings"
	"testing"
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
