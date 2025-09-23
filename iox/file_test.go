package iox_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go4x/goal/iox"
	"github.com/go4x/got"
)

func TestExistsFile(t *testing.T) {
	lg := got.New(t, "test ExistsFile")

	lg.Case("give an existing file")
	f := "./file_test.go"
	lg.Require(iox.File.Exists(f), "should exist")

	lg.Case("give an existing dir, but is not a file")
	f = "."
	lg.Require(!iox.File.Exists(f), "should not exist")

	lg.Case("give a non-existing path")
	f = "/non/existent/path/that/should/not/exist"
	lg.Require(!iox.File.Exists(f), "should not exist")

	lg.Case("give an empty string")
	f = ""
	lg.Require(!iox.File.Exists(f), "empty string should not exist")
}

func TestFileInfo(t *testing.T) {
	lg := got.New(t, "test FileInfo")

	lg.Case("get info of existing file")
	f := "./file_test.go"
	info, err := iox.File.Info(f)
	lg.Require(err == nil, "should not have error")
	lg.Require(info != nil, "info should not be nil")
	lg.Require(!info.IsDir(), "should not be a directory")
	lg.Require(info.Mode().IsRegular(), "should be a regular file")

	lg.Case("get info of non-existing file")
	f = "/non/existent/path/that/should/not/exist"
	info, err = iox.File.Info(f)
	lg.Require(err != nil, "should have error")
	lg.Require(info == nil, "info should be nil")

	lg.Case("get info of directory (should fail)")
	f = "."
	info, err = iox.File.Info(f)
	lg.Require(err == nil, "should not have error") // Info() uses os.Stat which doesn't fail for directories
	lg.Require(info != nil, "info should not be nil")
	lg.Require(info.IsDir(), "should be a directory")
}

// TestFileOperations_EdgeCases tests edge cases for file operations
func TestFileOperations_EdgeCases(t *testing.T) {
	lg := got.New(t, "test file operations edge cases")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	// Create a temporary file
	tempFile := filepath.Join(tempDir, "test_file.txt")
	err = os.WriteFile(tempFile, []byte("test content"), 0644)
	lg.Require(err == nil, "should create temp file")

	lg.Case("test File.Exists with temp file")
	lg.Require(iox.File.Exists(tempFile), "temp file should exist")

	lg.Case("test File.Exists with temp dir")
	lg.Require(!iox.File.Exists(tempDir), "temp dir should not exist as file")

	lg.Case("test File.Info with temp file")
	info, err := iox.File.Info(tempFile)
	lg.Require(err == nil, "should not have error")
	lg.Require(info != nil, "info should not be nil")
	lg.Require(info.Size() > 0, "file should have content")

	lg.Case("test File.Info with temp dir")
	info, err = iox.File.Info(tempDir)
	lg.Require(err == nil, "should not have error") // Info() doesn't distinguish between files and dirs
	lg.Require(info != nil, "info should not be nil")
	lg.Require(info.IsDir(), "should be a directory")

	// Test with a file that doesn't exist
	nonExistentFile := filepath.Join(tempDir, "non_existent.txt")
	lg.Case("test File.Exists with non-existent file")
	lg.Require(!iox.File.Exists(nonExistentFile), "non-existent file should not exist")

	lg.Case("test File.Info with non-existent file")
	info, err = iox.File.Info(nonExistentFile)
	lg.Require(err != nil, "should have error")
	lg.Require(info == nil, "info should be nil")
}
