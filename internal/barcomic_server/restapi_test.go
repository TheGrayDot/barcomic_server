package barcomic_server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	table := []struct {
		body       string
		method     string
		statusCode int
	}{
		{`OK`, http.MethodGet, 200},
		{`ERROR`, http.MethodPost, 400},
		{`ERROR`, http.MethodDelete, 400},
	}

	for _, v := range table {
		t.Run(v.body, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(v.method, "/heath", nil)

			healthHandler(w, r)

			if w.Code != v.statusCode {
				t.Fatalf("Expected status code: %d, but got: %d", v.statusCode, w.Code)
			}

			body := strings.TrimSpace(w.Body.String())

			if body != v.body {
				t.Fatalf("Expected body to be: '%s', but got: '%s'", v.body, body)
			}
		})
	}
}

func TestBarcodeHandler(t *testing.T) {
	table := []struct {
		responseBody string
		requestBody  string
		method       string
		statusCode   int
	}{
		// Valid requests: upc, upc+ean2, upc+ean5
		{`759606096015`, `759606096015`, http.MethodPost, 200},
		{`75960609601501`, `75960609601501`, http.MethodPost, 200},
		{`75960609601501211`, `75960609601501211`, http.MethodPost, 200},
		// Invalid requests: body type, too long, too short, invalid barcode
		{`ERROR`, `notanumber`, http.MethodPost, 400},
		{`ERROR`, `7596060960150121101211`, http.MethodPost, 400},
		{`ERROR`, `75960609601`, http.MethodPost, 400},
		{`ERROR`, `759606096010`, http.MethodPost, 400},
		// Invalid requests: HTTP methods not allowed
		{`ERROR`, `75960609601501211`, http.MethodGet, 400},
		{`ERROR`, `75960609601501211`, http.MethodDelete, 400},
	}

	for _, v := range table {
		t.Run(v.requestBody, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(v.method, "/barcode", strings.NewReader(v.requestBody))

			barcodeHandler(w, r)

			if w.Code != v.statusCode {
				t.Fatalf("Expected status code: %d, but got: %d", v.statusCode, w.Code)
			}

			body := strings.TrimSpace(w.Body.String())

			if body != v.responseBody {
				t.Fatalf("Expected body to be: '%s', but got: '%s'", v.responseBody, body)
			}
		})
	}
}

func TestOtherHandler(t *testing.T) {
	table := []struct {
		url        string
		method     string
		statusCode int
	}{
		{`/doesntexist`, http.MethodGet, 404},
		{`/doesntexist`, http.MethodPost, 404},
	}

	for _, v := range table {
		t.Run(v.url, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(v.method, v.url, nil)

			otherHandler(w, r)

			if w.Code != v.statusCode {
				t.Fatalf("Expected status code: %d, but got: %d", v.statusCode, w.Code)
			}
		})
	}
}

func TestValidateUpc(t *testing.T) {
	table := []struct {
		barcode string
		result  bool
	}{
		{`844284008570`, true},
		{`844284008571`, false},
		{`84428400857000011`, true},
	}

	for _, v := range table {
		t.Run(v.barcode, func(t *testing.T) {
			testResult := validateUpc(v.barcode)

			if testResult != v.result {
				t.Fatalf("Expected status code: %t, but got: %t", v.result, testResult)
			}
		})
	}
}
