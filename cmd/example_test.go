package cmd_test

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go4x/goal/cmd"
)

// ExampleExec demonstrates basic command execution
func ExampleExec() {
	output, err := cmd.Exec("echo", "Hello World")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(strings.TrimSpace(output))
	// Output: Hello World
}

// ExampleExecWithTimeout demonstrates command execution with timeout
func ExampleExecWithTimeout() {
	// This will timeout
	output, err := cmd.ExecWithTimeout(1*time.Second, "sleep", "5")
	if err != nil {
		fmt.Printf("Command timed out: %v\n", err)
	}
	fmt.Printf("Output: %s", output)
	// Output: Command timed out: signal: killed
	// Output:
}

// ExampleExecWithDir demonstrates command execution in a specific directory
func ExampleExecWithDir() {
	output, err := cmd.ExecWithDir("/tmp", "pwd")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(strings.TrimSpace(output))
	// Output: /tmp
}

// ExampleExecWithEnv demonstrates command execution with custom environment
func ExampleExecWithEnv() {
	env := map[string]string{
		"TEST_VAR": "test_value",
	}
	output, err := cmd.ExecWithEnv(env, "echo", "$TEST_VAR")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Print(output)
	// Output: $TEST_VAR
}

// ExampleExecSeparate demonstrates separate stdout and stderr capture
func ExampleExecSeparate() {
	stdout, stderr, err := cmd.ExecSeparate("bash", "-c", "echo 'stdout message'; echo 'stderr message' >&2")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Stdout: %s", stdout)
	fmt.Printf("Stderr: %s", stderr)
	// Output:
	// Stdout: stdout message
	// Stderr: stderr message
}

// ExampleExecStream demonstrates streaming command output
func ExampleExecStream() {
	var buf bytes.Buffer
	err := cmd.ExecStream(&buf, "echo", "streamed output")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Print(buf.String())
	// Output: streamed output
}

// ExampleExecWithContext demonstrates command execution with context cancellation
func ExampleExecWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cancel the context after a short delay
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	output, err := cmd.ExecWithContext(ctx, "sleep", "5")
	if err != nil {
		fmt.Printf("Command cancelled: %v\n", err)
	}
	fmt.Printf("Output: %s", output)
	// Output: Command cancelled: signal: killed
	// Output:
}

// ExampleExecWithRetry demonstrates command execution with retry mechanism
func ExampleExecWithRetry() {
	output, err := cmd.ExecWithRetry(3, 100*time.Millisecond, "echo", "retry test")
	if err != nil {
		fmt.Printf("Error after retries: %v\n", err)
		return
	}
	fmt.Println(strings.TrimSpace(output))
	// Output: retry test
}

// ExampleExecWithValidation demonstrates command execution with output validation
func ExampleExecWithValidation() {
	validator := func(output string) error {
		if !strings.Contains(output, "success") {
			return fmt.Errorf("output does not contain 'success'")
		}
		return nil
	}

	output, err := cmd.ExecWithValidation(validator, "echo", "operation success")
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		return
	}
	fmt.Println(strings.TrimSpace(output))
	// Output: operation success
}

// ExampleExecWithOutputFormat demonstrates command execution with output formatting
func ExampleExecWithOutputFormat() {
	formatter := func(output string) string {
		lines := strings.Split(strings.TrimSpace(output), "\n")
		return fmt.Sprintf("Command output (%d lines):\n%s", len(lines), output)
	}

	formatted, err := cmd.ExecWithOutputFormat(formatter, "echo", "line1\nline2\nline3")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Print(formatted)
	// Output:
	// Command output (3 lines):
	// line1
	// line2
	// line3
}

// ExampleExecAsync demonstrates asynchronous command execution
func ExampleExecAsync() {
	outputChan, errChan := cmd.ExecAsync("echo", "async test")

	select {
	case output := <-outputChan:
		fmt.Println(strings.TrimSpace(output))
	case err := <-errChan:
		fmt.Printf("Error: %v\n", err)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
	}
	// Output: async test
}

// ExampleExecBatch demonstrates batch command execution
func ExampleExecBatch() {
	commands := []struct {
		Shell string
		Args  []string
	}{
		{"echo", []string{"Hello"}},
		{"echo", []string{"World"}},
		{"echo", []string{"Batch"}},
	}

	outputs, errors := cmd.ExecBatch(commands)

	for i, output := range outputs {
		if errors[i] != nil {
			fmt.Printf("Command %d failed: %v\n", i, errors[i])
		} else {
			fmt.Printf("Command %d: %s", i, strings.TrimSpace(output))
		}
	}
	// Output:
	// Command 0: HelloCommand 1: WorldCommand 2: Batch
}

// ExampleExecWithPipe demonstrates piped command execution
func ExampleExecWithPipe() {
	output, err := cmd.ExecWithPipe("echo", []string{"line1\nline2\nline3"}, "grep", []string{"line2"})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Print(output)
	// Output: line2
}
