package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPages(w, r, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// extracting ID from the GET request
	artistID := strings.TrimPrefix(string(r.URL.Path), "/artist")

	// general artist data
	aristData, err := GetData("https://groupietrackers.herokuapp.com/api/artists" + artistID) // fetching data from the endpoint
	if err != nil {
		if err.Error() == "404" {
			ErrorPages(w, r, "404", http.StatusNotFound)
			log.Println(err)
			return
		} else {
			ErrorPages(w, r, "500", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	var artist Artists

	err = json.Unmarshal(aristData, &artist) // unmarshalling the json data into the struct we created
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if artist.Id == 0 {
		ErrorPages(w, r, "400", http.StatusBadRequest)
		log.Println("Error: Artist not found")
		return
	}

	// locations data
	locData, err := GetData("https://groupietrackers.herokuapp.com/api/locations" + artistID) // fetching data from the endpoint
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var locations Locations

	err = json.Unmarshal(locData, &locations) // unmarshalling the json data into the struct we created
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//getting the latitude and longitude of each location and sotring
	var cords []Cords
	for _, loc := range locations.Location {
		lat, lng, err := GetCords(loc) //calling function for each location
		if err != nil {
			ErrorPages(w, r, "500", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		//mapping and storing each value in its propper field
		cords = append(cords, Cords{Name: loc, Lat: lat, Lng: lng})
	}

	// dates data
	datesData, err := GetData("https://groupietrackers.herokuapp.com/api/dates" + artistID) // fetching data from the endpoint
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var dates Dates

	err = json.Unmarshal(datesData, &dates) // unmarshalling the json data into the struct we created
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// relations data
	relationsData, err := GetData("https://groupietrackers.herokuapp.com/api/relation" + artistID)
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var relations Relations

	err = json.Unmarshal(relationsData, &relations) // unmarshalling the json data into the struct we created
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// deleting stars from the dates

	for i := range dates.Date {
		dates.Date[i] = strings.TrimPrefix(dates.Date[i], "*")
	}

	// combining all the data that we fetched into a single struct so we can use it with the html file
	Details := AllDetails{
		Artist:    artist,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
		Cords:     cords,
	}

	temp, err := template.ParseFiles("templates/details.html")
	if err != nil {
		log.Println(err)
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		return
	}

	// executing the html template and passing the artist array to it if there are no errors
	err = temp.Execute(w, Details)
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		return
	}
}
