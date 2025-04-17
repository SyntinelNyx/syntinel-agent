package setup

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func CreateDirectory() {
	path := filepath.Join("/etc", "syntinel")

	err := os.Mkdir(path, 0755)
	if err != nil {
		if os.IsExist(err) {
			logger.Info("Directory %s already exists.", path)
		} else {
			logger.Fatal("Failed to create directory %s: %v", path, err)
		}
	} else {
		logger.Info("Syntinel directory created successfully.")
	}

	originalPath := os.Getenv("PATH")
	newPath := strings.Join([]string{path, originalPath}, ":")
	if err := os.Setenv("PATH", newPath); err != nil {
		logger.Fatal("Failed to add %s to PATH: %v", path, err)
	}
}
