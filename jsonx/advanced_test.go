package jsonx_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/go4x/goal/jsonx"
	"github.com/stretchr/testify/assert"
)

// Test streaming functions
func TestMarshalStream(t *testing.T) {
	data := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"tags": []string{"golang", "developer"},
	}

	var buf bytes.Buffer
	err := jsonx.MarshalStream(&buf, data, jsonx.Indent)
	assert.NoError(t, err)

	result := buf.String()
	assert.Contains(t, result, "John Doe")
	assert.Contains(t, result, "\n")
	assert.Contains(t, result, "  ")
}

func TestUnmarshalStream(t *testing.T) {
	jsonStr := `{"name":"Jane Doe","age":25,"email":"jane@example.com"}`
	reader := strings.NewReader(jsonStr)

	var user map[string]interface{}
	result, err := jsonx.UnmarshalStream(reader, &user)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", (*result)["name"])
	assert.Equal(t, float64(25), (*result)["age"])
}

// Test formatting options
func TestIndentWith(t *testing.T) {
	data := map[string]interface{}{
		"name":  "Test",
		"value": 42,
	}

	jsonStr, err := jsonx.Marshal(data, jsonx.IndentWith("\t", "\t"))
	assert.NoError(t, err)

	// Should contain tabs
	assert.Contains(t, jsonStr, "\t")
	assert.Contains(t, jsonStr, "\t\"name\"")
}

func TestCompact(t *testing.T) {
	data := map[string]interface{}{
		"name":  "Test",
		"value": 42,
	}

	// First create indented JSON
	indentedJSON, err := jsonx.Marshal(data, jsonx.Indent)
	assert.NoError(t, err)
	assert.Contains(t, indentedJSON, "\n")

	// Then compact it
	compactJSON, err := jsonx.Marshal(data, jsonx.Compact)
	assert.NoError(t, err)
	assert.NotContains(t, compactJSON, "\n")
	assert.NotContains(t, compactJSON, "  ")
}

func TestSortKeys(t *testing.T) {
	data := map[string]interface{}{
		"zebra":  "last",
		"apple":  "first",
		"banana": "middle",
	}

	jsonStr, err := jsonx.Marshal(data, jsonx.SortKeys)
	assert.NoError(t, err)

	// Keys should be sorted alphabetically
	applePos := strings.Index(jsonStr, "apple")
	bananaPos := strings.Index(jsonStr, "banana")
	zebraPos := strings.Index(jsonStr, "zebra")

	assert.True(t, applePos < bananaPos)
	assert.True(t, bananaPos < zebraPos)
}

// Test JSONPath functions
func TestGetPath(t *testing.T) {
	jsonData := []byte(`{
		"users": [
			{"name": "John", "age": 30},
			{"name": "Jane", "age": 25}
		],
		"settings": {
			"theme": "dark",
			"language": "en"
		}
	}`)

	// Test simple key access
	theme, err := jsonx.GetPath(jsonData, "settings.theme")
	assert.NoError(t, err)
	assert.Equal(t, "dark", theme)

	// Test array index access
	userName, err := jsonx.GetPath(jsonData, "users[0].name")
	assert.NoError(t, err)
	assert.Equal(t, "John", userName)

	// Test wildcard access
	allUsers, err := jsonx.GetPath(jsonData, "users[*]")
	assert.NoError(t, err)
	users, ok := allUsers.([]interface{})
	assert.True(t, ok)
	assert.Len(t, users, 2)

	// Test nested access
	firstUserAge, err := jsonx.GetPath(jsonData, "users[0].age")
	assert.NoError(t, err)
	assert.Equal(t, float64(30), firstUserAge)
}

func TestSetPath(t *testing.T) {
	jsonData := []byte(`{
		"users": [
			{"name": "John", "age": 30}
		],
		"settings": {
			"theme": "light"
		}
	}`)

	// Test setting a simple value
	newJSON, err := jsonx.SetPath(jsonData, "settings.theme", "dark")
	assert.NoError(t, err)

	var result map[string]interface{}
	_, err = jsonx.Unmarshal(newJSON, &result)
	assert.NoError(t, err)

	settings := result["settings"].(map[string]interface{})
	assert.Equal(t, "dark", settings["theme"])

	// Test setting array element
	newJSON2, err := jsonx.SetPath(jsonData, "users[0].age", 31)
	assert.NoError(t, err)

	var result2 map[string]interface{}
	_, err = jsonx.Unmarshal(newJSON2, &result2)
	assert.NoError(t, err)

	users := result2["users"].([]interface{})
	user := users[0].(map[string]interface{})
	assert.Equal(t, float64(31), user["age"])
}

func TestGetPathErrors(t *testing.T) {
	jsonData := []byte(`{"users": [{"name": "John"}]}`)

	// Test array index out of range
	_, err := jsonx.GetPath(jsonData, "users[5]")
	assert.Error(t, err)

	// Test accessing non-array with index
	_, err = jsonx.GetPath(jsonData, "users[0].name[0]")
	assert.Error(t, err)

	// Test accessing non-existent nested key (this returns nil, not error)
	result, err := jsonx.GetPath(jsonData, "users[0].nonexistent")
	assert.NoError(t, err)
	assert.Nil(t, result)
}

// Test type conversion functions
func TestConvert(t *testing.T) {
	// Test compatible conversions - numbers to numbers
	num, err := jsonx.Convert[float64](123)
	assert.NoError(t, err)
	assert.Equal(t, 123.0, num)

	// Test int to float
	floatFromInt, err := jsonx.Convert[float64](456)
	assert.NoError(t, err)
	assert.Equal(t, 456.0, floatFromInt)

	// Test compatible struct conversion

	// Test struct conversion
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := map[string]interface{}{
		"name": "Alice",
		"age":  30,
	}

	person, err := jsonx.Convert[Person](data)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", person.Name)
	assert.Equal(t, 30, person.Age)
}

func TestForceConvert(t *testing.T) {
	// Test successful conversion
	num := jsonx.ForceConvert[float64](123)
	assert.Equal(t, 123.0, num)

	// Test conversion that would fail
	assert.Panics(t, func() {
		jsonx.ForceConvert[int]("not a number")
	})
}

func TestConvertSlice(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5.0}

	numbers, err := jsonx.ConvertSlice[float64](data)
	assert.NoError(t, err)
	assert.Equal(t, []float64{1, 2, 3, 4, 5}, numbers)

	// Test with mixed types that can't convert
	mixedData := []interface{}{"1", "invalid", 3}
	_, err = jsonx.ConvertSlice[int](mixedData)
	assert.Error(t, err)
}

func TestConvertMap(t *testing.T) {
	data := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3.0,
	}

	numbers, err := jsonx.ConvertMap[float64](data)
	assert.NoError(t, err)
	assert.Equal(t, map[string]float64{"a": 1.0, "b": 2.0, "c": 3.0}, numbers)
}

// Test complex scenarios
func TestComplexJSONPath(t *testing.T) {
	jsonData := []byte(`{
		"company": {
			"employees": [
				{
					"id": 1,
					"name": "Alice",
					"department": "Engineering",
					"skills": ["Go", "Python", "JavaScript"]
				},
				{
					"id": 2,
					"name": "Bob",
					"department": "Marketing",
					"skills": ["Photoshop", "Writing"]
				}
			],
			"departments": {
				"Engineering": {"budget": 1000000},
				"Marketing": {"budget": 500000}
			}
		}
	}`)

	// Test nested path access
	firstEmployeeName, err := jsonx.GetPath(jsonData, "company.employees[0].name")
	assert.NoError(t, err)
	assert.Equal(t, "Alice", firstEmployeeName)

	// Test department budget access
	engBudget, err := jsonx.GetPath(jsonData, "company.departments.Engineering.budget")
	assert.NoError(t, err)
	assert.Equal(t, float64(1000000), engBudget)

	// Test setting nested value
	newJSON, err := jsonx.SetPath(jsonData, "company.employees[0].name", "Alice Smith")
	assert.NoError(t, err)

	updatedName, err := jsonx.GetPath(newJSON, "company.employees[0].name")
	assert.NoError(t, err)
	assert.Equal(t, "Alice Smith", updatedName)
}

func TestMultipleOptions(t *testing.T) {
	data := map[string]interface{}{
		"zebra":  "last",
		"apple":  "first",
		"banana": "middle",
	}

	// Test combining multiple options
	jsonStr, err := jsonx.Marshal(data, jsonx.SortKeys, jsonx.IndentWith("  ", "  "))
	assert.NoError(t, err)

	// Should be sorted and indented
	assert.Contains(t, jsonStr, "\n")
	assert.Contains(t, jsonStr, "  ")

	// Keys should be in alphabetical order
	lines := strings.Split(jsonStr, "\n")
	var keys []string
	for _, line := range lines {
		if strings.Contains(line, ":") {
			key := strings.TrimSpace(strings.Split(line, ":")[0])
			key = strings.Trim(key, `"`)
			if key != "" {
				keys = append(keys, key)
			}
		}
	}

	assert.Equal(t, []string{"apple", "banana", "zebra"}, keys)
}

func TestStreamingWithOptions(t *testing.T) {
	data := map[string]interface{}{
		"message":   "Hello, World!",
		"timestamp": time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		"count":     42,
	}

	var buf bytes.Buffer
	err := jsonx.MarshalStream(&buf, data, jsonx.SortKeys, jsonx.Indent)
	assert.NoError(t, err)

	result := buf.String()
	assert.Contains(t, result, "Hello, World!")
	assert.Contains(t, result, "\n")
	assert.Contains(t, result, "  ")

	// Keys should be sorted
	countPos := strings.Index(result, "count")
	messagePos := strings.Index(result, "message")
	timestampPos := strings.Index(result, "timestamp")

	assert.True(t, countPos < messagePos)
	assert.True(t, messagePos < timestampPos)
}

func TestNestedArraySupport(t *testing.T) {
	jsonData := []byte(`{
		"matrix": [
			[1, 2, 3],
			[4, 5, 6],
			[7, 8, 9]
		],
		"users": [
			{
				"name": "John",
				"skills": ["Go", "Python", "JavaScript"]
			}
		]
	}`)

	// Test nested array access
	value, err := jsonx.GetPath(jsonData, "matrix[0][1]")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), value)

	// Test mixed object and array access
	skill, err := jsonx.GetPath(jsonData, "users[0].skills[1]")
	assert.NoError(t, err)
	assert.Equal(t, "Python", skill)

	// Test setting nested array values
	newJSON, err := jsonx.SetPath(jsonData, "matrix[0][1]", 99)
	assert.NoError(t, err)

	updatedValue, err := jsonx.GetPath(newJSON, "matrix[0][1]")
	assert.NoError(t, err)
	assert.Equal(t, float64(99), updatedValue)
}
