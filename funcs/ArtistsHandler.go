package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPages(w, r, "405", http.StatusMethodNotAllowed)
		return
	}

	//fetching artist related data
	data, err := GetData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//unmarshalling data into the struct we created
	var artists []Artists
	err = json.Unmarshal(data, &artists)
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//fetching locations related data
	locationData, err := GetData("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//unmarshalling into struct
	var locations LocationsForMainData
	err = json.Unmarshal(locationData, &locations)
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//maps to store unique locations, firstalbum, and creation dates
	uniqueLocations := make(map[string]struct{})
	uniqueSearchLocations := make(map[string]struct{})
	uniqueFirstAlbums := make(map[string]struct{})
	uniqueCreationDates := make(map[int]struct{})

	//finding unique locations
	for _, index := range locations.Index {
		for _, loc := range index.Locations {
			//for filter options removing underscores and dashes to make them more readable
			modifiedLoc := strings.ReplaceAll(loc, "_", " ")
			modifiedLoc = strings.ReplaceAll(modifiedLoc, "-", ", ")

			if _, found := uniqueLocations[modifiedLoc]; !found {
				uniqueLocations[modifiedLoc] = struct{}{}
			}

			//for location suggestions
			if _, found := uniqueSearchLocations[loc]; !found {
				uniqueSearchLocations[loc] = struct{}{}
			}
		}
	}

	//finding unique artist creation dates and first album dates
	for _, artist := range artists {
		if _, found := uniqueFirstAlbums[artist.FirstAlbum]; !found {
			uniqueFirstAlbums[artist.FirstAlbum] = struct{}{}
		}

		if _, found := uniqueCreationDates[artist.CreationDate]; !found {
			uniqueCreationDates[artist.CreationDate] = struct{}{}
		}
	}

	//variable to store new non duplicated data
	var sortedLocations, SearchLocations, FirstAlbums []string
	var CreationDates []int

	//looping over filter locations and storing in slice
	for loc := range uniqueLocations {
		sortedLocations = append(sortedLocations, loc)
	}
	sort.Strings(sortedLocations) // sorting in alphabatical order

	//looping over search suggestion locations and storing in slice
	for loc := range uniqueSearchLocations {
		SearchLocations = append(SearchLocations, loc)
	}
	//looping over first album dates and storing in slice
	for album := range uniqueFirstAlbums {
		FirstAlbums = append(FirstAlbums, album)
	}

	//looping over creation dates and storing in slice
	for date := range uniqueCreationDates {
		CreationDates = append(CreationDates, date)
	}

	numOfMembers := r.URL.Query()[("numMembers")]
	minCreationDate := r.FormValue("minCreationDate")
	maxCreationDate := r.FormValue("maxCreationDate")
	minFirstAlbum := r.FormValue("minFirstAlbum")
	maxFirstAlbum := r.FormValue("maxFirstAlbum")
	location := r.FormValue("locations")
	search := r.FormValue("search")

	//function to filter the artists, on initial load when no filters are chosen it will load all artists
	filteredArtists, err := Filter(numOfMembers, minCreationDate, maxCreationDate, minFirstAlbum, maxFirstAlbum, location, search, artists, locations)
	if err != nil {
		//calling error page with the proper error retured from the function
		ErrorPages(w, r, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	allData := ArtistPageData{
		Artists:            filteredArtists,
		NumOfMembers:       [8]string{"1", "2", "3", "4", "5", "6", "7", "8"},
		Locations:          sortedLocations,
		SearchLocations:    SearchLocations,
		SearchFirstAlbum:   FirstAlbums,
		SearchCreationDate: CreationDates,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, allData)
	if err != nil {
		ErrorPages(w, r, "500", http.StatusInternalServerError)
		log.Println(err)
	}
}
