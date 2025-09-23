package iox

import (
	"fmt"
	"os"
	"strings"
)

// ExampleWalkDir demonstrates basic usage with simple filters
func ExampleWalkDir() {
	// Simple usage: find all .go files that are not hidden
	files, err := WalkDir("/path/to/project",
		FilterByExtension(".go"),
		FilterHidden,
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}
	// Output would show all .go files that don't start with '.'
}

// ExampleWalkDirWithFilters demonstrates complex filter combinations
func ExampleWalkDirWithFilters() {
	// Complex scenario: Find files that match either:
	// 1. Go files OR Markdown files (but not hidden)
	// 2. Large text files (>1KB)

	// Group 1: (.go OR .md) AND not hidden
	group1 := NewFilterGroup(FilterAnd,
		func(entry os.DirEntry, path string) bool {
			// (.go OR .md) AND not hidden
			orGroup := NewFilterGroup(FilterOr,
				FilterByExtension(".go"),
				FilterByExtension(".md"),
			)
			return orGroup.Apply(entry, path) && FilterHidden(entry, path)
		},
	)

	// Group 2: Large text files
	group2 := NewFilterGroup(FilterAnd,
		FilterByExtension(".txt"),
		FilterBySize(1024, 1024*1024), // 1KB to 1MB
	)

	files, err := WalkDirWithFilters("/path/to/project", group1, group2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}
}

// ExampleWalkDir_projectStructure demonstrates finding specific project files
func ExampleWalkDir_projectStructure() {
	// Find all source code files in a project
	sourceExtensions := []string{".go", ".java", ".py", ".js", ".ts", ".cpp", ".c"}

	filterGroup := NewFilterGroup(FilterAnd,
		func(entry os.DirEntry, path string) bool {
			// Source file with supported extension
			orGroup := NewFilterGroup(FilterOr)
			for _, ext := range sourceExtensions {
				orGroup.Filters = append(orGroup.Filters, FilterByExtension(ext))
			}
			return orGroup.Apply(entry, path)
		},
		FilterHidden, // Not hidden
		func(entry os.DirEntry, path string) bool {
			// Not in vendor, node_modules, or .git directories
			excludedPaths := []string{"vendor/", "node_modules/", ".git/"}
			for _, excluded := range excludedPaths {
				if strings.Contains(path, excluded) {
					return false
				}
			}
			return true
		},
	)

	files, err := WalkDirWithFilters("/path/to/project", filterGroup)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d source files\n", len(files))
	for _, file := range files {
		fmt.Println(file)
	}
}

// ExampleWalkDir_documentation demonstrates finding documentation files
func ExampleWalkDir_documentation() {
	// Find documentation files: README, docs, or markdown files
	docGroup := NewFilterGroup(FilterOr,
		FilterByName("README"),
		FilterByExtension(".md"),
		FilterByPathPattern("docs/"),
		FilterByExtension(".rst"),
		FilterByExtension(".txt"),
	)

	files, err := WalkDirWithFilters("/path/to/project", docGroup)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d documentation files\n", len(files))
	for _, file := range files {
		fmt.Println(file)
	}
}

// ExampleWalkDir_cleanup demonstrates finding files for cleanup
func ExampleWalkDir_cleanup() {
	// Find temporary files, logs, and cache files for cleanup
	cleanupGroup := NewFilterGroup(FilterOr,
		// Temporary files
		FilterByExtension(".tmp"),
		FilterByExtension(".temp"),
		FilterByName("temp"),

		// Log files
		FilterByExtension(".log"),
		FilterByName("log"),

		// Cache files
		FilterByExtension(".cache"),
		FilterByPathPattern("cache/"),

		// Build artifacts
		FilterByExtension(".o"),
		FilterByExtension(".obj"),
		FilterByExtension(".exe"),
		FilterByPathPattern("build/"),
		FilterByPathPattern("dist/"),
		FilterByPathPattern("target/"),
	)

	files, err := WalkDirWithFilters("/path/to/project", cleanupGroup)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d files that can be cleaned up\n", len(files))
	for _, file := range files {
		fmt.Println(file)
	}
}

// ExampleWalkDir_sizeAnalysis demonstrates analyzing file sizes
func ExampleWalkDir_sizeAnalysis() {
	// Analyze files by size categories
	sizeCategories := map[string]FilterGroup{
		"Small files (<1KB)": NewFilterGroup(FilterAnd,
			FilterBySize(0, 1024),
			FilterFilesOnly,
		),
		"Medium files (1KB-1MB)": NewFilterGroup(FilterAnd,
			FilterBySize(1024, 1024*1024),
			FilterFilesOnly,
		),
		"Large files (>1MB)": NewFilterGroup(FilterAnd,
			FilterBySize(1024*1024, 1024*1024*1024*10), // Up to 10GB
			FilterFilesOnly,
		),
	}

	for category, filterGroup := range sizeCategories {
		files, err := WalkDirWithFilters("/path/to/project", filterGroup)
		if err != nil {
			fmt.Printf("Error analyzing %s: %v\n", category, err)
			continue
		}
		fmt.Printf("%s: %d files\n", category, len(files))
	}
}

// ExampleWalkDir_directoryStructure demonstrates finding directory structure
func ExampleWalkDir_directoryStructure() {
	// Find all directories that contain source code
	sourceDirs := NewFilterGroup(FilterAnd,
		FilterDirectoriesOnly,
		FilterHidden, // Not hidden directories like .git
		func(entry os.DirEntry, path string) bool {
			// Directory contains source files
			// This is a simplified check - in practice you might want to
			// recursively check if the directory contains source files
			return true // For this example, include all non-hidden directories
		},
	)

	dirs, err := WalkDirWithFilters("/path/to/project", sourceDirs)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d directories\n", len(dirs))
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

// ExampleWalkDir_customFilter demonstrates creating custom filters
func ExampleWalkDir_customFilter() {
	// Custom filter: files modified in the last 7 days
	recentFiles := func(entry os.DirEntry, path string) bool {
		if entry.IsDir() {
			return true // Always include directories for traversal
		}

		info, err := entry.Info()
		if err != nil {
			return false
		}

		// Check if file was modified in the last 7 days
		// This is a simplified example - you'd use time.Now() and proper date comparison
		return info.ModTime().Unix() > 0 // Placeholder logic
	}

	// Combine with other filters
	recentGoFiles := NewFilterGroup(FilterAnd,
		FilterByExtension(".go"),
		recentFiles,
		FilterHidden,
	)

	files, err := WalkDirWithFilters("/path/to/project", recentGoFiles)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d recent Go files\n", len(files))
	for _, file := range files {
		fmt.Println(file)
	}
}
