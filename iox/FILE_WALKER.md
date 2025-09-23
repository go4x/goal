# File Walker - Flexible Directory Filtering System

The `file_walker.go` provides a powerful and flexible system for walking through directories with complex filtering capabilities. It supports both simple filter combinations and advanced multi-group filtering with AND/OR logic.

## Table of Contents

- [Overview](#overview)
- [Basic Usage](#basic-usage)
- [Filter Types](#filter-types)
- [Filter Combinations](#filter-combinations)
- [Advanced Usage](#advanced-usage)
- [Predefined Filters](#predefined-filters)
- [Examples](#examples)
- [API Reference](#api-reference)

## Overview

The file walker system consists of:

1. **WalkFilter**: A function type that determines whether a directory entry should be included
2. **FilterCombiner**: Defines how multiple filters are combined (AND/OR)
3. **FilterGroup**: A collection of filters with a specific combination logic
4. **Two main functions**: `WalkDir` (simple) and `WalkDirWithFilters` (advanced)

### Key Features

- **Flexible Filtering**: Support for AND/OR combinations within and between filter groups
- **Multiple Filter Groups**: Each group can have different combination logic
- **Predefined Filters**: Common filters for file extensions, names, sizes, etc.
- **Custom Filters**: Easy to create custom filter functions
- **Backward Compatibility**: Simple `WalkDir` function for basic use cases

## Basic Usage

### Simple Filtering with WalkDir

```go
// Find all .go files that are not hidden
files, err := iox.WalkDir("/path/to/project",
    iox.FilterByExtension(".go"),
    iox.FilterHidden,
)
```

### Advanced Filtering with WalkDirWithFilters

```go
// Complex filtering: (.go OR .md) AND not hidden
group := iox.NewFilterGroup(iox.FilterAnd,
    iox.NewFilterGroup(iox.FilterOr,
        iox.FilterByExtension(".go"),
        iox.FilterByExtension(".md"),
    ).Apply,
    iox.FilterHidden,
)

files, err := iox.WalkDirWithFilters("/path/to/project", group)
```

## Filter Types

### FilterCombiner

```go
type FilterCombiner int

const (
    FilterAnd FilterCombiner = iota  // All filters must return true
    FilterOr                         // At least one filter must return true
)
```

### FilterGroup

```go
type FilterGroup struct {
    Filters  []WalkFilter
    Combiner FilterCombiner
}
```

### WalkFilter

```go
type WalkFilter func(entry os.DirEntry, path string) bool
```

## Filter Combinations

### Within a FilterGroup

Filters within a group are combined according to the group's `Combiner`:

- **FilterAnd**: All filters must return `true`
- **FilterOr**: At least one filter must return `true`

### Between FilterGroups

Multiple `FilterGroup`s are combined with OR logic - if any group passes, the entry is included.

```go
// Group 1: Go files AND not hidden
group1 := NewFilterGroup(FilterAnd,
    FilterByExtension(".go"),
    FilterHidden,
)

// Group 2: Markdown files AND not hidden  
group2 := NewFilterGroup(FilterAnd,
    FilterByExtension(".md"),
    FilterHidden,
)

// Either group can match (OR between groups)
files, err := WalkDirWithFilters("/path", group1, group2)
```

### Complex Combinations

```go
// Group 1: (.go OR .java) AND not hidden AND not in vendor
group1 := NewFilterGroup(FilterAnd,
    NewFilterGroup(FilterOr,
        FilterByExtension(".go"),
        FilterByExtension(".java"),
    ).Apply,
    FilterHidden,
    func(entry os.DirEntry, path string) bool {
        return !strings.Contains(path, "vendor/")
    },
)

// Group 2: Large text files
group2 := NewFilterGroup(FilterAnd,
    FilterByExtension(".txt"),
    FilterBySize(1024, 1024*1024), // 1KB to 1MB
)

files, err := WalkDirWithFilters("/path", group1, group2)
```

## Advanced Usage

### Creating Custom Filters

```go
// Custom filter for files modified in the last 7 days
recentFiles := func(entry os.DirEntry, path string) bool {
    if entry.IsDir() {
        return true // Always include directories for traversal
    }
    
    info, err := entry.Info()
    if err != nil {
        return false
    }
    
    // Check if modified in last 7 days
    return time.Since(info.ModTime()) < 7*24*time.Hour
}

// Use custom filter
group := NewFilterGroup(FilterAnd,
    FilterByExtension(".go"),
    recentFiles,
    FilterHidden,
)
```

### Nested Filter Groups

```go
// Complex nested logic: ((.go OR .java) AND not hidden) OR (large .txt files)
nestedGroup := NewFilterGroup(FilterOr,
    // Sub-group 1: Source files
    NewFilterGroup(FilterAnd,
        NewFilterGroup(FilterOr,
            FilterByExtension(".go"),
            FilterByExtension(".java"),
        ).Apply,
        FilterHidden,
    ).Apply,
    
    // Sub-group 2: Large text files
    NewFilterGroup(FilterAnd,
        FilterByExtension(".txt"),
        FilterBySize(1024, 1024*1024),
    ).Apply,
)
```

## Predefined Filters

### File Type Filters

```go
// Extension-based filtering
FilterByExtension(".go", ".java", ".py")

// Name-based filtering
FilterByName("README")
FilterByPathPattern("docs/")

// Type-based filtering
FilterDirectoriesOnly()  // Only directories
FilterFilesOnly()        // Only files
FilterHidden()           // Exclude hidden files (starting with '.')
```

### Size-based Filtering

```go
// Files between 1KB and 1MB
FilterBySize(1024, 1024*1024)

// Files larger than 100MB
FilterBySize(100*1024*1024, 1024*1024*1024*10) // Up to 10GB
```

## Examples

### Project Source Code Analysis

```go
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
        // Not in excluded directories
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
```

### Documentation Files

```go
// Find documentation files
docGroup := NewFilterGroup(FilterOr,
    FilterByName("README"),
    FilterByExtension(".md", ".rst"),
    FilterByPathPattern("docs/"),
    FilterByExtension(".txt"),
)

files, err := WalkDirWithFilters("/path/to/project", docGroup)
```

### Cleanup Operations

```go
// Find files for cleanup
cleanupGroup := NewFilterGroup(FilterOr,
    // Temporary files
    FilterByExtension(".tmp", ".temp"),
    FilterByName("temp"),
    
    // Log files
    FilterByExtension(".log"),
    FilterByName("log"),
    
    // Cache files
    FilterByExtension(".cache"),
    FilterByPathPattern("cache/"),
    
    // Build artifacts
    FilterByExtension(".o", ".obj", ".exe"),
    FilterByPathPattern("build/", "dist/", "target/"),
)

files, err := WalkDirWithFilters("/path/to/project", cleanupGroup)
```

### Size Analysis

```go
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
        FilterBySize(1024*1024, 1024*1024*1024*10),
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
```

### Code Review Preparation

```go
// Advanced scenario: Find files for code review
codeReviewGroup := NewFilterGroup(FilterOr,
    // Go source files (not generated, not in vendor)
    NewFilterGroup(FilterAnd,
        FilterByExtension(".go"),
        FilterHidden,
        func(entry os.DirEntry, path string) bool {
            return !strings.Contains(path, "generated") && 
                   !strings.Contains(path, "vendor/")
        },
    ),
    
    // Test files
    NewFilterGroup(FilterAnd,
        func(entry os.DirEntry, path string) bool {
            name := strings.ToLower(entry.Name())
            return strings.HasSuffix(name, "_test.go")
        },
        FilterHidden,
    ),
    
    // Configuration files (in root or config directories)
    NewFilterGroup(FilterAnd,
        FilterByExtension(".json", ".yaml", ".yml"),
        FilterHidden,
        func(entry os.DirEntry, path string) bool {
            return !strings.Contains(path, "/") || 
                   strings.Contains(path, "/config/")
        },
    ),
    
    // Documentation
    NewFilterGroup(FilterAnd,
        FilterByExtension(".md"),
        FilterHidden,
    ),
)

files, err := WalkDirWithFilters("/path/to/project", codeReviewGroup)
```

## API Reference

### Functions

#### WalkDir(dir string, filters ...WalkFilter) ([]string, error)
Simple directory walking with basic filter combination (AND logic).

#### WalkDirWithFilters(dir string, filterGroups ...FilterGroup) ([]string, error)
Advanced directory walking with flexible filter group combinations.

#### NewFilterGroup(combiner FilterCombiner, filters ...WalkFilter) FilterGroup
Creates a new filter group with the specified combination logic.

### Predefined Filters

#### FilterByExtension(extensions ...string) WalkFilter
Matches files with the specified extensions.

#### FilterByName(pattern string) WalkFilter
Matches entries with names containing the specified pattern.

#### FilterBySize(minSize, maxSize int64) WalkFilter
Matches files within the specified size range.

#### FilterByPathPattern(pattern string) WalkFilter
Matches entries whose path contains the specified pattern.

#### FilterDirectoriesOnly(entry os.DirEntry, path string) bool
Only includes directories.

#### FilterFilesOnly(entry os.DirEntry, path string) bool
Only includes files.

#### FilterHidden(entry os.DirEntry, path string) bool
Excludes hidden files/directories (starting with '.').

### Methods

#### (fg FilterGroup) Apply(entry os.DirEntry, path string) bool
Applies the filter group to a directory entry.

## Best Practices

1. **Use Simple Filters When Possible**: For basic filtering, use `WalkDir` with predefined filters.

2. **Group Related Filters**: Use `FilterGroup` to logically group related filtering conditions.

3. **Avoid Deep Nesting**: While the system supports complex nesting, keep filter logic readable.

4. **Custom Filters for Complex Logic**: Create custom `WalkFilter` functions for complex business logic.

5. **Test Filter Logic**: Always test your filter combinations to ensure they work as expected.

6. **Performance Considerations**: 
   - Avoid expensive operations in filters (like file I/O)
   - Use early returns in custom filters
   - Consider the order of filters (put faster filters first in AND combinations)

7. **Error Handling**: Always check for errors when calling walk functions.

## Migration Guide

### From Simple Filters to Filter Groups

**Before:**
```go
files, err := WalkDir("/path", 
    FilterByExtension(".go"),
    FilterHidden,
)
```

**After (same result):**
```go
group := NewFilterGroup(FilterAnd,
    FilterByExtension(".go"),
    FilterHidden,
)
files, err := WalkDirWithFilters("/path", group)
```

### Adding OR Logic

**Before (AND only):**
```go
files, err := WalkDir("/path", 
    FilterByExtension(".go"),
    FilterByExtension(".java"), // This would never match (AND logic)
)
```

**After (OR logic):**
```go
group := NewFilterGroup(FilterOr,
    FilterByExtension(".go"),
    FilterByExtension(".java"),
)
files, err := WalkDirWithFilters("/path", group)
```

This flexible filtering system provides powerful capabilities for directory traversal while maintaining simplicity for common use cases.
