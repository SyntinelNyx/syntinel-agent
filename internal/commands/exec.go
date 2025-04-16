package commands

import (
	"os/exec"
	"strings"
    "path/filepath"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func Exec(args ...string) string {
	if len(args) == 0 {
		logger.Error("No arguments provided")
		return ""
	}
	
    // Parse the first argument with space delimiter
    parsedArgs := strings.Split(args[0], " ")

    // Remaining arguments
    commandArgs := parsedArgs[1:]

    binaryPath, err := exec.LookPath(parsedArgs[0])
    if err == nil {
        cmd := exec.Command(binaryPath, commandArgs...)
        output, err := cmd.CombinedOutput() // Captures stdout and stderr
        if err != nil {
            logger.Error("Error:", err)
        }

        logger.Info(string(output))
        return string(output)

    } else {
        binaryPath := filepath.Join("/etc", "syntinel", parsedArgs[0])

        cmd := exec.Command(binaryPath, commandArgs...)
        output, err := cmd.CombinedOutput() // Captures stdout and stderr
        if err != nil {
            logger.Error("Error:", err)
        }

        logger.Info(string(output))
        return string(output)
    }

}
