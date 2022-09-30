package server

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
		body       string
		method     string
		statusCode int
	}{
		{``, http.MethodPost, 200},
		{`75960609601501211`, http.MethodPost, 200},
		{`ERROR`, http.MethodGet, 400},
		{`ERROR`, http.MethodDelete, 400},
	}

	for _, v := range table {
		t.Run(v.body, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(v.method, "/barcode", strings.NewReader(v.body))

			barcodeHandler(w, r)

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
