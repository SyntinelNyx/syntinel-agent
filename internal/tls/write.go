package tls

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func WriteCert(path string, cert *x509.Certificate) {
	logger.Info("Generating TLS certificate...")

	f, err := os.Create(path)
	if err != nil {
		logger.Fatal("Failed to create certificate file: %v", err)
	}
	defer f.Close()

	if err := pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}); err != nil {
		logger.Fatal("Failed to encode certificate file: %v", err)
	}
}

func WriteKey(path string, key *ecdsa.PrivateKey) {
	logger.Info("Generating TLS key...")

	f, err := os.Create(path)
	if err != nil {
		logger.Fatal("Failed to create key file: %v", err)
	}
	defer f.Close()

	b, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		logger.Fatal("Failed to marshal EC private key: %v", err)
	}

	if err := pem.Encode(f, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}); err != nil {
		logger.Fatal("Failed to encode private key file: %v", err)
	}
}
