package uuid

import (
	"strings"

	"github.com/go4x/goal/valuex"
	"github.com/google/uuid"
)

func UUID() string {
	return valuex.Must(uuid.NewUUID()).String()
}

func UUID32() string {
	uid := UUID()
	return strings.ReplaceAll(uid, "-", "")
}
