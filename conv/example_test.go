package conv_test

import (
	"fmt"
	"log"

	"github.com/go4x/goal/conv"
)

// Example demonstrates basic string to int conversion with error handling.
func ExampleStrToInt() {
	value, err := conv.StrToInt("123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	// Output: 123
}

// ExampleStrToInt_error demonstrates error handling for invalid input.
func ExampleStrToInt_error() {
	_, err := conv.StrToInt("abc")
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: strconv.Atoi: parsing "abc": invalid syntax
}

// ExampleStrToIntSafe demonstrates safe conversion with default value.
func ExampleStrToIntSafe() {
	value1 := conv.StrToIntSafe("123", 0)
	value2 := conv.StrToIntSafe("abc", 999)
	fmt.Println(value1, value2)
	// Output: 123 999
}

// ExampleStrToFloat64 demonstrates string to float64 conversion.
func ExampleStrToFloat64() {
	value, err := conv.StrToFloat64("3.14159")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f\n", value)
	// Output: 3.14
}

// ExampleHexToInt64_demo demonstrates hexadecimal string to int64 conversion.
func ExampleHexToInt64_demo() {
	value, err := conv.HexToInt64("FF")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	// Output: 255
}

// ExampleStrToBool demonstrates string to bool conversion.
func ExampleStrToBool() {
	value1, _ := conv.StrToBool("true")
	value2, _ := conv.StrToBool("1")
	value3, _ := conv.StrToBool("false")
	fmt.Println(value1, value2, value3)
	// Output: true true false
}

// ExampleStrsToInt demonstrates string slice to int64 slice conversion.
func ExampleStrsToInt() {
	values, err := conv.StrsToInt([]string{"1", "2", "3"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)
	// Output: [1 2 3]
}

// ExampleStrsToInt_error demonstrates error handling for invalid string in slice.
func ExampleStrsToInt_error() {
	_, err := conv.StrsToInt([]string{"1", "abc", "3"})
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: strconv.ParseInt: parsing "abc": invalid syntax
}

// ExampleIntToStr demonstrates int to string conversion.
func ExampleIntToStr() {
	str := conv.IntToStr(123)
	fmt.Println(str)
	// Output: 123
}

// ExampleInt64ToHex_demo demonstrates int64 to hexadecimal string conversion.
func ExampleInt64ToHex_demo() {
	hex := conv.Int64ToHex(255)
	fmt.Println(hex)
	// Output: ff
}

// ExampleIntsToStr demonstrates int64 slice to string slice conversion.
func ExampleIntsToStr() {
	strings := conv.IntsToStr([]int64{1, 2, 3})
	fmt.Println(strings)
	// Output: [1 2 3]
}

// ExampleStrToUint64Safe demonstrates safe unsigned integer conversion.
func ExampleStrToUint64Safe() {
	value1 := conv.StrToUint64Safe("123", 0)
	value2 := conv.StrToUint64Safe("-1", 999)  // negative number
	value3 := conv.StrToUint64Safe("abc", 999) // invalid string
	fmt.Println(value1, value2, value3)
	// Output: 123 999 999
}

// ExampleStrToBoolSafe demonstrates safe boolean conversion.
func ExampleStrToBoolSafe() {
	value1 := conv.StrToBoolSafe("true", false)
	value2 := conv.StrToBoolSafe("abc", true) // invalid string
	value3 := conv.StrToBoolSafe("", false)   // empty string
	fmt.Println(value1, value2, value3)
	// Output: true true false
}

// ExampleStrToFloat64Safe demonstrates safe float conversion.
func ExampleStrToFloat64Safe() {
	value1 := conv.StrToFloat64Safe("3.14", 0.0)
	value2 := conv.StrToFloat64Safe("abc", 99.99)
	fmt.Println(value1, value2)
	// Output: 3.14 99.99
}

// ExampleHexToInt64Safe demonstrates safe hexadecimal conversion.
func ExampleHexToInt64Safe() {
	value1 := conv.HexToInt64Safe("FF", 0)
	value2 := conv.HexToInt64Safe("GG", 999) // invalid hex
	fmt.Println(value1, value2)
	// Output: 255 999
}
