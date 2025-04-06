package trivy

import (
	"fmt"
	"os/exec"
)

//variable
var binaryPath = "./data/dependencies/trivy"

func DeepScan(path string) string{
	// Execute the binary
	cmd := exec.Command(binaryPath, "fs", "-f", "json", "--scanners", "vuln", path)
	// fmt.Println("Executing command:", cmd.String())	
	output, err := cmd.CombinedOutput() // Captures stdout and stderr
	if err != nil {
		fmt.Println("Error:", err)
	}
	// fmt.Println(string(output))

	return (string(output))
}
