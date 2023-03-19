package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, html.EscapeString(r.Method))
	fmt.Fprintln(w, html.EscapeString(r.URL.String()))
	fmt.Fprintln(w, html.EscapeString(r.Proto))
	fmt.Fprintln(w, html.EscapeString(r.Host))
}

func main() {
	// http.HandleFuncのpatternに"/"をセットすると全パターンにマッチする。
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
