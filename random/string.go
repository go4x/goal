package random

import (
	"fmt"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const alphanumeric = letters + numbers
const hexStr = "0123456789abcdef"

// Additional character sets
const lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
const uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const symbols = "!@#$%^&*()_+-=[]{}|;':\",./<>?"
const punctuation = ".,!?;:"
const brackets = "()[]{}<>"
const hexUpper = "0123456789ABCDEF"
const alphanumericSymbols = letters + numbers + symbols
const strongPasswordChars = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
const digits = "0123456789"
const readableChars = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"

func Alphabetic(len int) string {
	return String([]byte(letters), len)
}

func Numeric(len int) string {
	return String([]byte(numbers), len)
}

func Alphanumeric(length int) string {
	return String([]byte(alphanumeric), length)
}

func Hex(length int) string {
	return String([]byte(hexStr), length)
}

func String(chars []byte, length int) string {
	selCharLen := len(chars)
	if selCharLen < 2 {
		panic("illegal argument: length of given chars need > 1")
	}
	if length <= 0 {
		panic("illegal argument: length need > 0")
	}
	builder := strings.Builder{}

	for i := 0; i < length; i++ {
		idx := globalRand.Intn(selCharLen)
		builder.WriteByte(chars[idx])
	}
	return builder.String()
}

// Lowercase generates a random lowercase string
func Lowercase(length int) string {
	return String([]byte(lowercaseLetters), length)
}

// Uppercase generates a random uppercase string
func Uppercase(length int) string {
	return String([]byte(uppercaseLetters), length)
}

// Symbols generates a random string with symbols
func Symbols(length int) string {
	return String([]byte(symbols), length)
}

// AlphanumericSymbols generates a random string with letters, numbers and symbols
func AlphanumericSymbols(length int) string {
	return String([]byte(alphanumericSymbols), length)
}

// StrongPassword generates a random string suitable for strong passwords
func StrongPassword(length int) string {
	return String([]byte(strongPasswordChars), length)
}

// Digits generates a random numeric string
func Digits(length int) string {
	return String([]byte(digits), length)
}

// Readable generates a random string with easily readable characters
func Readable(length int) string {
	return String([]byte(readableChars), length)
}

// HexUpper generates a random uppercase hexadecimal string
func HexUpper(length int) string {
	return String([]byte(hexUpper), length)
}

// ShortID generates a short random ID suitable for URLs
func ShortID(length int) string {
	return String([]byte(readableChars), length)
}

// NonRepeatingString generates a string without consecutive repeated characters
func NonRepeatingString(chars []byte, length int) string {
	if len(chars) < 2 {
		panic("chars must have at least 2 characters")
	}
	if length <= 0 {
		panic("length must be greater than 0")
	}

	builder := strings.Builder{}
	lastChar := byte(0)

	for i := 0; i < length; i++ {
		var char byte
		for {
			char = chars[globalRand.Intn(len(chars))]
			if char != lastChar {
				break
			}
		}
		builder.WriteByte(char)
		lastChar = char
	}

	return builder.String()
}

// WeightedChar represents a character with weight
type WeightedChar struct {
	Char   byte
	Weight int
}

// WeightedString generates a string using weighted character selection
func WeightedString(chars []WeightedChar, length int) string {
	if len(chars) == 0 {
		panic("chars cannot be empty")
	}
	if length <= 0 {
		panic("length must be greater than 0")
	}

	// Calculate total weight
	var totalWeight int
	for _, char := range chars {
		if char.Weight < 0 {
			panic("weight cannot be negative")
		}
		totalWeight += char.Weight
	}

	if totalWeight == 0 {
		panic("total weight cannot be zero")
	}

	builder := strings.Builder{}

	for i := 0; i < length; i++ {
		r := globalRand.Intn(totalWeight)
		cumulative := 0

		for _, char := range chars {
			cumulative += char.Weight
			if r < cumulative {
				builder.WriteByte(char.Char)
				break
			}
		}
	}

	return builder.String()
}

// PatternString generates a string based on a pattern
// Pattern characters:
//
//	a = lowercase letter
//	A = uppercase letter
//	n = number
//	s = symbol
//	x = any alphanumeric
//	? = any character
func PatternString(pattern string) string {
	if len(pattern) == 0 {
		panic("pattern cannot be empty")
	}

	builder := strings.Builder{}

	for _, char := range pattern {
		switch char {
		case 'a':
			builder.WriteByte(lowercaseLetters[globalRand.Intn(len(lowercaseLetters))])
		case 'A':
			builder.WriteByte(uppercaseLetters[globalRand.Intn(len(uppercaseLetters))])
		case 'n':
			builder.WriteByte(digits[globalRand.Intn(len(digits))])
		case 's':
			builder.WriteByte(symbols[globalRand.Intn(len(symbols))])
		case 'x':
			builder.WriteByte(alphanumeric[globalRand.Intn(len(alphanumeric))])
		case '?':
			builder.WriteByte(alphanumericSymbols[globalRand.Intn(len(alphanumericSymbols))])
		default:
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

// Password generates a random password with optional symbols
func Password(length int, includeSymbols bool) string {
	if length <= 0 {
		panic("length must be greater than 0")
	}

	var chars string
	if includeSymbols {
		chars = alphanumericSymbols
	} else {
		chars = alphanumeric
	}

	return String([]byte(chars), length)
}

// Username generates a random username
func Username(length int) string {
	if length <= 0 {
		panic("length must be greater than 0")
	}

	usernameChars := lowercaseLetters + uppercaseLetters + digits + "_"
	return String([]byte(usernameChars), length)
}

// Email generates a random email prefix
func Email(length int) string {
	if length <= 0 {
		panic("length must be greater than 0")
	}

	emailChars := lowercaseLetters + uppercaseLetters + digits + "._"
	return String([]byte(emailChars), length)
}

// Token generates a random security token
func Token(length int) string {
	if length <= 0 {
		panic("length must be greater than 0")
	}

	return String([]byte(alphanumeric), length)
}

// ColorHex generates a random color hex code
func ColorHex() string {
	return "#" + Hex(6)
}

// ColorRGB generates a random RGB color string
func ColorRGB() string {
	r := globalRand.Intn(256)
	g := globalRand.Intn(256)
	b := globalRand.Intn(256)
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// MACAddress generates a random MAC address
func MACAddress() string {
	parts := make([]string, 6)
	for i := range parts {
		parts[i] = HexUpper(2)
	}
	return strings.Join(parts, ":")
}

// IPAddress generates a random IP address
func IPAddress() string {
	parts := make([]string, 4)
	for i := range parts {
		parts[i] = fmt.Sprintf("%d", globalRand.Intn(256))
	}
	return strings.Join(parts, ".")
}
