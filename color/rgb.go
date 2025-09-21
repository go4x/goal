// Package color provides utilities for working with RGB color values.
// It supports conversion between RGB and hexadecimal color representations.
//
// The package includes:
//   - RGB struct for representing color values
//   - NewRGB constructor for creating RGB instances
//   - Hex() method for converting RGB to hexadecimal format
//   - Hex2RGB function for converting hexadecimal to RGB
//
// Example:
//
//	// Create RGB color and convert to hex
//	rgb := color.NewRGB(255, 0, 0) // Red color
//	hex := rgb.Hex()               // Returns "FF0000"
//
//	// Convert hex back to RGB
//	rgb2 := color.Hex2RGB("00FF00") // Green color
package color

import (
	"errors"
	"strconv"
	"strings"
)

type RGB struct {
	Red, Green, Blue uint8
}

func NewRGB(r, g, b uint8) RGB {
	return RGB{r, g, b}
}

func t2x(t uint8) string {
	result := strconv.FormatInt(int64(t), 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func (color RGB) Hex() string {
	r := t2x(color.Red)
	g := t2x(color.Green)
	b := t2x(color.Blue)
	return r + g + b
}

// Hex2RGB converts a hexadecimal color string to RGB.
// The hex string should be in format "RRGGBB" (6 characters).
// Returns an error if the input format is invalid.
func Hex2RGB(hex string) (RGB, error) {
	// Remove # prefix if present
	hex = strings.TrimPrefix(hex, "#")

	// Validate hex string format
	if len(hex) != 6 {
		return RGB{}, errors.New("hex string must be 6 characters long")
	}

	// Parse red component
	r64, err := strconv.ParseInt(hex[:2], 16, 10)
	if err != nil {
		return RGB{}, errors.New("invalid hex format for red component")
	}
	r := uint8(r64)

	// Parse green component
	g64, err := strconv.ParseInt(hex[2:4], 16, 10)
	if err != nil {
		return RGB{}, errors.New("invalid hex format for green component")
	}
	g := uint8(g64)

	// Parse blue component
	b64, err := strconv.ParseInt(hex[4:], 16, 10)
	if err != nil {
		return RGB{}, errors.New("invalid hex format for blue component")
	}
	b := uint8(b64)

	return RGB{r, g, b}, nil
}
