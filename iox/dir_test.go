package iox_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go4x/goal/iox"
	"github.com/go4x/got"
)

func TestDirExists(t *testing.T) {
	lg := got.New(t, "test DirExists")

	lg.Case("give an existing dir")
	dir := ".."
	exists, err := iox.Dir.Exists(dir)
	lg.Require(err == nil, "should not have error")
	lg.Require(exists, "should exist")

	lg.Case("give an existing file, but is not a dir")
	dir = "./file_test.go"
	exists, err = iox.Dir.Exists(dir)
	lg.Require(err == nil, "should not have error")
	lg.Require(!exists, "should not exist")

	lg.Case("give a none exists path")
	dir = "/nonexistent/path/that/should/not/exist"
	exists, err = iox.Dir.Exists(dir)
	lg.Require(err == nil, "should not have error")
	lg.Require(!exists, "should not exist")

	lg.Case("give an empty string")
	dir = ""
	exists, err = iox.Dir.Exists(dir)
	lg.Require(err == nil, "should not have error")
	lg.Require(!exists, "should not exist")
}

func TestAppendSeparator(t *testing.T) {
	lg := got.New(t, "test AppendSeparator")

	lg.Case("append separator to path without trailing separator")
	dir := "/path/to/dir"
	result := iox.Dir.AppendSeparator(dir)
	expected := dir + string(filepath.Separator)
	lg.Require(result == expected, "should append separator")

	lg.Case("do not append separator if already has one")
	dir = "/path/to/dir/"
	result = iox.Dir.AppendSeparator(dir)
	lg.Require(result == dir, "should not append separator")

	lg.Case("handle empty string")
	dir = ""
	result = iox.Dir.AppendSeparator(dir)
	lg.Require(result == "", "should return empty string")

	lg.Case("handle root path")
	dir = "/"
	result = iox.Dir.AppendSeparator(dir)
	lg.Require(result == "/", "should return root path")
}

func TestCreate(t *testing.T) {
	lg := got.New(t, "test Create")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	lg.Case("create a new directory")
	newDir := filepath.Join(tempDir, "new_dir")
	err = iox.Dir.Create(newDir)
	lg.Require(err == nil, "should create directory")
	lg.Require(iox.Exists(newDir), "directory should exist")

	lg.Case("create nested directories")
	nestedDir := filepath.Join(tempDir, "level1", "level2", "level3")
	err = iox.Dir.Create(nestedDir)
	lg.Require(err == nil, "should create nested directories")
	lg.Require(iox.Exists(nestedDir), "nested directory should exist")

	lg.Case("create directory that already exists")
	err = iox.Dir.Create(newDir)
	lg.Require(err == nil, "should not error when directory exists")
}

func TestCreateIfNotExists(t *testing.T) {
	lg := got.New(t, "test CreateIfNotExists")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	lg.Case("create directory that doesn't exist")
	newDir := filepath.Join(tempDir, "new_dir")
	err = iox.Dir.CreateIfNotExists(newDir)
	lg.Require(err == nil, "should create directory")
	lg.Require(iox.Exists(newDir), "directory should exist")

	lg.Case("try to create directory that already exists")
	err = iox.Dir.CreateIfNotExists(newDir)
	lg.Require(err != nil, "should error when directory exists")
	lg.Require(err.Error() == "dir already exists", "should have correct error message")
}

func TestIsEmpty(t *testing.T) {
	lg := got.New(t, "test IsEmpty")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	lg.Case("check empty directory")
	isEmpty, err := iox.Dir.IsEmpty(tempDir)
	lg.Require(err == nil, "should not have error")
	lg.Require(isEmpty, "should be empty")

	lg.Case("check directory with file")
	tempFile := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(tempFile, []byte("test"), 0644)
	lg.Require(err == nil, "should create file")

	isEmpty, err = iox.Dir.IsEmpty(tempDir)
	lg.Require(err == nil, "should not have error")
	lg.Require(!isEmpty, "should not be empty")

	lg.Case("check non-existent directory")
	isEmpty, err = iox.Dir.IsEmpty("/non/existent/dir")
	lg.Require(err != nil, "should have error")
	lg.Require(!isEmpty, "should not be empty")
}

func TestDelete(t *testing.T) {
	lg := got.New(t, "test Delete")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")

	lg.Case("delete directory")
	err = iox.Dir.Delete(tempDir)
	lg.Require(err == nil, "should delete directory")
	lg.Require(!iox.Exists(tempDir), "directory should not exist")

	lg.Case("delete non-existent directory")
	err = iox.Dir.Delete("/non/existent/dir")
	lg.Require(err == nil, "should not error when deleting non-existent directory")

	// Test deleting directory with content
	tempDir, err = os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")

	tempFile := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(tempFile, []byte("test"), 0644)
	lg.Require(err == nil, "should create file")

	lg.Case("delete directory with content")
	err = iox.Dir.Delete(tempDir)
	lg.Require(err == nil, "should delete directory and content")
	lg.Require(!iox.Exists(tempDir), "directory should not exist")
}

func TestDeleteIfExists(t *testing.T) {
	lg := got.New(t, "test DeleteIfExists")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")

	lg.Case("delete existing directory")
	err = iox.Dir.DeleteIfExists(tempDir)
	lg.Require(err == nil, "should delete directory")
	lg.Require(!iox.Exists(tempDir), "directory should not exist")

	lg.Case("try to delete non-existent directory")
	err = iox.Dir.DeleteIfExists("/non/existent/dir")
	lg.Require(err != nil, "should error when directory doesn't exist")
	lg.Require(err.Error() == "dir does not exist", "should have correct error message")
}

func TestDeleteIfEmpty(t *testing.T) {
	lg := got.New(t, "test DeleteIfEmpty")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")

	lg.Case("delete empty directory")
	err = iox.Dir.DeleteIfEmpty(tempDir)
	lg.Require(err == nil, "should delete empty directory")
	lg.Require(!iox.Exists(tempDir), "directory should not exist")

	// Create directory with content
	tempDir, err = os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(tempFile, []byte("test"), 0644)
	lg.Require(err == nil, "should create file")

	lg.Case("try to delete non-empty directory")
	err = iox.Dir.DeleteIfEmpty(tempDir)
	lg.Require(err != nil, "should error when directory is not empty")
	lg.Require(err.Error() == "dir is not empty", "should have correct error message")

	lg.Case("try to delete non-existent directory")
	err = iox.Dir.DeleteIfEmpty("/non/existent/dir")
	lg.Require(err != nil, "should error when directory doesn't exist")
}

func TestWalk(t *testing.T) {
	lg := got.New(t, "test Walk")

	// Create a temporary directory structure for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	// Create subdirectories and files
	subDir1 := filepath.Join(tempDir, "subdir1")
	subDir2 := filepath.Join(tempDir, "subdir2")
	err = os.MkdirAll(subDir1, 0755)
	lg.Require(err == nil, "should create subdir1")
	err = os.MkdirAll(subDir2, 0755)
	lg.Require(err == nil, "should create subdir2")

	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(subDir1, "file2.txt")
	file3 := filepath.Join(subDir2, "file3.txt")

	err = os.WriteFile(file1, []byte("test1"), 0644)
	lg.Require(err == nil, "should create file1")
	err = os.WriteFile(file2, []byte("test2"), 0644)
	lg.Require(err == nil, "should create file2")
	err = os.WriteFile(file3, []byte("test3"), 0644)
	lg.Require(err == nil, "should create file3")

	lg.Case("walk directory")
	files, err := iox.Dir.Walk(tempDir)
	lg.Require(err == nil, "should not have error")
	lg.Require(len(files) == 3, "should find 3 files")

	// Check that all expected files are found
	foundFiles := make(map[string]bool)
	for _, file := range files {
		foundFiles[file] = true
	}
	lg.Require(foundFiles[file1], "should find file1")
	lg.Require(foundFiles[file2], "should find file2")
	lg.Require(foundFiles[file3], "should find file3")

	lg.Case("walk single file")
	files, err = iox.Dir.Walk(file1)
	lg.Require(err == nil, "should not have error")
	lg.Require(len(files) == 1, "should find 1 file")
	lg.Require(files[0] == file1, "should find the correct file")

	lg.Case("walk non-existent path")
	files, err = iox.Dir.Walk("/non/existent/path")
	lg.Require(err != nil, "should have error")
	lg.Require(len(files) == 0, "should return empty slice")
}
