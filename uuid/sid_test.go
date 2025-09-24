package uuid

import (
	"regexp"
	"testing"
)

func TestUint64ToBase62(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want string
	}{
		{name: "zero", n: 0, want: "0"},
		{name: "one", n: 1, want: "1"},
		{name: "nine", n: 9, want: "9"},
		{name: "ten", n: 10, want: "a"},
		{name: "thirty-five", n: 35, want: "z"},
		{name: "thirty-six", n: 36, want: "A"},
		{name: "sixty-one", n: 61, want: "Z"},
		{name: "sixty-two", n: 62, want: "10"},
		{name: "one hundred twenty-three", n: 123, want: "1Z"},
		{name: "three thousand eight hundred forty-three", n: 3843, want: "ZZ"},
		{name: "two hundred thirty-eight thousand three hundred twenty-eight", n: 238328, want: "1000"},
		{name: "large number", n: 18446744073709551615, want: "LygHa16ahYG"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uint64ToBase62(tt.n)
			if got != tt.want {
				t.Errorf("uint64ToBase62(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestSid_GenString(t *testing.T) {
	sid := NewSid()

	// Test that we can generate a string ID
	str, err := sid.GenString()
	if err != nil {
		t.Errorf("GenString() error = %v", err)
		return
	}

	if str == "" {
		t.Error("GenString() returned empty string")
	}

	// Test that we can generate multiple different IDs
	str2, err := sid.GenString()
	if err != nil {
		t.Errorf("GenString() error = %v", err)
		return
	}

	if str == str2 {
		t.Error("GenString() returned the same ID twice")
	}
}

func TestSid_GenUint64(t *testing.T) {
	sid := NewSid()

	// Test that we can generate a uint64 ID
	id, err := sid.GenUint64()
	if err != nil {
		t.Errorf("GenUint64() error = %v", err)
		return
	}

	if id == 0 {
		t.Error("GenUint64() returned zero")
	}

	// Test that we can generate multiple different IDs
	id2, err := sid.GenUint64()
	if err != nil {
		t.Errorf("GenUint64() error = %v", err)
		return
	}

	if id == id2 {
		t.Error("GenUint64() returned the same ID twice")
	}
}

// Benchmark tests
func BenchmarkUint64ToBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uint64ToBase62(123456789)
	}
}

func BenchmarkSid_GenString(b *testing.B) {
	sid := NewSid()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sid.GenString()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSid_GenUint64(b *testing.B) {
	sid := NewSid()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sid.GenUint64()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Test UUID functions
func TestUUID(t *testing.T) {
	// Test that UUID generates a valid UUID format
	uuid := UUID()

	// UUID should be 36 characters long (32 hex chars + 4 hyphens)
	if len(uuid) != 36 {
		t.Errorf("UUID() length = %d, want 36", len(uuid))
	}

	// UUID should match the standard format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("UUID() = %s, does not match UUID format", uuid)
	}

	// Test that multiple calls generate different UUIDs
	uuid2 := UUID()
	if uuid == uuid2 {
		t.Error("UUID() returned the same UUID twice")
	}
}

func TestUUID32(t *testing.T) {
	// Test that UUID32 generates a valid 32-character string
	uuid32 := UUID32()

	// UUID32 should be 32 characters long
	if len(uuid32) != 32 {
		t.Errorf("UUID32() length = %d, want 32", len(uuid32))
	}

	// UUID32 should contain only hexadecimal characters
	hexRegex := regexp.MustCompile(`^[0-9a-f]{32}$`)
	if !hexRegex.MatchString(uuid32) {
		t.Errorf("UUID32() = %s, does not match hex format", uuid32)
	}

	// Test that multiple calls generate different UUIDs
	uuid32_2 := UUID32()
	if uuid32 == uuid32_2 {
		t.Error("UUID32() returned the same UUID twice")
	}
}

func TestUUID32Consistency(t *testing.T) {
	// Test that UUID32 generates valid 32-character hex strings
	for i := 0; i < 10; i++ {
		uuid32 := UUID32()

		// UUID32 should be 32 characters long
		if len(uuid32) != 32 {
			t.Errorf("UUID32() length = %d, want 32", len(uuid32))
		}

		// UUID32 should contain only hexadecimal characters
		hexRegex := regexp.MustCompile(`^[0-9a-f]{32}$`)
		if !hexRegex.MatchString(uuid32) {
			t.Errorf("UUID32() = %s, does not match hex format", uuid32)
		}
	}
}

// Benchmark UUID functions
func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}

func BenchmarkUUID32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID32()
	}
}

// Additional tests for edge cases
func TestSid_GenString_Uniqueness(t *testing.T) {
	sid := NewSid()
	generated := make(map[string]bool)
	count := 1000

	for i := 0; i < count; i++ {
		id, err := sid.GenString()
		if err != nil {
			t.Errorf("GenString() error = %v", err)
			return
		}
		if generated[id] {
			t.Errorf("GenString() generated duplicate: %s", id)
		}
		generated[id] = true
	}

	if len(generated) != count {
		t.Errorf("Expected %d unique IDs, got %d", count, len(generated))
	}
}

func TestSid_GenUint64_Uniqueness(t *testing.T) {
	sid := NewSid()
	generated := make(map[uint64]bool)
	count := 1000

	for i := 0; i < count; i++ {
		id, err := sid.GenUint64()
		if err != nil {
			t.Errorf("GenUint64() error = %v", err)
			return
		}
		if generated[id] {
			t.Errorf("GenUint64() generated duplicate: %d", id)
		}
		generated[id] = true
	}

	if len(generated) != count {
		t.Errorf("Expected %d unique IDs, got %d", count, len(generated))
	}
}

func TestSid_GenString_Concurrent(t *testing.T) {
	sid := NewSid()
	results := make(chan string, 100)
	errors := make(chan error, 100)

	// Generate IDs concurrently
	for i := 0; i < 100; i++ {
		go func() {
			id, err := sid.GenString()
			if err != nil {
				errors <- err
				return
			}
			results <- id
		}()
	}

	// Collect results
	generated := make(map[string]bool)
	for i := 0; i < 100; i++ {
		select {
		case id := <-results:
			if generated[id] {
				t.Errorf("Concurrent GenString() generated duplicate: %s", id)
			}
			generated[id] = true
		case err := <-errors:
			t.Errorf("Concurrent GenString() error = %v", err)
		}
	}

	if len(generated) != 100 {
		t.Errorf("Expected 100 unique IDs, got %d", len(generated))
	}
}

func TestSid_GenUint64_Concurrent(t *testing.T) {
	sid := NewSid()
	results := make(chan uint64, 100)
	errors := make(chan error, 100)

	// Generate IDs concurrently
	for i := 0; i < 100; i++ {
		go func() {
			id, err := sid.GenUint64()
			if err != nil {
				errors <- err
				return
			}
			results <- id
		}()
	}

	// Collect results
	generated := make(map[uint64]bool)
	for i := 0; i < 100; i++ {
		select {
		case id := <-results:
			if generated[id] {
				t.Errorf("Concurrent GenUint64() generated duplicate: %d", id)
			}
			generated[id] = true
		case err := <-errors:
			t.Errorf("Concurrent GenUint64() error = %v", err)
		}
	}

	if len(generated) != 100 {
		t.Errorf("Expected 100 unique IDs, got %d", len(generated))
	}
}
