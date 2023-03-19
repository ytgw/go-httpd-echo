package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	fmt.Fprintln(w, r.URL.String())
	fmt.Fprintln(w, r.Proto)
	fmt.Fprintln(w, r.Host)
}

func main() {
	// http.HandleFuncのpatternに"/"をセットすると全パターンにマッチする。
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
