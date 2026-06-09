package main

import (
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Input  string
	Banner string
	Result string
}

var temp = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {

	err := temp.Execute(w, nil)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {

	banner := r.FormValue("banner")
	input := r.FormValue("input")

	input = strings.ReplaceAll(input, "\r", "")//this

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

	err = temp.Execute(w, content)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}
