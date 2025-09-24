# Stringx Package

A comprehensive and powerful string manipulation library for Go that provides a rich set of utilities for string processing, validation, transformation, and building.

## Features

- **String Utilities**: Comprehensive set of string manipulation functions
- **String Builder**: Advanced string building with error handling and method chaining
- **String Validation**: Various validation functions for different string patterns
- **String Transformation**: Case conversion, formatting, and encoding utilities
- **String Constants**: Extensive collection of commonly used string constants
- **Pattern Matching**: Efficient multi-pattern string replacement using Aho-Corasick algorithm
- **Unicode Support**: Full Unicode and multi-byte character support
- **Performance Optimized**: Built on Go's standard library for optimal performance

## Installation

```bash
go get github.com/go4x/goal/stringx
```

## Quick Start

### Basic String Operations

```go
package main

import (
    "fmt"
    "github.com/go4x/goal/stringx"
)

func main() {
    // String validation
    fmt.Println(stringx.IsEmail("user@example.com"))        // true
    fmt.Println(stringx.IsNumeric("12345"))                 // true
    fmt.Println(stringx.IsAlpha("Hello"))                  // true
    
    // String transformation
    fmt.Println(stringx.ToSnakeCase("HelloWorld"))          // "hello_world"
    fmt.Println(stringx.ToCamelCase("hello_world"))        // "helloWorld"
    fmt.Println(stringx.Reverse("Hello"))                   // "olleH"
    
    // String utilities
    fmt.Println(stringx.Cut(10, "This is a long string"))   // "This is a ..."
    fmt.Println(stringx.BlurEmail("john@example.com"))      // "jo****@example.com"
}
```

### Advanced String Building

```go
// Using the enhanced Builder
builder := stringx.NewBuilder()
result := builder.
    WriteString("Hello").
    WriteSpace().
    WriteQuoted("World").
    WriteNewline().
    WriteIndent("  ", 1).
    WriteString("Indented content").
    String()

fmt.Println(result)
// Output:
// Hello "World"
//   Indented content
```

### Pattern Replacement

```go
// Multi-pattern string replacement
replacer := stringx.NewReplacer(map[string]string{
    "hello": "hi",
    "world": "universe",
    "goodbye": "bye",
})

text := "hello world, goodbye world"
result := replacer.Replace(text)
fmt.Println(result) // "hi universe, bye universe"
```

## API Reference

### String Utilities

#### Validation Functions
- `IsNumeric(s string) bool` - Check if string contains only numeric characters
- `IsAlpha(s string) bool` - Check if string contains only alphabetic characters
- `IsAlphaNumeric(s string) bool` - Check if string contains only alphanumeric characters
- `IsEmail(s string) bool` - Basic email format validation
- `IsEmpty(str string) bool` - Check if string is empty or whitespace only
- `IsSpace(s string) bool` - Check if string contains only spaces
- `HasLen(s string) bool` - Check if string has meaningful content

#### Transformation Functions
- `ToSnakeCase(s string) string` - Convert to snake_case
- `ToKebabCase(s string) string` - Convert to kebab-case
- `ToPascalCase(s string) string` - Convert to PascalCase
- `ToCamelCase(s string) string` - Convert to camelCase
- `ToTitle(s string) string` - Convert to Title Case
- `Reverse(s string) string` - Reverse the string
- `RemoveDuplicates(s string) string` - Remove consecutive duplicate characters

#### String Manipulation
- `Cut(max int, s string) string` - Truncate string with ellipsis
- `Trim(s, cut string) string` - Trim characters from both ends
- `TrimLeft(s, cut string) string` - Trim characters from left
- `TrimRight(s, cut string) string` - Trim characters from right
- `TrimSpace(s string) string` - Trim whitespace
- `RemSpace(s string) string` - Remove all spaces
- `Replace(s, old, new string, n int) string` - Replace occurrences
- `ReplaceAll(s, old, new string) string` - Replace all occurrences

#### Padding Functions
- `PadLeft(s string, length int, padChar rune) string` - Left pad
- `PadRight(s string, length int, padChar rune) string` - Right pad
- `PadCenter(s string, length int, padChar rune) string` - Center pad

#### String Analysis
- `CountWords(s string) int` - Count words in string
- `CountLines(s string) int` - Count lines in string
- `CountOccurrences(s, substr string) int` - Count substring occurrences
- `FindAll(s, substr string) []int` - Find all substring positions
- `ContainsAny(s string, substrings ...string) bool` - Check if contains any substring
- `ContainsAll(s string, substrings ...string) bool` - Check if contains all substrings

#### String Splitting and Joining
- `SplitAndTrim(s, sep string) []string` - Split and trim whitespace
- `JoinNonEmpty(sep string, strs ...string) string` - Join non-empty strings
- `Chunk(s string, size int) []string` - Split into chunks
- `Wrap(s string, width int) []string` - Wrap text to width

#### Privacy Functions
- `BlurEmail(email string) string` - Blur email address
- `Blur(str string, start, end int, sep string, num int) string` - Blur string content

### String Builder

The `Builder` type provides advanced string building capabilities with error handling and method chaining.

#### Basic Operations
- `NewBuilder() *Builder` - Create new builder
- `WriteString(s string) *Builder` - Append string
- `WriteRune(r rune) *Builder` - Append rune
- `WriteByte(c byte) *Builder` - Append byte
- `Write(p []byte) *Builder` - Append byte slice

#### Formatting Operations
- `Writef(format string, args ...interface{}) *Builder` - Formatted write
- `WriteLine(s string) *Builder` - Write string with newline
- `WriteLinef(format string, args ...interface{}) *Builder` - Formatted write with newline

#### Conditional Operations
- `WriteIf(condition bool, s string) *Builder` - Conditional write
- `WriteIfElse(condition bool, ifTrue, ifFalse string) *Builder` - Conditional choice

#### Repetition and Joining
- `WriteRepeat(s string, n int) *Builder` - Repeat string
- `WriteJoin(sep string, strs ...string) *Builder` - Join strings

#### Whitespace Operations
- `WriteSpace() *Builder` - Write space
- `WriteTab() *Builder` - Write tab
- `WriteNewline() *Builder` - Write newline
- `WriteIndent(indent string, n int) *Builder` - Write indentation

#### Wrapping Operations
- `WriteWrap(prefix, content, suffix string) *Builder` - Wrap content
- `WriteQuoted(content string) *Builder` - Wrap with double quotes
- `WriteSingleQuoted(content string) *Builder` - Wrap with single quotes
- `WriteBacktickQuoted(content string) *Builder` - Wrap with backticks
- `WriteBrackets(content string) *Builder` - Wrap with brackets
- `WriteParentheses(content string) *Builder` - Wrap with parentheses
- `WriteBraces(content string) *Builder` - Wrap with braces

#### State Operations
- `Len() int` - Get current length
- `Cap() int` - Get current capacity
- `Reset() *Builder` - Reset builder
- `Grow(n int) *Builder` - Pre-allocate capacity
- `Error() error` - Get any error
- `String() string` - Get final string

### String Constants

The package provides extensive string constants organized by category:

#### Punctuation and Symbols
```go
stringx.Exclamation     // "!"
stringx.AtSign          // "@"
stringx.HashTag         // "#"
stringx.DollarSign      // "$"
stringx.PercentSign     // "%"
stringx.Caret           // "^"
stringx.AmpersandSign   // "&"
stringx.StarSign        // "*"
stringx.PlusSign        // "+"
stringx.MinusSign       // "-"
stringx.EqualsSign      // "="
stringx.UnderscoreSign  // "_"
stringx.PipeSign        // "|"
stringx.BackslashSign   // "\\"
stringx.ForwardSlash    // "/"
stringx.ColonSign       // ":"
stringx.SemicolonSign   // ";"
stringx.CommaSign       // ","
stringx.DotSign         // "."
stringx.QuestionSign    // "?"
```

#### Brackets and Quotes
```go
stringx.LeftParen       // "("
stringx.RightParen      // ")"
stringx.LeftBracket     // "["
stringx.RightBracket    // "]"
stringx.LeftBrace       // "{"
stringx.RightBrace      // "}"
stringx.LeftAngle       // "<"
stringx.RightAngle      // ">"
stringx.DoubleQuote     // "\""
stringx.SingleQuote     // "'"
stringx.BacktickQuote   // "`"
```

#### Whitespace Characters
```go
stringx.SpaceChar       // " "
stringx.TabChar         // "\t"
stringx.NewlineChar     // "\n"
stringx.CarriageReturn  // "\r"
stringx.FormFeed        // "\f"
stringx.VerticalTab     // "\v"
```

#### Common Separators
```go
stringx.CommaSpace      // ", "
stringx.SemicolonSpace  // "; "
stringx.ColonSpace      // ": "
stringx.PipeSpace       // " | "
stringx.SlashSpace      // " / "
stringx.BackslashSpace  // " \\ "
```

#### Boolean Values
```go
stringx.BooleanTrue     // "true"
stringx.BooleanFalse    // "false"
stringx.BooleanYes      // "yes"
stringx.BooleanNo       // "no"
stringx.BooleanOn       // "on"
stringx.BooleanOff      // "off"
stringx.BooleanEnabled  // "enabled"
stringx.BooleanDisabled // "disabled"
```

#### File Extensions
```go
stringx.ExtJSON         // ".json"
stringx.ExtXML          // ".xml"
stringx.ExtHTML         // ".html"
stringx.ExtCSS          // ".css"
stringx.ExtJS           // ".js"
stringx.ExtGo           // ".go"
stringx.ExtTxt          // ".txt"
stringx.ExtLog          // ".log"
stringx.ExtYAML         // ".yaml"
stringx.ExtYML          // ".yml"
stringx.ExtTOML         // ".toml"
stringx.ExtINI          // ".ini"
stringx.ExtConf         // ".conf"
stringx.ExtConfig       // ".config"
```

#### Protocols
```go
stringx.ProtocolHTTP            // "http"
stringx.ProtocolHTTPS           // "https"
stringx.ProtocolFTP             // "ftp"
stringx.ProtocolFTPS            // "ftps"
stringx.ProtocolSFTP            // "sftp"
stringx.ProtocolSSH             // "ssh"
stringx.ProtocolTCP             // "tcp"
stringx.ProtocolUDP             // "udp"
stringx.ProtocolWebSocket       // "ws"
stringx.ProtocolWebSocketSecure // "wss"
```

#### Encoding Types
```go
stringx.EncodingUTF8   // "utf-8"
stringx.EncodingUTF16  // "utf-16"
stringx.EncodingASCII  // "ascii"
stringx.EncodingBase64 // "base64"
stringx.EncodingHex    // "hex"
stringx.EncodingURL    // "url"
```

#### Time Zones
```go
stringx.TimezoneUTC // "UTC"
stringx.TimezoneGMT // "GMT"
stringx.TimezoneEST // "EST"
stringx.TimezonePST // "PST"
stringx.TimezoneCST // "CST"
stringx.TimezoneMST // "MST"
```

#### Units
```go
stringx.UnitBytes // "bytes"
stringx.UnitKB    // "KB"
stringx.UnitMB    // "MB"
stringx.UnitGB    // "GB"
stringx.UnitTB    // "TB"
stringx.UnitPB    // "PB"
```

## Examples

### JSON Generation
```go
builder := stringx.NewBuilder()
json := builder.
    WriteBraces("").
    WriteNewline().
    WriteIndent("  ", 1).
    WriteQuoted("name").
    WriteString(stringx.Colon + stringx.Space).
    WriteQuoted("John").
    WriteString(stringx.Comma).
    WriteNewline().
    WriteIndent("  ", 1).
    WriteQuoted("age").
    WriteString(stringx.Colon + stringx.Space).
    WriteString("25").
    WriteNewline().
    WriteString("}").
    String()

fmt.Println(json)
// Output:
// {
//   "name": "John",
//   "age": 25
// }
```

### URL Building
```go
url := stringx.ProtocolHTTPS + 
    stringx.Colon + 
    stringx.ForwardSlash + 
    stringx.ForwardSlash + 
    "api.example.com" + 
    stringx.Slash + 
    "v1" + 
    stringx.Slash + 
    "users"
```

### File Path Construction
```go
configPath := "config" + stringx.Slash + "app" + stringx.ExtJSON
logPath := "logs" + stringx.Slash + "app" + stringx.ExtLog
```

### String Validation Pipeline
```go
func validateUserInput(input string) error {
    if stringx.IsEmpty(input) {
        return errors.New("input cannot be empty")
    }
    
    if !stringx.IsAlphaNumeric(input) {
        return errors.New("input must be alphanumeric")
    }
    
    if stringx.CountWords(input) > 10 {
        return errors.New("input too long")
    }
    
    return nil
}
```

### Text Processing
```go
func processText(text string) string {
    // Remove duplicates and normalize
    processed := stringx.RemoveDuplicates(text)
    
    // Convert to title case
    processed = stringx.ToTitle(processed)
    
    // Wrap long lines
    lines := stringx.Wrap(processed, 80)
    
    // Join with proper formatting
    return stringx.JoinNonEmpty(stringx.NewlineChar, lines...)
}
```

## Performance Considerations

- **Memory Efficiency**: The package is designed for minimal memory allocation
- **Unicode Support**: Full support for multi-byte characters and Unicode
- **Builder Pattern**: Efficient string building with pre-allocation
- **Algorithm Optimization**: Uses efficient algorithms like Aho-Corasick for pattern matching
- **Zero-Copy Operations**: Many operations avoid unnecessary string copying

## Best Practices

1. **Use Constants**: Prefer string constants over hardcoded strings
2. **Builder for Complex Strings**: Use Builder for complex string construction
3. **Validate Early**: Use validation functions early in your pipeline
4. **Handle Errors**: Always check for errors when using Builder
5. **Unicode Awareness**: Be aware of multi-byte characters in string operations

## License

This package is part of the goal project. See the main project license for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Changelog

### v1.0.0
- Initial release with comprehensive string utilities
- Advanced Builder with method chaining
- Extensive string constants collection
- Pattern matching with Aho-Corasick algorithm
- Full Unicode support
- Comprehensive test coverage
