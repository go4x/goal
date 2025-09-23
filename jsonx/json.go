// Package jsonx provides enhanced JSON marshaling and unmarshaling utilities.
// It extends the standard library's encoding/json package with additional features
// such as pretty printing, number handling, and strict field validation.
//
// The package offers both safe functions that return errors and panic functions
// that panic on failure, following Go naming conventions.
//
// Example usage:
//
//	// Basic marshaling
//	jsonStr, err := jsonx.Marshal(data)
//
//	// Pretty printed marshaling
//	prettyJSON, err := jsonx.Marshal(data, jsonx.Indent)
//
//	// Safe unmarshaling with options
//	result, err := jsonx.Unmarshal([]byte(jsonStr), &target, jsonx.UseNumber)
//
//	// Panic on failure (for cases where you're certain it will succeed)
//	result := jsonx.MustUnmarshal([]byte(jsonStr), &target)
package jsonx

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

// MarshalOption is a function type that can be used to modify JSON bytes after marshaling.
// Common options include pretty printing with Indent.
type MarshalOption func(bs []byte) []byte

// UnmarshalOption is a function type that can be used to configure the JSON decoder.
// Common options include UseNumber and DisallowUnknownFields.
type UnmarshalOption func(d *json.Decoder)

// Indent returns a MarshalOption that pretty-prints JSON with 2-space indentation.
// It panics if the input is not valid JSON.
//
// Example:
//
//	prettyJSON, err := jsonx.Marshal(data, jsonx.Indent)
func Indent(bs []byte) []byte {
	var buf bytes.Buffer
	err := json.Indent(&buf, bs, "", "  ")
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// IndentWith returns a MarshalOption that pretty-prints JSON with custom prefix and indent.
// It panics if the input is not valid JSON.
//
// Example:
//
//	prettyJSON, err := jsonx.Marshal(data, jsonx.IndentWith("\t", "\t"))
func IndentWith(prefix, indent string) MarshalOption {
	return func(bs []byte) []byte {
		var buf bytes.Buffer
		err := json.Indent(&buf, bs, prefix, indent)
		if err != nil {
			panic(err)
		}
		return buf.Bytes()
	}
}

// Compact returns a MarshalOption that compacts JSON by removing unnecessary whitespace.
// It panics if the input is not valid JSON.
//
// Example:
//
//	compactJSON, err := jsonx.Marshal(data, jsonx.Compact)
func Compact(bs []byte) []byte {
	var buf bytes.Buffer
	err := json.Compact(&buf, bs)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// SortKeys returns a MarshalOption that sorts JSON object keys alphabetically.
// It panics if the input is not valid JSON.
//
// Example:
//
//	sortedJSON, err := jsonx.Marshal(data, jsonx.SortKeys)
func SortKeys(bs []byte) []byte {
	var data interface{}
	err := json.Unmarshal(bs, &data)
	if err != nil {
		panic(err)
	}

	sorted, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return sorted
}

// UseNumber returns an UnmarshalOption that instructs the decoder to unmarshal
// numbers into json.Number instead of float64. This preserves precision for
// large integers and prevents scientific notation for floats.
//
// Example:
//
//	result, err := jsonx.Unmarshal([]byte(jsonStr), &target, jsonx.UseNumber)
func UseNumber(d *json.Decoder) {
	d.UseNumber()
}

// DisallowUnknownFields returns an UnmarshalOption that causes the decoder to
// return an error if the JSON contains unknown fields that don't correspond
// to any exported or non-ignored fields in the destination.
//
// Example:
//
//	result, err := jsonx.Unmarshal([]byte(jsonStr), &target, jsonx.DisallowUnknownFields)
func DisallowUnknownFields(d *json.Decoder) {
	d.DisallowUnknownFields()
}

// Marshal returns the JSON encoding of v as a string.
// It applies the provided options to modify the output (e.g., pretty printing).
//
// Example:
//
//	jsonStr, err := jsonx.Marshal(data)
//	prettyJSON, err := jsonx.Marshal(data, jsonx.Indent)
func Marshal(v any, options ...MarshalOption) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	for _, option := range options {
		bs = option(bs)
	}
	return string(bs), nil
}

// MustMarshal is like Marshal but panics if marshaling fails.
// It applies the provided options to modify the output.
//
// Example:
//
//	jsonStr := jsonx.MustMarshal(data)
//	prettyJSON := jsonx.MustMarshal(data, jsonx.Indent)
func MustMarshal(v any, options ...MarshalOption) string {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	for _, option := range options {
		bs = option(bs)
	}
	return string(bs)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// It applies the provided options to configure the decoder.
//
// Example:
//
//	var user User
//	result, err := jsonx.Unmarshal([]byte(jsonStr), &user)
//	result, err := jsonx.Unmarshal([]byte(jsonStr), &user, jsonx.UseNumber)
func Unmarshal[T any](data []byte, v *T, options ...UnmarshalOption) (*T, error) {
	err := createDecoder(data, options).Decode(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// UnmarshalString is like Unmarshal but accepts a string instead of []byte.
//
// Example:
//
//	var user User
//	result, err := jsonx.UnmarshalString(jsonStr, &user)
func UnmarshalString[T any](data string, v *T, options ...UnmarshalOption) (*T, error) {
	return Unmarshal([]byte(data), v, options...)
}

// MustUnmarshal is like Unmarshal but panics if unmarshaling fails.
// It applies the provided options to configure the decoder.
//
// Example:
//
//	var user User
//	result := jsonx.MustUnmarshal([]byte(jsonStr), &user)
func MustUnmarshal[T any](data []byte, v *T, options ...UnmarshalOption) *T {
	err := createDecoder(data, options).Decode(v)
	if err != nil {
		panic(err)
	}
	return v
}

// MustUnmarshalString is like MustUnmarshal but accepts a string instead of []byte.
//
// Example:
//
//	var user User
//	result := jsonx.MustUnmarshalString(jsonStr, &user)
func MustUnmarshalString[T any](data string, v *T, options ...UnmarshalOption) *T {
	return MustUnmarshal([]byte(data), v, options...)
}

// GetPath retrieves a value from JSON data using a JSONPath-like syntax.
// Supported paths:
//   - "key" - simple key access
//   - "key.subkey" - nested key access
//   - "array[0]" - array index access
//   - "array[*]" - wildcard access (returns entire array)
//   - "matrix[0][1]" - nested array access
//   - "users[0].skills[1]" - mixed object and array access
//
// Example:
//
//	value, err := jsonx.GetPath([]byte(jsonStr), "users[0].name")
//	allNames, err := jsonx.GetPath([]byte(jsonStr), "users[*].name")
//	nestedValue, err := jsonx.GetPath([]byte(jsonStr), "matrix[0][1]")
//	deepValue, err := jsonx.GetPath([]byte(jsonStr), "company.departments[0].teams[0].members[1]")
func GetPath(data []byte, path string) (interface{}, error) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	return getValueByPath(jsonData, path)
}

// SetPath sets a value in JSON data using a JSONPath-like syntax.
// Returns the modified JSON as bytes.
// Supports the same path syntax as GetPath.
//
// Example:
//
//	newJSON, err := jsonx.SetPath([]byte(jsonStr), "users[0].name", "New Name")
//	newJSON, err := jsonx.SetPath([]byte(jsonStr), "matrix[0][1]", 99)
//	newJSON, err := jsonx.SetPath([]byte(jsonStr), "company.departments[0].teams[0].members[1]", "Robert")
func SetPath(data []byte, path string, value interface{}) ([]byte, error) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	err = setValueByPath(jsonData, path, value)
	if err != nil {
		return nil, err
	}

	return json.Marshal(jsonData)
}

// getValueByPath retrieves a value from a JSON object using path
func getValueByPath(data interface{}, path string) (interface{}, error) {
	parts := parsePath(path)
	current := data

	for _, part := range parts {
		switch p := part.(type) {
		case string:
			// Object key access
			if obj, ok := current.(map[string]interface{}); ok {
				current = obj[p]
			} else {
				return nil, &PathError{Path: path, Message: "not an object"}
			}
		case int:
			// Array index access
			if arr, ok := current.([]interface{}); ok {
				if p >= 0 && p < len(arr) {
					current = arr[p]
				} else {
					return nil, &PathError{Path: path, Message: "array index out of range"}
				}
			} else {
				return nil, &PathError{Path: path, Message: "not an array"}
			}
		case wildcard:
			// Wildcard access - return array of values
			if arr, ok := current.([]interface{}); ok {
				return arr, nil
			} else {
				return nil, &PathError{Path: path, Message: "not an array"}
			}
		}
	}

	return current, nil
}

// setValueByPath sets a value in a JSON object using path
func setValueByPath(data interface{}, path string, value interface{}) error {
	parts := parsePath(path)
	if len(parts) == 0 {
		return &PathError{Path: path, Message: "empty path"}
	}

	current := data

	// Navigate to the parent of the target
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		switch p := part.(type) {
		case string:
			if obj, ok := current.(map[string]interface{}); ok {
				if _, exists := obj[p]; !exists {
					obj[p] = make(map[string]interface{})
				}
				current = obj[p]
			} else {
				return &PathError{Path: path, Message: "not an object"}
			}
		case int:
			if arr, ok := current.([]interface{}); ok {
				if p >= 0 && p < len(arr) {
					current = arr[p]
				} else {
					return &PathError{Path: path, Message: "array index out of range"}
				}
			} else {
				return &PathError{Path: path, Message: "not an array"}
			}
		default:
			return &PathError{Path: path, Message: "invalid path segment"}
		}
	}

	// Set the value
	lastPart := parts[len(parts)-1]
	switch p := lastPart.(type) {
	case string:
		if obj, ok := current.(map[string]interface{}); ok {
			obj[p] = value
		} else {
			return &PathError{Path: path, Message: "not an object"}
		}
	case int:
		if arr, ok := current.([]interface{}); ok {
			if p >= 0 && p < len(arr) {
				arr[p] = value
			} else {
				return &PathError{Path: path, Message: "array index out of range"}
			}
		} else {
			return &PathError{Path: path, Message: "not an array"}
		}
	default:
		return &PathError{Path: path, Message: "invalid path segment"}
	}

	return nil
}

type wildcard struct{}

// parsePath parses a JSONPath-like string into path segments
func parsePath(path string) []interface{} {
	var parts []interface{}

	// Enhanced parser that supports nested arrays like array[0][1]
	// Supports: key, key.subkey, array[0], array[*], array[0][1], key[0][1].subkey
	segments := strings.Split(path, ".")

	for _, segment := range segments {
		// Handle multiple array indices in one segment like "matrix[0][1]"
		if strings.Contains(segment, "[") && strings.Contains(segment, "]") {
			// Extract the key part (everything before the first [)
			keyPart := segment[:strings.Index(segment, "[")]
			if keyPart != "" {
				parts = append(parts, keyPart)
			}

			// Extract all array indices from the segment
			remaining := segment[strings.Index(segment, "["):]
			for len(remaining) > 0 && remaining[0] == '[' {
				// Find the closing ]
				closeIndex := strings.Index(remaining, "]")
				if closeIndex == -1 {
					break // Invalid syntax
				}

				// Extract the index
				indexStr := remaining[1:closeIndex]
				if indexStr == "*" {
					parts = append(parts, wildcard{})
				} else {
					if index, err := strconv.Atoi(indexStr); err == nil {
						parts = append(parts, index)
					}
				}

				// Move to the next array index
				remaining = remaining[closeIndex+1:]
			}
		} else {
			// Object key
			parts = append(parts, segment)
		}
	}

	return parts
}

// PathError represents an error in JSON path operations
type PathError struct {
	Path    string
	Message string
}

func (e *PathError) Error() string {
	return "jsonx path error at '" + e.Path + "': " + e.Message
}

// Convert converts data to the specified type T.
// It uses JSON marshaling/unmarshaling for type conversion.
//
// Example:
//
//	value := "123"
//	num, err := jsonx.Convert[int](value)
//	str, err := jsonx.Convert[string](42)
func Convert[T any](data interface{}) (T, error) {
	var result T

	// Marshal the input data
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	// Unmarshal into the target type
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// ForceConvert converts data to the specified type T, ignoring precision loss.
// It panics if conversion fails.
//
// Example:
//
//	value := "123"
//	num := jsonx.ForceConvert[int](value)
//	str := jsonx.ForceConvert[string](42)
func ForceConvert[T any](data interface{}) T {
	result, err := Convert[T](data)
	if err != nil {
		panic("jsonx: force convert failed: " + err.Error())
	}
	return result
}

// ConvertSlice converts a slice of data to a slice of type T.
//
// Example:
//
//	values := []interface{}{1, 2, 3, "4", 5}
//	numbers, err := jsonx.ConvertSlice[int](values)
func ConvertSlice[T any](data []interface{}) ([]T, error) {
	var result []T

	for _, item := range data {
		converted, err := Convert[T](item)
		if err != nil {
			return nil, err
		}
		result = append(result, converted)
	}

	return result, nil
}

// ConvertMap converts a map[string]interface{} to a map[string]T.
//
// Example:
//
//	data := map[string]interface{}{"a": 1, "b": "2", "c": 3.0}
//	numbers, err := jsonx.ConvertMap[int](data)
func ConvertMap[T any](data map[string]interface{}) (map[string]T, error) {
	result := make(map[string]T)

	for key, value := range data {
		converted, err := Convert[T](value)
		if err != nil {
			return nil, err
		}
		result[key] = converted
	}

	return result, nil
}

// createDecoder creates a JSON decoder with the given options applied.
func createDecoder(data []byte, options []UnmarshalOption) *json.Decoder {
	d := json.NewDecoder(bytes.NewReader(data))
	for _, option := range options {
		option(d)
	}
	return d
}

// MarshalStream marshals v to JSON and writes it to w.
// It applies the provided options to modify the output.
//
// Example:
//
//	var buf bytes.Buffer
//	err := jsonx.MarshalStream(&buf, data, jsonx.Indent)
//	if err != nil {
//	    log.Fatal(err)
//	}
func MarshalStream(w io.Writer, v any, options ...MarshalOption) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	for _, option := range options {
		bs = option(bs)
	}
	_, err = w.Write(bs)
	return err
}

// UnmarshalStream reads JSON from r and unmarshals it into v.
// It applies the provided options to configure the decoder.
//
// Example:
//
//	var user User
//	err := jsonx.UnmarshalStream(reader, &user, jsonx.UseNumber)
//	if err != nil {
//	    log.Fatal(err)
//	}
func UnmarshalStream[T any](r io.Reader, v *T, options ...UnmarshalOption) (*T, error) {
	d := json.NewDecoder(r)
	for _, option := range options {
		option(d)
	}
	err := d.Decode(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
