package envx

import (
	"reflect"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	t.Setenv("GOAL_ENVX_GET", "value")

	if got := Get("GOAL_ENVX_GET"); got != "value" {
		t.Fatalf("Get() = %q, want %q", got, "value")
	}
}

func TestExists(t *testing.T) {
	t.Setenv("GOAL_ENVX_EXISTS", "")

	if !Exists("GOAL_ENVX_EXISTS") {
		t.Fatal("Exists() = false, want true for empty but set variable")
	}
	if Exists("GOAL_ENVX_MISSING") {
		t.Fatal("Exists() = true, want false for missing variable")
	}
}

func TestGetDefault(t *testing.T) {
	t.Setenv("GOAL_ENVX_DEFAULT", "")

	if got := GetDefault("GOAL_ENVX_DEFAULT", "fallback"); got != "" {
		t.Fatalf("GetDefault() = %q, want empty string for set variable", got)
	}
	if got := GetDefault("GOAL_ENVX_MISSING", "fallback"); got != "fallback" {
		t.Fatalf("GetDefault() = %q, want fallback", got)
	}
}

func TestRequire(t *testing.T) {
	t.Setenv("GOAL_ENVX_REQUIRE", "value")

	got, err := Require("GOAL_ENVX_REQUIRE")
	if err != nil {
		t.Fatalf("Require() error = %v", err)
	}
	if got != "value" {
		t.Fatalf("Require() = %q, want %q", got, "value")
	}

	if _, err := Require("GOAL_ENVX_MISSING"); err == nil {
		t.Fatal("Require() error = nil, want error for missing variable")
	}
}

func TestGetInt(t *testing.T) {
	t.Setenv("GOAL_ENVX_INT", "42")
	t.Setenv("GOAL_ENVX_INT_BAD", "abc")

	if got := GetInt("GOAL_ENVX_INT", 7); got != 42 {
		t.Fatalf("GetInt() = %d, want 42", got)
	}
	if got := GetInt("GOAL_ENVX_INT_BAD", 7); got != 7 {
		t.Fatalf("GetInt() = %d, want fallback", got)
	}
	if got := GetInt("GOAL_ENVX_MISSING", 7); got != 7 {
		t.Fatalf("GetInt() = %d, want fallback", got)
	}
}

func TestGetBool(t *testing.T) {
	t.Setenv("GOAL_ENVX_BOOL", "true")
	t.Setenv("GOAL_ENVX_BOOL_BAD", "maybe")

	if got := GetBool("GOAL_ENVX_BOOL", false); !got {
		t.Fatal("GetBool() = false, want true")
	}
	if got := GetBool("GOAL_ENVX_BOOL_BAD", true); !got {
		t.Fatal("GetBool() = false, want fallback true")
	}
	if got := GetBool("GOAL_ENVX_MISSING", true); !got {
		t.Fatal("GetBool() = false, want fallback true")
	}
}

func TestGetDuration(t *testing.T) {
	t.Setenv("GOAL_ENVX_DURATION", "150ms")
	t.Setenv("GOAL_ENVX_DURATION_BAD", "soon")

	if got := GetDuration("GOAL_ENVX_DURATION", time.Second); got != 150*time.Millisecond {
		t.Fatalf("GetDuration() = %v, want 150ms", got)
	}
	if got := GetDuration("GOAL_ENVX_DURATION_BAD", time.Second); got != time.Second {
		t.Fatalf("GetDuration() = %v, want fallback", got)
	}
	if got := GetDuration("GOAL_ENVX_MISSING", time.Second); got != time.Second {
		t.Fatalf("GetDuration() = %v, want fallback", got)
	}
}

func TestGetSlice(t *testing.T) {
	t.Setenv("GOAL_ENVX_SLICE", "api, worker, , scheduler")
	t.Setenv("GOAL_ENVX_SLICE_SEMICOLON", "a;b;c")

	tests := []struct {
		name     string
		key      string
		sep      string
		fallback []string
		want     []string
	}{
		{
			name: "comma",
			key:  "GOAL_ENVX_SLICE",
			sep:  ",",
			want: []string{"api", "worker", "scheduler"},
		},
		{
			name: "custom separator",
			key:  "GOAL_ENVX_SLICE_SEMICOLON",
			sep:  ";",
			want: []string{"a", "b", "c"},
		},
		{
			name:     "missing",
			key:      "GOAL_ENVX_MISSING",
			sep:      ",",
			fallback: []string{"fallback"},
			want:     []string{"fallback"},
		},
		{
			name: "empty separator defaults to comma",
			key:  "GOAL_ENVX_SLICE",
			sep:  "",
			want: []string{"api", "worker", "scheduler"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSlice(tt.key, tt.sep, tt.fallback)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetSlice() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
