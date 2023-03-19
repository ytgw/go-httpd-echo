package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

func makeTableElement(keyValues [][2]string) string {
	var tableElement string = ""
	tableElement += "<table><tbody>\n"
	for _, keyValue := range keyValues {
		key := html.EscapeString(keyValue[0])
		val := html.EscapeString(keyValue[1])
		tableElement += fmt.Sprintf("<tr><td>%s:</td><td>%s</td></tr>\n", key, val)
	}
	tableElement += "</tbody></table>"
	return tableElement
}

func makeHTMLBody(r *http.Request, h1 string) string {
	var htmlBody string = ""
	htmlBody += fmt.Sprintln("<h1>" + html.EscapeString(h1) + "</h1>")

	// Basic Information
	htmlBody += fmt.Sprintln("<h2>Basic Request Information</h2>")
	basicInfos := [][2]string{
		{"Method", r.Method},
		{"URL", r.URL.String()},
		{"Protocol", r.Proto},
		{"Host", r.Host},
	}
	htmlBody += fmt.Sprintln(makeTableElement(basicInfos))

	// Other Headers
	htmlBody += fmt.Sprintln("<h2>Other Request Headers</h2>")
	var headers = [][2]string{}
	for key, values := range r.Header {
		headerDict := [2]string{key, strings.Join(values, ", ")}
		headers = append(headers, headerDict)
	}
	sort.Slice(headers, func(i, j int) bool { return headers[i][0] < headers[j][0] })
	htmlBody += fmt.Sprintln(makeTableElement(headers))

	// Body
	htmlBody += fmt.Sprintln("<h2>Request Body</h2>")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		htmlBody += fmt.Sprintln("Failed to read request body.")
	} else {
		requestBody := string(body)
		if len(requestBody) > 0 {
			htmlBody += fmt.Sprintln(html.EscapeString(requestBody))
		} else {
			htmlBody += fmt.Sprintln("Empty request body.")
		}
	}

	// return
	return htmlBody
}

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Your HTTP Request Information"
	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head><title>"+html.EscapeString(title)+"</title></head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, makeHTMLBody(r, title))
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

func main() {
	// http.HandleFuncのpatternに"/"をセットすると全パターンにマッチする。
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
