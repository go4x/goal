package iox

import (
	"bytes"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// Path provides convenient methods for path operations.
// It's a global instance that can be used directly for path-related operations.
//
// Example:
//
//	if iox.Path.PathExists("/path/to/check") {
//	    execPath := iox.Path.ExecPath()
//	    currentPath := iox.Path.CurrentPath()
//	    projectPath := iox.Path.ProjectPath()
//	    fmt.Printf("Executable: %s\n", execPath)
//	    fmt.Printf("Current: %s\n", currentPath)
//	    fmt.Printf("Project: %s\n", projectPath)
//	}
var Path = &paths{}

// paths provides methods for path operations.
type paths struct {
}

// ExecPath returns the directory path of the current executable file.
// It uses os.Executable() to get the full path to the executable and returns its directory.
// It panics if the executable path cannot be determined.
//
// Example:
//
//	execDir := iox.Path.ExecPath()
//	fmt.Printf("Executable directory: %s\n", execDir)
func (ps *paths) ExecPath() string {
	_path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(_path)
}

// CurrentPath returns the directory path of the source file that calls this method.
// It uses runtime.Caller(1) to get the caller's file path and returns its directory.
//
// Example:
//
//	currentDir := iox.Path.CurrentPath()
//	fmt.Printf("Current source directory: %s\n", currentDir)
func (ps *paths) CurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// PathExists checks if the given path exists (file or directory).
// It returns true if the path exists, false otherwise.
//
// Example:
//
//	if iox.Path.PathExists("/path/to/check") {
//	    fmt.Println("Path exists")
//	}
func (ps *paths) PathExists(path string) bool {
	return Exists(path)
}

// IsFile checks if the given path is a regular file (not a directory).
// It returns true if the path exists and is a regular file, false otherwise.
//
// Example:
//
//	if iox.Path.IsFile("/path/to/file.txt") {
//	    fmt.Println("It's a regular file")
//	}
func (ps *paths) IsFile(path string) bool {
	return IsRegularFile(path)
}

// IsDir checks if the given path is a directory.
// It returns true if the path exists and is a directory, false otherwise.
//
// Example:
//
//	if iox.Path.IsDir("/path/to/directory") {
//	    fmt.Println("It's a directory")
//	}
func (ps *paths) IsDir(path string) bool {
	return IsDir(path)
}

// ProjectPath returns the root path of the current Go project.
// It first tries to get the go.mod file path using 'go env GOMOD' command.
// If go.mod exists and modules are enabled, it returns the directory containing go.mod.
// If go.mod doesn't exist or modules are disabled, it falls back to ExecPath().
//
// Example:
//
//	projectRoot := iox.Path.ProjectPath()
//	fmt.Printf("Project root: %s\n", projectRoot)
func (ps *paths) ProjectPath() string {
	// Execute 'go env GOMOD' to get the go.mod file path
	stdout, _ := exec.Command("go", "env", "GOMOD").Output()
	p := string(bytes.TrimSpace(stdout))

	// If GOMOD returns a valid path (not "/dev/null" or empty), extract directory
	if p != "/dev/null" && p != "" {
		ss := strings.Split(p, "/")
		ss = ss[:len(ss)-1] // Remove the "go.mod" filename
		return strings.Join(ss, "/")
	}

	// Fallback to executable path if no go.mod found
	return ps.ExecPath()
}
