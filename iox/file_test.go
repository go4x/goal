package iox_test

import (
	"testing"

	"github.com/go4x/goal/iox"
	"github.com/go4x/got"
)

func TestExistsFile(t *testing.T) {
	lg := got.New(t, "test ExistsFile")

	lg.Case("give an existing file")
	f := "./file_test.go"
	lg.Require(iox.File.Exists(f), "should exist")

	lg.Case("give an existing dir, but is not a file")
	f = "."
	lg.Require(!iox.File.Exists(f), "should not exist")
}
