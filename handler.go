package main

import (
	"html/template"
	"net/http"
	"time"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
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

func AddFilmHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(time.Second * 1)
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "film-list-element", Film{title, director})
}
