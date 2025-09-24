package stringx_test

import (
	"testing"

	"github.com/go4x/goal/stringx"
)

func TestBuilderBasicOperations(t *testing.T) {
	b := stringx.NewBuilder()

	// Test basic string writing
	result := b.WriteString("Hello").WriteString(" ").WriteString("World").String()
	expected := "Hello World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test error handling
	if b.Error() != nil {
		t.Errorf("Expected no error, got %v", b.Error())
	}
}

func TestBuilderWriteRune(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteRune('H').WriteRune('e').WriteRune('l').WriteRune('l').WriteRune('o').String()
	expected := "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteByte(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteByte('H').WriteByte('e').WriteByte('l').WriteByte('l').WriteByte('o').String()
	expected := "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWrite(t *testing.T) {
	b := stringx.NewBuilder()
	data := []byte("Hello")
	result := b.Write(data).String()
	expected := "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWritef(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.Writef("Hello %s, you are %d years old", "World", 25).String()
	expected := "Hello World, you are 25 years old"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteLine(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteLine("Hello").WriteLine("World").String()
	expected := "Hello\nWorld\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteLinef(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteLinef("Hello %s", "World").WriteLinef("Count: %d", 42).String()
	expected := "Hello World\nCount: 42\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteIf(t *testing.T) {
	b := stringx.NewBuilder()

	// Test true condition
	result := b.WriteString("Start").WriteIf(true, "True").WriteIf(false, "False").String()
	expected := "StartTrue"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test false condition
	b2 := stringx.NewBuilder()
	result2 := b2.WriteString("Start").WriteIf(false, "True").WriteIf(true, "False").String()
	expected2 := "StartFalse"
	if result2 != expected2 {
		t.Errorf("Expected %q, got %q", expected2, result2)
	}
}

func TestBuilderWriteIfElse(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteString("Start").WriteIfElse(true, "True", "False").WriteIfElse(false, "True", "False").String()
	expected := "StartTrueFalse"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteRepeat(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteString("Start").WriteRepeat("X", 3).WriteString("End").String()
	expected := "StartXXXEnd"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test zero repeat
	b2 := stringx.NewBuilder()
	result2 := b2.WriteString("Start").WriteRepeat("X", 0).WriteString("End").String()
	expected2 := "StartEnd"
	if result2 != expected2 {
		t.Errorf("Expected %q, got %q", expected2, result2)
	}
}

func TestBuilderWriteJoin(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteString("Start").WriteJoin(", ", "a", "b", "c").WriteString("End").String()
	expected := "Starta, b, cEnd"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderLen(t *testing.T) {
	b := stringx.NewBuilder()
	if b.Len() != 0 {
		t.Errorf("Expected length 0, got %d", b.Len())
	}

	b.WriteString("Hello")
	if b.Len() != 5 {
		t.Errorf("Expected length 5, got %d", b.Len())
	}
}

func TestBuilderCap(t *testing.T) {
	b := stringx.NewBuilder()
	initialCap := b.Cap()

	b.WriteString("Hello")
	if b.Cap() < initialCap {
		t.Errorf("Capacity should not decrease")
	}
}

func TestBuilderReset(t *testing.T) {
	b := stringx.NewBuilder()
	b.WriteString("Hello")
	if b.Len() != 5 {
		t.Errorf("Expected length 5 before reset")
	}

	b.Reset()
	if b.Len() != 0 {
		t.Errorf("Expected length 0 after reset, got %d", b.Len())
	}
	if b.Error() != nil {
		t.Errorf("Expected no error after reset, got %v", b.Error())
	}
}

func TestBuilderGrow(t *testing.T) {
	b := stringx.NewBuilder()
	initialCap := b.Cap()

	b.Grow(100)
	if b.Cap() < initialCap+100 {
		t.Errorf("Capacity should grow by at least 100")
	}
}

func TestBuilderWriteSpace(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteString("Hello").WriteSpace().WriteString("World").String()
	expected := "Hello World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteTab(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteString("Hello").WriteTab().WriteString("World").String()
	expected := "Hello\tWorld"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteNewline(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteString("Hello").WriteNewline().WriteString("World").String()
	expected := "Hello\nWorld"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteIndent(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteString("Start").WriteIndent("  ", 2).WriteString("Indented").String()
	expected := "Start    Indented"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteWrap(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteWrap("<", "content", ">").String()
	expected := "<content>"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteQuoted(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteQuoted("Hello World").String()
	expected := `"Hello World"`
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteSingleQuoted(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteSingleQuoted("Hello World").String()
	expected := `'Hello World'`
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteBacktickQuoted(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteBacktickQuoted("Hello World").String()
	expected := "`Hello World`"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteBrackets(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteBrackets("Hello World").String()
	expected := "[Hello World]"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteParentheses(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteParentheses("Hello World").String()
	expected := "(Hello World)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderWriteBraces(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.WriteBraces("Hello World").String()
	expected := "{Hello World}"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderChaining(t *testing.T) {
	b := stringx.NewBuilder()
	result := b.
		WriteString("Hello").
		WriteSpace().
		WriteQuoted("World").
		WriteNewline().
		WriteIndent("  ", 1).
		WriteString("Indented").
		String()

	expected := "Hello \"World\"\n  Indented"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderErrorHandling(t *testing.T) {
	b := stringx.NewBuilder()

	// Simulate an error by writing to a nil builder (this won't actually cause an error in practice)
	// But we can test the error handling mechanism
	if b.Error() != nil {
		t.Errorf("Expected no error initially")
	}

	// Test that operations continue to work after potential errors
	b.WriteString("Hello")
	if b.Error() != nil {
		t.Errorf("Expected no error after WriteString")
	}
}

func TestBuilderComplexExample(t *testing.T) {
	b := stringx.NewBuilder()

	// Build a JSON-like structure
	result := b.
		WriteString("{").
		WriteNewline().
		WriteIndent("  ", 1).
		WriteQuoted("name").
		WriteString(": ").
		WriteQuoted("John").
		WriteString(",").
		WriteNewline().
		WriteIndent("  ", 1).
		WriteQuoted("age").
		WriteString(": ").
		WriteString("25").
		WriteNewline().
		WriteString("}").
		String()

	expected := `{
  "name": "John",
  "age": 25
}`
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
