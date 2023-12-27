package main

import "os/exec"

// ExecCmd shell: bash or sh
func ExecCmd(shell, cmdStr string) (stdoutStderr string, err error) {
	out, err := exec.Command(shell, "-c", cmdStr).CombinedOutput()
	return string(out), err
}
