package server

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var success []byte = []byte("OK")
var error []byte = []byte("ERROR")

func server() {
	fmt.Println("[*] Generating TLS certificate...")
	tlsCert := GenerateTLSCertificate()

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		// Other options
	}

	server := http.Server{
		Addr:      ConnHost + ":" + ConnPort,
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/barcode", barcodeHandler)

	fmt.Println("[*] Starting HTTP server...")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write(success)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(error)
	}
}

func barcodeHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.Body = http.MaxBytesReader(w, req.Body, 10000)
		buffer, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		bufferString := string(buffer)

		if flag.Lookup("test.v") == nil {
			// Only send to keyvent if not unit testing
			SendKeys(bufferString)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buffer)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(error)
	}
}
