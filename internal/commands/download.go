package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func DownloadFile(name string, data []byte) (string, error) {
	path := filepath.Join("/etc", "syntinel", name)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %v", path, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to write to file %s: %v", path, err)
	}

	return fmt.Sprintf("chunk written to file %s successfully", name), nil
}
