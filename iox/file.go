package iox

import (
	"os"
)

// File provides convenient methods for file operations.
// It's a global instance that can be used directly for file-related operations.
//
// Example:
//
//	if iox.File.Exists("/path/to/file") {
//	    info, err := iox.File.Info("/path/to/file")
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("File size: %d bytes\n", info.Size())
//	}
var File = &files{}

// files provides methods for file operations.
type files struct {
}

// Exists checks if a file (not directory) exists at the given path.
// It returns true if the path exists and is a regular file, false otherwise.
//
// Example:
//
//	if iox.File.Exists("/path/to/file.txt") {
//	    fmt.Println("File exists")
//	}
func (f *files) Exists(file string) bool {
	b := Exists(file)
	if b {
		f, _ := os.Stat(file)
		return !f.IsDir()
	}
	return b
}

// Info returns the FileInfo for the given file path.
// It returns an error if the file doesn't exist or cannot be accessed.
//
// Example:
//
//	info, err := iox.File.Info("/path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("File size: %d bytes\n", info.Size())
func (f *files) Info(file string) (os.FileInfo, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	return fi, nil
}

// Copy copies a file or directory from src to dst using the global Copy function.
// It supports the same operations as iox.Copy:
// - File to file: copies src file to dst file
// - File to directory: copies src file to dst directory (keeping original filename)
// - Directory to directory: recursively copies src directory to dst directory
//
// Example:
//
//	err := iox.File.Copy("/path/to/source.txt", "/path/to/destination.txt")
//	err := iox.File.Copy("/path/to/source.txt", "/path/to/directory/")
//	err := iox.File.Copy("/path/to/source_dir/", "/path/to/destination_dir/")
func (f *files) Copy(src, dst string) error {
	return Copy(src, dst)
}
