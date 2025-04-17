package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

func Exec(command string) (string, error) {
	if len(command) == 0 {
		return "", fmt.Errorf("not enough argument")
	}

	args := strings.Fields(command)
	execPath, err := exec.LookPath(args[0])
	if err != nil {
		return "", fmt.Errorf("Error resolving path for %s: %v", args[0], err)
	}

	cmd := exec.Command(execPath, args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Error executing command: %v", err)
	}

	return string(output), nil
}
