package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", HomePageHandler)
    http.HandleFunc("/add-film/", AddFilmHandler)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
