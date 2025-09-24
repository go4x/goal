# printer - Enhanced Printing Utilities for Go

A comprehensive Go package providing enhanced printing utilities with formatting, table support, and structured output capabilities.

## Features

- **Enhanced Formatting**: Advanced column formatting with width control
- **Table Support**: Built-in table printing with headers and rows
- **Structured Output**: JSON-like and struct-like output formatting
- **Error Handling**: Proper error handling for all operations
- **Type Safety**: Support for various data types with automatic formatting
- **Performance Optimized**: Efficient string building and formatting

## Installation

```bash
go get github.com/go4x/goal/printer
```

## Quick Start

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // Basic printing
    printer.Println("Hello %s, you are %d years old", "Alice", 25)
    
    // Column formatting
    printer.Printwln(10, "Name", "Age", "City")
    printer.Printwln(10, "Alice", 25, "New York")
    
    // Table printing
    headers := []string{"Name", "Age", "Score"}
    rows := [][]any{
        {"Alice", 25, 95.5},
        {"Bob", 30, 87.2},
    }
    printer.PrintTable(headers, rows, 12)
}
```

## API Reference

### Basic Printing

#### `NewLine()`
Prints a newline character.

```go
printer.NewLine()  // Prints \n
```

#### `NewSepLine()`
Prints a separator line with 80 equal signs.

```go
printer.NewSepLine()  // Prints ========================================================================================
```

#### `Printf(format string, args ...any)`
Wrapper around fmt.Printf for consistency.

```go
printer.Printf("Hello %s, you are %d years old\n", "Alice", 25)
```

#### `Println(format string, args ...any)`
Prints formatted text with automatic newline.

```go
printer.Println("Hello %s, you are %d years old", "Alice", 25)
// Output: Hello Alice, you are 25 years old
```

### Column Formatting

#### `Printw(width int, cols ...any) error`
Prints formatted columns with specified width.

```go
err := printer.Printw(10, "Name", "Age", "City")
if err != nil {
    // Handle error
}
```

#### `Printwln(width int, cols ...any) error`
Prints formatted columns with specified width and newline.

```go
err := printer.Printwln(10, "Name", "Age", "City")
if err != nil {
    // Handle error
}
```

### Table Printing

#### `PrintTable(headers []string, rows [][]any, colWidth int) error`
Prints a formatted table with headers and rows.

```go
headers := []string{"Name", "Age", "Score"}
rows := [][]any{
    {"Alice", 25, 95.5},
    {"Bob", 30, 87.2},
}

err := printer.PrintTable(headers, rows, 12)
if err != nil {
    // Handle error
}
```

### Structured Output

#### `PrintJSON(data map[string]any)`
Prints data in JSON-like format.

```go
data := map[string]any{
    "name": "Alice",
    "age":  25,
    "city": "New York",
}
printer.PrintJSON(data)
```

#### `PrintStruct(name string, fields map[string]any)`
Prints a struct in readable format.

```go
fields := map[string]any{
    "Name": "Alice",
    "Age":  25,
    "City": "New York",
}
printer.PrintStruct("Person", fields)
```

## Usage Examples

### Basic Usage

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // Basic printing
    printer.Println("Welcome to the printer package!")
    printer.NewSepLine()
    
    // Column formatting
    printer.Printwln(15, "Product", "Price", "Stock")
    printer.Printwln(15, "Laptop", 999.99, 10)
    printer.Printwln(15, "Mouse", 29.99, 50)
}
```

### Table Printing

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // Create a sales report
    headers := []string{"Product", "Quantity", "Price", "Total"}
    rows := [][]any{
        {"Laptop", 5, 999.99, 4999.95},
        {"Mouse", 20, 29.99, 599.80},
        {"Keyboard", 15, 79.99, 1199.85},
    }
    
    printer.Println("=== Sales Report ===")
    printer.PrintTable(headers, rows, 15)
}
```

### Structured Data

```go
package main

import (
    "github.com/go4x/goal/printer"
)

func main() {
    // Print user information as JSON
    user := map[string]any{
        "name": "Alice",
        "age":  25,
        "email": "alice@example.com",
        "active": true,
    }
    
    printer.Println("User Information:")
    printer.PrintJSON(user)
    
    // Print as struct
    printer.Println("\nUser Details:")
    printer.PrintStruct("User", user)
}
```

### Error Handling

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/printer"
)

func main() {
    headers := []string{"Name", "Age"}
    rows := [][]any{
        {"Alice", 25, "extra"}, // This will cause an error
    }
    
    if err := printer.PrintTable(headers, rows, 10); err != nil {
        fmt.Printf("Error: %v\n", err)
        // Handle the error appropriately
    }
}
```

## Type Support

The printer package supports various data types with automatic formatting:

- **Strings**: Left-aligned with specified width
- **Integers**: Right-aligned with specified width
- **Floats**: Right-aligned with 2 decimal places
- **Other types**: Formatted using %v with specified width

## Performance

The printer package is optimized for performance:

- **Efficient string building** using strings.Builder
- **Minimal memory allocations** for common operations
- **Fast column formatting** with pre-computed format strings
- **Optimized table printing** with batch operations

## Testing

Run the tests:

```bash
go test ./printer
```

Run with coverage:

```bash
go test ./printer -cover
```

Run examples:

```bash
go test ./printer -run Example
```

Run benchmarks:

```bash
go test ./printer -bench=.
```

## License

This package is part of the `goal` project and follows the same license terms.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
