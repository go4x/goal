package stringx

import (
	"fmt"
	"strings"
)

// Builder is a wrapper around strings.Builder that provides error handling.
// It allows chaining operations and tracks any errors that occur during string building.
type Builder struct {
	w   strings.Builder // Underlying strings.Builder
	err error           // Any error that occurred during building
}

// NewBuilder creates a new Builder instance.
func NewBuilder() *Builder {
	return &Builder{}
}

// WriteString appends the given string to the builder.
// Returns the builder for method chaining. If an error occurs, subsequent calls are ignored.
func (b *Builder) WriteString(s string) *Builder {
	if b.err != nil {
		return b
	}
	if _, err := b.w.WriteString(s); err != nil {
		b.err = err
	}
	return b
}

// Error returns any error that occurred during string building.
func (b *Builder) Error() error {
	return b.err
}

// String returns the accumulated string.
func (b *Builder) String() string {
	return b.w.String()
}

// WriteRune appends a single rune to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteRune(r rune) *Builder {
	if b.err != nil {
		return b
	}
	if _, err := b.w.WriteRune(r); err != nil {
		b.err = err
	}
	return b
}

// WriteByte appends a single byte to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteByte(c byte) *Builder {
	if b.err != nil {
		return b
	}
	if err := b.w.WriteByte(c); err != nil {
		b.err = err
	}
	return b
}

// Write appends a byte slice to the builder.
// Returns the builder for method chaining.
func (b *Builder) Write(p []byte) *Builder {
	if b.err != nil {
		return b
	}
	if _, err := b.w.Write(p); err != nil {
		b.err = err
	}
	return b
}

// Writef appends a formatted string to the builder using fmt.Sprintf.
// Returns the builder for method chaining.
func (b *Builder) Writef(format string, args ...interface{}) *Builder {
	if b.err != nil {
		return b
	}
	if _, err := fmt.Fprintf(&b.w, format, args...); err != nil {
		b.err = err
	}
	return b
}

// WriteLine appends a string followed by a newline to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteLine(s string) *Builder {
	return b.WriteString(s).WriteString("\n")
}

// WriteLinef appends a formatted string followed by a newline to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteLinef(format string, args ...interface{}) *Builder {
	return b.Writef(format, args...).WriteString("\n")
}

// WriteIf appends a string to the builder if the condition is true.
// Returns the builder for method chaining.
func (b *Builder) WriteIf(condition bool, s string) *Builder {
	if condition {
		return b.WriteString(s)
	}
	return b
}

// WriteIfElse appends one string if condition is true, another if false.
// Returns the builder for method chaining.
func (b *Builder) WriteIfElse(condition bool, ifTrue, ifFalse string) *Builder {
	if condition {
		return b.WriteString(ifTrue)
	}
	return b.WriteString(ifFalse)
}

// WriteRepeat appends a string repeated n times to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteRepeat(s string, n int) *Builder {
	if b.err != nil {
		return b
	}
	if n <= 0 {
		return b
	}
	for i := 0; i < n; i++ {
		if _, err := b.w.WriteString(s); err != nil {
			b.err = err
			return b
		}
	}
	return b
}

// WriteJoin appends strings joined by a separator to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteJoin(sep string, strs ...string) *Builder {
	if b.err != nil {
		return b
	}
	joined := strings.Join(strs, sep)
	return b.WriteString(joined)
}

// Len returns the number of accumulated bytes.
func (b *Builder) Len() int {
	return b.w.Len()
}

// Cap returns the capacity of the builder's underlying byte slice.
func (b *Builder) Cap() int {
	return b.w.Cap()
}

// Reset resets the builder to be empty.
func (b *Builder) Reset() *Builder {
	b.w.Reset()
	b.err = nil
	return b
}

// Grow grows the builder's capacity by at least n bytes.
// Returns the builder for method chaining.
func (b *Builder) Grow(n int) *Builder {
	if b.err != nil {
		return b
	}
	b.w.Grow(n)
	return b
}

// WriteSpace appends a space character to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteSpace() *Builder {
	return b.WriteString(" ")
}

// WriteTab appends a tab character to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteTab() *Builder {
	return b.WriteString("\t")
}

// WriteNewline appends a newline character to the builder.
// Returns the builder for method chaining.
func (b *Builder) WriteNewline() *Builder {
	return b.WriteString("\n")
}

// WriteIndent appends an indentation string repeated n times.
// Returns the builder for method chaining.
func (b *Builder) WriteIndent(indent string, n int) *Builder {
	return b.WriteRepeat(indent, n)
}

// WriteWrap wraps the content with prefix and suffix.
// Returns the builder for method chaining.
func (b *Builder) WriteWrap(prefix, content, suffix string) *Builder {
	return b.WriteString(prefix).WriteString(content).WriteString(suffix)
}

// WriteQuoted wraps the content with quotes.
// Returns the builder for method chaining.
func (b *Builder) WriteQuoted(content string) *Builder {
	return b.WriteString(`"`).WriteString(content).WriteString(`"`)
}

// WriteSingleQuoted wraps the content with single quotes.
// Returns the builder for method chaining.
func (b *Builder) WriteSingleQuoted(content string) *Builder {
	return b.WriteString(`'`).WriteString(content).WriteString(`'`)
}

// WriteBacktickQuoted wraps the content with backticks.
// Returns the builder for method chaining.
func (b *Builder) WriteBacktickQuoted(content string) *Builder {
	return b.WriteString("`").WriteString(content).WriteString("`")
}

// WriteBrackets wraps the content with brackets.
// Returns the builder for method chaining.
func (b *Builder) WriteBrackets(content string) *Builder {
	return b.WriteString("[").WriteString(content).WriteString("]")
}

// WriteParentheses wraps the content with parentheses.
// Returns the builder for method chaining.
func (b *Builder) WriteParentheses(content string) *Builder {
	return b.WriteString("(").WriteString(content).WriteString(")")
}

// WriteBraces wraps the content with braces.
// Returns the builder for method chaining.
func (b *Builder) WriteBraces(content string) *Builder {
	return b.WriteString("{").WriteString(content).WriteString("}")
}
