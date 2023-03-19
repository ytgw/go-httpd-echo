package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head><title>Your HTTP Request Information</title></head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, html.EscapeString(r.Method))
	fmt.Fprintln(w, "<br>")
	fmt.Fprintln(w, html.EscapeString(r.URL.String()))
	fmt.Fprintln(w, "<br>")
	fmt.Fprintln(w, html.EscapeString(r.Proto))
	fmt.Fprintln(w, "<br>")
	fmt.Fprintln(w, html.EscapeString(r.Host))
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

func main() {
	// http.HandleFuncのpatternに"/"をセットすると全パターンにマッチする。
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
