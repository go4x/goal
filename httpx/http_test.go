package httpx

import (
	"net/url"
	"strings"
	"testing"
)

func TestPackageLevelMethods(t *testing.T) {
	// Test Get method
	t.Run("Get", func(t *testing.T) {
		params := url.Values{}
		params.Set("test", "value")

		resp, err := Get("https://httpbin.org/get", params)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	// Test Get without params
	t.Run("Get without params", func(t *testing.T) {
		resp, err := Get("https://httpbin.org/get", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelGetWithBody(t *testing.T) {
	t.Run("GetWithBody", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"test": "value"}`)

		resp, err := GetWithBody("https://httpbin.org/get", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("GetWithBody with nil body", func(t *testing.T) {
		resp, err := GetWithBody("https://httpbin.org/get", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelGetJson(t *testing.T) {
	t.Run("GetJson", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"query": "test"}`)

		resp, err := GetJson("https://httpbin.org/get", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelGetForm(t *testing.T) {
	t.Run("GetForm", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("field1", "value1")
		formData.Set("field2", "value2")
		formBody := strings.NewReader(formData.Encode())

		resp, err := GetForm("https://httpbin.org/get", formBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPost(t *testing.T) {
	t.Run("Post", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"test": "data"}`)

		resp, err := Post("https://httpbin.org/post", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("Post with nil body", func(t *testing.T) {
		resp, err := Post("https://httpbin.org/post", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPostJson(t *testing.T) {
	t.Run("PostJson", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"name": "test", "value": 123}`)

		resp, err := PostJson("https://httpbin.org/post", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPostForm(t *testing.T) {
	t.Run("PostForm", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("name", "test")
		formData.Set("value", "123")
		formBody := strings.NewReader(formData.Encode())

		resp, err := PostForm("https://httpbin.org/post", formBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPut(t *testing.T) {
	t.Run("Put", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"id": 1, "name": "updated"}`)

		resp, err := Put("https://httpbin.org/put", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPutJson(t *testing.T) {
	t.Run("PutJson", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"id": 1, "name": "updated"}`)

		resp, err := PutJson("https://httpbin.org/put", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPutForm(t *testing.T) {
	t.Run("PutForm", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("id", "1")
		formData.Set("name", "updated")
		formBody := strings.NewReader(formData.Encode())

		resp, err := PutForm("https://httpbin.org/put", formBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPatch(t *testing.T) {
	t.Run("Patch", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"name": "patched"}`)

		resp, err := Patch("https://httpbin.org/patch", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPatchJson(t *testing.T) {
	t.Run("PatchJson", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"name": "patched"}`)

		resp, err := PatchJson("https://httpbin.org/patch", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelPatchForm(t *testing.T) {
	t.Run("PatchForm", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("name", "patched")
		formBody := strings.NewReader(formData.Encode())

		resp, err := PatchForm("https://httpbin.org/patch", formBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelDelete(t *testing.T) {
	t.Run("Delete", func(t *testing.T) {
		resp, err := Delete("https://httpbin.org/delete")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelDeleteJson(t *testing.T) {
	t.Run("DeleteJson", func(t *testing.T) {
		jsonBody := strings.NewReader(`{"confirmCode": "ABC123"}`)

		resp, err := DeleteJson("https://httpbin.org/delete", jsonBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelDeleteForm(t *testing.T) {
	t.Run("DeleteForm", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("confirmCode", "ABC123")
		formBody := strings.NewReader(formData.Encode())

		resp, err := DeleteForm("https://httpbin.org/delete", formBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

func TestPackageLevelMethods_EdgeCases(t *testing.T) {
	t.Run("nil body handling", func(t *testing.T) {
		resp, err := Post("https://httpbin.org/post", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("empty body handling", func(t *testing.T) {
		emptyBody := strings.NewReader("")
		resp, err := Post("https://httpbin.org/post", emptyBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})

	t.Run("large body handling", func(t *testing.T) {
		// Create a large JSON body
		largeData := strings.Repeat(`{"field": "value", "number": 123},`, 1000)
		largeBody := strings.NewReader(`[` + largeData[:len(largeData)-1] + `]`)

		resp, err := Post("https://httpbin.org/post", largeBody)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer func() { _ = resp.Close() }()

		if !resp.IsSuccess() {
			t.Errorf("Expected success response, got status %d", resp.StatusCode)
		}
	})
}

// Benchmark tests for package-level methods
func BenchmarkPackageLevel_Get(b *testing.B) {
	params := url.Values{}
	params.Set("test", "benchmark")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := Get("https://httpbin.org/get", params)
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Close()
	}
}

func BenchmarkPackageLevel_PostJson(b *testing.B) {
	jsonBody := strings.NewReader(`{"test": "benchmark"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := PostJson("https://httpbin.org/post", jsonBody)
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Close()
	}
}

func BenchmarkPackageLevel_Delete(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := Delete("https://httpbin.org/delete")
		if err != nil {
			b.Fatal(err)
		}
		_ = resp.Close()
	}
}
