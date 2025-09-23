// Package iox provides utilities for file and directory operations.
// It includes functions for checking file/directory existence, path operations,
// and text file manipulation with flexible filtering capabilities.
//
// The package is designed to be simple yet powerful, offering both basic
// functionality and advanced features like the flexible file walker system
// that supports complex filter combinations.
package iox

import (
	"io"
	"os"
	"path/filepath"
)

// Exists checks if a file or directory exists at the given path.
// It returns true if the path exists, false otherwise.
//
// Example:
//
//	if iox.Exists("/path/to/file") {
//	    fmt.Println("File exists")
//	}
func Exists(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir checks if the given path is a directory.
// It returns true if the path exists and is a directory, false otherwise.
//
// Example:
//
//	if iox.IsDir("/path/to/directory") {
//	    fmt.Println("It's a directory")
//	}
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return fi.IsDir()
}

// IsRegularFile checks if the given path is a regular file (not a directory or symlink).
// It returns true if the path exists and is a regular file, false otherwise.
//
// Example:
//
//	if iox.IsRegularFile("/path/to/file.txt") {
//	    fmt.Println("It's a regular file")
//	}
func IsRegularFile(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return fi.Mode().IsRegular()
}

// Copy copies a file or directory from src to dst.
// It supports the following operations:
// - File to file: copies src file to dst file
// - File to directory: copies src file to dst directory (keeping original filename)
// - Directory to directory: recursively copies src directory to dst directory
//
// Example:
//
//	err := iox.Copy("/path/to/source.txt", "/path/to/destination.txt")
//	err := iox.Copy("/path/to/source.txt", "/path/to/directory/")
//	err := iox.Copy("/path/to/source_dir/", "/path/to/destination_dir/")
func Copy(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		return copyDir(src, dst)
	} else {
		return copyFile(src, dst)
	}
}

// copyFile copies a file from src to dst.
// If dst is a directory, the file is copied into that directory with its original name.
// If dst is a file, the file is copied to that location.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Check if dst is a directory
	dstInfo, err := os.Stat(dst)
	if err == nil && dstInfo.IsDir() {
		// dst is a directory, append source filename
		srcName := filepath.Base(src)
		dst = filepath.Join(dst, srcName)
	}

	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Copy file permissions
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

// copyDir recursively copies a directory from src to dst.
func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Create destination directory
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	// Read source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Copy each entry
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// Rename renames a file or directory from oldPath to newPath.
// It returns an error if the operation fails.
//
// Example:
//
//	err := iox.Rename("/path/to/old_name", "/path/to/new_name")
//	err := iox.Rename("/path/to/old_dir/", "/path/to/new_dir/")
func Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
