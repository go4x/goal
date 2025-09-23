package iox

import (
	"os"
	"path/filepath"
	"strings"
)

// WalkFilter is a function type that determines whether a directory entry should be included
// during directory traversal. It receives the directory entry and full path as parameters.
// Returns true to include the entry, false to exclude it.
//
// Example:
//
//	filter := func(entry os.DirEntry, path string) bool {
//	    return !entry.IsDir() && strings.HasSuffix(path, ".go")
//	}
type WalkFilter func(entry os.DirEntry, path string) bool

// FilterCombiner defines how multiple filters within a FilterGroup should be combined.
// It determines the logical operation used when evaluating multiple filters.
type FilterCombiner int

const (
	// FilterAnd requires all filters to return true (logical AND operation)
	FilterAnd FilterCombiner = iota
	// FilterOr requires at least one filter to return true (logical OR operation)
	FilterOr
)

// FilterGroup represents a group of filters with a specific combination logic.
// Filters within a group are combined according to the Combiner field.
// Multiple FilterGroups are combined with OR logic (if any group passes, the entry is included).
//
// Example:
//
//	group := iox.NewFilterGroup(iox.FilterAnd,
//	    iox.FilterByExtension(".go"),
//	    iox.FilterHidden,
//	)
type FilterGroup struct {
	Filters  []WalkFilter
	Combiner FilterCombiner
}

// NewFilterGroup creates a new FilterGroup with the specified filters and combiner.
// The filters will be combined using the specified FilterCombiner logic.
//
// Example:
//
//	group := iox.NewFilterGroup(iox.FilterAnd,
//	    iox.FilterByExtension(".go"),
//	    iox.FilterHidden,
//	)
func NewFilterGroup(combiner FilterCombiner, filters ...WalkFilter) FilterGroup {
	return FilterGroup{
		Filters:  filters,
		Combiner: combiner,
	}
}

// Apply applies the filter group to a directory entry.
func (fg FilterGroup) Apply(entry os.DirEntry, path string) bool {
	if len(fg.Filters) == 0 {
		return true // No filters means include everything
	}

	switch fg.Combiner {
	case FilterAnd:
		// All filters must return true
		for _, filter := range fg.Filters {
			if !filter(entry, path) {
				return false
			}
		}
		return true

	case FilterOr:
		// At least one filter must return true
		for _, filter := range fg.Filters {
			if filter(entry, path) {
				return true
			}
		}
		return false

	default:
		return true
	}
}

// walkDir recursively walks through a directory and returns all files that match the filter groups.
func walkDir(dir string, filterGroups []FilterGroup) ([]string, error) {
	dir = Dir.AppendSeparator(dir)
	es, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	var files []string
	for _, entry := range es {
		fullPath := filepath.Join(dir, entry.Name())

		// Apply all filter groups - if any group passes, include the entry
		// If no filter groups, include everything
		shouldInclude := len(filterGroups) == 0
		for _, group := range filterGroups {
			if group.Apply(entry, fullPath) {
				shouldInclude = true
				break
			}
		}

		if !shouldInclude {
			continue
		}

		if entry.IsDir() {
			// Recursively walk subdirectories
			subFiles, err := walkDir(fullPath, filterGroups)
			if err != nil {
				return []string{}, err
			}
			files = append(files, subFiles...)
		} else {
			// Add file to results
			files = append(files, fullPath)
		}
	}

	return files, nil
}

// WalkDirWithFilters walks through a directory with flexible filter combinations.
// You can pass multiple FilterGroup instances to create complex filtering logic.
func WalkDirWithFilters(dir string, filterGroups ...FilterGroup) ([]string, error) {
	return walkDir(dir, filterGroups)
}

// WalkDir walks through a directory with simple filter functions (legacy compatibility).
// Multiple filters are combined with AND logic by default.
func WalkDir(dir string, filters ...WalkFilter) ([]string, error) {
	if len(filters) == 0 {
		// No filters, include everything
		return walkDir(dir, []FilterGroup{})
	}

	// Convert simple filters to a filter group with AND logic
	filterGroup := NewFilterGroup(FilterAnd, filters...)
	return walkDir(dir, []FilterGroup{filterGroup})
}

// Predefined filter functions for common use cases

// FilterByExtension returns a filter that matches files with the specified extensions.
func FilterByExtension(extensions ...string) WalkFilter {
	return func(entry os.DirEntry, path string) bool {
		if entry.IsDir() {
			return true // Include directories
		}

		ext := strings.ToLower(filepath.Ext(path))
		for _, targetExt := range extensions {
			if strings.ToLower(targetExt) == ext {
				return true
			}
		}
		return false
	}
}

// FilterByName returns a filter that matches entries with names containing the specified pattern.
func FilterByName(pattern string) WalkFilter {
	return func(entry os.DirEntry, path string) bool {
		return strings.Contains(strings.ToLower(entry.Name()), strings.ToLower(pattern))
	}
}

// FilterBySize returns a filter that matches files within the specified size range.
func FilterBySize(minSize, maxSize int64) WalkFilter {
	return func(entry os.DirEntry, path string) bool {
		if entry.IsDir() {
			return true // Include directories
		}

		info, err := entry.Info()
		if err != nil {
			return false
		}

		size := info.Size()
		return size >= minSize && size <= maxSize
	}
}

// FilterDirectoriesOnly returns a filter that only includes directories.
func FilterDirectoriesOnly(entry os.DirEntry, path string) bool {
	return entry.IsDir()
}

// FilterFilesOnly returns a filter that only includes files.
func FilterFilesOnly(entry os.DirEntry, path string) bool {
	return !entry.IsDir()
}

// FilterHidden returns a filter that excludes hidden files/directories (starting with '.').
func FilterHidden(entry os.DirEntry, path string) bool {
	return !strings.HasPrefix(entry.Name(), ".")
}

// FilterByPathPattern returns a filter that matches entries whose path contains the specified pattern.
func FilterByPathPattern(pattern string) WalkFilter {
	return func(entry os.DirEntry, path string) bool {
		return strings.Contains(strings.ToLower(path), strings.ToLower(pattern))
	}
}
