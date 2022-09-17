package server

import (
	"flag"
	"io"
	"net/http"
)

var success []byte = []byte("OK")
var error []byte = []byte("ERROR")

func server() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/barcode", barcodeHandler)
	var hostname = ConnHost + ":" + ConnPort
	http.ListenAndServe(hostname, nil)
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
		req.Body = http.MaxBytesReader(w, req.Body, 256)
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
