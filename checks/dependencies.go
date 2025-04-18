package checks

import (
  "os/exec"
  "fmt"
)


func isCommandAvailable(command string) bool {
  cmd := exec.Command("which", command)
  err := cmd.Run()
  return err == nil
}



func checkDependencies() error {
  jmtpfs := isCommandAvailable("jmtpfs")
  if !jmtpfs {
    return fmt.Errorf("Command jmtpfs is not available")
  }

  fusermount := isCommandAvailable("fusermount")
  if !fusermount {
    return fmt.Errorf("Command fusermount is not available")
  }
  return nil
}
