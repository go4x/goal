package value

// Must returns the value if error is nil, otherwise panics.
// This is useful for cases where you're certain the operation will succeed.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// If returns v1 if condition is true, otherwise returns v2.
// This is a generic conditional operator.
func If[T any](b bool, v1 T, v2 T) T {
	if b {
		return v1
	}
	return v2
}
