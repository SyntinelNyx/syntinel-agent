package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func DownloadFile(name string, data []byte) string {
	path := filepath.Join("/etc", "syntinel", name)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		logger.Error("Failed to open file %s: %v", path, err)
		return fmt.Sprintf("Failed to open file %s: %v", path, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		logger.Error("Failed to write to file %s: %v", path, err)
		return fmt.Sprintf("Failed to write to file %s: %v", path, err)
	}

	logger.Info("Chunk written to file %s successfully", name)
	return fmt.Sprintf("Chunk written to file %s successfully", name)
}
