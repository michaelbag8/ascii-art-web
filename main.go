package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("server is runing.....http://localhost:8080")
	http.HandleFunc("GET /", homeHandler)
	http.HandleFunc("POST /ascii-art", asciiArtHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
