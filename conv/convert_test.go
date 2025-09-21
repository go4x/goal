package conv_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/go4x/goal/conv"
)

func ExampleInt64ToHex() {
	var x int64 = math.MinInt64
	var y int64 = math.MaxInt64
	var z int64 = math.MaxInt32

	fmt.Println(x, y, z) // -9223372036854775808 9223372036854775807 2147483647

	xs := conv.Int64ToHex(x)
	ys := conv.Int64ToHex(y)
	zs := conv.Int64ToHex(z)
	fmt.Println(xs, ys, zs) // -8000000000000000 7fffffffffffffff 7fffffff
}

func ExampleHexToInt64() {
	xs := "-8000000000000000"
	ys := "7fffffffffffffff"
	zs := "7fffffff"

	x, _ := conv.HexToInt64(xs)
	y, _ := conv.HexToInt64(ys)
	z, _ := conv.HexToInt64(zs)
	fmt.Println(x, y, z) // -9223372036854775808 9223372036854775807 2147483647
}

func TestHexToInt64(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "string param: 0", args: args{src: "0"}, want: 0},
		{name: "string param: -1", args: args{src: "-1"}, want: -1},
		{name: "string param: 10", args: args{src: "A"}, want: 10},
		{name: "string param: 11", args: args{src: "B"}, want: 11},
		{name: "string param: 17", args: args{src: "11"}, want: 17},
		{name: "string param: -8000000000000000", args: args{src: "-8000000000000000"}, want: math.MinInt64},
		{name: "string param: 7fffffffffffffff", args: args{src: "7fffffffffffffff"}, want: math.MaxInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.HexToInt64(tt.args.src)
			if err != nil {
				t.Errorf("HexToInt64() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("HexToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32ToStr(t *testing.T) {
	type args struct {
		src int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "2147483647"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-2147483648"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.Int32ToStr(tt.args.src); got != tt.want {
				t.Errorf("Int32ToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToHex(t *testing.T) {
	type args struct {
		src int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "7fffffff"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-80000000"},
		{name: "param: 9223372036854775807", args: args{src: math.MaxInt64}, want: "7fffffffffffffff"},
		{name: "param: -9223372036854775808", args: args{src: math.MinInt64}, want: "-8000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.Int64ToHex(tt.args.src); got != tt.want {
				t.Errorf("Int64ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToStr(t *testing.T) {
	type args struct {
		src int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "2147483647"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-2147483648"},
		{name: "param: 9223372036854775807", args: args{src: math.MaxInt64}, want: "9223372036854775807"},
		{name: "param: -9223372036854775808", args: args{src: math.MinInt64}, want: "-9223372036854775808"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.Int64ToStr(tt.args.src); got != tt.want {
				t.Errorf("Int64ToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToStr(t *testing.T) {
	type args struct {
		src int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "2147483647"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-2147483648"},
		{name: "param: 9223372036854775807", args: args{src: math.MaxInt64}, want: "9223372036854775807"},
		{name: "param: -9223372036854775808", args: args{src: math.MinInt64}, want: "-9223372036854775808"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.IntToStr(tt.args.src); got != tt.want {
				t.Errorf("IntToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntsToStr(t *testing.T) {
	type args struct {
		is []int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "case 1", args: args{is: []int64{-1, 0, 1}}, want: []string{"-1", "0", "1"}},
		{name: "case 2", args: args{is: []int64{math.MinInt64, 0, math.MaxInt64}}, want: []string{"-9223372036854775808", "0", "9223372036854775807"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.IntsToStr(tt.args.is); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntsToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToFloat64(t *testing.T) {
	type args struct {
		amount string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "case 1", args: args{amount: "0"}, want: 0},
		{name: "case 2", args: args{amount: "1"}, want: 1},
		{name: "case 3", args: args{amount: "-1"}, want: -1},
		{name: "case 4", args: args{amount: "3.1415"}, want: 3.1415},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrToFloat64(tt.args.amount)
			if err != nil {
				t.Errorf("StrToFloat64() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("StrToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test safe version functions
func TestStrToIntSafe(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		defaultValue int
		want         int
	}{
		{name: "valid string", str: "123", defaultValue: 0, want: 123},
		{name: "invalid string", str: "abc", defaultValue: 999, want: 999},
		{name: "empty string", str: "", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToIntSafe(tt.str, tt.defaultValue); got != tt.want {
				t.Errorf("StrToIntSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test error handling
func TestStrToIntError(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		wantErr bool
	}{
		{name: "valid string", str: "123", wantErr: false},
		{name: "invalid string", str: "abc", wantErr: true},
		{name: "empty string", str: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := conv.StrToInt(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToInt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test all string to int conversion functions
func TestStrToInt32(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    int32
		wantErr bool
	}{
		{name: "valid string", str: "123", want: 123, wantErr: false},
		{name: "invalid string", str: "abc", want: 0, wantErr: true},
		{name: "empty string", str: "", want: 0, wantErr: true},
		{name: "max int32", str: "2147483647", want: math.MaxInt32, wantErr: false},
		{name: "min int32", str: "-2147483648", want: math.MinInt32, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrToInt32(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToInt64(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    int64
		wantErr bool
	}{
		{name: "valid string", str: "123", want: 123, wantErr: false},
		{name: "invalid string", str: "abc", want: 0, wantErr: true},
		{name: "empty string", str: "", want: 0, wantErr: true},
		{name: "max int64", str: "9223372036854775807", want: math.MaxInt64, wantErr: false},
		{name: "min int64", str: "-9223372036854775808", want: math.MinInt64, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrToInt64(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToUint(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    uint
		wantErr bool
	}{
		{name: "valid string", str: "123", want: 123, wantErr: false},
		{name: "invalid string", str: "abc", want: 0, wantErr: true},
		{name: "negative string", str: "-1", want: 0, wantErr: true},
		{name: "empty string", str: "", want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrToUint(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToUint64(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    uint64
		wantErr bool
	}{
		{name: "valid string", str: "123", want: 123, wantErr: false},
		{name: "invalid string", str: "abc", want: 0, wantErr: true},
		{name: "negative string", str: "-1", want: 0, wantErr: true},
		{name: "empty string", str: "", want: 0, wantErr: true},
		{name: "max uint64", str: "18446744073709551615", want: math.MaxUint64, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrToUint64(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToUint32(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    uint32
		wantErr bool
	}{
		{name: "valid string", str: "123", want: 123, wantErr: false},
		{name: "invalid string", str: "abc", want: 0, wantErr: true},
		{name: "negative string", str: "-1", want: 0, wantErr: true},
		{name: "empty string", str: "", want: 0, wantErr: true},
		{name: "max uint32", str: "4294967295", want: math.MaxUint32, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrToUint32(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToBool(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    bool
		wantErr bool
	}{
		{name: "true string", str: "true", want: true, wantErr: false},
		{name: "false string", str: "false", want: false, wantErr: false},
		{name: "1 string", str: "1", want: true, wantErr: false},
		{name: "0 string", str: "0", want: false, wantErr: false},
		{name: "invalid string", str: "abc", want: false, wantErr: true},
		{name: "empty string", str: "", want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrToBool(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrsToInt(t *testing.T) {
	tests := []struct {
		name    string
		ss      []string
		want    []int64
		wantErr bool
	}{
		{name: "valid strings", ss: []string{"1", "2", "3"}, want: []int64{1, 2, 3}, wantErr: false},
		{name: "empty slice", ss: []string{}, want: nil, wantErr: false},
		{name: "nil slice", ss: nil, want: nil, wantErr: false},
		{name: "invalid string in slice", ss: []string{"1", "abc", "3"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StrsToInt(tt.ss)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrsToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrsToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test all safe version functions
func TestStrToInt32Safe(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		defaultValue int32
		want         int32
	}{
		{name: "valid string", str: "123", defaultValue: 0, want: 123},
		{name: "invalid string", str: "abc", defaultValue: 999, want: 999},
		{name: "empty string", str: "", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToInt32Safe(tt.str, tt.defaultValue); got != tt.want {
				t.Errorf("StrToInt32Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToInt64Safe(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		defaultValue int64
		want         int64
	}{
		{name: "valid string", str: "123", defaultValue: 0, want: 123},
		{name: "invalid string", str: "abc", defaultValue: 999, want: 999},
		{name: "empty string", str: "", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToInt64Safe(tt.str, tt.defaultValue); got != tt.want {
				t.Errorf("StrToInt64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToUintSafe(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		defaultValue uint
		want         uint
	}{
		{name: "valid string", str: "123", defaultValue: 0, want: 123},
		{name: "invalid string", str: "abc", defaultValue: 999, want: 999},
		{name: "negative string", str: "-1", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToUintSafe(tt.str, tt.defaultValue); got != tt.want {
				t.Errorf("StrToUintSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToUint64Safe(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		defaultValue uint64
		want         uint64
	}{
		{name: "valid string", str: "123", defaultValue: 0, want: 123},
		{name: "invalid string", str: "abc", defaultValue: 999, want: 999},
		{name: "negative string", str: "-1", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToUint64Safe(tt.str, tt.defaultValue); got != tt.want {
				t.Errorf("StrToUint64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToUint32Safe(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		defaultValue uint32
		want         uint32
	}{
		{name: "valid string", str: "123", defaultValue: 0, want: 123},
		{name: "invalid string", str: "abc", defaultValue: 999, want: 999},
		{name: "negative string", str: "-1", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToUint32Safe(tt.str, tt.defaultValue); got != tt.want {
				t.Errorf("StrToUint32Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToFloat64Safe(t *testing.T) {
	tests := []struct {
		name         string
		amount       string
		defaultValue float64
		want         float64
	}{
		{name: "valid string", amount: "123.45", defaultValue: 0, want: 123.45},
		{name: "invalid string", amount: "abc", defaultValue: 999.99, want: 999.99},
		{name: "empty string", amount: "", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToFloat64Safe(tt.amount, tt.defaultValue); got != tt.want {
				t.Errorf("StrToFloat64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToInt64Safe(t *testing.T) {
	tests := []struct {
		name         string
		src          string
		defaultValue int64
		want         int64
	}{
		{name: "valid hex", src: "FF", defaultValue: 0, want: 255},
		{name: "invalid hex", src: "GG", defaultValue: 999, want: 999},
		{name: "empty string", src: "", defaultValue: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.HexToInt64Safe(tt.src, tt.defaultValue); got != tt.want {
				t.Errorf("HexToInt64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToBoolSafe(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		defaultValue bool
		want         bool
	}{
		{name: "valid true", str: "true", defaultValue: false, want: true},
		{name: "valid false", str: "false", defaultValue: true, want: false},
		{name: "invalid string", str: "abc", defaultValue: true, want: true},
		{name: "empty string", str: "", defaultValue: false, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToBoolSafe(tt.str, tt.defaultValue); got != tt.want {
				t.Errorf("StrToBoolSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
