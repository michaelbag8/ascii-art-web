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

	//'h''e''\r''\n''m''e'
	//input = strings.ReplaceAll(input, "\r\n", "\n")
	var cleanedInput strings.Builder 
	for _, c := range input{
		if c == '\r' || c == '\n'{
			continue
		}
		
		cleanedInput.WriteRune(c)
		//fmt.Printf("%q", c)

	}
//fmt.Println(cleanedInput.String())
	err :=ValidateInput(cleanedInput.String(), banner)
	if err!=nil {
		renderError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	bannerMap, err := LoadBanner("banner/" + banner + ".txt")
	if err != nil {
		renderError(w, http.StatusNotFound, "Not Found")
		return
	}

	result := Generate(cleanedInput.String(), bannerMap)

	content := PageData{
		Input:  cleanedInput.String(),
		Banner: banner,
		Result: result,
	}

	err = temp.Execute(w, content)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

}
