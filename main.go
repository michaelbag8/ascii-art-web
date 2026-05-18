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
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return

	}
	banner, input := r.FormValue("banner"), r.FormValue("input")

	err := ValidateInput(input, banner)
	if err != nil {
		renderError(w,http.StatusBadRequest, "Bad Request")
		return
	}

	bannerMap, err := LoadBanner("banner/" + banner + ".txt")
	if err != nil {
		renderError(w, http.StatusNotFound, "Not Found")
		return
	}

	result := Generate(input, bannerMap)

	content := PageData{
		Input:  input,
		Banner: banner,
		Result: result,
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = templ.Execute(w, content)
	if err != nil {
		renderError(w,http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

type PageData struct {
	Input  string
	Banner string
	Result string
}

type ErrorData struct {
	Code    int
	Message string
}

func renderError(w http.ResponseWriter, code int, message string) {
	templ, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)

	data := ErrorData{
		Code:    code,
		Message: message,
	}

	err = templ.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "InternalServerError")
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
