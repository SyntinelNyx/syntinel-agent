package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/google/uuid"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/SyntinelNyx/syntinel-agent/internal/tls"
)

func main() {
	caCertPath, caKeyPath := parseFlags()
	caCert, caKey := loadTls(caCertPath, caKeyPath)

	id := uuid.New().String()
	agentID := "agent-" + id
	cert, key := tls.CreateAgentCert(agentID, caCert, caKey)

	caCertName := "ca-cert.pem"
	certName := agentID + ".crt"
	keyName := agentID + ".key"

	dataPath := filepath.Join("internal", "data")

	idPath := filepath.Join(dataPath, "agent-id")
	newCaCertPath := filepath.Join(dataPath, caCertName)
	certPath := filepath.Join(dataPath, certName)
	keyPath := filepath.Join(dataPath, keyName)

	if err := os.WriteFile(idPath, []byte(id), 0644); err != nil {
		logger.Fatal("Failed to write id to file: %v", err)
	}

	if err := copyFile(caCertPath, newCaCertPath); err != nil {
		fmt.Println("Error:", err)
	}

	tls.WriteCert(certPath, cert)
	tls.WriteKey(keyPath, key)

	execTemplate(caCertName, certName, keyName)
	buildAgent(id)
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("could not create destination file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("copy failed: %w", err)
	}

	return nil
}

func parseFlags() (string, string) {
	caCertPath := flag.String("ca-cert", "", "Path to CA certificate PEM file")
	caKeyPath := flag.String("ca-key", "", "Path to CA private key PEM file")
	flag.Parse()

	if *caCertPath == "" || *caKeyPath == "" {
		logger.Fatal("Both --ca-cert and --ca-key must be provided as arguments")
	}

	return *caCertPath, *caKeyPath
}

func loadTls(caCertPath string, caKeyPath string) (*x509.Certificate, *ecdsa.PrivateKey) {
	caCert, err := tls.LoadCert(caCertPath)
	if err != nil {
		logger.Fatal("Failed to load CA cert: %v", err)
	}

	caKey, err := tls.LoadKey(caKeyPath)
	if err != nil {
		logger.Fatal("Failed to load CA key: %v", err)
	}

	return caCert, caKey
}

func execTemplate(caCertName, certName, keyName string) {
	tmplData := map[string]string{
		"CaCertPath": caCertName,
		"CertPath":   certName,
		"KeyPath":    keyName,
	}

	tmplContent, err := os.ReadFile("internal/data/embed.go.tmpl")
	if err != nil {
		logger.Fatal("Failed to read template: %v", err)
	}

	tmpl, err := template.New("embed").Parse(string(tmplContent))
	if err != nil {
		logger.Fatal("Failed to parse template: %v", err)
	}

	out, err := os.Create("internal/data/embed.go")
	if err != nil {
		logger.Fatal("Failed to create output file: %v", err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, tmplData); err != nil {
		logger.Fatal("Failed to execute template: %v", err)
	}

	logger.Info("File embed.go generated successfully.")
}

func buildAgent(id string) {
	cmd := exec.Command("make", fmt.Sprintf("id=%s", strings.Split(id, "-")[0]))
	_, err := cmd.Output()

	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Syntinel agent built successfully.")
}
