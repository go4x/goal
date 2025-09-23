package jsonx_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go4x/goal/jsonx"
	"github.com/go4x/got"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	var mp = make(map[string]any)
	var smp = make(map[string]any)
	mp["a"] = 1
	mp["b"] = 2

	smp["s1"] = 11
	smp["s2"] = 22
	mp["c"] = smp

	json, err := jsonx.Marshal(mp)
	assert.NoError(t, err)
	assert.Equal(t, json, "{\"a\":1,\"b\":2,\"c\":{\"s1\":11,\"s2\":22}}")
}

type User struct {
	Name      string    // `json:"name"`
	Age       int       // `json:"age"`
	BirthDate time.Time // `json:"birthDate"`
	Other     float64   // 其他字段，科学计数法的情况
}

func TestMarshal1(t *testing.T) {
	var user = User{
		Name:      "hank",
		Age:       20,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		Other:     171.65,
	}
	json, err := jsonx.Marshal(user)
	assert.NoError(t, err)
	assert.Equal(t, json, "{\"Name\":\"hank\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00Z\",\"Other\":171.65}")
}

func TestMarshalf(t *testing.T) {
	var user = User{
		Name:      "hank",
		Age:       20,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Other:     171.65,
	}
	json, err := jsonx.Marshal(user, jsonx.Indent)
	assert.NoError(t, err)
	assert.Equal(t, json, "{\n  \"Name\": \"hank\",\n  \"Age\": 20,\n  \"BirthDate\": \"2000-01-01T00:00:00Z\",\n  \"Other\": 171.65\n}")
}

func TestUnmarshal(t *testing.T) {
	s := "{\"Name\":\"hank\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00Z\",\"Other\":171.65}"
	u, err := jsonx.Unmarshal([]byte(s), &User{})
	assert.NoError(t, err)
	assert.Equal(t, u.Name, "hank")
	assert.Equal(t, u.Age, 20)
	assert.Equal(t, u.BirthDate, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, u.Other, 171.65)
}

func TestUnmarshal1(t *testing.T) {
	logger := got.New(t, "test Parse")

	logger.Case("parse user json")
	s := "{\n  \"Name\": \"hank\",\n  \"Age\": 20,\n  \"BirthDate\": \"2000-01-01T00:00:00Z\",\n  \"Other\":17165123123}"
	u, err := jsonx.Unmarshal([]byte(s), &User{})
	assert.NoError(t, err)
	assert.Equal(t, u.Name, "hank")
	assert.Equal(t, u.Age, 20)
	assert.Equal(t, u.BirthDate, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, u.Other, 17165123123.0)
}

func TestUnmarshalUseNumber(t *testing.T) {
	s := "{\"Name\":\"hank\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00Z\",\"Other\":17165123123}"
	var mp = make(map[string]any)
	_, err := jsonx.Unmarshal([]byte(s), &mp, jsonx.UseNumber)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, mp["Name"], "hank")
	ageNumber, ok := mp["Age"].(json.Number)
	assert.True(t, ok)
	age, err := ageNumber.Int64()
	assert.NoError(t, err)
	assert.Equal(t, age, int64(20))

	assert.Equal(t, mp["BirthDate"], "2000-01-01T00:00:00Z")
	otherNumber, ok := mp["Other"].(json.Number)
	assert.True(t, ok)
	other, err := otherNumber.Int64()
	assert.NoError(t, err)
	assert.Equal(t, other, int64(17165123123))
}

func TestUnmarshalDisallowUnknownFields(t *testing.T) {
	logger := got.New(t, "test ParseDisallowUnknownFields")

	logger.Case("add a gender field which will cause error")
	// 增加 gender 属性
	s := "{\"Name\":\"hank\",\"gender\":1,\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00Z\",\"Other\":171.65}"
	_, err := jsonx.Unmarshal([]byte(s), &User{}, jsonx.DisallowUnknownFields)
	assert.True(t, err != nil)
}

// Test new string-based functions
func TestUnmarshalString(t *testing.T) {
	s := "{\"Name\":\"hank\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00Z\",\"Other\":171.65}"
	u, err := jsonx.UnmarshalString(s, &User{})
	assert.NoError(t, err)
	assert.Equal(t, u.Name, "hank")
	assert.Equal(t, u.Age, 20)
	assert.Equal(t, u.BirthDate, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, u.Other, 171.65)
}

func TestMustUnmarshalString(t *testing.T) {
	s := "{\"Name\":\"hank\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00Z\",\"Other\":171.65}"
	u := jsonx.MustUnmarshalString(s, &User{})
	assert.Equal(t, u.Name, "hank")
	assert.Equal(t, u.Age, 20)
	assert.Equal(t, u.BirthDate, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, u.Other, 171.65)
}

func TestMustMarshalWithIndent(t *testing.T) {
	user := User{
		Name:      "hank",
		Age:       20,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Other:     171.65,
	}
	json := jsonx.MustMarshal(user, jsonx.Indent)
	expected := "{\n  \"Name\": \"hank\",\n  \"Age\": 20,\n  \"BirthDate\": \"2000-01-01T00:00:00Z\",\n  \"Other\": 171.65\n}"
	assert.Equal(t, json, expected)
}

// Test edge cases and error conditions
func TestMarshalNil(t *testing.T) {
	json, err := jsonx.Marshal(nil)
	assert.NoError(t, err)
	assert.Equal(t, json, "null")
}

func TestMarshalEmptyMap(t *testing.T) {
	json, err := jsonx.Marshal(map[string]interface{}{})
	assert.NoError(t, err)
	assert.Equal(t, json, "{}")
}

func TestMarshalEmptySlice(t *testing.T) {
	json, err := jsonx.Marshal([]interface{}{})
	assert.NoError(t, err)
	assert.Equal(t, json, "[]")
}

func TestUnmarshalInvalidJSON(t *testing.T) {
	_, err := jsonx.Unmarshal([]byte("invalid json"), &User{})
	assert.Error(t, err)
}

func TestUnmarshalStringInvalidJSON(t *testing.T) {
	_, err := jsonx.UnmarshalString("invalid json", &User{})
	assert.Error(t, err)
}

func TestMustUnmarshalInvalidJSON(t *testing.T) {
	assert.Panics(t, func() {
		jsonx.MustUnmarshal([]byte("invalid json"), &User{})
	})
}

func TestMustUnmarshalStringInvalidJSON(t *testing.T) {
	assert.Panics(t, func() {
		jsonx.MustUnmarshalString("invalid json", &User{})
	})
}

func TestUnmarshalToWrongType(t *testing.T) {
	s := "{\"Name\":\"hank\",\"Age\":\"invalid_age\"}"
	user, err := jsonx.Unmarshal([]byte(s), &User{})
	assert.Error(t, err)
	assert.True(t, user == nil)
}

// Test complex nested structures
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address Address  `json:"address"`
	Hobbies []string `json:"hobbies"`
}

func TestMarshalComplexStructure(t *testing.T) {
	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
		},
		Hobbies: []string{"reading", "swimming", "coding"},
	}

	json, err := jsonx.Marshal(person)
	assert.NoError(t, err)

	// Unmarshal back and verify
	result, err := jsonx.Unmarshal([]byte(json), &Person{})
	assert.NoError(t, err)
	assert.Equal(t, person.Name, result.Name)
	assert.Equal(t, person.Age, result.Age)
	assert.Equal(t, person.Address.Street, result.Address.Street)
	assert.Equal(t, person.Address.City, result.Address.City)
	assert.Equal(t, person.Hobbies, result.Hobbies)
}

func TestMarshalWithMultipleOptions(t *testing.T) {
	data := map[string]interface{}{
		"name":  "test",
		"value": 42,
	}

	json, err := jsonx.Marshal(data, jsonx.Indent)
	assert.NoError(t, err)

	// Should be pretty printed
	assert.Contains(t, json, "\n")
	assert.Contains(t, json, "  ")
}

func TestUnmarshalWithMultipleOptions(t *testing.T) {
	s := "{\"name\":\"test\",\"age\":25,\"unknown_field\":\"should_cause_error\"}"

	// Test with UseNumber only - should work
	result1, err := jsonx.Unmarshal([]byte(s), &map[string]interface{}{}, jsonx.UseNumber)
	assert.NoError(t, err)
	assert.Equal(t, "test", (*result1)["name"])

	// Test with DisallowUnknownFields - should fail
	var result2 User
	_, err = jsonx.Unmarshal([]byte(s), &result2, jsonx.DisallowUnknownFields)
	assert.Error(t, err)
}

// Test empty string handling
func TestUnmarshalEmptyString(t *testing.T) {
	_, err := jsonx.Unmarshal([]byte(""), &User{})
	assert.Error(t, err)
}

func TestUnmarshalStringEmptyString(t *testing.T) {
	_, err := jsonx.UnmarshalString("", &User{})
	assert.Error(t, err)
}
