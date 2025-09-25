# IOx - File and Directory Operations Package

IOx is a comprehensive Go package that provides convenient utilities for file and directory operations. It offers both basic functionality and advanced features like flexible file walking with complex filter combinations.

## Features

### 🔧 Core Operations
- **File Operations**: Check existence, get file info, handle file types
- **Directory Operations**: Create, delete, check emptiness, recursive walking
- **Path Utilities**: Get executable path, current path, project root
- **Text File Handling**: Buffered reading/writing with convenient methods

### 🎯 Advanced Features
- **Flexible File Walker**: Recursive directory traversal with powerful filtering
- **Complex Filter Combinations**: AND/OR logic with multiple filter groups
- **Performance Optimized**: Efficient algorithms with benchmark testing
- **Cross-Platform**: Works on Windows, macOS, and Linux

### 📊 Quality Assurance
- **High Test Coverage**: 61.4% statement coverage with comprehensive test suite
- **Performance Benchmarked**: All major functions have performance benchmarks
- **Well Documented**: Complete API documentation with examples
- **Error Handling**: Robust error handling with clear error messages

## Installation

```bash
go get github.com/go4x/goal/iox
```

## Quick Start

### Basic File Operations

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // Check if file exists
    if iox.Exists("/path/to/file.txt") {
        fmt.Println("File exists!")
    }

    // Check if directory exists
    if iox.IsDir("/path/to/directory") {
        fmt.Println("It's a directory!")
    }

    // Check if it's a regular file
    if iox.IsRegularFile("/path/to/file.txt") {
        fmt.Println("It's a regular file!")
    }
}
```

### Directory Operations

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // Create directory
    err := iox.Dir.Create("/path/to/new/directory")
    if err != nil {
        fmt.Printf("Error creating directory: %v\n", err)
        return
    }

    // Check if directory is empty
    isEmpty, err := iox.Dir.IsEmpty("/path/to/directory")
    if err != nil {
        fmt.Printf("Error checking directory: %v\n", err)
        return
    }
    
    if isEmpty {
        fmt.Println("Directory is empty")
    }

    // Walk directory recursively
    files, err := iox.Dir.Walk("/path/to/directory")
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
        return
    }
    
    for _, file := range files {
        fmt.Println("Found file:", file)
    }
}
```

### Text File Operations

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // Create a new text file
    tf, err := iox.NewTxtFile("/path/to/file.txt")
    if err != nil {
        fmt.Printf("Error creating text file: %v\n", err)
        return
    }
    defer tf.Close()

    // Write lines to file
    _, err = tf.WriteLine("Hello, World!")
    if err != nil {
        fmt.Printf("Error writing line: %v\n", err)
        return
    }
    
    _, err = tf.WriteLine("This is line 2")
    if err != nil {
        fmt.Printf("Error writing line: %v\n", err)
        return
    }

    // Flush buffer to ensure data is written
    err = tf.Flush()
    if err != nil {
        fmt.Printf("Error flushing: %v\n", err)
        return
    }

    // Read all lines from file
    lines, err := tf.ReadAll()
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    for _, line := range lines {
        fmt.Println("Line:", line)
    }
}
```

### Advanced File Walking with Filters

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // Simple filtering - find all .go files
    goFiles, err := iox.WalkDir("/path/to/project", iox.FilterByExtension(".go"))
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
        return
    }
    
    fmt.Printf("Found %d Go files\n", len(goFiles))
    for _, file := range goFiles {
        fmt.Println("Go file:", file)
    }

    // Complex filtering with multiple filter groups
    goGroup := iox.NewFilterGroup(iox.FilterAnd,
        iox.FilterByExtension(".go"),
        iox.FilterHidden,
    )
    
    txtGroup := iox.NewFilterGroup(iox.FilterAnd,
        iox.FilterByExtension(".txt"),
        iox.FilterByName("README"),
    )
    
    // Files matching either group will be included (OR logic between groups)
    files, err := iox.WalkDirWithFilters("/path/to/project", goGroup, txtGroup)
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
        return
    }
    
    fmt.Printf("Found %d matching files\n", len(files))
    for _, file := range files {
        fmt.Println("Matching file:", file)
    }
}
```

### Path Utilities

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/iox"
)

func main() {
    // Get executable directory
    execPath := iox.Path.ExecPath()
    fmt.Printf("Executable directory: %s\n", execPath)

    // Get current source file directory
    currentPath := iox.Path.CurrentPath()
    fmt.Printf("Current directory: %s\n", currentPath)

    // Get Go project root
    projectPath := iox.Path.ProjectPath()
    fmt.Printf("Project root: %s\n", projectPath)

    // Check if path is a file
    if iox.Path.IsFile("/path/to/file.txt") {
        fmt.Println("It's a file!")
    }

    // Check if path is a directory
    if iox.Path.IsDir("/path/to/directory") {
        fmt.Println("It's a directory!")
    }
}
```

## API Reference

### Core Functions

#### File and Directory Existence
```go
// Check if file or directory exists
func Exists(file string) bool

// Check if path exists and return whether it's a directory
func PathExists(path string) (bool, error)

// Check if path is a directory
func IsDir(path string) bool

// Check if path is a regular file
func IsRegularFile(path string) bool
```

#### Global Instances
```go
// File operations
var File *files

// Directory operations  
var Dir *dirs

// Path utilities
var Path *paths
```

### File Operations

#### File Instance Methods
```go
// Check if file exists (not directory)
func (f *files) Exists(file string) bool

// Get file information
func (f *files) Info(file string) (os.FileInfo, error)
```

### Directory Operations

#### Directory Instance Methods
```go
// Check if directory exists
func (d *dirs) Exists(dir string) (bool, error)

// Append path separator if needed
func (d *dirs) AppendSeparator(dir string) string

// Create directory and parent directories
func (d *dirs) Create(dir string) error

// Create directory only if it doesn't exist
func (d *dirs) CreateIfNotExists(dir string) error

// Check if directory is empty
func (d *dirs) IsEmpty(dir string) (bool, error)

// Delete directory recursively
func (d *dirs) Delete(dir string) error

// Delete directory only if it exists
func (d *dirs) DeleteIfExists(dir string) error

// Delete directory only if it's empty
func (d *dirs) DeleteIfEmpty(dir string) error

// Walk directory recursively (returns all files)
func (d *dirs) Walk(dir string) ([]string, error)
```

### Path Utilities

#### Path Instance Methods
```go
// Get executable directory path
func (ps *paths) ExecPath() string

// Get current source file directory
func (ps *paths) CurrentPath() string

// Check if path exists
func (ps *paths) PathExists(path string) bool

// Check if path is a file
func (ps *paths) IsFile(path string) bool

// Check if path is a directory
func (ps *paths) IsDir(path string) bool

// Get Go project root path
func (ps *paths) ProjectPath() string
```

### Text File Operations

#### TxtFile Methods
```go
// Create new text file
func NewTxtFile(f string) (*TxtFile, error)

// Write line to file (buffered)
func (tf *TxtFile) WriteLine(s string) (*TxtFile, error)

// Flush buffered data to disk
func (tf *TxtFile) Flush() error

// Close file and cleanup
func (tf *TxtFile) Close() error

// Read all lines from file
func (tf *TxtFile) ReadAll() ([]string, error)
```

### File Walking and Filtering

#### Walking Functions
```go
// Walk directory with simple filters
func WalkDir(dir string, filters ...WalkFilter) ([]string, error)

// Walk directory with complex filter groups
func WalkDirWithFilters(dir string, filterGroups ...FilterGroup) ([]string, error)
```

#### Filter Functions
```go
// Filter by file extensions
func FilterByExtension(extensions ...string) WalkFilter

// Filter by name pattern (contains)
func FilterByName(pattern string) WalkFilter

// Filter by size range
func FilterBySize(minSize, maxSize int64) WalkFilter

// Filter by path pattern (regex)
func FilterByPathPattern(pattern string) WalkFilter

// Include only directories
func FilterDirectoriesOnly(entry os.DirEntry, path string) bool

// Include only files
func FilterFilesOnly(entry os.DirEntry, path string) bool

// Exclude hidden files
func FilterHidden(entry os.DirEntry, path string) bool
```

#### Filter Group Operations
```go
// Create new filter group
func NewFilterGroup(combiner FilterCombiner, filters ...WalkFilter) FilterGroup

// Apply filter group to directory entry
func (fg FilterGroup) Apply(entry os.DirEntry, path string) bool
```

#### Filter Combiners
```go
// Logical AND combination
const FilterAnd FilterCombiner

// Logical OR combination  
const FilterOr FilterCombiner
```

## Advanced Usage

### Complex Filter Combinations

The file walker supports sophisticated filtering with multiple filter groups:

```go
// Create filter groups with different logic
goFiles := iox.NewFilterGroup(iox.FilterAnd,
    iox.FilterByExtension(".go"),
    iox.FilterHidden, // Exclude hidden files
)

largeFiles := iox.NewFilterGroup(iox.FilterAnd,
    iox.FilterBySize(1024, 1024*1024), // 1KB to 1MB
    iox.FilterFilesOnly,
)

// Files matching either group will be included
files, err := iox.WalkDirWithFilters("/project", goFiles, largeFiles)
```

### Performance Considerations

The package is optimized for performance:

```go
// Benchmark results (typical):
// Exists: ~1,175 ns/op
// IsDir: ~1,177 ns/op  
// WalkDir: ~25,000 ns/op (100 files)
// WriteLine: ~71 ns/op
// ReadAll: ~25,382 ns/op (1000 lines)
```

### Error Handling

All functions return appropriate errors:

```go
// Always check errors
files, err := iox.WalkDir("/path")
if err != nil {
    log.Fatalf("WalkDir failed: %v", err)
}

// Handle specific error types
if iox.Exists("/path") {
    // Path exists, safe to proceed
} else {
    // Path doesn't exist, handle accordingly
}
```

## Testing

The package includes comprehensive tests:

```bash
# Run all tests
go test ./iox

# Run tests with coverage
go test -cover ./iox

# Run benchmarks
go test -bench=. ./iox

# Run specific test
go test -run TestWalkDir ./iox
```

### Test Coverage

- **Statement Coverage**: 61.4%
- **Test Types**: Unit tests, integration tests, edge case tests, performance benchmarks
- **Test Files**: 7 test files with 25+ test functions

## Performance Benchmarks

Key performance metrics:

| Function | Performance | Notes |
|----------|-------------|-------|
| `Exists` | ~1,175 ns/op | Basic file existence check |
| `IsDir` | ~1,177 ns/op | Directory type check |
| `WalkDir` | ~25,000 ns/op | 100 files traversal |
| `WriteLine` | ~71 ns/op | Buffered text writing |
| `ReadAll` | ~25,382 ns/op | 1000 lines reading |
| `ExecPath` | ~66 ns/op | Executable path |
| `ProjectPath` | ~6.8 ms/op | Includes `go env GOMOD` |

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

### Development Guidelines

- Follow Go conventions
- Add comprehensive tests
- Update documentation
- Run benchmarks for performance-critical code
- Maintain backward compatibility

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Changelog

### v1.0.0
- Initial release
- Core file and directory operations
- Text file handling
- Advanced file walking with filters
- Comprehensive test suite
- Performance benchmarks
- Complete documentation