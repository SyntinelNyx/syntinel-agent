package commands

import (
    "fmt"
    "os/exec"
)

func RunScript(scriptPath string, args ...string) string {
    scriptPath = "/etc/syntinel/" + scriptPath

    // Execute the script
    cmd := exec.Command("bash", append([]string{scriptPath}, args...)...)
    output, err := cmd.CombinedOutput() // Captures stdout and stderr
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(string(output))

    return string(output)
}

func RunBinary(binaryPath string, args ...string) string{
    binaryPath = "/etc/syntinel/" + binaryPath

    // Execute the binary
    cmd := exec.Command(binaryPath, args...)
    output, err := cmd.CombinedOutput() // Captures stdout and stderr
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(string(output))

    return string(output)
}
