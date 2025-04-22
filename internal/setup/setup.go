package setup

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/SyntinelNyx/syntinel-agent/internal/data"
	https "github.com/SyntinelNyx/syntinel-agent/internal/http"
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/SyntinelNyx/syntinel-agent/internal/sysinfo"
)

type EnrollRequest struct {
	UUID      string           `json:"uuid"`
	Info      sysinfo.HostInfo `json:"hostInfo"`
	Root_User string           `json:"rootUser"`
}

func Start() {
	ip := flag.String("ip", "", "Server IP and port (e.g. 127.0.0.1:8443)")
	root := flag.String("root", "", "Root username")

	flag.Parse()

	initFile := filepath.Join("/etc", "syntinel", "init")
	if fileExists(initFile) {
		logger.Info("Agent already initialized, skipping enrollment...")
		return
	}

	if *ip == "" || *root == "" {
		logger.Fatal("Both --ip and --root flags are required")
	}

	commands := []string{"bash", "kopia", "trivy"}
	if err := checkCommands(commands); err != nil {
		logger.Fatal("Failed to resolve dependencies: %v", err)
	}
	logger.Info("All dependencies found.")

	if err := createDirectory(); err != nil {
		logger.Fatal("Failed to create directory: %v", err)
	}
	logger.Info("Syntinel directory created successfully.")

	if err := enroll(*ip, *root); err != nil {
		logger.Fatal("Enroll failed: %v", err)
	}
	logger.Info("Enrollment successful!")

	if _, err := os.Create(initFile); err != nil {
		logger.Fatal("Failed to create file at %s, please create it manually.", initFile)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func checkCommands(commands []string) error {
	for _, cmd := range commands {
		if _, err := exec.LookPath(cmd); err != nil {
			return fmt.Errorf("dependency '%s' is not installed or is not in path. Please install the dependency and try again", cmd)
		}
	}
	return nil
}

func createDirectory() error {
	path := filepath.Join("/etc", "syntinel")
	uploadPath := filepath.Join("/etc", "syntinel", "upload")

	err := os.Mkdir(path, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory %s: %v", path, err)
	}

	err = os.Mkdir(uploadPath, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory %s: %v", uploadPath, err)
	}

	originalPath := os.Getenv("PATH")
	newPath := strings.Join([]string{uploadPath, originalPath}, ":")
	if err := os.Setenv("PATH", newPath); err != nil {
		return fmt.Errorf("failed to add %s to path: %v", uploadPath, err)
	}

	return nil
}

func enroll(ip string, rootUser string) error {
	info, err := sysinfo.System()
	if err != nil {
		return fmt.Errorf("failed to collect system info: %w", err)
	}

	reqBody := EnrollRequest{
		UUID:      string(data.ID),
		Info:      *info,
		Root_User: rootUser,
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	endpoint := fmt.Sprintf("https://%s/v1/api/agent/enroll", ip)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client, err := https.NewHTTPClientWithCustomCA()
	if err != nil {
		return fmt.Errorf("failed to add custom ca: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("enroll request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("enroll failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
