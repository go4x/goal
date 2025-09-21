package ciphers_test

import (
	"fmt"
	"testing"

	"github.com/go4x/goal/ciphers"
	"github.com/go4x/goal/random"
)

var plainText = "exampleplaintext" // 16 bytes
var key = []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}

func TestAES(t *testing.T) {
	// ECB
	s, _ := ciphers.AES.Encrypt([]byte(plainText), key, ciphers.ECB, nil)
	fmt.Printf("%x\n", s) // fixed ciphertext: 40d32c0de54cc6b82c39f22e641455d3a254be88e037ddd9d79fb6411c3f9df8
	s, _ = ciphers.AES.Decrypt(s, key, ciphers.ECB, nil)
	fmt.Println(string(s))

	// CBC
	// random iv, length is 16 bytes, each ciphertext is different
	r := random.Hex(len(key))
	fmt.Println("random iv:", r)
	iv := []byte(r)
	s, _ = ciphers.AES.Encrypt([]byte(plainText), key, ciphers.CBC, iv)
	fmt.Printf("%x\n", s) // random ciphertext
	s, _ = ciphers.AES.Decrypt(s, key, ciphers.CBC, iv)
	fmt.Println(string(s))

	// fixed iv, each ciphertext is the same
	iv = key
	s, _ = ciphers.AES.Encrypt([]byte(plainText), key, ciphers.CBC, iv)
	fmt.Printf("%x\n", s) // fixed ciphertext: 0735437968e811771051aa81734b1098b8353285c0c9517a752a429a3efc44fe
	s, _ = ciphers.AES.Decrypt(s, key, ciphers.CBC, iv)
	fmt.Println(string(s))
}

func TestAES_ErrorCases(t *testing.T) {
	// Test invalid key length
	invalidKey := []byte{1, 2, 3} // too short
	_, err := ciphers.AES.Encrypt([]byte("test"), invalidKey, ciphers.ECB, nil)
	if err == nil {
		t.Error("Expected error for invalid key length")
	}
	fmt.Printf("Invalid key error: %v\n", err)

	// Test invalid IV length for CBC
	shortIV := []byte{1, 2, 3} // too short
	_, err = ciphers.AES.Encrypt([]byte("test"), key, ciphers.CBC, shortIV)
	if err == nil {
		t.Error("Expected error for invalid IV length")
	}
	fmt.Printf("Invalid IV error: %v\n", err)

	// Test decryption with invalid key
	_, err = ciphers.AES.Decrypt([]byte("test"), invalidKey, ciphers.ECB, nil)
	if err == nil {
		t.Error("Expected error for invalid key length in decrypt")
	}
	fmt.Printf("Invalid key decrypt error: %v\n", err)

	// Test decryption with invalid IV
	_, err = ciphers.AES.Decrypt([]byte("test"), key, ciphers.CBC, shortIV)
	if err == nil {
		t.Error("Expected error for invalid IV length in decrypt")
	}
	fmt.Printf("Invalid IV decrypt error: %v\n", err)
}

func TestAES_EdgeCases(t *testing.T) {
	// Test with different key sizes
	key16 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	key24 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	key32 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Test AES-128 (16 bytes)
	encrypted, err := ciphers.AES.Encrypt([]byte("test"), key16, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("AES-128 encryption failed: %v", err)
	}
	decrypted, err := ciphers.AES.Decrypt(encrypted, key16, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("AES-128 decryption failed: %v", err)
	}
	if string(decrypted) != "test" {
		t.Errorf("AES-128 round trip failed: got %s, want test", string(decrypted))
	}

	// Test AES-192 (24 bytes)
	encrypted, err = ciphers.AES.Encrypt([]byte("test"), key24, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("AES-192 encryption failed: %v", err)
	}
	decrypted, err = ciphers.AES.Decrypt(encrypted, key24, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("AES-192 decryption failed: %v", err)
	}
	if string(decrypted) != "test" {
		t.Errorf("AES-192 round trip failed: got %s, want test", string(decrypted))
	}

	// Test AES-256 (32 bytes)
	encrypted, err = ciphers.AES.Encrypt([]byte("test"), key32, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("AES-256 encryption failed: %v", err)
	}
	decrypted, err = ciphers.AES.Decrypt(encrypted, key32, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("AES-256 decryption failed: %v", err)
	}
	if string(decrypted) != "test" {
		t.Errorf("AES-256 round trip failed: got %s, want test", string(decrypted))
	}

	// Test CBC with different key sizes
	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	// AES-128 CBC
	encrypted, err = ciphers.AES.Encrypt([]byte("test"), key16, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("AES-128 CBC encryption failed: %v", err)
	}
	decrypted, err = ciphers.AES.Decrypt(encrypted, key16, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("AES-128 CBC decryption failed: %v", err)
	}
	if string(decrypted) != "test" {
		t.Errorf("AES-128 CBC round trip failed: got %s, want test", string(decrypted))
	}

	// AES-192 CBC
	encrypted, err = ciphers.AES.Encrypt([]byte("test"), key24, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("AES-192 CBC encryption failed: %v", err)
	}
	decrypted, err = ciphers.AES.Decrypt(encrypted, key24, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("AES-192 CBC decryption failed: %v", err)
	}
	if string(decrypted) != "test" {
		t.Errorf("AES-192 CBC round trip failed: got %s, want test", string(decrypted))
	}

	// AES-256 CBC
	encrypted, err = ciphers.AES.Encrypt([]byte("test"), key32, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("AES-256 CBC encryption failed: %v", err)
	}
	decrypted, err = ciphers.AES.Decrypt(encrypted, key32, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("AES-256 CBC decryption failed: %v", err)
	}
	if string(decrypted) != "test" {
		t.Errorf("AES-256 CBC round trip failed: got %s, want test", string(decrypted))
	}
}

func TestAES_PaddingEdgeCases(t *testing.T) {
	// Test with data that requires different padding amounts
	testCases := []struct {
		name string
		data string
	}{
		{"empty", ""},
		{"single byte", "a"},
		{"15 bytes", "123456789012345"},
		{"16 bytes", "1234567890123456"},
		{"17 bytes", "12345678901234567"},
		{"31 bytes", "1234567890123456789012345678901"},
		{"32 bytes", "12345678901234567890123456789012"},
		{"33 bytes", "123456789012345678901234567890123"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test ECB
			encrypted, err := ciphers.AES.Encrypt([]byte(tc.data), key, ciphers.ECB, nil)
			if err != nil {
				t.Errorf("ECB encryption failed for %s: %v", tc.name, err)
				return
			}
			decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.ECB, nil)
			if err != nil {
				t.Errorf("ECB decryption failed for %s: %v", tc.name, err)
				return
			}
			if string(decrypted) != tc.data {
				t.Errorf("ECB round trip failed for %s: got %s, want %s", tc.name, string(decrypted), tc.data)
			}

			// Test CBC
			iv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
			encrypted, err = ciphers.AES.Encrypt([]byte(tc.data), key, ciphers.CBC, iv)
			if err != nil {
				t.Errorf("CBC encryption failed for %s: %v", tc.name, err)
				return
			}
			decrypted, err = ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
			if err != nil {
				t.Errorf("CBC decryption failed for %s: %v", tc.name, err)
				return
			}
			if string(decrypted) != tc.data {
				t.Errorf("CBC round trip failed for %s: got %s, want %s", tc.name, string(decrypted), tc.data)
			}
		})
	}
}

func TestAES_InvalidCiphertext(t *testing.T) {
	// Test decryption with invalid ciphertext
	invalidCiphertext := []byte{1, 2, 3, 4, 5} // not multiple of block size
	_, err := ciphers.AES.Decrypt(invalidCiphertext, key, ciphers.ECB, nil)
	if err == nil {
		t.Error("Expected error for invalid ciphertext length")
	}
	fmt.Printf("Invalid ciphertext error: %v\n", err)

	// Test decryption with invalid padding
	// Create a valid length ciphertext but with invalid padding
	invalidPadding := make([]byte, 16) // 16 bytes, multiple of block size
	for i := range invalidPadding {
		invalidPadding[i] = 0xFF // invalid padding value
	}
	_, err = ciphers.AES.Decrypt(invalidPadding, key, ciphers.ECB, nil)
	if err == nil {
		t.Error("Expected error for invalid padding")
	}
	fmt.Printf("Invalid padding error: %v\n", err)

	// Test decryption with empty ciphertext
	_, err = ciphers.AES.Decrypt([]byte{}, key, ciphers.ECB, nil)
	if err == nil {
		t.Error("Expected error for empty ciphertext")
	}
	fmt.Printf("Empty ciphertext error: %v\n", err)
}

func TestAES_MoreErrorCases(t *testing.T) {
	// Test with unsupported mode (should not happen with current implementation)
	// but we can test the error handling in the switch statements

	// Test CBC with nil IV (should fail validation)
	_, err := ciphers.AES.Encrypt([]byte("test"), key, ciphers.CBC, nil)
	if err == nil {
		t.Error("Expected error for nil IV in CBC mode")
	}
	fmt.Printf("Nil IV error: %v\n", err)

	// Test decryption with nil IV
	_, err = ciphers.AES.Decrypt([]byte("test"), key, ciphers.CBC, nil)
	if err == nil {
		t.Error("Expected error for nil IV in CBC decrypt")
	}
	fmt.Printf("Nil IV decrypt error: %v\n", err)

	// Test with very long IV (should fail validation)
	longIV := make([]byte, 32) // too long
	_, err = ciphers.AES.Encrypt([]byte("test"), key, ciphers.CBC, longIV)
	if err == nil {
		t.Error("Expected error for long IV")
	}
	fmt.Printf("Long IV error: %v\n", err)

	// Test decryption with long IV
	_, err = ciphers.AES.Decrypt([]byte("test"), key, ciphers.CBC, longIV)
	if err == nil {
		t.Error("Expected error for long IV in decrypt")
	}
	fmt.Printf("Long IV decrypt error: %v\n", err)
}

func TestAES_Performance(t *testing.T) {
	// Test with larger data to ensure performance is reasonable
	largeData := make([]byte, 1024*1024) // 1MB
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// Test ECB with large data
	encrypted, err := ciphers.AES.Encrypt(largeData, key, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("ECB encryption failed for large data: %v", err)
	}
	decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.ECB, nil)
	if err != nil {
		t.Errorf("ECB decryption failed for large data: %v", err)
	}
	if len(decrypted) != len(largeData) {
		t.Errorf("ECB round trip failed for large data: got %d bytes, want %d", len(decrypted), len(largeData))
	}

	// Test CBC with large data
	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	encrypted, err = ciphers.AES.Encrypt(largeData, key, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("CBC encryption failed for large data: %v", err)
	}
	decrypted, err = ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
	if err != nil {
		t.Errorf("CBC decryption failed for large data: %v", err)
	}
	if len(decrypted) != len(largeData) {
		t.Errorf("CBC round trip failed for large data: got %d bytes, want %d", len(decrypted), len(largeData))
	}
}

func TestAES_UnsupportedMode(t *testing.T) {
	// Test with unsupported mode (this should not happen in practice)
	// but we can test the error handling in the switch statements

	// Create a custom mode that's not supported
	unsupportedMode := ciphers.AESMode(999)

	_, err := ciphers.AES.Encrypt([]byte("test"), key, unsupportedMode, nil)
	if err == nil {
		t.Error("Expected error for unsupported mode")
	}
	fmt.Printf("Unsupported mode encrypt error: %v\n", err)

	_, err = ciphers.AES.Decrypt([]byte("test"), key, unsupportedMode, nil)
	if err == nil {
		t.Error("Expected error for unsupported mode")
	}
	fmt.Printf("Unsupported mode decrypt error: %v\n", err)
}

func TestAES_BlockSizeEdgeCases(t *testing.T) {
	// Test specific padding edge cases that might not be covered
	testCases := []struct {
		name string
		data []byte
	}{
		{"exactly one block", make([]byte, 16)},
		{"exactly two blocks", make([]byte, 32)},
		{"one byte short of block", make([]byte, 15)},
		{"one byte over block", make([]byte, 17)},
		{"two bytes short of block", make([]byte, 14)},
		{"two bytes over block", make([]byte, 18)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Fill with some data
			for i := range tc.data {
				tc.data[i] = byte(i % 256)
			}

			// Test ECB
			encrypted, err := ciphers.AES.Encrypt(tc.data, key, ciphers.ECB, nil)
			if err != nil {
				t.Errorf("ECB encryption failed for %s: %v", tc.name, err)
				return
			}
			decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.ECB, nil)
			if err != nil {
				t.Errorf("ECB decryption failed for %s: %v", tc.name, err)
				return
			}
			if len(decrypted) != len(tc.data) {
				t.Errorf("ECB round trip failed for %s: got %d bytes, want %d", tc.name, len(decrypted), len(tc.data))
			}

			// Test CBC
			iv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
			encrypted, err = ciphers.AES.Encrypt(tc.data, key, ciphers.CBC, iv)
			if err != nil {
				t.Errorf("CBC encryption failed for %s: %v", tc.name, err)
				return
			}
			decrypted, err = ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, iv)
			if err != nil {
				t.Errorf("CBC decryption failed for %s: %v", tc.name, err)
				return
			}
			if len(decrypted) != len(tc.data) {
				t.Errorf("CBC round trip failed for %s: got %d bytes, want %d", tc.name, len(decrypted), len(tc.data))
			}
		})
	}
}
