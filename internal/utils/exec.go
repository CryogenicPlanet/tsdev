package utils

import (
	"os"
	"os/exec"
)

func ExecWithOutput(path string, name string, args ...string) error {

	cmd := exec.Command(name, args...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}

	cmd.Wait()
	return nil
}
