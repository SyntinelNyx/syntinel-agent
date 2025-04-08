package setup

import (
	"os/exec"

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
