package http

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/SyntinelNyx/syntinel-agent/internal/data"
)

func NewHTTPClientWithCustomCA() (*http.Client, error) {
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(data.CaCert) {
		return nil, fmt.Errorf("failed to add CA cert to pool")
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}, nil
}
