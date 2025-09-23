package iox_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go4x/goal/iox"
	"github.com/go4x/got"
)

func TestPathPathExists(t *testing.T) {
	lg := got.New(t, "test Path.PathExists")

	lg.Case("give an exists path")
	path := ".." // Use relative path that should exist
	lg.Require(iox.Path.PathExists(path), "given path should exist")

	lg.Case("give an none exists path")
	path = "/nonexistent/path/that/should/not/exist"
	lg.Require(!iox.Path.PathExists(path), "given path should not exist")

	lg.Case("give an empty string")
	path = ""
	lg.Require(!iox.Path.PathExists(path), "empty string should not exist")
}

func TestIsFile(t *testing.T) {
	lg := got.New(t, "test IsFile")

	lg.Case("give an existing file")
	f := "./path_test.go"
	lg.Require(iox.Path.IsFile(f), "should be a file")

	lg.Case("give an existing directory")
	f = ".."
	lg.Require(!iox.Path.IsFile(f), "should not be a file")

	lg.Case("give a non-existing path")
	f = "/non/existent/path"
	lg.Require(!iox.Path.IsFile(f), "should not be a file")

	lg.Case("give an empty string")
	f = ""
	lg.Require(!iox.Path.IsFile(f), "empty string should not be a file")
}

func TestPathIsDir(t *testing.T) {
	lg := got.New(t, "test Path.IsDir")

	lg.Case("give an existing directory")
	d := ".."
	lg.Require(iox.Path.IsDir(d), "should be a directory")

	lg.Case("give an existing file")
	d = "./path_test.go"
	lg.Require(!iox.Path.IsDir(d), "should not be a directory")

	lg.Case("give a non-existing path")
	d = "/non/existent/path"
	lg.Require(!iox.Path.IsDir(d), "should not be a directory")

	lg.Case("give an empty string")
	d = ""
	lg.Require(!iox.Path.IsDir(d), "empty string should not be a directory")
}

func TestExecPath(t *testing.T) {
	lg := got.New(t, "test ExecPath")

	lg.Case("get executable path")
	execPath := iox.Path.ExecPath()
	lg.Require(execPath != "", "should not be empty")

	// Verify it's a valid directory
	lg.Require(iox.Exists(execPath), "should be a valid path")
	lg.Require(iox.IsDir(execPath), "should be a directory")
}

func TestCurrentPath(t *testing.T) {
	lg := got.New(t, "test CurrentPath")

	lg.Case("get current path")
	currentPath := iox.Path.CurrentPath()
	lg.Require(currentPath != "", "should not be empty")

	// Verify it's a valid directory
	lg.Require(iox.Exists(currentPath), "should be a valid path")
	lg.Require(iox.IsDir(currentPath), "should be a directory")
}

func TestProjectPath(t *testing.T) {
	lg := got.New(t, "test ProjectPath")

	lg.Case("get project path")
	projectPath := iox.Path.ProjectPath()
	lg.Require(projectPath != "", "should not be empty")

	// Verify it's a valid directory
	lg.Require(iox.Exists(projectPath), "should be a valid path")
	lg.Require(iox.IsDir(projectPath), "should be a directory")
}

// TestPathOperations_EdgeCases tests edge cases for path operations
func TestPathOperations_EdgeCases(t *testing.T) {
	lg := got.New(t, "test path operations edge cases")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	// Create a temporary file
	tempFile := filepath.Join(tempDir, "test_file.txt")
	err = os.WriteFile(tempFile, []byte("test content"), 0644)
	lg.Require(err == nil, "should create temp file")

	lg.Case("test PathExists with temp file")
	lg.Require(iox.Path.PathExists(tempFile), "temp file should exist")

	lg.Case("test PathExists with temp dir")
	lg.Require(iox.Path.PathExists(tempDir), "temp dir should exist")

	lg.Case("test IsFile with temp file")
	lg.Require(iox.Path.IsFile(tempFile), "temp file should be a file")

	lg.Case("test IsFile with temp dir")
	lg.Require(!iox.Path.IsFile(tempDir), "temp dir should not be a file")

	lg.Case("test IsDir with temp dir")
	lg.Require(iox.Path.IsDir(tempDir), "temp dir should be a directory")

	lg.Case("test IsDir with temp file")
	lg.Require(!iox.Path.IsDir(tempFile), "temp file should not be a directory")

	// Test with non-existent paths
	nonExistentFile := filepath.Join(tempDir, "non_existent.txt")
	lg.Case("test PathExists with non-existent file")
	lg.Require(!iox.Path.PathExists(nonExistentFile), "non-existent file should not exist")

	lg.Case("test IsFile with non-existent file")
	lg.Require(!iox.Path.IsFile(nonExistentFile), "non-existent file should not be a file")

	lg.Case("test IsDir with non-existent file")
	lg.Require(!iox.Path.IsDir(nonExistentFile), "non-existent file should not be a directory")
}
