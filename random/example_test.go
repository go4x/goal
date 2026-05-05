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

}

func ExampleLowercase() {
	// Generate lowercase random strings
	fmt.Printf("Lowercase: %s\n", random.Lowercase(8))

}

func ExampleUppercase() {
	// Generate uppercase random strings
	fmt.Printf("Uppercase: %s\n", random.Uppercase(8))

}

func ExampleDigits() {
	// Generate digit-only random strings
	fmt.Printf("Digits: %s\n", random.Digits(6))

}

func ExampleSymbols() {
	// Generate symbol-only random strings
	fmt.Printf("Symbols: %s\n", random.Symbols(5))

}

func ExampleHexUpper() {
	// Generate uppercase hexadecimal strings
	fmt.Printf("Hex Upper: %s\n", random.HexUpper(8))

}

func ExampleAlphanumericSymbols() {
	// Generate alphanumeric with symbols
	fmt.Printf("Alphanumeric + Symbols: %s\n", random.AlphanumericSymbols(10))

}

func ExampleStrongPassword() {
	// Generate strong password (no confusing characters)
	fmt.Printf("Strong Password: %s\n", random.StrongPassword(12))

}

func ExampleReadable() {
	// Generate readable strings (no confusing characters)
	fmt.Printf("Readable: %s\n", random.Readable(10))

}

func ExampleShortID() {
	// Generate short IDs for URLs
	fmt.Printf("Short ID: %s\n", random.ShortID(8))

}

func ExamplePassword() {
	// Generate passwords with and without symbols
	fmt.Printf("Password with symbols: %s\n", random.Password(12, true))
	fmt.Printf("Password without symbols: %s\n", random.Password(12, false))

}

func ExampleUsername() {
	// Generate usernames
	fmt.Printf("Username: %s\n", random.Username(8))

}

func ExampleEmail() {
	// Generate email prefixes
	fmt.Printf("Email prefix: %s@example.com\n", random.Email(8))

}

func ExampleToken() {
	// Generate security tokens
	fmt.Printf("Token: %s\n", random.Token(32))

}

func ExampleColorHex() {
	// Generate random color hex codes
	fmt.Printf("Color Hex: %s\n", random.ColorHex())

}

func ExampleColorRGB() {
	// Generate random RGB color strings
	fmt.Printf("Color RGB: %s\n", random.ColorRGB())

}

func ExampleMACAddress() {
	// Generate random MAC addresses
	fmt.Printf("MAC Address: %s\n", random.MACAddress())

}

func ExampleIPAddress() {
	// Generate random IP addresses
	fmt.Printf("IP Address: %s\n", random.IPAddress())

}

func ExampleWeightedString() {
	// Generate weighted random strings
	chars := []random.WeightedChar{
		{Char: 'a', Weight: 1},
		{Char: 'b', Weight: 2},
		{Char: 'c', Weight: 3},
	}
	fmt.Printf("Weighted String: %s\n", random.WeightedString(chars, 10))

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

}
