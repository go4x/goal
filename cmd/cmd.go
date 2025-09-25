// Package cmd provides utilities for executing shell commands.
//
// This package offers a simple interface for running external commands
// and capturing their output, with proper error handling.
//
// Example:
//
//	output, err := cmd.Exec("echo", "Hello World")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(output) // Output: Hello World
package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Exec executes a shell command with the given arguments and returns its combined output.
//
// The function takes a shell command name and variable arguments, then executes
// the command using os/exec.Command. It captures both stdout and stderr output
// using CombinedOutput().
//
// Parameters:
//   - shell: The name of the command to execute (e.g., "ls", "echo", "bash")
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The combined output (stdout + stderr) from the command execution
//   - error: Any error that occurred during command execution, including:
//   - Command not found errors
//   - Permission denied errors
//   - Command execution failures (non-zero exit codes)
//   - I/O errors during output capture
//
// Note: Even if the command fails (returns non-zero exit code), this function
// will still return the output along with the error. This allows callers to
// examine both the error and any output the command may have produced.
//
// Example usage:
//
//	// Successful command
//	output, err := cmd.Exec("echo", "Hello World")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
//	fmt.Printf("Output: %s", output) // Output: Hello World
//
//	// Command with error
//	output, err := cmd.Exec("ls", "/nonexistent/directory")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//		fmt.Printf("Error output: %s", output) // Will contain error message
//	}
//
//	// Bash command with shell features
//	output, err := cmd.Exec("bash", "-c", "echo 'Hello' && echo 'World'")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
//	fmt.Printf("Output: %s", output) // Output: Hello\nWorld
//
//	// Command with no arguments
//	output, err := cmd.Exec("whoami")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
//	fmt.Printf("Current user: %s", strings.TrimSpace(output))
func Exec(shell string, args ...string) (string, error) {
	cmd := exec.Command(shell, args...)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		return string(bs), err
	}
	return string(bs), nil
}

// ExecWithTimeout executes a shell command with a timeout and returns its combined output.
// If the command takes longer than the specified timeout, it will be killed and an error returned.
//
// Parameters:
//   - timeout: The maximum time to wait for the command to complete
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The combined output (stdout + stderr) from the command execution
//   - error: Any error that occurred during command execution or timeout
//
// Example:
//
//	output, err := cmd.ExecWithTimeout(5*time.Second, "sleep", "10")
//	if err != nil {
//		log.Printf("Command timed out or failed: %v", err)
//	}
func ExecWithTimeout(timeout time.Duration, shell string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, shell, args...)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		return string(bs), err
	}
	return string(bs), nil
}

// ExecWithDir executes a shell command in the specified working directory.
//
// Parameters:
//   - dir: The working directory for the command
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The combined output (stdout + stderr) from the command execution
//   - error: Any error that occurred during command execution
//
// Example:
//
//	output, err := cmd.ExecWithDir("/tmp", "ls", "-la")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
func ExecWithDir(dir string, shell string, args ...string) (string, error) {
	cmd := exec.Command(shell, args...)
	cmd.Dir = dir
	bs, err := cmd.CombinedOutput()
	if err != nil {
		return string(bs), err
	}
	return string(bs), nil
}

// ExecWithEnv executes a shell command with custom environment variables.
//
// Parameters:
//   - env: A map of environment variables to set for the command
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The combined output (stdout + stderr) from the command execution
//   - error: Any error that occurred during command execution
//
// Example:
//
//	env := map[string]string{
//		"PATH": "/usr/local/bin:/usr/bin",
//		"DEBUG": "1",
//	}
//	output, err := cmd.ExecWithEnv(env, "echo", "$DEBUG")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
func ExecWithEnv(env map[string]string, shell string, args ...string) (string, error) {
	cmd := exec.Command(shell, args...)

	// Set environment variables
	for key, value := range env {
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", key, value))
	}

	bs, err := cmd.CombinedOutput()
	if err != nil {
		return string(bs), err
	}
	return string(bs), nil
}

// ExecSeparate executes a shell command and returns stdout and stderr separately.
//
// Parameters:
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The stdout output from the command execution
//   - string: The stderr output from the command execution
//   - error: Any error that occurred during command execution
//
// Example:
//
//	stdout, stderr, err := cmd.ExecSeparate("ls", "/nonexistent")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//		log.Printf("Stdout: %s", stdout)
//		log.Printf("Stderr: %s", stderr)
//	}
func ExecSeparate(shell string, args ...string) (string, string, error) {
	cmd := exec.Command(shell, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", "", err
	}

	if err := cmd.Start(); err != nil {
		return "", "", err
	}

	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		return "", "", err
	}

	stderrBytes, err := io.ReadAll(stderr)
	if err != nil {
		return "", "", err
	}

	if err := cmd.Wait(); err != nil {
		return string(stdoutBytes), string(stderrBytes), err
	}

	return string(stdoutBytes), string(stderrBytes), nil
}

// ExecStream executes a shell command and streams its output to the provided writer.
//
// Parameters:
//   - writer: The io.Writer to stream output to
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - error: Any error that occurred during command execution
//
// Example:
//
//	var buf bytes.Buffer
//	err := cmd.ExecStream(&buf, "echo", "Hello World")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
//	fmt.Println(buf.String()) // Output: Hello World
func ExecStream(writer io.Writer, shell string, args ...string) error {
	cmd := exec.Command(shell, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer

	return cmd.Run()
}

// ExecWithContext executes a shell command with a context for cancellation.
//
// Parameters:
//   - ctx: The context for cancellation
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The combined output (stdout + stderr) from the command execution
//   - error: Any error that occurred during command execution
//
// Example:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	output, err := cmd.ExecWithContext(ctx, "sleep", "10")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
func ExecWithContext(ctx context.Context, shell string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, shell, args...)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		return string(bs), err
	}
	return string(bs), nil
}

// ExecWithRetry executes a shell command with retry mechanism.
//
// Parameters:
//   - maxRetries: Maximum number of retry attempts
//   - delay: Delay between retries
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The combined output (stdout + stderr) from the command execution
//   - error: Any error that occurred during command execution after all retries
//
// Example:
//
//	output, err := cmd.ExecWithRetry(3, time.Second, "curl", "http://example.com")
//	if err != nil {
//		log.Printf("Command failed after retries: %v", err)
//	}
func ExecWithRetry(maxRetries int, delay time.Duration, shell string, args ...string) (string, error) {
	var lastErr error
	var output string

	for i := 0; i <= maxRetries; i++ {
		output, lastErr = Exec(shell, args...)
		if lastErr == nil {
			return output, nil
		}

		if i < maxRetries {
			time.Sleep(delay)
		}
	}

	return output, lastErr
}

// ExecWithValidation executes a shell command and validates the output.
//
// Parameters:
//   - validator: A function to validate the command output
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The combined output (stdout + stderr) from the command execution
//   - error: Any error that occurred during command execution or validation
//
// Example:
//
//	validator := func(output string) error {
//		if !strings.Contains(output, "success") {
//			return fmt.Errorf("output does not contain 'success'")
//		}
//		return nil
//	}
//
//	output, err := cmd.ExecWithValidation(validator, "echo", "operation success")
//	if err != nil {
//		log.Printf("Command failed or validation failed: %v", err)
//	}
func ExecWithValidation(validator func(string) error, shell string, args ...string) (string, error) {
	output, err := Exec(shell, args...)
	if err != nil {
		return output, err
	}

	if err := validator(output); err != nil {
		return output, fmt.Errorf("validation failed: %w", err)
	}

	return output, nil
}

// ExecWithOutputFormat executes a shell command and formats the output.
//
// Parameters:
//   - formatter: A function to format the command output
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - string: The formatted output from the command execution
//   - error: Any error that occurred during command execution
//
// Example:
//
//	formatter := func(output string) string {
//		lines := strings.Split(strings.TrimSpace(output), "\n")
//		return fmt.Sprintf("Command output (%d lines):\n%s", len(lines), output)
//	}
//
//	formatted, err := cmd.ExecWithOutputFormat(formatter, "ls", "-la")
//	if err != nil {
//		log.Printf("Command failed: %v", err)
//	}
//	fmt.Println(formatted)
func ExecWithOutputFormat(formatter func(string) string, shell string, args ...string) (string, error) {
	output, err := Exec(shell, args...)
	if err != nil {
		return output, err
	}

	return formatter(output), nil
}

// ExecAsync executes a shell command asynchronously and returns a channel for the result.
//
// Parameters:
//   - shell: The name of the command to execute
//   - args: Variable number of string arguments to pass to the command
//
// Returns:
//   - <-chan string: A channel that will receive the command output
//   - <-chan error: A channel that will receive any error
//
// Example:
//
//	outputChan, errChan := cmd.ExecAsync("sleep", "5")
//
//	select {
//	case output := <-outputChan:
//		fmt.Printf("Command completed: %s", output)
//	case err := <-errChan:
//		log.Printf("Command failed: %v", err)
//	}
func ExecAsync(shell string, args ...string) (<-chan string, <-chan error) {
	outputChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		output, err := Exec(shell, args...)
		if err != nil {
			errChan <- err
			return
		}
		outputChan <- output
	}()

	return outputChan, errChan
}

// ExecBatch executes multiple shell commands in sequence.
//
// Parameters:
//   - commands: A slice of command structures
//
// Returns:
//   - []string: The outputs from all commands
//   - []error: The errors from all commands
//
// Example:
//
//	commands := []struct {
//		Shell string
//		Args  []string
//	}{
//		{"echo", []string{"Hello"}},
//		{"echo", []string{"World"}},
//		{"ls", []string{"-la"}},
//	}
//
//	outputs, errors := cmd.ExecBatch(commands)
//	for i, output := range outputs {
//		if errors[i] != nil {
//			log.Printf("Command %d failed: %v", i, errors[i])
//		} else {
//			fmt.Printf("Command %d output: %s", i, output)
//		}
//	}
func ExecBatch(commands []struct {
	Shell string
	Args  []string
}) ([]string, []error) {
	outputs := make([]string, len(commands))
	errors := make([]error, len(commands))

	for i, cmd := range commands {
		outputs[i], errors[i] = Exec(cmd.Shell, cmd.Args...)
	}

	return outputs, errors
}

// ExecWithPipe executes a shell command and pipes its output to another command.
//
// Parameters:
//   - firstShell: The first command to execute
//   - firstArgs: Arguments for the first command
//   - secondShell: The second command to execute
//   - secondArgs: Arguments for the second command
//
// Returns:
//   - string: The combined output from the piped commands
//   - error: Any error that occurred during command execution
//
// Example:
//
//	output, err := cmd.ExecWithPipe("ls", []string{"-la"}, "grep", []string{"txt"})
//	if err != nil {
//		log.Printf("Piped command failed: %v", err)
//	}
func ExecWithPipe(firstShell string, firstArgs []string, secondShell string, secondArgs []string) (string, error) {
	firstCmd := exec.Command(firstShell, firstArgs...)
	secondCmd := exec.Command(secondShell, secondArgs...)

	// Create a pipe between the commands
	pipe, err := firstCmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	secondCmd.Stdin = pipe
	secondCmd.Stdout = &strings.Builder{}

	// Start the second command
	if err := secondCmd.Start(); err != nil {
		return "", err
	}

	// Start the first command
	if err := firstCmd.Start(); err != nil {
		return "", err
	}

	// Wait for the first command to complete
	if err := firstCmd.Wait(); err != nil {
		return "", err
	}

	// Close the pipe
	_ = pipe.Close()

	// Wait for the second command to complete
	if err := secondCmd.Wait(); err != nil {
		return "", err
	}

	// Get the output from the second command
	output := secondCmd.Stdout.(*strings.Builder).String()
	return output, nil
}
