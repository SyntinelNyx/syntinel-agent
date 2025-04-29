package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func expandHome(path string) string {
	if strings.HasPrefix(path, "~") {
		if home, err := os.UserHomeDir(); err == nil {
			return strings.Replace(path, "~", home, 1)
		}
	}
	return path
}

func Exec(command string) (string, error) {
	if len(command) == 0 {
		return "", fmt.Errorf("not enough argument")
	}

	args := strings.Fields(command)
	for i, arg := range args {
		args[i] = expandHome(arg)
	}

	execPath, err := exec.LookPath(args[0])
	if err != nil {
		return "", fmt.Errorf("error resolving path for %s: %v", args[0], err)
	}

	cmd := exec.Command(execPath, args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %v\n%s", err, string(output))
	}

	return string(output), nil
}
