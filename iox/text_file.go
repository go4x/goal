package iox

import (
	"bufio"
	"io"
	"os"
)

// TxtFile provides a convenient interface for reading and writing text files.
// It uses buffered I/O for efficient file operations and provides methods
// for both reading and writing text content line by line.
//
// Example:
//
//	tf, err := iox.NewTxtFile("/path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer tf.Close()
//
//	// Write some lines
//	tf.WriteLine("Hello, World!")
//	tf.WriteLine("This is line 2")
//	tf.Flush()
//
//	// Read all lines
//	lines := tf.ReadAll()
//	for _, line := range lines {
//	    fmt.Println(line)
//	}
type TxtFile struct {
	file string
	f    *os.File
	bufw *bufio.Writer
}

// NewTxtFile creates a new TxtFile instance for the given file path.
// The file is opened in read-write mode, created if it doesn't exist, and opened in append mode.
// It returns an error if the file cannot be opened or created.
//
// Example:
//
//	tf, err := iox.NewTxtFile("/path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer tf.Close()
func NewTxtFile(f string) (*TxtFile, error) {
	tf := &TxtFile{
		file: f,
	}
	var err error
	tf.f, err = tf.createAndOpen(f)
	if err != nil {
		return nil, err
	}
	tf.bufw = bufio.NewWriter(tf.f)
	return tf, nil
}

// createAndOpen opens a file in read-write mode, creating it if it doesn't exist.
// The file is opened in append mode, so new content will be added to the end.
func (tf *TxtFile) createAndOpen(fpath string) (*os.File, error) {
	return os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

// WriteLine writes a string to the file followed by a newline character.
// It returns the TxtFile instance and any error that occurred during writing.
// The content is buffered and will be written to disk when Flush() is called.
//
// Example:
//
//	_, err := tf.WriteLine("Hello, World!")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (tf *TxtFile) WriteLine(s string) (*TxtFile, error) {
	if tf.bufw == nil {
		tf.bufw = bufio.NewWriter(tf.f)
	}
	_, err := tf.bufw.WriteString(s + "\n")
	return tf, err
}

// Flush writes any buffered data to the underlying file.
// It should be called before closing the file to ensure all data is written.
//
// Example:
//
//	err := tf.Flush()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (tf *TxtFile) Flush() error {
	return tf.bufw.Flush()
}

// Close closes the file and flushes any remaining buffered data.
// It should always be called when done with the file to ensure proper cleanup.
//
// Example:
//
//	defer tf.Close() // Ensure file is closed
func (tf *TxtFile) Close() error {
	if tf.bufw != nil {
		if err := tf.bufw.Flush(); err != nil {
			return err
		}
	}
	if tf.f != nil {
		return tf.f.Close()
	}
	return nil
}

// ReadAll reads all lines from the file and returns them as a slice of strings.
// It resets the file pointer to the beginning before reading.
// It returns an error if the file cannot be read.
//
// Example:
//
//	lines, err := tf.ReadAll()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, line := range lines {
//	    fmt.Println(line)
//	}
func (tf *TxtFile) ReadAll() ([]string, error) {
	var content []string

	// Reset file pointer to beginning
	_, err := tf.f.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	var buf = bufio.NewReader(tf.f)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		content = append(content, string(line))
	}
	return content, nil
}
