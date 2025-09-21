package stringx_test

import (
	"testing"

	"github.com/go4x/goal/assert"
	"github.com/go4x/goal/stringx"
	"github.com/go4x/got"
)

func TestBlurEmail(t *testing.T) {
	type Case struct {
		email  string
		expect string
	}
	cases := []Case{
		{"12345678@qq.com", "12****8qq.com"},
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
	assert.True(stringx.EndsWith("", ""))
	assert.True(stringx.EndsWith("a", ""))
	assert.True(!stringx.EndsWith("", "a"))

	s := "aaabb123b"
	assert.True(stringx.EndsWith(s, "b"))
	assert.True(stringx.EndsWith(s, "3b"))
	assert.True(stringx.EndsWith(s, "23b"))
	assert.True(stringx.EndsWith(s, "123b"))
	assert.True(!stringx.EndsWith(s, "a"))

	assert.True(stringx.StartsWith("", ""))
	assert.True(stringx.StartsWith("a", ""))
	assert.True(!stringx.StartsWith("", "a"))

	assert.True(stringx.StartsWith(s, "a"))
	assert.True(stringx.StartsWith(s, "aa"))
	assert.True(stringx.StartsWith(s, "aaa"))
	assert.True(stringx.StartsWith(s, "aaab"))
	assert.True(!stringx.StartsWith(s, "aaab1"))
	assert.True(!stringx.StartsWith(s, "1aaab1"))
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

	tl := got.New(t, "test CamelCaseToUnderscore")
	tl.Case("camelcase to underscore")

	for _, c := range cs {
		r := stringx.CamelCaseToUnderscore(c[0])
		tl.Require(r == c[1], "expect result is: %v, but is: %v", c[1], r)
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

	tl := got.New(t, "test UnderscoreToCamelCase")
	tl.Case("camelcase to underscore")

	for _, c := range cs {
		r := stringx.UnderscoreToCamelCase(c[1])
		tl.Require(r == c[0], "expect result is: %v, but is: %v", c[0], r)
	}
}
