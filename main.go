package main

import (
	"fmt"
	"log"
	"net/http"

	funcs "groupie/funcs"
)

func main() {
	// for css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// main handlers
	http.HandleFunc("/", funcs.HomeHandler)
	http.HandleFunc("/artist", funcs.ArtistsHandler)
	http.HandleFunc("/artist/", funcs.DetailsHandler)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
