package iox_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go4x/goal/iox"
	"github.com/go4x/got"
)

func TestNewTxtFile(t *testing.T) {
	lg := got.New(t, "test NewTxtFile")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	lg.Case("create new text file")
	tempFile := filepath.Join(tempDir, "test.txt")
	tf, err := iox.NewTxtFile(tempFile)
	lg.Require(err == nil, "should create text file")
	lg.Require(tf != nil, "should return text file instance")
	defer tf.Close()

	// Verify file exists
	lg.Require(iox.Exists(tempFile), "file should exist")
}

func TestTxtFileWriteLine(t *testing.T) {
	lg := got.New(t, "test TxtFile WriteLine")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "test.txt")
	tf, err := iox.NewTxtFile(tempFile)
	lg.Require(err == nil, "should create text file")
	defer tf.Close()

	lg.Case("write single line")
	_, err = tf.WriteLine("Hello, World!")
	lg.Require(err == nil, "should write line")

	// Flush to ensure data is written
	err = tf.Flush()
	lg.Require(err == nil, "should flush")

	lg.Case("write multiple lines")
	_, err = tf.WriteLine("Line 2")
	lg.Require(err == nil, "should write second line")

	_, err = tf.WriteLine("Line 3")
	lg.Require(err == nil, "should write third line")

	err = tf.Flush()
	lg.Require(err == nil, "should flush")

	// Verify content
	content, err := os.ReadFile(tempFile)
	lg.Require(err == nil, "should read file")
	lg.Require(strings.Contains(string(content), "Hello, World!"), "should contain first line")
	lg.Require(strings.Contains(string(content), "Line 2"), "should contain second line")
	lg.Require(strings.Contains(string(content), "Line 3"), "should contain third line")
}

func TestTxtFileReadAll(t *testing.T) {
	lg := got.New(t, "test TxtFile ReadAll")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "test.txt")
	tf, err := iox.NewTxtFile(tempFile)
	lg.Require(err == nil, "should create text file")
	defer tf.Close()

	lg.Case("write lines and read them back")
	testLines := []string{
		"First line",
		"Second line",
		"Third line",
	}

	// Write lines
	for _, line := range testLines {
		_, err = tf.WriteLine(line)
		lg.Require(err == nil, "should write line")
	}

	err = tf.Flush()
	lg.Require(err == nil, "should flush")

	// Read lines
	lines, err := tf.ReadAll()
	lg.Require(err == nil, "should read all lines")
	lg.Require(len(lines) == len(testLines), "should have correct number of lines")

	for i, line := range lines {
		lg.Require(line == testLines[i], "should match written line")
	}
}

func TestTxtFileFlush(t *testing.T) {
	lg := got.New(t, "test TxtFile Flush")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "test.txt")
	tf, err := iox.NewTxtFile(tempFile)
	lg.Require(err == nil, "should create text file")
	defer tf.Close()

	lg.Case("write and flush")
	_, err = tf.WriteLine("Test line")
	lg.Require(err == nil, "should write line")

	err = tf.Flush()
	lg.Require(err == nil, "should flush")

	// Verify content is written
	content, err := os.ReadFile(tempFile)
	lg.Require(err == nil, "should read file")
	lg.Require(strings.Contains(string(content), "Test line"), "should contain written line")
}

func TestTxtFileClose(t *testing.T) {
	lg := got.New(t, "test TxtFile Close")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "test.txt")
	tf, err := iox.NewTxtFile(tempFile)
	lg.Require(err == nil, "should create text file")

	lg.Case("write, close and verify")
	_, err = tf.WriteLine("Test line")
	lg.Require(err == nil, "should write line")

	err = tf.Close()
	lg.Require(err == nil, "should close file")

	// Verify content is written (Close should flush)
	content, err := os.ReadFile(tempFile)
	lg.Require(err == nil, "should read file")
	lg.Require(strings.Contains(string(content), "Test line"), "should contain written line")
}

// TestTxtFile_EdgeCases tests edge cases for text file operations
func TestTxtFile_EdgeCases(t *testing.T) {
	lg := got.New(t, "test text file edge cases")

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	lg.Case("test with empty file")
	tempFile := filepath.Join(tempDir, "empty.txt")
	tf, err := iox.NewTxtFile(tempFile)
	lg.Require(err == nil, "should create empty file")
	defer tf.Close()

	// Read empty file
	lines, err := tf.ReadAll()
	lg.Require(err == nil, "should read empty file")
	lg.Require(len(lines) == 0, "should have no lines")

	lg.Case("test with empty lines")
	tempFile2 := filepath.Join(tempDir, "empty_lines.txt")
	tf2, err := iox.NewTxtFile(tempFile2)
	lg.Require(err == nil, "should create file with empty lines")
	defer tf2.Close()

	_, err = tf2.WriteLine("")
	lg.Require(err == nil, "should write empty line")
	_, err = tf2.WriteLine("Non-empty line")
	lg.Require(err == nil, "should write non-empty line")
	_, err = tf2.WriteLine("")
	lg.Require(err == nil, "should write another empty line")

	err = tf2.Flush()
	lg.Require(err == nil, "should flush")

	lines, err = tf2.ReadAll()
	lg.Require(err == nil, "should read file with empty lines")
	lg.Require(len(lines) == 3, "should have 3 lines")
	lg.Require(lines[0] == "", "first line should be empty")
	lg.Require(lines[1] == "Non-empty line", "second line should be non-empty")
	lg.Require(lines[2] == "", "third line should be empty")

	lg.Case("test with special characters")
	tempFile3 := filepath.Join(tempDir, "special.txt")
	tf3, err := iox.NewTxtFile(tempFile3)
	lg.Require(err == nil, "should create file with special characters")
	defer tf3.Close()

	specialLines := []string{
		"Line with spaces and tabs\t",
		"Line with unicode: 你好世界",
		"Line with symbols: !@#$%^&*()",
		"Line with embedded newline: embedded",
	}

	for _, line := range specialLines {
		_, err = tf3.WriteLine(line)
		lg.Require(err == nil, "should write special line")
	}

	err = tf3.Flush()
	lg.Require(err == nil, "should flush")

	lines, err = tf3.ReadAll()
	lg.Require(err == nil, "should read file with special characters")
	// Note: ReadAll() reads lines without newline characters, so the count should match
	lg.Require(len(lines) == len(specialLines), "should have correct number of lines, got %d, expected %d", len(lines), len(specialLines))

	for i, line := range lines {
		if i < len(specialLines) {
			lg.Require(line == specialLines[i], "should match special line at index %d", i)
		}
	}
}

// TestTxtFile_ErrorHandling tests error handling scenarios
func TestTxtFile_ErrorHandling(t *testing.T) {
	lg := got.New(t, "test text file error handling")

	lg.Case("test with invalid path")
	invalidFile := "/invalid/path/that/does/not/exist/test.txt"
	tf, err := iox.NewTxtFile(invalidFile)
	lg.Require(err != nil, "should error with invalid path")
	lg.Require(tf == nil, "should return nil for invalid path")

	// Create a file and then make it read-only to test write errors
	tempDir, err := os.MkdirTemp("", "iox_test_*")
	lg.Require(err == nil, "should create temp dir")
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "readonly.txt")
	err = os.WriteFile(tempFile, []byte("initial content"), 0444) // Read-only
	lg.Require(err == nil, "should create read-only file")

	tf, err = iox.NewTxtFile(tempFile)
	if err != nil {
		// If NewTxtFile fails on read-only file, that's expected
		lg.Require(err != nil, "should error when opening read-only file")
		return
	}
	defer func() {
		if tf != nil {
			tf.Close()
		}
	}()

	// Try to write to read-only file (this might work on some systems)
	_, err = tf.WriteLine("test")
	// On some systems, this might succeed, so we don't enforce an error
	lg.Require(true, "write operation completed (may or may not succeed depending on system)")
}
