package commands

import (
    "fmt"
    "os"
	"path/filepath"

    "github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func DownloadFile(name string, data byte) string {
	path := filepath.Join("/etc", "syntinel", name)

    err := os.WriteFile(path, []byte{data}, 0644)

    if err != nil {
        logger.Error("Failed to write file %s: %v", path, err)
        return fmt.Sprintf("Failed to write file %s: %v", path, err)
    }

    logger.Info("File %s downloaded successfully", name)

    return fmt.Sprintf("File %s downloaded successfully", name)
}