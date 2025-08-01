package iox_test

import (
	"testing"

	"github.com/gophero/goal/iox"
	"github.com/gophero/got"
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
