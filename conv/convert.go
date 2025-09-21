// Package conv provides type conversion utilities for Go.
//
// This package offers safe and efficient conversion functions between different
// data types, including string-to-number conversions, number-to-string conversions,
// and hexadecimal conversions. All functions follow Go language best practices
// with proper error handling.
//
// The package provides two types of functions:
//  1. Standard conversion functions that return (result, error) - for explicit error handling
//  2. Safe conversion functions that return default values on error - for convenience
//
// Example usage:
//
//	// Standard conversion with error handling
//	value, err := conv.StrToInt("123")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Safe conversion with default value
//	value := conv.StrToIntSafe("abc", 0) // returns 0 on error
//
//	// Never-failing conversions
//	str := conv.IntToStr(123) // always succeeds
package conv

import (
	"strconv"
)

// StrToInt converts string to int.
// Returns an error if the string cannot be parsed as an integer.
//
// Example:
//
//	value, err := conv.StrToInt("123")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(value) // Output: 123
func StrToInt(str string) (int, error) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// StrToInt32 converts string to int32.
// Returns an error if the string cannot be parsed as a 32-bit integer.
func StrToInt32(str string) (int32, error) {
	n, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

// StrToInt64 converts string to int64.
// Returns an error if the string cannot be parsed as a 64-bit integer.
func StrToInt64(str string) (int64, error) {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// IntToStr converts int to string.
// This function never fails and always returns a valid string representation.
func IntToStr(src int) string {
	return strconv.Itoa(src)
}

// Int64ToStr converts int64 to string.
// This function never fails and always returns a valid string representation.
func Int64ToStr(src int64) string {
	return strconv.FormatInt(src, 10)
}

// Int32ToStr converts int32 to string.
// This function never fails and always returns a valid string representation.
func Int32ToStr(src int32) string {
	return strconv.FormatInt(int64(src), 10)
}

// StrToUint converts string to uint.
// Returns an error if the string cannot be parsed as an unsigned integer.
func StrToUint(str string) (uint, error) {
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}

// StrToUint64 converts string to uint64.
// Returns an error if the string cannot be parsed as a 64-bit unsigned integer.
func StrToUint64(str string) (uint64, error) {
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// StrToUint32 converts string to uint32.
// Returns an error if the string cannot be parsed as a 32-bit unsigned integer.
func StrToUint32(str string) (uint32, error) {
	n, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(n), nil
}

// UintToStr converts uint to string.
// This function never fails and always returns a valid string representation.
func UintToStr(src uint) string {
	return strconv.FormatUint(uint64(src), 10)
}

// Uint64ToStr converts uint64 to string.
// This function never fails and always returns a valid string representation.
func Uint64ToStr(src uint64) string {
	return strconv.FormatUint(src, 10)
}

// Uint32ToStr converts uint32 to string.
// This function never fails and always returns a valid string representation.
func Uint32ToStr(src uint32) string {
	return strconv.FormatUint(uint64(src), 10)
}

// StrToFloat64 converts string to float64.
// Returns an error if the string cannot be parsed as a floating-point number.
//
// Example:
//
//	value, err := conv.StrToFloat64("3.14159")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(value) // Output: 3.14159
func StrToFloat64(amount string) (float64, error) {
	float, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}
	return float, nil
}

// Int64ToHex converts int64 to hexadecimal string.
// This function never fails and always returns a valid hexadecimal representation.
func Int64ToHex(src int64) string {
	return strconv.FormatInt(src, 16)
}

// HexToInt64 converts hexadecimal string to int64.
// Returns an error if the string cannot be parsed as a hexadecimal number.
//
// Example:
//
//	value, err := conv.HexToInt64("FF")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(value) // Output: 255
func HexToInt64(src string) (int64, error) {
	id, err := strconv.ParseInt(src, 16, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// IntsToStr converts an int64 slice to a string slice.
// This function never fails and always returns a valid string slice.
func IntsToStr(is []int64) []string {
	if len(is) == 0 {
		return nil
	}
	var ss = make([]string, len(is))
	for i, v := range is {
		ss[i] = strconv.FormatInt(v, 10)
	}
	return ss
}

// StrsToInt converts a string slice to an int64 slice.
// Returns an error if any string in the slice cannot be parsed as an integer.
//
// Example:
//
//	values, err := conv.StrsToInt([]string{"1", "2", "3"})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(values) // Output: [1 2 3]
func StrsToInt(ss []string) ([]int64, error) {
	if len(ss) == 0 {
		return nil, nil
	}
	var is = make([]int64, 0, len(ss))
	for _, s := range ss {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		is = append(is, i)
	}
	return is, nil
}

// StrToBool converts string to bool.
// Returns an error if the string cannot be parsed as a boolean value.
//
// Example:
//
//	value, err := conv.StrToBool("true")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(value) // Output: true
func StrToBool(str string) (bool, error) {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return b, nil
}

// Safe version functions - return default values on conversion failure

// StrToIntSafe converts string to int, returns default value on error.
// This is a safe version that never panics and always returns a valid integer.
//
// Example:
//
//	value := conv.StrToIntSafe("123", 0)  // returns 123
//	value := conv.StrToIntSafe("abc", 0)  // returns 0 (default value)
func StrToIntSafe(str string, defaultValue int) int {
	if n, err := StrToInt(str); err == nil {
		return n
	}
	return defaultValue
}

// StrToInt32Safe converts string to int32, returns default value on error.
// This is a safe version that never panics and always returns a valid int32.
func StrToInt32Safe(str string, defaultValue int32) int32 {
	if n, err := StrToInt32(str); err == nil {
		return n
	}
	return defaultValue
}

// StrToInt64Safe converts string to int64, returns default value on error.
// This is a safe version that never panics and always returns a valid int64.
func StrToInt64Safe(str string, defaultValue int64) int64 {
	if n, err := StrToInt64(str); err == nil {
		return n
	}
	return defaultValue
}

// StrToUintSafe converts string to uint, returns default value on error.
// This is a safe version that never panics and always returns a valid uint.
func StrToUintSafe(str string, defaultValue uint) uint {
	if n, err := StrToUint(str); err == nil {
		return n
	}
	return defaultValue
}

// StrToUint64Safe converts string to uint64, returns default value on error.
// This is a safe version that never panics and always returns a valid uint64.
func StrToUint64Safe(str string, defaultValue uint64) uint64 {
	if n, err := StrToUint64(str); err == nil {
		return n
	}
	return defaultValue
}

// StrToUint32Safe converts string to uint32, returns default value on error.
// This is a safe version that never panics and always returns a valid uint32.
func StrToUint32Safe(str string, defaultValue uint32) uint32 {
	if n, err := StrToUint32(str); err == nil {
		return n
	}
	return defaultValue
}

// StrToFloat64Safe converts string to float64, returns default value on error.
// This is a safe version that never panics and always returns a valid float64.
func StrToFloat64Safe(amount string, defaultValue float64) float64 {
	if f, err := StrToFloat64(amount); err == nil {
		return f
	}
	return defaultValue
}

// HexToInt64Safe converts hexadecimal string to int64, returns default value on error.
// This is a safe version that never panics and always returns a valid int64.
func HexToInt64Safe(src string, defaultValue int64) int64 {
	if n, err := HexToInt64(src); err == nil {
		return n
	}
	return defaultValue
}

// StrToBoolSafe converts string to bool, returns default value on error.
// This is a safe version that never panics and always returns a valid bool.
func StrToBoolSafe(str string, defaultValue bool) bool {
	if b, err := StrToBool(str); err == nil {
		return b
	}
	return defaultValue
}
