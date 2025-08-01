package base64x_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gophero/goal/ciphers/base64x"
	"github.com/gophero/got"
)

const (
	plaintext = "abcdefghijjklmnopqrstuvwxyz0123456789`~-_=+[]\\{}|;':\",./<>?"
	base64enc = "YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB-LV89K1tdXHt9fDsnOiIsLi88Pj8"
)

func Test_Base64UrlEncode(t *testing.T) {
	lg := got.New(t, "Base64UrlEncode")
	lg.Case("testing base64s.URLEncoding.Encode")
	enc := base64x.RawURLEncoding.Encode([]byte(plaintext))
	fmt.Println(enc)
	lg.Require(base64enc == enc, "result should match")
}

func Test_Base64UrlDecode(t *testing.T) {
	lg := got.New(t, "Base64UrlDecode")
	lg.Case("testing base64s.URLEncoding.Decode")
	dec, err := base64x.RawURLEncoding.Decode(base64enc, true)
	lg.Require(err == nil, "requires no error")
	lg.Require(reflect.DeepEqual([]byte(plaintext), dec), "results should be matched")
}

func Test_Base64Url_Encode(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal string",
			args: args{input: []byte("hello world")},
			want: "aGVsbG8gd29ybGQ=",
		},
		{
			name: "empty string",
			args: args{input: []byte("")},
			want: "",
		},
		{
			name: "with special chars",
			args: args{input: []byte("foo+bar/baz?")},
			want: "Zm9vK2Jhci9iYXo_",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := base64x.URLEncoding.Encode(tt.args.input)
			if got != tt.want {
				t.Errorf("URLEncoding.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Base64Url_Decode(t *testing.T) {
	type args struct {
		input  string
		strict []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "normal string",
			args:    args{input: "aGVsbG8gd29ybGQ=", strict: []bool{true}},
			want:    []byte("hello world"),
			wantErr: false,
		},
		{
			name:    "empty string",
			args:    args{input: "", strict: []bool{true}},
			want:    []byte(""),
			wantErr: false,
		},
		{
			name:    "with special chars",
			args:    args{input: "Zm9vK2Jhci9iYXo_", strict: []bool{true}},
			want:    []byte("foo+bar/baz?"),
			wantErr: false,
		},
		{
			name:    "invalid base64 string",
			args:    args{input: "!!!", strict: []bool{true}},
			want:    []byte{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base64x.URLEncoding.Decode(tt.args.input, tt.args.strict...)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLEncoding.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLEncoding.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RawBase64Url_Encode(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal string",
			args: args{input: []byte("hello world")},
			want: "aGVsbG8gd29ybGQ",
		},
		{
			name: "empty string",
			args: args{input: []byte("")},
			want: "",
		},
		{
			name: "with special chars",
			args: args{input: []byte("foo+bar/baz?")},
			want: "Zm9vK2Jhci9iYXo_",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := base64x.RawURLEncoding.Encode(tt.args.input)
			if got != tt.want {
				t.Errorf("RawURLEncoding.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RawBase64Url_Decode(t *testing.T) {
	type args struct {
		input  string
		strict []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "normal string",
			args:    args{input: "aGVsbG8gd29ybGQ", strict: []bool{true}},
			want:    []byte("hello world"),
			wantErr: false,
		},
		{
			name:    "empty string",
			args:    args{input: "", strict: []bool{true}},
			want:    []byte(""),
			wantErr: false,
		},
		{
			name:    "with special chars",
			args:    args{input: "Zm9vK2Jhci9iYXo_", strict: []bool{true}},
			want:    []byte("foo+bar/baz?"),
			wantErr: false,
		},
		{
			name:    "invalid base64 string",
			args:    args{input: "!!!", strict: []bool{true}},
			want:    []byte{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base64x.RawURLEncoding.Decode(tt.args.input, tt.args.strict...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLEncoding.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawURLEncoding.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
