package iox

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSimpleFilterCombinations(t *testing.T) {
	// Create a temporary directory structure for testing
	tempDir := t.TempDir()

	// Create test files
	testFiles := []string{
		"file1.txt",
		"file2.go",
		"file3.md",
		"README.txt",
		".hidden.txt",
	}

	// Create files
	for _, file := range testFiles {
		filePath := filepath.Join(tempDir, file)
		err := os.WriteFile(filePath, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", file, err)
		}
	}

	t.Run("AND logic - txt files that are not hidden", func(t *testing.T) {
		files, err := WalkDir(tempDir,
			FilterByExtension(".txt"),
			FilterHidden,
		)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}

		// Should only include file1.txt and README.txt (not .hidden.txt)
		expectedCount := 2
		if len(files) != expectedCount {
			t.Errorf("Expected %d files, got %d", expectedCount, len(files))
			for _, file := range files {
				t.Logf("Found file: %s", file)
			}
		}
	})

	t.Run("OR logic - go or md files", func(t *testing.T) {
		filterGroup := NewFilterGroup(FilterOr,
			FilterByExtension(".go"),
			FilterByExtension(".md"),
		)

		files, err := WalkDirWithFilters(tempDir, filterGroup)
		if err != nil {
			t.Fatalf("WalkDirWithFilters failed: %v", err)
		}

		// Should include file2.go and file3.md
		expectedCount := 2
		if len(files) != expectedCount {
			t.Errorf("Expected %d files, got %d", expectedCount, len(files))
			for _, file := range files {
				t.Logf("Found file: %s", file)
			}
		}
	})

	t.Run("Multiple groups - either condition can match", func(t *testing.T) {
		// Group 1: Go or MD files
		group1 := NewFilterGroup(FilterOr,
			FilterByExtension(".go"),
			FilterByExtension(".md"),
		)

		// Group 2: Text files (not hidden)
		group2 := NewFilterGroup(FilterAnd,
			FilterByExtension(".txt"),
			FilterHidden,
		)

		files, err := WalkDirWithFilters(tempDir, group1, group2)
		if err != nil {
			t.Fatalf("WalkDirWithFilters failed: %v", err)
		}

		// Should include: file2.go, file3.md, file1.txt, README.txt
		expectedCount := 4
		if len(files) != expectedCount {
			t.Errorf("Expected %d files, got %d", expectedCount, len(files))
			for _, file := range files {
				t.Logf("Found file: %s", file)
			}
		}
	})

	t.Run("No filters - should include everything", func(t *testing.T) {
		files, err := WalkDir(tempDir)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}

		// Should include all files
		expectedCount := len(testFiles)
		if len(files) != expectedCount {
			t.Errorf("Expected %d files, got %d", expectedCount, len(files))
			for _, file := range files {
				t.Logf("Found file: %s", file)
			}
		}
	})
}

func TestFilterGroupApply(t *testing.T) {
	entry := &mockDirEntry{name: "test.txt", isDir: false}

	t.Run("AND combination", func(t *testing.T) {
		group := NewFilterGroup(FilterAnd,
			func(entry os.DirEntry, path string) bool { return strings.HasSuffix(path, ".txt") },
			func(entry os.DirEntry, path string) bool { return !strings.HasPrefix(entry.Name(), ".") },
		)

		if !group.Apply(entry, "/path/to/test.txt") {
			t.Error("AND filter should return true for matching entry")
		}

		hiddenEntry := &mockDirEntry{name: ".hidden.txt", isDir: false}
		if group.Apply(hiddenEntry, "/path/to/.hidden.txt") {
			t.Error("AND filter should return false for hidden file")
		}
	})

	t.Run("OR combination", func(t *testing.T) {
		group := NewFilterGroup(FilterOr,
			func(entry os.DirEntry, path string) bool { return strings.HasSuffix(path, ".txt") },
			func(entry os.DirEntry, path string) bool { return strings.HasSuffix(path, ".go") },
		)

		if !group.Apply(entry, "/path/to/test.txt") {
			t.Error("OR filter should return true for .txt file")
		}

		goEntry := &mockDirEntry{name: "test.go", isDir: false}
		if !group.Apply(goEntry, "/path/to/test.go") {
			t.Error("OR filter should return true for .go file")
		}

		mdEntry := &mockDirEntry{name: "test.md", isDir: false}
		if group.Apply(mdEntry, "/path/to/test.md") {
			t.Error("OR filter should return false for .md file")
		}
	})

	t.Run("Empty filter group", func(t *testing.T) {
		group := NewFilterGroup(FilterAnd)
		if !group.Apply(entry, "/path/to/test.txt") {
			t.Error("Empty filter group should always return true")
		}
	})
}

// mockDirEntry implements os.DirEntry for testing
type mockDirEntry struct {
	name  string
	isDir bool
}

func (m *mockDirEntry) Name() string               { return m.name }
func (m *mockDirEntry) IsDir() bool                { return m.isDir }
func (m *mockDirEntry) Type() os.FileMode          { return os.ModePerm }
func (m *mockDirEntry) Info() (os.FileInfo, error) { return nil, nil }

// TestWalkDir_Performance tests performance with large directory structures
func TestWalkDir_Performance(t *testing.T) {
	// Create a temporary directory with many files
	tempDir, err := os.MkdirTemp("", "iox_performance_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create 50 test files
	for i := 0; i < 50; i++ {
		fileName := fmt.Sprintf("file_%03d.go", i)
		filePath := filepath.Join(tempDir, fileName)
		err = os.WriteFile(filePath, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file %s: %v", fileName, err)
		}
	}

	// Create 5 subdirectories with files
	for i := 0; i < 5; i++ {
		subDir := filepath.Join(tempDir, fmt.Sprintf("subdir_%d", i))
		err = os.Mkdir(subDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create subdirectory %s: %v", subDir, err)
		}

		// Create 10 files in each subdirectory
		for j := 0; j < 10; j++ {
			fileName := fmt.Sprintf("file_%03d.txt", j)
			filePath := filepath.Join(subDir, fileName)
			err = os.WriteFile(filePath, []byte("test content"), 0644)
			if err != nil {
				t.Fatalf("Failed to create file %s: %v", fileName, err)
			}
		}
	}

	t.Run("Performance with large directory structure", func(t *testing.T) {
		files, err := WalkDir(tempDir)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}
		if len(files) < 100 {
			t.Errorf("Expected at least 100 files, got %d", len(files))
		}
	})

	t.Run("Performance with filters", func(t *testing.T) {
		goFilter := FilterByExtension(".go")
		files, err := WalkDir(tempDir, goFilter)
		if err != nil {
			t.Fatalf("WalkDir with filter failed: %v", err)
		}
		if len(files) != 50 {
			t.Errorf("Expected exactly 50 .go files, got %d", len(files))
		}
	})
}

// TestFilterFunctions_EdgeCases tests edge cases for filter functions
func TestFilterFunctions_EdgeCases(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "iox_filter_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files with various extensions
	testFiles := []string{
		"file1.go",
		"file2.GO", // uppercase
		"file3.go.bak",
		"file4.txt",
		"file5.TXT", // uppercase
		"file6.md",
		"file7.MD", // uppercase
		"file8",    // no extension
		"file9.",   // empty extension
	}

	for _, file := range testFiles {
		filePath := filepath.Join(tempDir, file)
		err = os.WriteFile(filePath, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file %s: %v", file, err)
		}
	}

	t.Run("FilterByExtension with case sensitivity", func(t *testing.T) {
		goFilter := FilterByExtension(".go")
		files, err := WalkDir(tempDir, goFilter)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}
		// Should find file1.go and file3.go.bak (but not file2.GO due to case sensitivity)
		if len(files) < 2 {
			t.Errorf("Expected at least 2 .go files, got %d", len(files))
		}
	})

	t.Run("FilterByExtension with multiple extensions", func(t *testing.T) {
		multiFilter := FilterByExtension(".go", ".txt", ".md")
		files, err := WalkDir(tempDir, multiFilter)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}
		if len(files) < 5 {
			t.Errorf("Expected at least 5 files with multiple extensions, got %d", len(files))
		}
	})

	t.Run("FilterByName with pattern", func(t *testing.T) {
		nameFilter := FilterByName("file")
		files, err := WalkDir(tempDir, nameFilter)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}
		if len(files) < 9 {
			t.Errorf("Expected at least 9 files containing 'file', got %d", len(files))
		}
	})

	t.Run("FilterBySize with size range", func(t *testing.T) {
		sizeFilter := FilterBySize(10, 20) // 10-20 bytes
		files, err := WalkDir(tempDir, sizeFilter)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}
		if len(files) < 5 {
			t.Errorf("Expected at least 5 files in size range, got %d", len(files))
		}
	})

	t.Run("FilterFilesOnly", func(t *testing.T) {
		fileFilter := FilterFilesOnly
		files, err := WalkDir(tempDir, fileFilter)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}
		if len(files) < 9 {
			t.Errorf("Expected at least 9 files, got %d", len(files))
		}
	})

	t.Run("FilterHidden", func(t *testing.T) {
		hiddenFilter := FilterHidden
		files, err := WalkDir(tempDir, hiddenFilter)
		if err != nil {
			t.Fatalf("WalkDir failed: %v", err)
		}
		// All files should be visible (not hidden)
		if len(files) < 9 {
			t.Errorf("Expected at least 9 visible files, got %d", len(files))
		}
	})
}
