package stringx_test

import (
	"testing"

	"github.com/go4x/goal/stringx"
)

func TestBlurEmail(t *testing.T) {
	type Case struct {
		email  string
		expect string
	}
	cases := []Case{
		{"12345678@qq.com", "1****8@qq.com"},
		{"abcd@126.com", "a****d@126.com"},
	}
	for _, c := range cases {
		dst := stringx.BlurEmail(c.email)
		if dst != c.expect {
			t.Errorf("test failed, expect: %v, but found: %v", c.expect, dst)
		}
	}
}

func TestEndsWith(t *testing.T) {
	if !stringx.EndsWith("", "") {
		t.Error("EndsWith(\"\", \"\") should be true")
	}
	if !stringx.EndsWith("a", "") {
		t.Error("EndsWith(\"a\", \"\") should be true")
	}
	if stringx.EndsWith("", "a") {
		t.Error("EndsWith(\"\", \"a\") should be false")
	}

	s := "aaabb123b"
	if !stringx.EndsWith(s, "b") {
		t.Error("EndsWith should be true for 'b'")
	}
	if !stringx.EndsWith(s, "3b") {
		t.Error("EndsWith should be true for '3b'")
	}
	if !stringx.EndsWith(s, "23b") {
		t.Error("EndsWith should be true for '23b'")
	}
	if !stringx.EndsWith(s, "123b") {
		t.Error("EndsWith should be true for '123b'")
	}
	if stringx.EndsWith(s, "a") {
		t.Error("EndsWith should be false for 'a'")
	}

	if !stringx.StartsWith("", "") {
		t.Error("StartsWith(\"\", \"\") should be true")
	}
	if !stringx.StartsWith("a", "") {
		t.Error("StartsWith(\"a\", \"\") should be true")
	}
	if stringx.StartsWith("", "a") {
		t.Error("StartsWith(\"\", \"a\") should be false")
	}

	if !stringx.StartsWith(s, "a") {
		t.Error("StartsWith should be true for 'a'")
	}
	if !stringx.StartsWith(s, "aa") {
		t.Error("StartsWith should be true for 'aa'")
	}
	if !stringx.StartsWith(s, "aaa") {
		t.Error("StartsWith should be true for 'aaa'")
	}
	if !stringx.StartsWith(s, "aaab") {
		t.Error("StartsWith should be true for 'aaab'")
	}
	if stringx.StartsWith(s, "aaab1") {
		t.Error("StartsWith should be false for 'aaab1'")
	}
	if stringx.StartsWith(s, "1aaab1") {
		t.Error("StartsWith should be false for '1aaab1'")
	}
}

func TestCamelCaseToUnderscore(t *testing.T) {
	cs := [][]string{
		{"HelloWorld", "hello_world"},
		{"helloWorld", "hello_world"},
		{"Helloworld", "helloworld"},
		{"AbcDEFGh", "abc_def_gh"},
		{"AbcDefGh", "abc_def_gh"},
		{"abcDefGh", "abc_def_gh"},
		{"abcDefGh😄", "abc_def_gh😄"},
	}

	for _, c := range cs {
		r := stringx.CamelCaseToUnderscore(c[0])
		if r != c[1] {
			t.Errorf("CamelCaseToUnderscore(%q) should be %q, got %q", c[0], c[1], r)
		}
	}
}

func TestUnderscoreToCamelCase(t *testing.T) {
	cs := [][]string{
		{"HelloWorld", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"Helloworld", "helloworld"},
		{"AbcDefGh", "abc_def_gh"},
		{"AbcDefGh中文", "abc_def_gh中文"},
	}

	for _, c := range cs {
		r := stringx.UnderscoreToCamelCase(c[1])
		if r != c[0] {
			t.Errorf("UnderscoreToCamelCase(%q) should be %q, got %q", c[1], c[0], r)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"0", true},
		{"", false},
		{"abc", false},
		{"12a", false},
		{"12.3", false},
		{"-123", false},
	}

	for _, test := range tests {
		result := stringx.IsNumeric(test.input)
		if result != test.expected {
			t.Errorf("IsNumeric(%q) should be %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"ABC", true},
		{"AbC", true},
		{"", false},
		{"123", false},
		{"abc123", false},
		{"a b", false},
	}

	for _, test := range tests {
		result := stringx.IsAlpha(test.input)
		if result != test.expected {
			t.Errorf("IsAlpha(%q) should be %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc123", true},
		{"ABC", true},
		{"123", true},
		{"", false},
		{"a b", false},
		{"a-b", false},
		{"a.b", false},
	}

	for _, test := range tests {
		result := stringx.IsAlphaNumeric(test.input)
		if result != test.expected {
			t.Errorf("IsAlphaNumeric(%q) should be %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"test@example.com", true},
		{"user@domain.org", true},
		{"", false},
		{"invalid", false},
		{"@domain.com", false},
		{"user@", false},
		{"user@domain", false},
	}

	for _, test := range tests {
		result := stringx.IsEmail(test.input)
		if result != test.expected {
			t.Errorf("IsEmail(%q) should be %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello_world"},
		{"helloWorld", "hello_world"},
		{"hello_world", "hello_world"},
		{"", ""},
		{"A", "a"},
	}

	for _, test := range tests {
		result := stringx.ToSnakeCase(test.input)
		if result != test.expected {
			t.Errorf("ToSnakeCase(%q) should be %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello-world"},
		{"helloWorld", "hello-world"},
		{"hello_world", "hello-world"},
		{"", ""},
		{"A", "a"},
	}

	for _, test := range tests {
		result := stringx.ToKebabCase(test.input)
		if result != test.expected {
			t.Errorf("ToKebabCase(%q) should be %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "HelloWorld"},
		{"hello-world", "HelloWorld"},
		{"helloWorld", "HelloWorld"},
		{"", ""},
		{"a", "A"},
	}

	for _, test := range tests {
		result := stringx.ToPascalCase(test.input)
		if result != test.expected {
			t.Errorf("ToPascalCase(%q) should be %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "helloWorld"},
		{"hello-world", "helloWorld"},
		{"HelloWorld", "helloWorld"},
		{"", ""},
		{"A", "a"},
	}

	for _, test := range tests {
		result := stringx.ToCamelCase(test.input)
		if result != test.expected {
			t.Errorf("ToCamelCase(%q) should be %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"世界", "界世"},
	}

	for _, test := range tests {
		result := stringx.Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) should be %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "helo"},
		{"", ""},
		{"a", "a"},
		{"aa", "a"},
		{"aabbcc", "abc"},
		{"hello world", "helo world"},
	}

	for _, test := range tests {
		result := stringx.RemoveDuplicates(test.input)
		if result != test.expected {
			t.Errorf("RemoveDuplicates(%q) should be %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		padChar  rune
		expected string
	}{
		{"hello", 10, '0', "00000hello"},
		{"hello", 5, '0', "hello"},
		{"hello", 3, '0', "hello"},
		{"", 5, '0', "00000"},
	}

	for _, test := range tests {
		result := stringx.PadLeft(test.input, test.length, test.padChar)
		if result != test.expected {
			t.Errorf("PadLeft(%q, %d, %c) should be %q, got %q", test.input, test.length, test.padChar, test.expected, result)
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		padChar  rune
		expected string
	}{
		{"hello", 10, '0', "hello00000"},
		{"hello", 5, '0', "hello"},
		{"hello", 3, '0', "hello"},
		{"", 5, '0', "00000"},
	}

	for _, test := range tests {
		result := stringx.PadRight(test.input, test.length, test.padChar)
		if result != test.expected {
			t.Errorf("PadRight(%q, %d, %c) should be %q, got %q", test.input, test.length, test.padChar, test.expected, result)
		}
	}
}

func TestPadCenter(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		padChar  rune
		expected string
	}{
		{"hello", 10, '0', "00hello000"},
		{"hello", 11, '0', "000hello000"},
		{"hello", 5, '0', "hello"},
		{"hello", 3, '0', "hello"},
		{"", 5, '0', "00000"},
	}

	for _, test := range tests {
		result := stringx.PadCenter(test.input, test.length, test.padChar)
		if result != test.expected {
			t.Errorf("PadCenter(%q, %d, %c) should be %q, got %q", test.input, test.length, test.padChar, test.expected, result)
		}
	}
}

func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"a, b, c", ",", []string{"a", "b", "c"}},
		{"a,,b, c", ",", []string{"a", "b", "c"}},
		{"", ",", []string{}},
		{"a", ",", []string{"a"}},
	}

	for _, test := range tests {
		result := stringx.SplitAndTrim(test.input, test.sep)
		if len(result) != len(test.expected) {
			t.Errorf("SplitAndTrim(%q, %q) should have length %d, got %d", test.input, test.sep, len(test.expected), len(result))
			continue
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("SplitAndTrim(%q, %q) should be %v, got %v", test.input, test.sep, test.expected, result)
				break
			}
		}
	}
}

func TestJoinNonEmpty(t *testing.T) {
	tests := []struct {
		sep      string
		strs     []string
		expected string
	}{
		{",", []string{"a", "", "b", "c"}, "a,b,c"},
		{",", []string{"", "a", "b", ""}, "a,b"},
		{",", []string{}, ""},
		{",", []string{"", "", ""}, ""},
	}

	for _, test := range tests {
		result := stringx.JoinNonEmpty(test.sep, test.strs...)
		if result != test.expected {
			t.Errorf("JoinNonEmpty(%q, %v) should be %q, got %q", test.sep, test.strs, test.expected, result)
		}
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		input    string
		size     int
		expected []string
	}{
		{"hello", 2, []string{"he", "ll", "o"}},
		{"hello", 3, []string{"hel", "lo"}},
		{"hello", 5, []string{"hello"}},
		{"hello", 10, []string{"hello"}},
		{"", 2, []string{""}},
		{"hello", 0, []string{"hello"}},
	}

	for _, test := range tests {
		result := stringx.Chunk(test.input, test.size)
		if len(result) != len(test.expected) {
			t.Errorf("Chunk(%q, %d) should have length %d, got %d", test.input, test.size, len(test.expected), len(result))
			continue
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("Chunk(%q, %d) should be %v, got %v", test.input, test.size, test.expected, result)
				break
			}
		}
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		expected []string
	}{
		{"hello world", 5, []string{"hello", "world"}},
		{"hello world", 20, []string{"hello world"}},
		{"", 10, []string{""}},
		{"hello", 3, []string{"hello"}},
	}

	for _, test := range tests {
		result := stringx.Wrap(test.input, test.width)
		if len(result) != len(test.expected) {
			t.Errorf("Wrap(%q, %d) should have length %d, got %d", test.input, test.width, len(test.expected), len(result))
			continue
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("Wrap(%q, %d) should be %v, got %v", test.input, test.width, test.expected, result)
				break
			}
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"hello", 1},
		{"", 0},
		{"  hello   world  ", 2},
	}

	for _, test := range tests {
		result := stringx.CountWords(test.input)
		if result != test.expected {
			t.Errorf("CountWords(%q) should be %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello\nworld", 2},
		{"hello", 1},
		{"", 0},
		{"hello\nworld\n", 3},
	}

	for _, test := range tests {
		result := stringx.CountLines(test.input)
		if result != test.expected {
			t.Errorf("CountLines(%q) should be %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected int
	}{
		{"hello world", "l", 3},
		{"hello world", "o", 2},
		{"hello world", "x", 0},
		{"hello world", "", 0},
		{"", "l", 0},
	}

	for _, test := range tests {
		result := stringx.CountOccurrences(test.input, test.substr)
		if result != test.expected {
			t.Errorf("CountOccurrences(%q, %q) should be %d, got %d", test.input, test.substr, test.expected, result)
		}
	}
}

func TestFindAll(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected []int
	}{
		{"hello world", "l", []int{2, 3, 9}},
		{"hello world", "o", []int{4, 7}},
		{"hello world", "x", []int{}},
		{"hello world", "", []int{}},
		{"", "l", []int{}},
	}

	for _, test := range tests {
		result := stringx.FindAll(test.input, test.substr)
		if len(result) != len(test.expected) {
			t.Errorf("FindAll(%q, %q) should have length %d, got %d", test.input, test.substr, len(test.expected), len(result))
			continue
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("FindAll(%q, %q) should be %v, got %v", test.input, test.substr, test.expected, result)
				break
			}
		}
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		input    string
		substrs  []string
		expected bool
	}{
		{"hello world", []string{"hello", "world"}, true},
		{"hello world", []string{"x", "y"}, false},
		{"hello world", []string{"x", "hello"}, true},
		{"", []string{"hello"}, false},
	}

	for _, test := range tests {
		result := stringx.ContainsAny(test.input, test.substrs...)
		if result != test.expected {
			t.Errorf("ContainsAny(%q, %v) should be %v, got %v", test.input, test.substrs, test.expected, result)
		}
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		input    string
		substrs  []string
		expected bool
	}{
		{"hello world", []string{"hello", "world"}, true},
		{"hello world", []string{"hello", "x"}, false},
		{"hello world", []string{"hello"}, true},
		{"", []string{"hello"}, false},
	}

	for _, test := range tests {
		result := stringx.ContainsAll(test.input, test.substrs...)
		if result != test.expected {
			t.Errorf("ContainsAll(%q, %v) should be %v, got %v", test.input, test.substrs, test.expected, result)
		}
	}
}
