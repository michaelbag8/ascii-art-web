package main

import (
	"log"
	"net/http"
	"html/template"
)

func main() {

func homeHandler(w http.ResponseWriter, r *http.Request) {
template.ParseFiles
}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
	
}
