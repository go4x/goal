package jsonx_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go4x/goal/jsonx"
)

// Benchmark data structures
type BenchmarkUser struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	Email       string                 `json:"email"`
	Age         int                    `json:"age"`
	CreatedAt   time.Time              `json:"created_at"`
	IsActive    bool                   `json:"is_active"`
	Score       float64                `json:"score"`
	Tags        []string               `json:"tags"`
	Preferences map[string]interface{} `json:"preferences"`
}

type BenchmarkData struct {
	Users []BenchmarkUser `json:"users"`
	Total int             `json:"total"`
}

func createBenchmarkData() BenchmarkData {
	users := make([]BenchmarkUser, 1000)
	for i := 0; i < 1000; i++ {
		users[i] = BenchmarkUser{
			ID:        i,
			Name:      "User " + string(rune(i%26+'A')),
			Email:     "user" + string(rune(i%26+'A')) + "@example.com",
			Age:       20 + (i % 50),
			CreatedAt: time.Now().Add(-time.Duration(i) * time.Hour),
			IsActive:  i%2 == 0,
			Score:     float64(i) * 0.1,
			Tags:      []string{"tag1", "tag2", "tag3"},
			Preferences: map[string]interface{}{
				"theme":         "dark",
				"lang":          "en",
				"notifications": true,
			},
		}
	}
	return BenchmarkData{
		Users: users,
		Total: 1000,
	}
}

// Benchmark standard library vs jsonx
func BenchmarkStandardMarshal(b *testing.B) {
	data := createBenchmarkData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonxMarshal(b *testing.B) {
	data := createBenchmarkData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStandardUnmarshal(b *testing.B) {
	data := createBenchmarkData()
	jsonBytes, _ := json.Marshal(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result BenchmarkData
		err := json.Unmarshal(jsonBytes, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonxUnmarshal(b *testing.B) {
	data := createBenchmarkData()
	jsonStr, _ := jsonx.Marshal(data)
	jsonBytes := []byte(jsonStr)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Unmarshal(jsonBytes, &BenchmarkData{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark different options
func BenchmarkMarshalWithIndent(b *testing.B) {
	data := createBenchmarkData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(data, jsonx.Indent)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalWithUseNumber(b *testing.B) {
	data := createBenchmarkData()
	jsonStr, _ := jsonx.Marshal(data)
	jsonBytes := []byte(jsonStr)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Unmarshal(jsonBytes, &BenchmarkData{}, jsonx.UseNumber)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalWithDisallowUnknownFields(b *testing.B) {
	data := createBenchmarkData()
	jsonStr, _ := jsonx.Marshal(data)
	jsonBytes := []byte(jsonStr)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Unmarshal(jsonBytes, &BenchmarkData{}, jsonx.DisallowUnknownFields)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark string-based functions
func BenchmarkUnmarshalString(b *testing.B) {
	data := createBenchmarkData()
	jsonStr, _ := jsonx.Marshal(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.UnmarshalString(jsonStr, &BenchmarkData{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMustUnmarshalString(b *testing.B) {
	data := createBenchmarkData()
	jsonStr, _ := jsonx.Marshal(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = jsonx.MustUnmarshalString(jsonStr, &BenchmarkData{})
	}
}

// Benchmark different data sizes
func BenchmarkMarshalSmall(b *testing.B) {
	data := BenchmarkUser{
		ID:        1,
		Name:      "Test User",
		Email:     "test@example.com",
		Age:       25,
		CreatedAt: time.Now(),
		IsActive:  true,
		Score:     95.5,
		Tags:      []string{"tag1"},
		Preferences: map[string]interface{}{
			"theme": "light",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalMedium(b *testing.B) {
	data := createBenchmarkData()
	// Use only first 100 users
	data.Users = data.Users[:100]
	data.Total = 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalLarge(b *testing.B) {
	data := createBenchmarkData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark error cases
func BenchmarkUnmarshalInvalidJSON(b *testing.B) {
	invalidJSON := []byte(`{"invalid": json}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Unmarshal(invalidJSON, &BenchmarkData{})
		if err == nil {
			b.Fatal("expected error")
		}
	}
}

func BenchmarkMarshalNil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalEmptySlice(b *testing.B) {
	data := []interface{}{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalEmptyMap(b *testing.B) {
	data := map[string]interface{}{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := jsonx.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
