package iox_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/go4x/goal/iox"
)

// BenchmarkExists benchmarks the Exists function
func BenchmarkExists(b *testing.B) {
	// Create a temporary file for benchmarking
	tempFile, err := os.CreateTemp("", "benchmark_test_*")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iox.Exists(tempFile.Name())
	}
}

// BenchmarkIsDir benchmarks the IsDir function
func BenchmarkIsDir(b *testing.B) {
	// Create a temporary directory for benchmarking
	tempDir, err := os.MkdirTemp("", "benchmark_test_*")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iox.IsDir(tempDir)
	}
}

// BenchmarkIsRegularFile benchmarks the IsRegularFile function
func BenchmarkIsRegularFile(b *testing.B) {
	// Create a temporary file for benchmarking
	tempFile, err := os.CreateTemp("", "benchmark_test_*")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iox.IsRegularFile(tempFile.Name())
	}
}

// BenchmarkWalkDir benchmarks the WalkDir function
func BenchmarkWalkDir(b *testing.B) {
	// Create a temporary directory structure for benchmarking
	tempDir, err := os.MkdirTemp("", "benchmark_walker_test_*")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create 100 test files in the directory
	for i := 0; i < 100; i++ {
		fileName := filepath.Join(tempDir, fmt.Sprintf("file_%03d.txt", i))
		err = os.WriteFile(fileName, []byte("test content"), 0644)
		if err != nil {
			b.Fatalf("Failed to create file: %v", err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := iox.WalkDir(tempDir)
		if err != nil {
			b.Fatalf("WalkDir failed: %v", err)
		}
	}
}

// BenchmarkWalkDirWithFilter benchmarks the WalkDir function with a filter
func BenchmarkWalkDirWithFilter(b *testing.B) {
	// Create a temporary directory structure for benchmarking
	tempDir, err := os.MkdirTemp("", "benchmark_walker_filter_test_*")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create 50 .go files and 50 .txt files
	for i := 0; i < 50; i++ {
		goFile := filepath.Join(tempDir, fmt.Sprintf("file_%03d.go", i))
		txtFile := filepath.Join(tempDir, fmt.Sprintf("file_%03d.txt", i))

		err = os.WriteFile(goFile, []byte("package main"), 0644)
		if err != nil {
			b.Fatalf("Failed to create .go file: %v", err)
		}

		err = os.WriteFile(txtFile, []byte("test content"), 0644)
		if err != nil {
			b.Fatalf("Failed to create .txt file: %v", err)
		}
	}

	// Create a filter for .go files
	goFilter := iox.FilterByExtension(".go")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := iox.WalkDir(tempDir, goFilter)
		if err != nil {
			b.Fatalf("WalkDir with filter failed: %v", err)
		}
	}
}

// BenchmarkWalkDirWithFilters benchmarks the WalkDirWithFilters function
func BenchmarkWalkDirWithFilters(b *testing.B) {
	// Create a temporary directory structure for benchmarking
	tempDir, err := os.MkdirTemp("", "benchmark_walker_filters_test_*")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create 50 .go files and 50 .txt files
	for i := 0; i < 50; i++ {
		goFile := filepath.Join(tempDir, fmt.Sprintf("file_%03d.go", i))
		txtFile := filepath.Join(tempDir, fmt.Sprintf("file_%03d.txt", i))

		err = os.WriteFile(goFile, []byte("package main"), 0644)
		if err != nil {
			b.Fatalf("Failed to create .go file: %v", err)
		}

		err = os.WriteFile(txtFile, []byte("test content"), 0644)
		if err != nil {
			b.Fatalf("Failed to create .txt file: %v", err)
		}
	}

	// Create filter groups
	goGroup := iox.NewFilterGroup(iox.FilterAnd, iox.FilterByExtension(".go"))
	txtGroup := iox.NewFilterGroup(iox.FilterAnd, iox.FilterByExtension(".txt"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := iox.WalkDirWithFilters(tempDir, goGroup, txtGroup)
		if err != nil {
			b.Fatalf("WalkDirWithFilters failed: %v", err)
		}
	}
}

// BenchmarkTxtFile_WriteLine benchmarks the TxtFile WriteLine method
func BenchmarkTxtFile_WriteLine(b *testing.B) {
	// Create a temporary file for benchmarking
	tempFile, err := os.CreateTemp("", "benchmark_txt_*")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	tf, err := iox.NewTxtFile(tempFile.Name())
	if err != nil {
		b.Fatalf("Failed to create TxtFile: %v", err)
	}
	defer tf.Close()

	testLine := "This is a test line for benchmarking purposes"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tf.WriteLine(testLine)
		if err != nil {
			b.Fatalf("WriteLine failed: %v", err)
		}
	}
	tf.Flush()
}

// BenchmarkTxtFile_ReadAll benchmarks the TxtFile ReadAll method
func BenchmarkTxtFile_ReadAll(b *testing.B) {
	// Create a temporary file with content for benchmarking
	tempFile, err := os.CreateTemp("", "benchmark_txt_read_*")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write 1000 lines to the file
	for i := 0; i < 1000; i++ {
		_, err = tempFile.WriteString(fmt.Sprintf("This is line %d for benchmarking\n", i))
		if err != nil {
			b.Fatalf("Failed to write to temp file: %v", err)
		}
	}
	tempFile.Close()

	tf, err := iox.NewTxtFile(tempFile.Name())
	if err != nil {
		b.Fatalf("Failed to create TxtFile: %v", err)
	}
	defer tf.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tf.ReadAll()
		if err != nil {
			b.Fatalf("ReadAll failed: %v", err)
		}
	}
}

// BenchmarkDir_Create benchmarks the Dir Create method
func BenchmarkDir_Create(b *testing.B) {
	baseDir, err := os.MkdirTemp("", "benchmark_dir_create_*")
	if err != nil {
		b.Fatalf("Failed to create base temp dir: %v", err)
	}
	defer os.RemoveAll(baseDir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dirPath := filepath.Join(baseDir, fmt.Sprintf("test_dir_%d", i))
		err := iox.Dir.Create(dirPath)
		if err != nil {
			b.Fatalf("Dir.Create failed: %v", err)
		}
	}
}

// BenchmarkPath_ExecPath benchmarks the Path ExecPath method
func BenchmarkPath_ExecPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = iox.Path.ExecPath()
	}
}

// BenchmarkPath_CurrentPath benchmarks the Path CurrentPath method
func BenchmarkPath_CurrentPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = iox.Path.CurrentPath()
	}
}

// BenchmarkPath_ProjectPath benchmarks the Path ProjectPath method
func BenchmarkPath_ProjectPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = iox.Path.ProjectPath()
	}
}

// BenchmarkCopy_File benchmarks the Copy function for file operations
func BenchmarkCopy_File(b *testing.B) {
	// Create a temporary file for benchmarking
	tempFile, err := os.CreateTemp("", "benchmark_copy_*")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some content to the file
	content := make([]byte, 1024) // 1KB file
	for i := range content {
		content[i] = byte(i % 256)
	}
	_, err = tempFile.Write(content)
	if err != nil {
		b.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dstFile := fmt.Sprintf("/tmp/benchmark_copy_dst_%d", i)
		err := iox.Copy(tempFile.Name(), dstFile)
		if err != nil {
			b.Fatalf("Copy failed: %v", err)
		}
		os.Remove(dstFile) // Clean up
	}
}

// BenchmarkCopy_Directory benchmarks the Copy function for directory operations
func BenchmarkCopy_Directory(b *testing.B) {
	// Create a temporary directory with files for benchmarking
	tempDir, err := os.MkdirTemp("", "benchmark_copy_dir_*")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create 10 files in the directory
	for i := 0; i < 10; i++ {
		fileName := filepath.Join(tempDir, fmt.Sprintf("file_%d.txt", i))
		content := make([]byte, 100) // 100 bytes per file
		for j := range content {
			content[j] = byte(j % 256)
		}
		err = os.WriteFile(fileName, content, 0644)
		if err != nil {
			b.Fatalf("Failed to create file: %v", err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dstDir := fmt.Sprintf("/tmp/benchmark_copy_dst_dir_%d", i)
		err := iox.Copy(tempDir, dstDir)
		if err != nil {
			b.Fatalf("Copy failed: %v", err)
		}
		os.RemoveAll(dstDir) // Clean up
	}
}

// BenchmarkFile_Copy benchmarks the File.Copy method
func BenchmarkFile_Copy(b *testing.B) {
	// Create a temporary file for benchmarking
	tempFile, err := os.CreateTemp("", "benchmark_file_copy_*")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some content to the file
	content := make([]byte, 512) // 512 bytes file
	for i := range content {
		content[i] = byte(i % 256)
	}
	_, err = tempFile.Write(content)
	if err != nil {
		b.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dstFile := fmt.Sprintf("/tmp/benchmark_file_copy_dst_%d", i)
		err := iox.File.Copy(tempFile.Name(), dstFile)
		if err != nil {
			b.Fatalf("File.Copy failed: %v", err)
		}
		os.Remove(dstFile) // Clean up
	}
}
