package color_test

import (
	"strings"
	"testing"

	"github.com/gophero/goal/color"
	"github.com/gophero/got"
)

func TestRgb(t *testing.T) {
	cases := map[string][]uint8{
		"000000": {0, 0, 0},
		"FFFFFF": {255, 255, 255},
		"FF0000": {255, 0, 0},
		"00FF00": {0, 255, 0},
		"0000FF": {0, 0, 255},
		"C8C8C8": {200, 200, 200},
	}
	logger := got.New(t, "Rgb")

	logger.Case("test rgb.Hex()")
	for k, v := range cases {
		rgb := color.NewRGB(v[0], v[1], v[2])
		hex := rgb.Hex()
		logger.Require(strings.ToUpper(hex) == k, "hex is correct. expects hex to be %s, found %s", hex, k)
	}

	logger.Case("test color.Hex2RGB()")
	for _, v := range cases {
		rgb := color.NewRGB(v[0], v[1], v[2])
		hex := rgb.Hex()
		rgb2, err := color.Hex2RGB(hex)
		logger.Require(err == nil, "Hex2RGB should not return error for valid hex: %s", hex)
		logger.Require(rgb2.Red == v[0] && rgb2.Green == v[1] && rgb2.Blue == v[2], "rgb result should be correct, dest rgb is: %v, found: %v", rgb2, v)
	}
}

func TestHex2RGBErrorHandling(t *testing.T) {
	logger := got.New(t, "Hex2RGBErrorHandling")

	logger.Case("test invalid hex formats")
	invalidCases := []struct {
		hex  string
		desc string
	}{
		{"", "empty string"},
		{"12345", "too short"},
		{"1234567", "too long"},
		{"GGGGGG", "invalid characters"},
		{"12345G", "mixed valid/invalid"},
		{"12", "way too short"},
	}

	for _, tc := range invalidCases {
		_, err := color.Hex2RGB(tc.hex)
		logger.Require(err != nil, "Hex2RGB should return error for %s: %s", tc.desc, tc.hex)
	}

	logger.Case("test hex with # prefix")
	hexWithPrefix := "#FF0000"
	rgb, err := color.Hex2RGB(hexWithPrefix)
	logger.Require(err == nil, "Hex2RGB should handle # prefix without error")
	logger.Require(rgb.Red == 255 && rgb.Green == 0 && rgb.Blue == 0, "should parse hex with # prefix correctly")
}

func TestRGBBoundaryValues(t *testing.T) {
	logger := got.New(t, "RGBBoundaryValues")

	logger.Case("test boundary RGB values")
	boundaryCases := []struct {
		r, g, b uint8
		hex     string
	}{
		{0, 0, 0, "000000"},
		{255, 255, 255, "FFFFFF"},
		{1, 1, 1, "010101"},
		{254, 254, 254, "FEFEFE"},
		{128, 128, 128, "808080"},
	}

	for _, tc := range boundaryCases {
		rgb := color.NewRGB(tc.r, tc.g, tc.b)
		hex := rgb.Hex()
		logger.Require(strings.ToUpper(hex) == tc.hex, "boundary value %d,%d,%d should produce hex %s, got %s", tc.r, tc.g, tc.b, tc.hex, hex)

		// Test round-trip conversion
		rgb2, err := color.Hex2RGB(hex)
		logger.Require(err == nil, "boundary value should convert back without error")
		logger.Require(rgb2.Red == tc.r && rgb2.Green == tc.g && rgb2.Blue == tc.b, "boundary value should round-trip correctly")
	}
}
