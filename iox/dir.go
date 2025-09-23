package iox

import (
	"os"
	"path/filepath"

	"github.com/go4x/goal/errorx"
)

// Dir provides convenient methods for directory operations.
// It's a global instance that can be used directly for directory-related operations.
//
// Example:
//
//	if exists, err := iox.Dir.Exists("/path/to/directory"); err == nil && exists {
//	    err := iox.Dir.Create("/path/to/new/directory")
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	}
var Dir = &dirs{}

// dirs provides methods for directory operations.
type dirs struct {
}

// Exists checks if a directory exists at the given path.
// It returns (true, nil) if the path exists and is a directory,
// (false, nil) if the path doesn't exist or is not a directory,
// and (false, error) if there was an error checking the path.
//
// Example:
//
//	exists, err := iox.Dir.Exists("/path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if exists {
//	    fmt.Println("Directory exists")
//	}
func (d *dirs) Exists(dir string) (bool, error) {
	fi, err := os.Stat(dir)
	if err == nil {
		return fi.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// AppendSeparator appends a path separator to the directory path if it doesn't already end with one.
// It returns the original string if it's empty, or the string with a trailing separator.
//
// Example:
//
//	path := iox.Dir.AppendSeparator("/path/to/dir")
//	// path is now "/path/to/dir/"
func (d *dirs) AppendSeparator(dir string) string {
	if dir == "" {
		return dir
	}
	s := string(filepath.Separator)
	if dir[len(dir)-1:] == s {
		return dir
	} else {
		return dir + s
	}
}

// Create creates a directory and any necessary parent directories.
// It returns an error if the directory cannot be created.
//
// Example:
//
//	err := iox.Dir.Create("/path/to/new/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (d *dirs) Create(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// CreateIfNotExists creates a directory only if it doesn't already exist.
// It returns an error if the directory already exists or cannot be created.
//
// Example:
//
//	err := iox.Dir.CreateIfNotExists("/path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (d *dirs) CreateIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return d.Create(dir)
	}
	return errorx.New("dir already exists")
}

// IsEmpty checks if a directory is empty (contains no files or subdirectories).
// It returns (true, nil) if the directory is empty, (false, nil) if not empty,
// and (false, error) if there was an error reading the directory.
//
// Example:
//
//	isEmpty, err := iox.Dir.IsEmpty("/path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if isEmpty {
//	    fmt.Println("Directory is empty")
//	}
func (d *dirs) IsEmpty(dir string) (bool, error) {
	es, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	return len(es) == 0, nil
}

// Delete removes a directory and all its contents recursively.
// It returns an error if the directory cannot be deleted.
//
// Example:
//
//	err := iox.Dir.Delete("/path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (d *dirs) Delete(dir string) error {
	return os.RemoveAll(dir)
}

// DeleteIfExists removes a directory only if it exists.
// It returns an error if the directory doesn't exist or cannot be deleted.
//
// Example:
//
//	err := iox.Dir.DeleteIfExists("/path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (d *dirs) DeleteIfExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return errorx.New("dir does not exist")
	}
	return d.Delete(dir)
}

// DeleteIfEmpty removes a directory only if it exists and is empty.
// It returns an error if the directory doesn't exist, is not empty, or cannot be deleted.
//
// Example:
//
//	err := iox.Dir.DeleteIfEmpty("/path/to/empty/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (d *dirs) DeleteIfEmpty(dir string) error {
	if b, err := d.IsEmpty(dir); err != nil {
		return err
	} else if !b {
		return errorx.New("dir is not empty")
	}
	return d.Delete(dir)
}

// Walk recursively walks through a directory and returns all file paths.
// If the given path is a file, it returns a slice containing just that file.
// If the given path is a directory, it returns all files in the directory and subdirectories.
//
// Example:
//
//	files, err := iox.Dir.Walk("/path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, file := range files {
//	    fmt.Println(file)
//	}
func (d *dirs) Walk(dir string, filterGroups ...FilterGroup) ([]string, error) {
	fi, err := os.Stat(dir)
	if err != nil {
		return []string{}, err
	}
	if fi.IsDir() {
		root := d.AppendSeparator(dir)
		return walkDir(root, filterGroups)
	} else {
		return []string{dir}, nil
	}
}
