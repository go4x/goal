package base64x_test

import (
	"reflect"
	"testing"

	"github.com/go4x/goal/ciphers/base64x"
	"github.com/go4x/got"
)

var s = "abcdefghijjklmnopqrstuvwxyz0123456789`~-_=+[]\\{}|;':\",./<>?"
var sw = "YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB+LV89K1tdXHt9fDsnOiIsLi88Pj8="
var rw = "YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB+LV89K1tdXHt9fDsnOiIsLi88Pj8"
var ch = "hello, 这是中文!"
var sw1 = "aGVsbG8sIOi/meaYr+S4reaWhyE="
var rw1 = "aGVsbG8sIOi/meaYr+S4reaWhyE"

func TestBase(t *testing.T) {
	tl := got.New(t, "Base64")

	tl.Case("StdEncoding.Encode")
	gotStd := base64x.StdEncoding.Encode([]byte(s))
	wantStd := sw
	if gotStd != wantStd {
		tl.Fail("StdEncoding.Encode failed")
	} else {
		tl.Pass("StdEncoding.Encode passed")
	}

	tl.Case("RawStdEncoding.Encode")
	gotRaw := base64x.RawStdEncoding.Encode([]byte(s))
	wantRaw := rw
	if gotRaw != wantRaw {
		tl.Fail("RawStdEncoding.Encode failed")
	} else {
		tl.Pass("RawStdEncoding.Encode passed")
	}

	tl.Case("StdEncoding.Encode with Chinese")
	gotStd1 := base64x.StdEncoding.Encode([]byte(ch))
	wantStd1 := sw1
	if gotStd1 != wantStd1 {
		tl.Fail("StdEncoding.Encode with Chinese failed")
	} else {
		tl.Pass("StdEncoding.Encode with Chinese passed")
	}

	tl.Case("RawStdEncoding.Encode with Chinese")
	gotRaw1 := base64x.RawStdEncoding.Encode([]byte(ch))
	wantRaw1 := rw1
	if gotRaw1 != wantRaw1 {
		tl.Fail("RawStdEncoding.Encode with Chinese failed")
	} else {
		tl.Pass("RawStdEncoding.Encode with Chinese passed")
	}
}

func Test_base64raw_Decode(t *testing.T) {
	type args struct {
		str    string
		strict []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "case1", args: args{str: rw, strict: []bool{true}}, want: []byte(s), wantErr: false},
		{name: "case2", args: args{str: rw1, strict: []bool{true}}, want: []byte(ch), wantErr: false},
		{name: "case3", args: args{str: rw, strict: []bool{false}}, want: []byte(s), wantErr: false},
		{name: "case4", args: args{str: rw, strict: []bool{}}, want: []byte(s), wantErr: false},
		{name: "case5", args: args{str: "!!!", strict: []bool{true}}, want: []byte{}, wantErr: true},
		{name: "case6", args: args{str: "!!!", strict: []bool{false}}, want: []byte{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.RawStdEncoding
			got, err := ba.Decode(tt.args.str, tt.args.strict...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64raw_Encode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{b: []byte(s)}, want: rw},
		{name: "case2", args: args{b: []byte(ch)}, want: rw1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.RawStdEncoding
			if got := ba.Encode(tt.args.b); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64std_Decode(t *testing.T) {
	type args struct {
		str    string
		strict []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "case1", args: args{str: sw, strict: []bool{true}}, want: []byte(s), wantErr: false},
		{name: "case2", args: args{str: sw1, strict: []bool{true}}, want: []byte(ch), wantErr: false},
		{name: "case3", args: args{str: sw, strict: []bool{false}}, want: []byte(s), wantErr: false},
		{name: "case4", args: args{str: sw, strict: []bool{}}, want: []byte(s), wantErr: false},
		{name: "case5", args: args{str: "!!!", strict: []bool{true}}, want: []byte{}, wantErr: true},
		{name: "case6", args: args{str: "!!!", strict: []bool{false}}, want: []byte{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.StdEncoding
			got, err := ba.Decode(tt.args.str, tt.args.strict...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64std_Encode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{b: []byte(s)}, want: sw},
		{name: "case2", args: args{b: []byte(ch)}, want: sw1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.StdEncoding
			if got := ba.Encode(tt.args.b); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_base64raw_Encode1(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{b: []byte("hello world")}, want: "aGVsbG8gd29ybGQ"},
		{name: "case2", args: args{b: []byte("foo")}, want: "Zm9v"},
		{name: "case3", args: args{b: []byte("")}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.RawStdEncoding
			if got := ba.Encode(tt.args.b); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64raw_Decode1(t *testing.T) {
	type args struct {
		str    string
		strict []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "case1", args: args{str: "aGVsbG8gd29ybGQ", strict: []bool{true}}, want: []byte("hello world"), wantErr: false},
		{name: "case2", args: args{str: "Zm9v", strict: []bool{true}}, want: []byte("foo"), wantErr: false},
		{name: "case3", args: args{str: "", strict: []bool{true}}, want: []byte(""), wantErr: false},
		{name: "case4", args: args{str: "!!!", strict: []bool{true}}, want: []byte(""), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.RawStdEncoding
			got, err := ba.Decode(tt.args.str, tt.args.strict...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
