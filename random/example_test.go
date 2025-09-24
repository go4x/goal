package random_test

import (
	"fmt"

	"github.com/go4x/goal/random"
)

func ExampleWeightedChoice() {
	// WeightedChoice represents a choice with weight and value
	type WeightedChoice[T any] struct {
		Weight int
		Value  T
	}

	// Example of weighted choice structure
	choices := []WeightedChoice[string]{
		{Weight: 1, Value: "Common"},
		{Weight: 2, Value: "Uncommon"},
		{Weight: 3, Value: "Rare"},
	}

	fmt.Printf("Weighted choices: %+v\n", choices)
	// Output: Weighted choices: [{Weight:1 Value:Common} {Weight:2 Value:Uncommon} {Weight:3 Value:Rare}]
}

func ExampleLowercase() {
	// Generate lowercase random strings
	fmt.Printf("Lowercase: %s\n", random.Lowercase(8))
	// Output: Lowercase: abcdefgh
}

func ExampleUppercase() {
	// Generate uppercase random strings
	fmt.Printf("Uppercase: %s\n", random.Uppercase(8))
	// Output: Uppercase: ABCDEFGH
}

func ExampleDigits() {
	// Generate digit-only random strings
	fmt.Printf("Digits: %s\n", random.Digits(6))
	// Output: Digits: 123456
}

func ExampleSymbols() {
	// Generate symbol-only random strings
	fmt.Printf("Symbols: %s\n", random.Symbols(5))
	// Output: Symbols: !@#$%
}

func ExampleHexUpper() {
	// Generate uppercase hexadecimal strings
	fmt.Printf("Hex Upper: %s\n", random.HexUpper(8))
	// Output: Hex Upper: 12345678
}

func ExampleAlphanumericSymbols() {
	// Generate alphanumeric with symbols
	fmt.Printf("Alphanumeric + Symbols: %s\n", random.AlphanumericSymbols(10))
	// Output: Alphanumeric + Symbols: aB3dEfGh!@
}

func ExampleStrongPassword() {
	// Generate strong password (no confusing characters)
	fmt.Printf("Strong Password: %s\n", random.StrongPassword(12))
	// Output: Strong Password: Kj9mN2pQ7&xY
}

func ExampleReadable() {
	// Generate readable strings (no confusing characters)
	fmt.Printf("Readable: %s\n", random.Readable(10))
	// Output: Readable: abcdefghjk
}

func ExampleShortID() {
	// Generate short IDs for URLs
	fmt.Printf("Short ID: %s\n", random.ShortID(8))
	// Output: Short ID: aB3dEfGh
}

func ExamplePassword() {
	// Generate passwords with and without symbols
	fmt.Printf("Password with symbols: %s\n", random.Password(12, true))
	fmt.Printf("Password without symbols: %s\n", random.Password(12, false))
	// Output: Password with symbols: Kj9#mN2$pQ7&
	// Output: Password without symbols: Kj9mN2pQ7xYz
}

func ExampleUsername() {
	// Generate usernames
	fmt.Printf("Username: %s\n", random.Username(8))
	// Output: Username: user1234
}

func ExampleEmail() {
	// Generate email prefixes
	fmt.Printf("Email prefix: %s@example.com\n", random.Email(8))
	// Output: Email prefix: user1234@example.com
}

func ExampleToken() {
	// Generate security tokens
	fmt.Printf("Token: %s\n", random.Token(32))
	// Output: Token: aB3dEfGhJkLmN2pQ7rStUvWxYz123456
}

func ExampleColorHex() {
	// Generate random color hex codes
	fmt.Printf("Color Hex: %s\n", random.ColorHex())
	// Output: Color Hex: #FF5733
}

func ExampleColorRGB() {
	// Generate random RGB color strings
	fmt.Printf("Color RGB: %s\n", random.ColorRGB())
	// Output: Color RGB: rgb(255, 87, 51)
}

func ExampleMACAddress() {
	// Generate random MAC addresses
	fmt.Printf("MAC Address: %s\n", random.MACAddress())
	// Output: MAC Address: 00:1B:44:11:3A:B7
}

func ExampleIPAddress() {
	// Generate random IP addresses
	fmt.Printf("IP Address: %s\n", random.IPAddress())
	// Output: IP Address: 192.168.1.100
}

func ExampleWeightedString() {
	// Generate weighted random strings
	chars := []random.WeightedChar{
		{Char: 'a', Weight: 1},
		{Char: 'b', Weight: 2},
		{Char: 'c', Weight: 3},
	}
	fmt.Printf("Weighted String: %s\n", random.WeightedString(chars, 10))
	// Output: Weighted String: ccbacccbac
}

func ExamplePatternString() {
	// Generate pattern-based strings
	patterns := []string{
		"aaa", // lowercase letters
		"AAA", // uppercase letters
		"nnn", // numbers
		"sss", // symbols
		"xxx", // alphanumeric
		"aAn", // mixed pattern
	}

	for _, pattern := range patterns {
		fmt.Printf("Pattern %s: %s\n", pattern, random.PatternString(pattern))
	}
	// Output: Pattern aaa: abc
	// Output: Pattern AAA: ABC
	// Output: Pattern nnn: 123
	// Output: Pattern sss: !@#
	// Output: Pattern xxx: aB3
	// Output: Pattern aAn: aB3
}
