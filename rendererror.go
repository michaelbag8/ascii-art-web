package main

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	Code    int
	Message string
}

func renderError(w http.ResponseWriter, code int, message string) {
	templ, err := template.ParseFiles("templates/errors.html")
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
