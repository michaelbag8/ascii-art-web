package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "error: parsing file", http.StatusInternalServerError)
		return
	}
	templ.Execute(w, nil)

}

func main() {
	

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
