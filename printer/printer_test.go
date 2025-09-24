package printer

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestNewLine(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	NewLine()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if output != "\n" {
		t.Errorf("NewLine() = %q, want %q", output, "\n")
	}
}

func TestNewSepLine(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	NewSepLine()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "========================================================================================\n"
	if output != expected {
		t.Errorf("NewSepLine() = %q, want %q", output, expected)
	}
}

func TestPrintf(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Printf("Hello %s", "World")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello World"
	if output != expected {
		t.Errorf("Printf() = %q, want %q", output, expected)
	}
}

func TestPrintln(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Println("Hello %s", "World")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello World\n"
	if output != expected {
		t.Errorf("Println() = %q, want %q", output, expected)
	}
}

func TestPrintlnWithNewline(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Println("Hello %s\n", "World")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello World\n"
	if output != expected {
		t.Errorf("Println() = %q, want %q", output, expected)
	}
}

func TestPrintw(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := Printw(10, "Hello", 42, 3.14)

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Errorf("Printw() error = %v, want nil", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that output contains the formatted values
	if !strings.Contains(output, "Hello") {
		t.Errorf("Printw() output should contain 'Hello'")
	}
	if !strings.Contains(output, "42") {
		t.Errorf("Printw() output should contain '42'")
	}
}

func TestPrintwln(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := Printwln(10, "Hello", 42, 3.14)

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Errorf("Printwln() error = %v, want nil", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that output ends with newline
	if !strings.HasSuffix(output, "\n") {
		t.Errorf("Printwln() output should end with newline")
	}
}

func TestPrintwEmpty(t *testing.T) {
	// Capture output
	old := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	err := Printw(10)

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Errorf("Printw() with no args error = %v, want nil", err)
	}
}

func TestPrintwlnEmpty(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := Printwln(10)

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Errorf("Printwln() with no args error = %v, want nil", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Should print just a newline
	if output != "\n" {
		t.Errorf("Printwln() with no args = %q, want %q", output, "\n")
	}
}

func TestBuildFormatString(t *testing.T) {
	tests := []struct {
		width  int
		cols   []any
		expect string
	}{
		{10, []any{"hello"}, "%-10s"},
		{10, []any{42}, "%-10d"},
		{10, []any{3.14}, "%-10.2f"},
		{10, []any{"hello", 42, 3.14}, "%-10s%-10d%-10.2f"},
		{5, []any{true}, "%-5v"},
	}

	for _, test := range tests {
		result := buildFormatString(test.width, test.cols)
		if result != test.expect {
			t.Errorf("buildFormatString(%d, %v) = %q, want %q", test.width, test.cols, result, test.expect)
		}
	}
}

func TestPrintTable(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	headers := []string{"Name", "Age", "Score"}
	rows := [][]any{
		{"Alice", 25, 95.5},
		{"Bob", 30, 87.2},
		{"Charlie", 35, 92.8},
	}

	err := PrintTable(headers, rows, 10)

	w.Close()
	os.Stdout = old

	if err != nil {
		t.Errorf("PrintTable() error = %v, want nil", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that output contains headers and data
	if !strings.Contains(output, "Name") {
		t.Errorf("PrintTable() output should contain 'Name'")
	}
	if !strings.Contains(output, "Alice") {
		t.Errorf("PrintTable() output should contain 'Alice'")
	}
}

func TestPrintTableEmptyHeaders(t *testing.T) {
	// Capture output
	old := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	headers := []string{}
	rows := [][]any{}

	err := PrintTable(headers, rows, 10)

	w.Close()
	os.Stdout = old

	if err == nil {
		t.Errorf("PrintTable() with empty headers should return error")
	}
}

func TestPrintTableMismatchedRows(t *testing.T) {
	// Capture output
	old := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	headers := []string{"Name", "Age"}
	rows := [][]any{
		{"Alice", 25, "extra"}, // Extra column
	}

	err := PrintTable(headers, rows, 10)

	w.Close()
	os.Stdout = old

	if err == nil {
		t.Errorf("PrintTable() with mismatched rows should return error")
	}
}

func TestPrintJSON(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	data := map[string]any{
		"name": "Alice",
		"age":  25,
		"city": "New York",
	}

	PrintJSON(data)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that output contains JSON-like structure
	if !strings.Contains(output, "{") {
		t.Errorf("PrintJSON() output should contain '{'")
	}
	if !strings.Contains(output, "}") {
		t.Errorf("PrintJSON() output should contain '}'")
	}
	if !strings.Contains(output, "name") {
		t.Errorf("PrintJSON() output should contain 'name'")
	}
}

func TestPrintStruct(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fields := map[string]any{
		"Name": "Alice",
		"Age":  25,
		"City": "New York",
	}

	PrintStruct("Person", fields)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that output contains struct-like structure
	if !strings.Contains(output, "Person {") {
		t.Errorf("PrintStruct() output should contain 'Person {'")
	}
	if !strings.Contains(output, "}") {
		t.Errorf("PrintStruct() output should contain '}'")
	}
	if !strings.Contains(output, "Name") {
		t.Errorf("PrintStruct() output should contain 'Name'")
	}
}
