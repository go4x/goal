package assert

import (
	"fmt"
	"reflect"
	"strings"
)

// basic methods

// Require is the basic method to assert a condition, if the condition is false, it will panic with the given format and arguments
// format is a string that will be formatted with the arguments
// v is a variadic parameter that will be used to format the string
// if the condition is true, the function will return normally
func Require(cond bool, format string, v ...any) {
	if !cond {
		panic(fmt.Sprintf(format, v...))
	}
}

// convenient methods

// True asserts that the given boolean value is true
func True(b bool) {
	Require(b, "assert failed, expects [%t] but found [%t]", true, b)
}

// Nil asserts that the given value is nil
func Nil(t any) {
	Require(t == nil, "assert failed, expects [nil] but found [not nil]: %v", t)
}

// NoneNil asserts that the given value is not nil
func NoneNil(t any) {
	Require(t != nil, "assert failed, expects [not nil] but found [nil]")
}

// Blank asserts that the given string is blank (empty or only contains whitespace).
func Blank(s string) {
	s = strings.TrimSpace(s)
	Require(s == "", "assert failed, expects [\"\"] but found [%s]", s)
}

// NotBlank asserts that the given string is not blank (not empty or not only contains whitespace).
func NotBlank(s string) {
	s = strings.TrimSpace(s)
	Require(s != "", "assert failed, expects [not empty] but found [\"\"]")
}

// HasElems asserts that the given collection is not nil and has elements. Param c should be a collection type such as array, map, slice, or channel,
// do nothing if not.
func HasElems(c any) {
	typ := reflect.TypeOf(c).Kind()
	if typ == reflect.Array || typ == reflect.Map || typ == reflect.Slice || typ == reflect.Chan {
		n := reflect.ValueOf(c).Len()
		Require(c != nil && n > 0, "assert failed, expects collection is none nil and has elements")
	}
}

// Equals asserts that the given two values are equal
func Equals(t1 any, t2 any) {
	Require(t1 == t2, "assert failed, expecting t1 equals t2 but not")
}

// DeepEquals asserts that the given two values are deeply equal, it will use reflect.DeepEqual to compare the two values.
func DeepEquals(t1 any, t2 any) {
	Require(reflect.DeepEqual(t1, t2), "assert failed, expects %v deep equals %v, but not", t1, t2)
}
