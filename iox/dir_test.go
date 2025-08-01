package iox_test

import (
	"path/filepath"
	"testing"

	"github.com/gophero/goal/iox"
	"github.com/gophero/got"
	"github.com/stretchr/testify/assert"
)

func TestExistsDir(t *testing.T) {
	lg := got.New(t, "test ExistsDir")

	lg.Case("give a existing dir")
	f := "/Users/home/goal/"
	b, err := iox.Dir.Exists(f)
	if err != nil {
		t.Error(err)
	}
	lg.Require(b, "should exist")

	lg.Case("give a existing file, but not a director")
	f = "/Users/home/goal/iox/file_test.go"
	b, err = iox.Dir.Exists(f)
	if err != nil {
		t.Error(err)
	}
	lg.Require(!b, "should not exist")
}

func TestAppendSep(t *testing.T) {
	s := "/a/b"
	r := iox.Dir.AppendSep(s)
	assert.True(t, r == s+string(filepath.Separator))
	s = "/a/b/c/"
	r = iox.Dir.AppendSep(s)
	assert.True(t, r == s)
}
