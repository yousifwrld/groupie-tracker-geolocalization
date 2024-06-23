package web

import (
	"html/template"
	"log"
	"net/http"
)

// function renders the error html page based on the error type
func ErrorPages(w http.ResponseWriter, r *http.Request, errMessage string, code int) {
	// sending http response with the status code to display the error code in the inspect console
	w.WriteHeader(code)

	type Error struct {
		Code       string
		Msg        string
		BackToPage string
		BackToMsg  string
	}

	switch errMessage {
	case "404":
		data := Error{
			Code:       "404",
			Msg:        "The page you are looking for could not be found.",
			BackToPage: "/",
			BackToMsg:  "Go back to home page",
		}
		temp, err := template.ParseFiles("templates/errors.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500: Internal Server Error!", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		return

	case "artist404":
		data := Error{
			Code:       "404",
			Msg:        "The Artist you are looking for could not be found.",
			BackToPage: "/artist",
			BackToMsg:  "Go back to artists page",
		}
		temp, err := template.ParseFiles("templates/errors.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500: Internal Server Error!", http.StatusNotFound)
			return
		}
		temp.Execute(w, data)
		return

	case "400":
		data := Error{
			Code:       "400",
			Msg:        "The server cannot process your request due to invalid or malformed data.",
			BackToPage: "/artist",
			BackToMsg:  "Go back to artist page",
		}
		temp, err := template.ParseFiles("templates/errors.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500: Internal Server Error!", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		return

	case "filter400":
		data := Error{
			Code:       "400",
			Msg:        "No results were found for filter query.",
			BackToPage: "/artist",
			BackToMsg:  "Go back to artists page",
		}
		temp, err := template.ParseFiles("templates/errors.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500: Internal Server Error!", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		return

	case "405":
		data := Error{
			Code:       "405",
			Msg:        "The requested method is not allowed.",
			BackToPage: "/",
			BackToMsg:  "Go back to home page",
		}

		temp, err := template.ParseFiles("templates/errors.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500: Internal Server Error!", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		return

	case "500":
		data := Error{
			Code:       "500",
			Msg:        "The server encounterd an error from our end and could not complete your request.",
			BackToPage: "/",
			BackToMsg:  "Go back to home page",
		}
		temp, err := template.ParseFiles("templates/errors.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500: Internal Server Error!", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		return
	}
}
