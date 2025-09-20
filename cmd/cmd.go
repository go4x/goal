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
	"os/exec"
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
