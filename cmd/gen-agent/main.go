package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/google/uuid"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/SyntinelNyx/syntinel-agent/internal/tls"
)

func main() {
	caCertPath, caKeyPath := parseFlags()
	caCert, caKey := loadTls(caCertPath, caKeyPath)

	agentID := "agent-" + uuid.New().String()
	cert, key := tls.CreateAgentCert(agentID, caCert, caKey)

	certName := agentID + ".crt"
	keyName := agentID + ".key"

	certPath := filepath.Join("internal", "data", certName)
	keyPath := filepath.Join("internal", "data", keyName)

	tls.WriteCert(certPath, cert)
	tls.WriteKey(keyPath, key)

	execTemplate(certName, keyName)
	buildAgent()
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

func execTemplate(certName, keyName string) {
	tmplData := map[string]string{
		"CertPath": certName,
		"KeyPath":  keyName,
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

func buildAgent() {
	cmd := exec.Command("make")
	_, err := cmd.Output()

	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Syntinel agent built successfully.")
}
