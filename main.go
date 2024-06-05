package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", MyHandler)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
