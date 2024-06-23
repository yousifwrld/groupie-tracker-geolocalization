package web

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPages(w, r, "405", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		ErrorPages(w, r, "404", http.StatusNotFound)
		return
	}

	// rendering the html file
	temp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println(err)
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		return
	}

	// executing the html template and passing the artist array to it if there are no errors
	err = temp.Execute(w, nil)
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		return
	}
}
