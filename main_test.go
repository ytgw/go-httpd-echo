package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func assertIsStatusOK(t *testing.T, actual int) {
	if actual != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", actual, http.StatusOK)
	}
}

func assertContainsBody(t *testing.T, body string, substring string) {
	if !strings.Contains(body, substring) {
		t.Errorf("unexpected body: got %v want %v", body, substring)
	}
}

func TestGetHandler(t *testing.T) {
	// setup
	request, err := http.NewRequest("GET", "http://example.com/", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	responseRecorder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(handler)

	// act
	testHandler.ServeHTTP(responseRecorder, request)

	// check
	assertIsStatusOK(t, responseRecorder.Code)
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Basic Request Information</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<tr><td>Method:</td><td>GET</td></tr>")
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Other Request Headers</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Request Body</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "Empty request body.")
}

func TestPostHandler(t *testing.T) {
	// setup
	request, err := http.NewRequest("POST", "http://example.com/", strings.NewReader("post data"))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	responseRecorder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(handler)

	// act
	testHandler.ServeHTTP(responseRecorder, request)

	// check
	assertIsStatusOK(t, responseRecorder.Code)
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Basic Request Information</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<tr><td>Method:</td><td>POST</td></tr>")
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Other Request Headers</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Request Body</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "post data")
}

func TestTCPExposerHandler(t *testing.T) {
	// setup
	request, err := http.NewRequest("GET", "http://echo.tcpexposer.com/", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Header.Add("X-Forwarded-Proto", "http")
	responseRecorder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(handler)

	// act
	testHandler.ServeHTTP(responseRecorder, request)

	// check
	assertIsStatusOK(t, responseRecorder.Code)
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Basic Request Information</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<tr><td>Method:</td><td>GET</td></tr>")

	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Request Headers Add by TCP Exposer</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<tr><td>X-Forwarded-Proto:</td><td>http</td></tr>")

	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Other Request Headers</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<tr><td>Accept:</td><td>text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8</td></tr>")

	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Request Body</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "Empty request body.")
}
