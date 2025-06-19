package http_client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURL(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	defer server.Close()

	body, err := GetURL(server.URL)
	if err != nil {
		t.Errorf("GetURL() returned error: %v", err)
		return
	}

	want := "Hello, World!"
	if body != want {
		t.Errorf("GetURL() = %q; want %q", body, want)
	}
}

func TestGetStatusCode(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	statusCode, err := GetStatusCode(server.URL)
	if err != nil {
		t.Errorf("GetStatusCode() returned error: %v", err)
		return
	}

	want := 200
	if statusCode != want {
		t.Errorf("GetStatusCode() = %d; want %d", statusCode, want)
	}
}

func TestCheckURL(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	if !CheckURL(server.URL) {
		t.Errorf("CheckURL() = false; want true")
	}
}

func TestGetHeaders(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	}))
	defer server.Close()

	headerValue, err := GetHeaders(server.URL, "Content-Type")
	if err != nil {
		t.Errorf("GetHeaders() returned error: %v", err)
		return
	}

	want := "application/json"
	if headerValue != want {
		t.Errorf("GetHeaders() = %q; want %q", headerValue, want)
	}
}
