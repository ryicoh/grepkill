package grepkill

import (
	"os"
)

func Kill(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	proc.Kill()
	return nil
}
