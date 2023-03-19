package main

import (
	"log"
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

func createRequest(method string, url string, body string, headers [][2]string) *http.Request {
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	for _, header := range headers {
		request.Header.Add(header[0], header[1])
	}
	return request
}

func assertContainCommonHeader(t *testing.T, body string) {
	assertContainsBody(t, body, "<h2>Other Request Headers</h2>")
	assertContainsBody(t, body, "<tr><td>Accept:</td><td>text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8</td></tr>")
	assertContainsBody(t, body, "<tr><td>Accept-Encoding:</td><td>gzip, deflate, br</td></tr>")
}

func TestGetHandler(t *testing.T) {
	// setup
	request := createRequest("GET", "http://example.com/", "", [][2]string{})
	responseRecorder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(handler)

	// act
	testHandler.ServeHTTP(responseRecorder, request)

	// check
	assertIsStatusOK(t, responseRecorder.Code)
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Basic Request Information</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<tr><td>Method:</td><td>GET</td></tr>")

	assertContainCommonHeader(t, responseRecorder.Body.String())

	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Request Body</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "Empty request body.")
}

func TestPostHandler(t *testing.T) {
	// setup
	request := createRequest("POST", "http://example.com/", "post data", [][2]string{})
	responseRecorder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(handler)

	// act
	testHandler.ServeHTTP(responseRecorder, request)

	// check
	assertIsStatusOK(t, responseRecorder.Code)
	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Basic Request Information</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "<tr><td>Method:</td><td>POST</td></tr>")

	assertContainCommonHeader(t, responseRecorder.Body.String())

	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Request Body</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "post data")
}

func TestTCPExposerHandler(t *testing.T) {
	// setup
	tcpExposerHeaders := [][2]string{{"X-Forwarded-Proto", "http"}, {"X-Forwarded-Port", "80"}}
	request := createRequest("GET", "http://echo.tcpexposer.com/", "", tcpExposerHeaders)
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

	assertContainCommonHeader(t, responseRecorder.Body.String())

	assertContainsBody(t, responseRecorder.Body.String(), "<h2>Request Body</h2>")
	assertContainsBody(t, responseRecorder.Body.String(), "Empty request body.")
}
