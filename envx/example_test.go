package envx_test

import (
	"fmt"
	"os"
	"time"

	"github.com/go4x/goal/envx"
)

func ExampleGetDefault() {
	defer unsetForExample("GOAL_ENVX_EXAMPLE_MISSING")()
	fmt.Println(envx.GetDefault("GOAL_ENVX_EXAMPLE_MISSING", "fallback"))
	// Output: fallback
}

func ExampleRequire() {
	defer unsetForExample("GOAL_ENVX_EXAMPLE_MISSING")()
	value, err := envx.Require("GOAL_ENVX_EXAMPLE_MISSING")
	fmt.Println(value == "")
	fmt.Println(err != nil)
	// Output:
	// true
	// true
}

func ExampleGetInt() {
	defer unsetForExample("GOAL_ENVX_EXAMPLE_PORT")()
	fmt.Println(envx.GetInt("GOAL_ENVX_EXAMPLE_PORT", 8080))
	// Output: 8080
}

func ExampleGetDuration() {
	defer unsetForExample("GOAL_ENVX_EXAMPLE_TIMEOUT")()
	fmt.Println(envx.GetDuration("GOAL_ENVX_EXAMPLE_TIMEOUT", 5*time.Second))
	// Output: 5s
}

func ExampleGetSlice() {
	defer unsetForExample("GOAL_ENVX_EXAMPLE_SERVICES")()
	services := envx.GetSlice("GOAL_ENVX_EXAMPLE_SERVICES", ",", []string{"api", "worker"})
	fmt.Println(services)
	// Output: [api worker]
}

func unsetForExample(key string) func() {
	old, had := os.LookupEnv(key)
	_ = os.Unsetenv(key)
	return func() {
		if had {
			_ = os.Setenv(key, old)
			return
		}
		_ = os.Unsetenv(key)
	}
}
