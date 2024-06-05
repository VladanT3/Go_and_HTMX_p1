package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", h1)

    log.Fatal(http.ListenAndServe(":8000", nil))
}

func h1(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world\n")
    io.WriteString(w, r.Method)
}
