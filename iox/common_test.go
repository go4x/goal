package iox_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go4x/goal/iox"
	"github.com/go4x/got"
)

func TestExists(t *testing.T) {
	lg := got.New(t, "test Exists")

	lg.Case("give a existing dir")
	f := ".."
	lg.Require(iox.Exists(f), "should exist")

	lg.Case("give a existing file")
	f = "./file_test.go"
	lg.Require(iox.Exists(f), "should exist")

	lg.Case("give a non-existing path")
	f = "/non/existent/path/that/should/not/exist"
	lg.Require(!iox.Exists(f), "should not exist")

	lg.Case("give an empty string")
	f = ""
	lg.Require(!iox.Exists(f), "empty string should not exist")
}

func TestIsDir(t *testing.T) {
	lg := got.New(t, "test IsDir")

	lg.Case("give a existing dir")
	f := "../iox"
	lg.Require(iox.IsDir(f), "is dir")

	lg.Case("give a existing file")
	f = "../iox/file_test.go"
	lg.Require(!iox.IsDir(f), "is not a dir")

	lg.Case("give a non-existing path")
	f = "/non/existent/path/that/should/not/exist"
	lg.Require(!iox.IsDir(f), "non-existing path is not a dir")

	lg.Case("give an empty string")
	f = ""
	lg.Require(!iox.IsDir(f), "empty string is not a dir")
}

func TestIsRegularFile(t *testing.T) {
	lg := got.New(t, "test IsRegularFile")

	lg.Case("give a existing dir")
	f := "../iox/"
	lg.Require(!iox.IsRegularFile(f), "is not a regular file")

	lg.Case("give a existing regular file")
	f = "../iox/file_test.go"
	lg.Require(iox.IsRegularFile(f), "is a regular file")

	lg.Case("give a soft symlink file")
	f = "/etc"
	lg.Require(!iox.IsRegularFile(f), "is not a regular file")

	lg.Case("give a non-existing path")
	f = "/non/existent/path/that/should/not/exist"
	lg.Require(!iox.IsRegularFile(f), "non-existing path is not a regular file")

	lg.Case("give an empty string")
	f = ""
	lg.Require(!iox.IsRegularFile(f), "empty string is not a regular file")
}

// TestCommonFunctions_EdgeCases tests edge cases for common functions
func TestCommonFunctions_EdgeCases(t *testing.T) {
	lg := got.New(t, "test common functions edge cases")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	// Create a temporary file
	tempFile := filepath.Join(tempDir, "test_file.txt")
	err = os.WriteFile(tempFile, []byte("test content"), 0644)
	lg.Require(err == nil, "should create temp file")

	lg.Case("test Exists with temp file")
	lg.Require(iox.Exists(tempFile), "temp file should exist")

	lg.Case("test Exists with temp dir")
	lg.Require(iox.Exists(tempDir), "temp dir should exist")

	lg.Case("test IsDir with temp dir")
	lg.Require(iox.IsDir(tempDir), "temp dir should be a directory")

	lg.Case("test IsDir with temp file")
	lg.Require(!iox.IsDir(tempFile), "temp file should not be a directory")

	lg.Case("test IsRegularFile with temp file")
	lg.Require(iox.IsRegularFile(tempFile), "temp file should be a regular file")

	lg.Case("test IsRegularFile with temp dir")
	lg.Require(!iox.IsRegularFile(tempDir), "temp dir should not be a regular file")

	// Test with symlink (if supported)
	symlinkFile := filepath.Join(tempDir, "symlink.txt")
	err = os.Symlink(tempFile, symlinkFile)
	if err == nil {
		defer os.Remove(symlinkFile)

		lg.Case("test Exists with symlink")
		lg.Require(iox.Exists(symlinkFile), "symlink should exist")

		lg.Case("test IsRegularFile with symlink")
		lg.Require(iox.IsRegularFile(symlinkFile), "symlink should be a regular file")

		lg.Case("test IsDir with symlink")
		lg.Require(!iox.IsDir(symlinkFile), "symlink should not be a directory")
	}

	// Test with special characters in path
	specialFile := filepath.Join(tempDir, "测试文件_特殊字符.txt")
	err = os.WriteFile(specialFile, []byte("test content"), 0644)
	lg.Require(err == nil, "should create file with special characters")

	lg.Case("test Exists with special characters")
	lg.Require(iox.Exists(specialFile), "file with special characters should exist")

	lg.Case("test IsRegularFile with special characters")
	lg.Require(iox.IsRegularFile(specialFile), "file with special characters should be regular")

	// Test with nested paths
	nestedDir := filepath.Join(tempDir, "nested", "deep", "directory")
	err = os.MkdirAll(nestedDir, 0755)
	lg.Require(err == nil, "should create nested directory")

	lg.Case("test Exists with nested path")
	lg.Require(iox.Exists(nestedDir), "nested path should exist")

	lg.Case("test IsDir with nested path")
	lg.Require(iox.IsDir(nestedDir), "nested path should be a directory")
}

// TestCommonFunctions_ErrorHandling tests error handling scenarios
func TestCommonFunctions_ErrorHandling(t *testing.T) {
	lg := got.New(t, "test common functions error handling")

	lg.Case("test with relative paths")
	lg.Require(iox.Exists("."), "current directory should exist")
	lg.Require(iox.IsDir("."), "current directory should be a directory")
	lg.Require(!iox.IsRegularFile("."), "current directory should not be a regular file")

	lg.Case("test with parent directory")
	lg.Require(iox.Exists(".."), "parent directory should exist")
	lg.Require(iox.IsDir(".."), "parent directory should be a directory")
	lg.Require(!iox.IsRegularFile(".."), "parent directory should not be a regular file")
}

func TestCopy_FileToFile(t *testing.T) {
	lg := got.New(t, "test Copy file to file")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	dstDir := filepath.Join(tempDir, "dst")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create src dir")

	// Create source file
	srcFile := filepath.Join(srcDir, "test.txt")
	content := "Hello, World! This is a test file."
	err = os.WriteFile(srcFile, []byte(content), 0644)
	lg.Require(err == nil, "should create source file")

	// Copy file to file
	dstFile := filepath.Join(dstDir, "copied.txt")
	err = iox.Copy(srcFile, dstFile)
	lg.Require(err == nil, "should copy file to file")

	// Verify destination file exists and has correct content
	lg.Require(iox.Exists(dstFile), "destination file should exist")

	copiedContent, err := os.ReadFile(dstFile)
	lg.Require(err == nil, "should read copied file")
	lg.Require(string(copiedContent) == content, "copied content should match original")
}

func TestCopy_FileToDirectory(t *testing.T) {
	lg := got.New(t, "test Copy file to directory")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	dstDir := filepath.Join(tempDir, "dst")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create src dir")
	lg.Require(os.MkdirAll(dstDir, 0755) == nil, "should create dst dir")

	// Create source file
	srcFile := filepath.Join(srcDir, "test.txt")
	content := "Hello, World! This is a test file."
	err = os.WriteFile(srcFile, []byte(content), 0644)
	lg.Require(err == nil, "should create source file")

	// Copy file to directory
	err = iox.Copy(srcFile, dstDir)
	lg.Require(err == nil, "should copy file to directory")

	// Verify file was copied to directory with original name
	expectedDstFile := filepath.Join(dstDir, "test.txt")
	lg.Require(iox.Exists(expectedDstFile), "destination file should exist in directory")

	copiedContent, err := os.ReadFile(expectedDstFile)
	lg.Require(err == nil, "should read copied file")
	lg.Require(string(copiedContent) == content, "copied content should match original")
}

func TestCopy_DirectoryToDirectory(t *testing.T) {
	lg := got.New(t, "test Copy directory to directory")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	dstDir := filepath.Join(tempDir, "dst")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create src dir")

	// Create source directory structure
	subDir := filepath.Join(srcDir, "subdir")
	lg.Require(os.MkdirAll(subDir, 0755) == nil, "should create subdir")

	// Create files in source directory
	files := map[string]string{
		"file1.txt":        "Content of file 1",
		"file2.txt":        "Content of file 2",
		"subdir/file3.txt": "Content of file 3",
	}

	for file, content := range files {
		filePath := filepath.Join(srcDir, file)
		fileDir := filepath.Dir(filePath)
		lg.Require(os.MkdirAll(fileDir, 0755) == nil, "should create file dir")
		err = os.WriteFile(filePath, []byte(content), 0644)
		lg.Require(err == nil, "should create file: %s", file)
	}

	// Copy directory to directory
	err = iox.Copy(srcDir, dstDir)
	lg.Require(err == nil, "should copy directory to directory")

	// Verify all files were copied
	for file, content := range files {
		dstFile := filepath.Join(dstDir, file)
		lg.Require(iox.Exists(dstFile), "destination file should exist: %s", file)

		copiedContent, err := os.ReadFile(dstFile)
		lg.Require(err == nil, "should read copied file: %s", file)
		lg.Require(string(copiedContent) == content, "copied content should match original: %s", file)
	}

	// Verify subdirectory was created
	expectedSubDir := filepath.Join(dstDir, "subdir")
	lg.Require(iox.IsDir(expectedSubDir), "subdirectory should exist")
}

func TestCopy_FilePermissions(t *testing.T) {
	lg := got.New(t, "test Copy preserves file permissions")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	dstDir := filepath.Join(tempDir, "dst")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create src dir")

	// Create source file with specific permissions
	srcFile := filepath.Join(srcDir, "test.txt")
	content := "Test file with permissions"
	err = os.WriteFile(srcFile, []byte(content), 0755) // Readable by all
	lg.Require(err == nil, "should create source file")

	// Copy file
	dstFile := filepath.Join(dstDir, "copied.txt")
	err = iox.Copy(srcFile, dstFile)
	lg.Require(err == nil, "should copy file")

	// Check permissions
	srcInfo, err := os.Stat(srcFile)
	lg.Require(err == nil, "should stat source file")

	dstInfo, err := os.Stat(dstFile)
	lg.Require(err == nil, "should stat destination file")

	lg.Require(srcInfo.Mode() == dstInfo.Mode(), "file permissions should be preserved")
}

func TestCopy_NonExistentSource(t *testing.T) {
	lg := got.New(t, "test Copy with non-existent source")

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	// Try to copy non-existent file
	srcFile := filepath.Join(tempDir, "non_existent.txt")
	dstFile := filepath.Join(tempDir, "destination.txt")

	err = iox.Copy(srcFile, dstFile)
	lg.Require(err != nil, "should return error for non-existent source")
}

func TestCopy_CreateDestinationDirectory(t *testing.T) {
	lg := got.New(t, "test Copy creates destination directory")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create src dir")

	// Create source file
	srcFile := filepath.Join(srcDir, "test.txt")
	content := "Test file"
	err = os.WriteFile(srcFile, []byte(content), 0644)
	lg.Require(err == nil, "should create source file")

	// Copy to non-existent directory path
	dstFile := filepath.Join(tempDir, "new", "deep", "path", "copied.txt")
	err = iox.Copy(srcFile, dstFile)
	lg.Require(err == nil, "should copy file and create directories")

	// Verify file was copied
	lg.Require(iox.Exists(dstFile), "destination file should exist")

	copiedContent, err := os.ReadFile(dstFile)
	lg.Require(err == nil, "should read copied file")
	lg.Require(string(copiedContent) == content, "copied content should match original")
}

func TestFile_Copy(t *testing.T) {
	lg := got.New(t, "test File.Copy method")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	dstDir := filepath.Join(tempDir, "dst")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create src dir")

	// Create source file
	srcFile := filepath.Join(srcDir, "test.txt")
	content := "Hello, World!"
	err = os.WriteFile(srcFile, []byte(content), 0644)
	lg.Require(err == nil, "should create source file")

	// Use File.Copy method
	dstFile := filepath.Join(dstDir, "copied.txt")
	err = iox.File.Copy(srcFile, dstFile)
	lg.Require(err == nil, "should copy file using File.Copy")

	// Verify file was copied
	lg.Require(iox.Exists(dstFile), "destination file should exist")

	copiedContent, err := os.ReadFile(dstFile)
	lg.Require(err == nil, "should read copied file")
	lg.Require(string(copiedContent) == content, "copied content should match original")
}

func TestCopy_ComplexDirectoryStructure(t *testing.T) {
	lg := got.New(t, "test Copy complex directory structure")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	srcDir := filepath.Join(tempDir, "src")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create src dir")

	// Create complex directory structure
	structure := map[string]string{
		"level1/file1.txt":               "Level 1 file 1",
		"level1/file2.txt":               "Level 1 file 2",
		"level1/level2/file3.txt":        "Level 2 file 3",
		"level1/level2/file4.txt":        "Level 2 file 4",
		"level1/level2/level3/file5.txt": "Level 3 file 5",
		"root.txt":                       "Root file",
	}

	for file, content := range structure {
		filePath := filepath.Join(srcDir, file)
		fileDir := filepath.Dir(filePath)
		lg.Require(os.MkdirAll(fileDir, 0755) == nil, "should create directory for: %s", file)
		err = os.WriteFile(filePath, []byte(content), 0644)
		lg.Require(err == nil, "should create file: %s", file)
	}

	// Copy directory
	dstDir := filepath.Join(tempDir, "dst")
	err = iox.Copy(srcDir, dstDir)
	lg.Require(err == nil, "should copy complex directory structure")

	// Verify all files were copied
	for file, content := range structure {
		dstFile := filepath.Join(dstDir, file)
		lg.Require(iox.Exists(dstFile), "destination file should exist: %s", file)

		copiedContent, err := os.ReadFile(dstFile)
		lg.Require(err == nil, "should read copied file: %s", file)
		lg.Require(string(copiedContent) == content, "copied content should match original: %s", file)
	}

	// Verify directory structure
	lg.Require(iox.IsDir(filepath.Join(dstDir, "level1")), "level1 directory should exist")
	lg.Require(iox.IsDir(filepath.Join(dstDir, "level1", "level2")), "level2 directory should exist")
	lg.Require(iox.IsDir(filepath.Join(dstDir, "level1", "level2", "level3")), "level3 directory should exist")
}

func TestCopy_EmptyDirectory(t *testing.T) {
	lg := got.New(t, "test Copy empty directory")

	// Create temporary directories
	tempDir, err := os.MkdirTemp("", "iox_copy_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	// Create empty source directory
	srcDir := filepath.Join(tempDir, "empty_src")
	lg.Require(os.MkdirAll(srcDir, 0755) == nil, "should create empty src dir")

	// Copy empty directory
	dstDir := filepath.Join(tempDir, "empty_dst")
	err = iox.Copy(srcDir, dstDir)
	lg.Require(err == nil, "should copy empty directory")

	// Verify destination directory exists
	lg.Require(iox.IsDir(dstDir), "destination directory should exist")

	// Verify directory is empty
	isEmpty, err := iox.Dir.IsEmpty(dstDir)
	lg.Require(err == nil, "should check if directory is empty")
	lg.Require(isEmpty, "destination directory should be empty")
}
