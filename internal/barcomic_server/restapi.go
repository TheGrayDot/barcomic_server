package barcomic_server

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
)

var success []byte = []byte("OK")
var error []byte = []byte("ERROR")

func startRestApi() {
	fmt.Println("[*] Generating TLS certificate...")
	tlsCert := GenerateTLSCertificate()

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}

	server := http.Server{
		Addr:      config.addr + ":" + config.port,
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/barcode", barcodeHandler)
	http.HandleFunc("/", otherHandler)

	fmt.Println("[*] Starting HTTP server...")
	fmt.Println()
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func verboseLoggingHandler(req *http.Request) {
	// Only log if not unit testing
	if flag.Lookup("test.v") == nil {
		fmt.Printf("INFO: %s %s %s\n", req.Method, req.RemoteAddr, req.RequestURI)
	}
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	verboseLoggingHandler(req)
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write(success)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(error)
		return
	}
}

func barcodeHandler(w http.ResponseWriter, req *http.Request) {
	verboseLoggingHandler(req)
	if req.Method == "POST" {
		req.Body = http.MaxBytesReader(w, req.Body, 10000)
		buffer, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		bufferString := string(buffer)

		// Only send to keyvent if not unit testing
		if flag.Lookup("test.v") == nil {
			robotgo.TypeStr(bufferString)
			robotgo.KeyTap("enter")
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buffer)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(error)
		return
	}
}

func otherHandler(w http.ResponseWriter, req *http.Request) {
	verboseLoggingHandler(req)
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
}
