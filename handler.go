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

var temp = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	err := temp.Execute(w, nil)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return

	}
	banner := r.FormValue("banner")
	input := r.FormValue("input")

	if input == "" {
		renderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		renderError(w, http.StatusInternalServerError, "empty banner")
		return
	}
	for _, ch := range input {
		if ch == '\n' {
			continue
		}
		for _, ch := range input {
			if ch < 32 || ch > 126 {
				renderError(w, http.StatusInternalServerError, "Non ascii character")
				return
			}
		}

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

	err = temp.Execute(w, content)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}
