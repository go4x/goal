package cmd

import (
	"os/exec"

	"github.com/gophero/logx"
)

func Exec(shell string, args ...string) string {
	cmd := exec.Command(shell, args...)
	logx.Debug("Execute command", "cmd", cmd.String())
	bs, err := cmd.CombinedOutput()
	if err != nil {
		logx.Error("Execute command failed", "cmd", cmd.String(), "error", err)
		return string(bs)
	}
	return string(bs)
}
