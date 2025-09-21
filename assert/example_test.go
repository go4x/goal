package assert_test

import (
	"testing"

	"github.com/go4x/goal/assert"
)

// Example demonstrates the correct usage of the assert package in tests.
func Example() {
	// This example shows how to use assert functions in test code.
	// DO NOT use these functions in production code or goroutines.
}

// TestCorrectUsage demonstrates the correct way to use assert functions.
func TestCorrectUsage(t *testing.T) {
	// ✅ CORRECT: Use in test functions
	assert.True(true)
	assert.Nil(nil)
	assert.NotBlank("hello")
	assert.Equals(1, 1)
	assert.DeepEquals([]int{1, 2}, []int{1, 2})
}

// TestIncorrectUsage demonstrates what NOT to do.
func TestIncorrectUsage(t *testing.T) {
	// ❌ INCORRECT: Don't use in goroutines
	// go func() {
	//     assert.True(false) // This will panic and terminate the entire program!
	// }()

	// ❌ INCORRECT: Don't use for input validation in production code
	// func ProcessUser(user *User) {
	//     assert.NoneNil(user) // This will panic and crash the program!
	// }

	// ✅ CORRECT: Use proper error handling in production code
	// func ProcessUser(user *User) error {
	//     if user == nil {
	//         return errors.New("user cannot be nil")
	//     }
	//     return nil
	// }
}

// TestPanicBehavior demonstrates that assert functions panic on failure.
func TestPanicBehavior(t *testing.T) {
	// This test shows that assert functions panic when conditions fail
	// In real tests, you would use defer recover() to test panic behavior

	// Example of testing panic behavior:
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic, but didn't get one")
			}
		}()
		assert.True(false) // This will panic
	}()
}
