package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ValidateInput(input string, banner string) error {
	if input == "" {
		return fmt.Errorf("input is empty")
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return fmt.Errorf("not a banner method")
	}
	for _, ch := range input {
		if ch == '\n' {
			continue
		}
		if ch < 32 || ch > 126 {
			return fmt.Errorf("non ascii character: %c", ch)
		}
	}
	return nil
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}
	banner, input := r.FormValue("banner"), r.FormValue("input")

	err := ValidateInput(input, banner)
	if err != nil{
		http.Error(w, "Bad request", http.StatusMethodNotAllowed)
		return
	}
}

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
