package checks

import (
	"github.com/aguirre-matteo/mtp-tui/errors"
	"os/exec"
)

func isCommandAvailable(command string) bool {
	cmd := exec.Command("which", command)
	err := cmd.Run()
	return err == nil
}

func checkDependencies() error {
	jmtpfs := isCommandAvailable("jmtpfs")
	if !jmtpfs {
		return errors.JmtpfsNotFound
	}

	fusermount := isCommandAvailable("fusermount")
	if !fusermount {
		return errors.FuseNotFound
	}
	return nil
}
