package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func RunScript(args ...string) string {
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return ""
	}

	// Parse the first argument with space delimiter
	parsedArgs := strings.Split(args[0], " ")

	scriptPath := "/etc/syntinel/" + parsedArgs[0]
	parsedArgs = parsedArgs[1:] // Remove the script name from the arguments

	// Execute the script
	cmd := exec.Command("bash", append([]string{scriptPath}, parsedArgs...)...)
	output, err := cmd.CombinedOutput() // Captures stdout and stderr
	if err != nil {
		logger.Error("Error executing script: %v", err)
	}
	logger.Info("Script output: %s", string(output))

	return string(output)
}

func RunBinary(args ...string) string {
	if len(args) == 0 {
		logger.Error("No arguments provided")
		return ""
	}

	// Parse the first argument with space delimiter
	parsedArgs := strings.Split(args[0], " ")

	binaryPath := "/etc/syntinel/" + parsedArgs[0]
	parsedArgs = parsedArgs[1:] // Remove the script name from the arguments

	// Execute the binary
	cmd := exec.Command(binaryPath, parsedArgs...)
	output, err := cmd.CombinedOutput() // Captures stdout and stderr
	if err != nil {
		logger.Error("Error:", err)
	}
	logger.Info(string(output))

	return string(output)
}
