package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HttpVersion %s", r.Proto)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9292", nil))
}
