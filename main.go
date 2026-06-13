package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /{$}", homeHandler)
	http.HandleFunc("POST /ascii-art", asciiArtHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("server is runing.....http://localhost:9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}

}
