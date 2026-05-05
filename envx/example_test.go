package envx_test

import (
	"fmt"
	"time"

	"github.com/go4x/goal/envx"
)

func ExampleGetDefault() {
	fmt.Println(envx.GetDefault("GOAL_ENVX_EXAMPLE_MISSING", "fallback"))
	// Output: fallback
}

func ExampleRequire() {
	value, err := envx.Require("GOAL_ENVX_EXAMPLE_MISSING")
	fmt.Println(value == "")
	fmt.Println(err != nil)
	// Output:
	// true
	// true
}

func ExampleGetInt() {
	fmt.Println(envx.GetInt("GOAL_ENVX_EXAMPLE_PORT", 8080))
	// Output: 8080
}

func ExampleGetDuration() {
	fmt.Println(envx.GetDuration("GOAL_ENVX_EXAMPLE_TIMEOUT", 5*time.Second))
	// Output: 5s
}

func ExampleGetSlice() {
	services := envx.GetSlice("GOAL_ENVX_EXAMPLE_SERVICES", ",", []string{"api", "worker"})
	fmt.Println(services)
	// Output: [api worker]
}
