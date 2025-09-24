package uuid

import (
	"strings"

	"github.com/google/uuid"
)

// UUIDSafe generates a new random UUID and returns it as a standard string format.
// The returned string follows the standard UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//
// This function uses the Google UUID library and will panic if UUID generation fails.
// For production use, consider using the Sid struct for distributed ID generation.
//
// Example:
//
//	id := uuid.UUIDSafe()
//	fmt.Println(id) // Output: "550e8400-e29b-41d4-a716-446655440000"
func UUIDSafe() string {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return id.String()
}

// UUID generates a new random UUID and returns it as a standard string format.
// The returned string follows the standard UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//
// This function uses the Google UUID library and will return an error if UUID generation fails.
// For production use, consider using the Sid struct for distributed ID generation.
//
// Example:
//
//	id, err := uuid.UUID()
//	fmt.Println(id) // Output: "550e8400-e29b-41d4-a716-446655440000"
func UUID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// UUID32 generates a new random UUID and returns it as a 32-character string without hyphens.
// This is useful when you need a compact UUID format for database keys or URLs.
//
// Example:
//
//	id := uuid.UUID32()
//	fmt.Println(id) // Output: "550e8400e29b41d4a716446655440000"
func UUID32() string {
	return strings.ReplaceAll(UUIDSafe(), "-", "")
}
