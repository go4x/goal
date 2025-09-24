package printer

import (
	"fmt"
	"strconv"
	"strings"
)

// NewLine prints a newline character
func NewLine() {
	fmt.Println()
}

// NewSepLine prints a separator line with 80 equal signs
func NewSepLine() {
	fmt.Println("========================================================================================")
}

// Printf is a wrapper around fmt.Printf for consistency
func Printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

// Println prints formatted text with automatic newline
func Println(format string, args ...any) {
	if len(format) > 0 && format[len(format)-1] != '\n' {
		format += "\n"
	}
	fmt.Printf(format, args...)
}

// Printw prints formatted columns with specified width
func Printw(width int, cols ...any) error {
	return printw(width, false, cols...)
}

// Printwln prints formatted columns with specified width and newline
func Printwln(width int, cols ...any) error {
	return printw(width, true, cols...)
}

// printw is the internal implementation for column printing
func printw(width int, addNewline bool, cols ...any) error {
	if len(cols) == 0 {
		if addNewline {
			fmt.Println()
		}
		return nil
	}

	format := buildFormatString(width, cols)
	if addNewline {
		format += "\n"
	}

	fmt.Printf(format, cols...)
	return nil
}

// buildFormatString creates a format string for the given columns
func buildFormatString(width int, cols []any) string {
	var sb strings.Builder

	for _, col := range cols {
		switch col.(type) {
		case string:
			sb.WriteString("%-" + strconv.Itoa(width) + "s")
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			sb.WriteString("%-" + strconv.Itoa(width) + "d")
		case float32, float64:
			// Use fixed decimal places for floats
			sb.WriteString("%-" + strconv.Itoa(width) + ".2f")
		default:
			// For unsupported types, use %v with width
			sb.WriteString("%-" + strconv.Itoa(width) + "v")
		}
	}

	return sb.String()
}

// PrintTable prints a table with headers and rows
func PrintTable(headers []string, rows [][]any, colWidth int) error {
	if len(headers) == 0 {
		return fmt.Errorf("headers cannot be empty")
	}

	// Print headers
	headerArgs := make([]any, len(headers))
	for i, h := range headers {
		headerArgs[i] = h
	}

	if err := Printwln(colWidth, headerArgs...); err != nil {
		return err
	}

	// Print separator line
	NewSepLine()

	// Print rows
	for _, row := range rows {
		if len(row) != len(headers) {
			return fmt.Errorf("row length %d does not match header length %d", len(row), len(headers))
		}

		if err := Printwln(colWidth, row...); err != nil {
			return err
		}
	}

	return nil
}

// PrintJSON prints data in JSON-like format
func PrintJSON(data map[string]any) {
	fmt.Println("{")
	for key, value := range data {
		fmt.Printf("  %q: %v,\n", key, value)
	}
	fmt.Println("}")
}

// PrintStruct prints a struct in a readable format
func PrintStruct(name string, fields map[string]any) {
	fmt.Printf("%s {\n", name)
	for key, value := range fields {
		fmt.Printf("  %s: %v\n", key, value)
	}
	fmt.Println("}")
}
