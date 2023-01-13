package barcomic

import (
	"crypto/x509"
	"testing"
)

func TestGenerateTLSCertificate(t *testing.T) {
	// Use cert.go to generate a TLS certificate
	tlsCert := GenerateTLSCertificate()

	// Parse the generated TLS certificate
	cert, err := x509.ParseCertificate(tlsCert.Certificate[0])
	if err != nil {
		panic(err)
	}

	// Check cert common name
	serverHost := config.addr + ":" + config.port
	if cert.Subject.CommonName != serverHost {
		t.Fatalf("Expected host: %s, but got: %d", serverHost, cert.IPAddresses)
	}
}
