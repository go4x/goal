package iox_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gophero/goal/iox"
	"github.com/gophero/got"
)

func TestPathExists(t *testing.T) {
	lg := got.New(t, "test PathExists")

	lg.Case("give an exists path")
	path := "../goal"
	lg.Require(iox.Path.PathExists(path), "given path should exist")

	lg.Case("give an none exists path")
	path = "/Users/haha"
	lg.Require(!iox.Path.PathExists(path), "given path should not exist")
}

func TestExecPath(t *testing.T) {
	execpath, err := os.Executable() // 获得程序路径
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(execpath)
	fmt.Println(dir)

	s, _ := os.Getwd()
	println(s)

	println(iox.Path.ExecPath())
	println(iox.Path.CurrentPath())
	println(iox.Path.ProjectPath())
}
