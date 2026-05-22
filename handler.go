package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Input  string
	Banner string
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	if r.Method != "GET" {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return

	}
	banner, input := r.FormValue("banner"), r.FormValue("input")

	err := ValidateInput(input, banner)
	if err != nil {
		renderError(w, http.StatusBadRequest, "Bad Request")
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
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}


