package main

import (
	"net/http"
	"html/template"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
    films := map[string][]Film{
        "Films": {
            {Title: "The Godfather", Director: "Francis Ford Coppola"},
            {Title: "Blade Runner", Director: "Ridley Scott"},
            {Title: "The Thing", Director: "John Carpenter"},
        },
    }

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, films)
}
