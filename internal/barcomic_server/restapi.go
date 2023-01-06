package barcomic_server

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	fmt.Printf("[*] Starting HTTP server...\n\n")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func verboseLoggingHandler(req *http.Request) {
	// Only log if not unit testing
	if flag.Lookup("test.v") == nil && config.verbose {
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
		// UPC + EAN5 is longest barcode with 17 chars
		// So set max byte length to 20
		req.Body = http.MaxBytesReader(w, req.Body, 20)
		buffer, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(error)
			return
		}

		bufferString := string(buffer)

		// Check request body is digits
		barcodeRegexp := regexp.MustCompile(`\d{12,17}`)
		match := barcodeRegexp.MatchString(bufferString)
		if !match {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(error)
			return
		}

		// Check request body is a valid UPC
		isValidUpc := validateUpc(bufferString)
		if !isValidUpc {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(error)
			return
		}

		// Only send to keyvent if not unit testing
		if flag.Lookup("test.v") == nil {
			robotgo.TypeStr(bufferString)
			robotgo.KeyTap("enter")
		}

		// Print barcode if verbose
		if config.verbose {
			fmt.Printf("INFO: %s\n", bufferString)
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

func validateUpc(barcode string) bool {
	// Extract first 11 digits of barcode
	barcodePrefix := barcode[0:11]
	// Extract last (check) digit from barcode
	checkDigit := barcode[11:12]

	// Sum all digits
	// Even digits are multiplied by 3
	sum := 0
	for i, v := range barcodePrefix {
		value, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			sum += 3 * value
		} else {
			sum += value
		}
	}

	result := (10 - sum%10) % 10
	checkDigitInt, _ := strconv.Atoi(string(checkDigit))
	return result == checkDigitInt
}
