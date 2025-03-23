package shx

import (
	"fmt"
	"os/exec"
)

func RunScript(scriptPath string) {
	// Execute the script
	cmd := exec.Command("bash", scriptPath)
	output, err := cmd.CombinedOutput() // Captures stdout and stderr
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(output))
}

func RunBinary(binaryPath string, args ...string) {
	// Execute the binary
	cmd := exec.Command(binaryPath, args...)
	output, err := cmd.CombinedOutput() // Captures stdout and stderr
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(output))
}