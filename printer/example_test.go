package printer_test

import (
	"fmt"

	"github.com/go4x/goal/printer"
)

func ExampleNewLine() {
	// Print a newline
	printer.NewLine()
	// Output:
	//
}

func ExampleNewSepLine() {
	// Print a separator line
	printer.NewSepLine()
	// Output:
	// ========================================================================================
}

func ExamplePrintf() {
	// Print formatted text
	printer.Printf("Hello %s, you are %d years old\n", "Alice", 25)
	// Output: Hello Alice, you are 25 years old
}

func ExamplePrintln() {
	// Print formatted text with automatic newline
	printer.Println("Hello %s, you are %d years old", "Alice", 25)
	// Output: Hello Alice, you are 25 years old
}

func ExamplePrintw() {
	// Print formatted columns with specified width
	printer.Printw(10, "Name", "Age", "City")
	printer.Printw(10, "Alice", 25, "New York")
	printer.Printw(10, "Bob", 30, "London")
	// Output:
	// Name      Age       City
	// Alice     25        New York
	// Bob       30        London
}

func ExamplePrintwln() {
	// Print formatted columns with newline
	printer.Printwln(10, "Name", "Age", "City")
	printer.Printwln(10, "Alice", 25, "New York")
	printer.Printwln(10, "Bob", 30, "London")
	// Output:
	// Name      Age       City
	// Alice     25        New York
	// Bob       30        London
}

func ExamplePrintTable() {
	// Print a formatted table
	headers := []string{"Name", "Age", "Score"}
	rows := [][]any{
		{"Alice", 25, 95.5},
		{"Bob", 30, 87.2},
		{"Charlie", 35, 92.8},
	}

	printer.PrintTable(headers, rows, 12)
	// Output:
	// Name         Age         Score
	// ========================================================================================
	// Alice        25          95.50
	// Bob          30          87.20
	// Charlie      35          92.80
}

func ExamplePrintJSON() {
	// Print data in JSON-like format
	data := map[string]any{
		"name": "Alice",
		"age":  25,
		"city": "New York",
	}

	printer.PrintJSON(data)
	// Output:
	// {
	//   "name": Alice,
	//   "age": 25,
	//   "city": New York,
	// }
}

func ExamplePrintStruct() {
	// Print a struct in readable format
	fields := map[string]any{
		"Name": "Alice",
		"Age":  25,
		"City": "New York",
	}

	printer.PrintStruct("Person", fields)
	// Output:
	// Person {
	//   Name: Alice
	//   Age: 25
	//   City: New York
	// }
}

func ExamplePrintTable_complex() {
	// Complex usage: creating a report
	printer.Println("=== Sales Report ===")
	printer.NewSepLine()

	// Print table headers
	printer.Printwln(15, "Product", "Quantity", "Price", "Total")

	// Print table data
	products := []struct {
		name     string
		quantity int
		price    float64
	}{
		{"Laptop", 5, 999.99},
		{"Mouse", 20, 29.99},
		{"Keyboard", 15, 79.99},
	}

	total := 0.0
	for _, p := range products {
		itemTotal := float64(p.quantity) * p.price
		total += itemTotal
		printer.Printwln(15, p.name, p.quantity, p.price, itemTotal)
	}

	printer.NewSepLine()
	printer.Printf("Total: $%.2f\n", total)

	// Print summary as JSON
	summary := map[string]any{
		"total_products": len(products),
		"total_amount":   total,
		"average_price":  total / float64(len(products)),
	}

	printer.Println("\nSummary:")
	printer.PrintJSON(summary)
	// Output:
	// === Sales Report ===
	// ========================================================================================
	// Product         Quantity      Price          Total
	// Laptop          5             999.99         4999.95
	// Mouse           20            29.99          599.80
	// Keyboard        15            79.99          1199.85
	// ========================================================================================
	// Total: $6799.60
	//
	// Summary:
	// {
	//   "total_products": 3,
	//   "total_amount": 6799.6,
	//   "average_price": 2266.5333333333333,
	// }
}

func ExamplePrintTable_error() {
	// Demonstrate error handling
	headers := []string{"Name", "Age"}
	rows := [][]any{
		{"Alice", 25, "extra"}, // This will cause an error
	}

	if err := printer.PrintTable(headers, rows, 10); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// Output: Error: row length 3 does not match header length 2
}

func ExamplePrintw_custom() {
	// Custom formatting with different widths
	printer.Println("=== Product Inventory ===")

	// Headers with different widths
	printer.Printw(20, "Product Name")
	printer.Printw(10, "Stock")
	printer.Printw(15, "Price")
	printer.NewLine()

	// Data
	printer.Printw(20, "Gaming Laptop")
	printer.Printw(10, 15)
	printer.Printw(15, 1299.99)
	printer.NewLine()

	printer.Printw(20, "Wireless Mouse")
	printer.Printw(10, 50)
	printer.Printw(15, 39.99)
	printer.NewLine()

	printer.Printw(20, "Mechanical Keyboard")
	printer.Printw(10, 25)
	printer.Printw(15, 149.99)
	printer.NewLine()
	// Output:
	// === Product Inventory ===
	// Product Name         Stock      Price
	// Gaming Laptop        15         1299.99
	// Wireless Mouse       50         39.99
	// Mechanical Keyboard  25         149.99
}
