package setup

import (
	"os/exec"
	"path/filepath"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func CheckCommands() {
	commands := []string{"bash", "kopia", "trivy"}
	for _, cmd := range commands {
		if _, err := exec.LookPath(cmd); err != nil {
			logger.Fatal("Dependency '%s' is not installed or is not in path. Please install the dependency and try again.", cmd)
		}
	}
	logger.Info("All dependencies found.")
}

func CreateDirectories() {
	path := filepath.Join("/etc", "syntinel")

    if err := exec.Command("mkdir", "-p", path).Run(); err != nil {
        logger.Fatal("Failed to create directory %s ", path)
    }
    logger.Info("Syntinel directory created successfully.")
}