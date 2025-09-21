package iox_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go4x/goal/iox"
)

func TestWalkAllFiles(t *testing.T) {
	fs := iox.WalkDir("/Users/home/Downloads")
	fmt.Println(strings.Join(fs, "\n"))
}
