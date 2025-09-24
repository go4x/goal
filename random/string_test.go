package random_test

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/go4x/goal/random"
	"github.com/go4x/got"
)

func TestRandomAlphabetic(t *testing.T) {
	tl := got.New(t, "test Alphabetic")
	tl.Case("loop 10 times to generate random string")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Alphabetic(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomNumber(t *testing.T) {
	tl := got.New(t, "test Numeric")
	tl.Case("loop 10 times to generate random number as string")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Numeric(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomAlphanumeric(t *testing.T) {
	tl := got.New(t, "test Alphanumeric")
	tl.Case("loop 10 times to generate random Alphanumeric")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Alphanumeric(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomHex(t *testing.T) {
	tl := got.New(t, "test Hex")
	tl.Case("loop 10 times to generate random hex string")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Hex(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

// Test new string generation functions
func TestLowercase(t *testing.T) {
	tl := got.New(t, "test Lowercase")
	tl.Case("generate lowercase strings")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Lowercase(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check all characters are lowercase
		for _, char := range s {
			tl.Require(char >= 'a' && char <= 'z', "all characters should be lowercase, got %c", char)
		}
	}
}

func TestUppercase(t *testing.T) {
	tl := got.New(t, "test Uppercase")
	tl.Case("generate uppercase strings")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Uppercase(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check all characters are uppercase
		for _, char := range s {
			tl.Require(char >= 'A' && char <= 'Z', "all characters should be uppercase, got %c", char)
		}
	}
}

func TestDigits(t *testing.T) {
	tl := got.New(t, "test Digits")
	tl.Case("generate digit strings")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Digits(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check all characters are digits
		for _, char := range s {
			tl.Require(char >= '0' && char <= '9', "all characters should be digits, got %c", char)
		}
	}
}

func TestSymbols(t *testing.T) {
	tl := got.New(t, "test Symbols")
	tl.Case("generate symbol strings")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Symbols(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check all characters are symbols
		for _, char := range s {
			tl.Require(!((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')),
				"all characters should be symbols, got %c", char)
		}
	}
}

func TestHexUpper(t *testing.T) {
	tl := got.New(t, "test HexUpper")
	tl.Case("generate uppercase hex strings")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.HexUpper(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check all characters are valid hex uppercase
		for _, char := range s {
			tl.Require((char >= '0' && char <= '9') || (char >= 'A' && char <= 'F'),
				"all characters should be uppercase hex, got %c", char)
		}
	}
}

func TestAlphanumericSymbols(t *testing.T) {
	tl := got.New(t, "test AlphanumericSymbols")
	tl.Case("generate alphanumeric with symbols")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.AlphanumericSymbols(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
	}
}

func TestStrongPassword(t *testing.T) {
	tl := got.New(t, "test StrongPassword")
	tl.Case("generate strong password strings")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.StrongPassword(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check no confusing characters (0, O, l, I, 1)
		for _, char := range s {
			tl.Require(char != '0' && char != 'O' && char != 'l' && char != 'I' && char != '1',
				"should not contain confusing characters, got %c", char)
		}
	}
}

func TestReadable(t *testing.T) {
	tl := got.New(t, "test Readable")
	tl.Case("generate readable strings")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Readable(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check no confusing characters
		for _, char := range s {
			tl.Require(char != '0' && char != 'O' && char != 'l' && char != 'I' && char != '1',
				"should not contain confusing characters, got %c", char)
		}
	}
}

func TestShortID(t *testing.T) {
	tl := got.New(t, "test ShortID")
	tl.Case("generate short IDs")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.ShortID(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
	}
}

func TestPassword(t *testing.T) {
	tl := got.New(t, "test Password")
	tl.Case("generate passwords with and without symbols")

	// Test with symbols
	for i := 0; i < 5; i++ {
		n := rand.Intn(20) + 8
		s := random.Password(n, true)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
	}

	// Test without symbols
	for i := 0; i < 5; i++ {
		n := rand.Intn(20) + 8
		s := random.Password(n, false)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check no symbols
		for _, char := range s {
			tl.Require((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9'),
				"should not contain symbols, got %c", char)
		}
	}
}

func TestUsername(t *testing.T) {
	tl := got.New(t, "test Username")
	tl.Case("generate usernames")
	for i := 0; i < 10; i++ {
		n := rand.Intn(20) + 3
		s := random.Username(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check valid username characters
		for _, char := range s {
			tl.Require((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '_',
				"should contain only valid username characters, got %c", char)
		}
	}
}

func TestEmail(t *testing.T) {
	tl := got.New(t, "test Email")
	tl.Case("generate email prefixes")
	for i := 0; i < 10; i++ {
		n := rand.Intn(20) + 3
		s := random.Email(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check valid email characters
		for _, char := range s {
			tl.Require((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '.' || char == '_',
				"should contain only valid email characters, got %c", char)
		}
	}
}

func TestToken(t *testing.T) {
	tl := got.New(t, "test Token")
	tl.Case("generate tokens")
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 8
		s := random.Token(n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
	}
}

func TestColorHex(t *testing.T) {
	tl := got.New(t, "test ColorHex")
	tl.Case("generate color hex strings")
	for i := 0; i < 10; i++ {
		s := random.ColorHex()
		tl.Require(len(s) == 7, "color hex should be 7 characters, got %d", len(s))
		tl.Require(s[0] == '#', "should start with #, got %c", s[0])
		// Check valid hex characters
		for i := 1; i < len(s); i++ {
			char := s[i]
			tl.Require((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F'),
				"should contain valid hex characters, got %c", char)
		}
	}
}

func TestColorRGB(t *testing.T) {
	tl := got.New(t, "test ColorRGB")
	tl.Case("generate RGB color strings")
	for i := 0; i < 10; i++ {
		s := random.ColorRGB()
		tl.Require(len(s) >= 10, "RGB string should be at least 10 characters, got %d", len(s))
		tl.Require(s[:4] == "rgb(", "should start with 'rgb(', got %s", s[:4])
		tl.Require(s[len(s)-1] == ')', "should end with ')', got %c", s[len(s)-1])
	}
}

func TestMACAddress(t *testing.T) {
	tl := got.New(t, "test MACAddress")
	tl.Case("generate MAC addresses")
	for i := 0; i < 10; i++ {
		s := random.MACAddress()
		tl.Require(len(s) == 17, "MAC address should be 17 characters, got %d", len(s))
		// Check format XX:XX:XX:XX:XX:XX
		for i := 0; i < 17; i += 3 {
			if i < 15 {
				tl.Require(s[i+2] == ':', "should have colons at positions 2,5,8,11,14, got %c at position %d", s[i+2], i+2)
			}
		}
	}
}

func TestIPAddress(t *testing.T) {
	tl := got.New(t, "test IPAddress")
	tl.Case("generate IP addresses")
	for i := 0; i < 10; i++ {
		s := random.IPAddress()
		tl.Require(len(s) >= 7, "IP address should be at least 7 characters, got %d", len(s))
		// Check format X.X.X.X
		parts := strings.Split(s, ".")
		tl.Require(len(parts) == 4, "should have 4 parts separated by dots, got %d", len(parts))
		for _, part := range parts {
			tl.Require(len(part) > 0, "each part should not be empty")
		}
	}
}

func TestWeightedString(t *testing.T) {
	tl := got.New(t, "test WeightedString")
	tl.Case("generate weighted strings")

	chars := []random.WeightedChar{
		{Char: 'a', Weight: 1},
		{Char: 'b', Weight: 2},
		{Char: 'c', Weight: 3},
	}

	for i := 0; i < 10; i++ {
		n := rand.Intn(20) + 1
		s := random.WeightedString(chars, n)
		tl.Require(n == len(s), "length should be %d, got %d", n, len(s))
		// Check all characters are from the weighted set
		for _, char := range s {
			tl.Require(char == 'a' || char == 'b' || char == 'c', "should only contain weighted characters, got %c", char)
		}
	}
}

func TestPatternString(t *testing.T) {
	tl := got.New(t, "test PatternString")
	tl.Case("generate pattern-based strings")

	patterns := []string{
		"aaa", // lowercase letters
		"AAA", // uppercase letters
		"nnn", // numbers
		"sss", // symbols
		"xxx", // alphanumeric
		"???", // any character
		"aAn", // mixed pattern
	}

	for _, pattern := range patterns {
		s := random.PatternString(pattern)
		tl.Require(len(s) == len(pattern), "length should match pattern length %d, got %d", len(pattern), len(s))
	}
}

// Test edge cases and error conditions
func TestStringEdgeCases(t *testing.T) {
	tl := got.New(t, "test String edge cases")

	tl.Case("test zero length")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for zero length")
		}()
		random.String([]byte("abc"), 0)
	}()

	tl.Case("test negative length")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for negative length")
		}()
		random.String([]byte("abc"), -1)
	}()

	tl.Case("test empty character set")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for empty character set")
		}()
		random.String([]byte{}, 5)
	}()

	tl.Case("test single character set")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for single character set")
		}()
		random.String([]byte("a"), 5)
	}()
}

func TestPasswordEdgeCases(t *testing.T) {
	tl := got.New(t, "test Password edge cases")

	tl.Case("test zero length password")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for zero length")
		}()
		random.Password(0, true)
	}()

	tl.Case("test negative length password")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for negative length")
		}()
		random.Password(-1, true)
	}()
}

func TestWeightedStringEdgeCases(t *testing.T) {
	tl := got.New(t, "test WeightedString edge cases")

	tl.Case("test zero weight")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for zero total weight")
		}()
		chars := []random.WeightedChar{
			{Char: 'a', Weight: 0},
			{Char: 'b', Weight: 0},
		}
		random.WeightedString(chars, 5)
	}()

	tl.Case("test empty weighted chars")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for empty weighted chars")
		}()
		random.WeightedString([]random.WeightedChar{}, 5)
	}()
}

func TestPatternStringEdgeCases(t *testing.T) {
	tl := got.New(t, "test PatternString edge cases")

	tl.Case("test empty pattern")
	func() {
		defer func() {
			err := recover()
			tl.Require(err != nil, "should panic for empty pattern")
		}()
		random.PatternString("")
	}()
}
