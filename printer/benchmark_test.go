package printer

import (
	"os"
	"testing"
)

func BenchmarkNewLine(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewLine()
	}
}

func BenchmarkNewSepLine(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewSepLine()
	}
}

func BenchmarkPrintf(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Printf("Hello %s, you are %d years old\n", "Alice", 25)
	}
}

func BenchmarkPrintln(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Println("Hello %s, you are %d years old", "Alice", 25)
	}
}

func BenchmarkPrintw(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Printw(10, "Name", "Age", "City")
	}
}

func BenchmarkPrintwln(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Printwln(10, "Name", "Age", "City")
	}
}

func BenchmarkBuildFormatString(b *testing.B) {
	cols := []any{"Name", 25, 3.14, true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buildFormatString(10, cols)
	}
}

func BenchmarkPrintTable(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	headers := []string{"Name", "Age", "Score"}
	rows := [][]any{
		{"Alice", 25, 95.5},
		{"Bob", 30, 87.2},
		{"Charlie", 35, 92.8},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintTable(headers, rows, 10)
	}
}

func BenchmarkPrintJSON(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	data := map[string]any{
		"name": "Alice",
		"age":  25,
		"city": "New York",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintJSON(data)
	}
}

func BenchmarkPrintStruct(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	fields := map[string]any{
		"Name": "Alice",
		"Age":  25,
		"City": "New York",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintStruct("Person", fields)
	}
}

func BenchmarkPrintwWithDifferentTypes(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Printw(10, "String", 42, 3.14, true, []int{1, 2, 3})
	}
}

func BenchmarkPrintwlnWithDifferentTypes(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Printwln(10, "String", 42, 3.14, true, []int{1, 2, 3})
	}
}

func BenchmarkPrintwLargeTable(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout.Close()
		os.Stdout = old
	}()

	// Create a large table
	headers := []string{"ID", "Name", "Age", "Score", "City", "Country"}
	rows := make([][]any, 1000)
	for i := 0; i < 1000; i++ {
		rows[i] = []any{i, "User", 20 + i%50, 50.0 + float64(i%50), "City", "Country"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintTable(headers, rows, 10)
	}
}
