package web

import (
	"fmt"
	"strconv"
	"strings"
)

func Filter(numOfMembers []string, minCreationDate, maxCreationDate string, minFirstAlbum, maxFirstAlbum string, location string, search string, artists []Artists, locations LocationsForMainData) ([]Artists, error) {
	filteredArtists := artists
	var tempFilter []Artists

	// members filter: if > 0 means that it was applied
	if len(numOfMembers) > 0 {
		//reseting temp filter value to store artists for the current filter only
		tempFilter = []Artists{}
		for _, artist := range artists {
			for _, num := range numOfMembers {
				// validating the numbers chosen
				numInt, err := strconv.Atoi(num)
				if err != nil || numInt < 1 || numInt > 8 {
					return nil, fmt.Errorf("400")
				}
				//checking if number of members matches the members and adding to array
				if len(artist.Members) == numInt {
					tempFilter = append(tempFilter, artist)
				}
			}
		}
		//storing value of tempfilter in filtered artist to use in next filter and final return
		filteredArtists = tempFilter
	}

	// creation date filter
	if minCreationDate != "" || maxCreationDate != "" {
		//reseting temp filter value to store artists for the current filter only
		tempFilter = []Artists{}
		//validating input
		if minCreationDate != "" {
			if minCreationDate < "1950" {
				return nil, fmt.Errorf("400")
			}
		}
		//validating input
		if maxCreationDate != "" {
			if maxCreationDate > "2024" {
				return nil, fmt.Errorf("400")
			}
		}
		//looping over filtered artist to filter results even more
		for _, artist := range filteredArtists {
			// this will check if both dates were applied or if only one of them was
			//if min date was set it will check if date is more than min and add, if max was set it will do the oposite
			//if both were set it will check for the values in between
			if (minCreationDate == "" || strconv.Itoa(artist.CreationDate) >= minCreationDate) &&
				(maxCreationDate == "" || strconv.Itoa(artist.CreationDate) <= maxCreationDate) {
				tempFilter = append(tempFilter, artist)
			}
		}
		//storing value of tempfilter in filtered artist to use in next filter and final return
		filteredArtists = tempFilter
	}

	// first album filter
	if minFirstAlbum != "" || maxFirstAlbum != "" {
		//reseting temp filter value to store artists for the current filter only
		tempFilter = []Artists{}
		//validation
		if minFirstAlbum != "" {
			if minFirstAlbum < "1950" {
				return nil, fmt.Errorf("400")
			}
		}
		//validation
		if maxFirstAlbum != "" {
			if maxFirstAlbum > "2024" {
				return nil, fmt.Errorf("400")
			}
		}
		//looping over filtered artist to filter results even more
		for _, artist := range filteredArtists {
			//extracting year only from the date to make comparison easier

			year := strings.Split(artist.FirstAlbum, "-")[2]
			// this will check if both dates were applied or if only one of them was
			//if min date was set it will check if date is more than min and add, if max was set it will do the oposite
			//if both were set it will check for the values in between
			if (minFirstAlbum == "" || year >= minFirstAlbum) &&
				(maxFirstAlbum == "" || year <= maxFirstAlbum) {
				tempFilter = append(tempFilter, artist)
			}
		}
		//storing value of tempfilter in filtered artist to use in next filter and final return
		filteredArtists = tempFilter
	}

	// location filter
	if location != "" {
		//reseting temp filter value to store artists for the current filter only
		tempFilter = []Artists{}
		artistLocations := make(map[int][]string)
		for _, index := range locations.Index {
			artistLocations[index.Id] = index.Locations
		}
		//looping over filtered artist to filter results even more
		for _, artist := range filteredArtists {
			if locs, exists := artistLocations[artist.Id]; exists {
				for _, loc := range locs {
					loc = strings.ReplaceAll(loc, "_", " ")
					loc = strings.ReplaceAll(loc, "-", ", ")
					if loc == location {
						tempFilter = append(tempFilter, artist)
					}
				}
			}
		}
		//storing value of tempfilter in filtered artist to use in next filter and final return
		filteredArtists = tempFilter
	}

	//for search query
	if search != "" {

		//turn query to lower case
		query := (strings.ToLower(search))
		// if search was an artist or band
		if strings.HasSuffix(search, "- Artist/Band") {

			name := strings.TrimSpace(strings.TrimSuffix(query, " - artist/band")) // extract name
			//reseting temp filter value to store artists for the current filter only
			tempFilter = []Artists{}

			for _, artist := range artists {
				if strings.Contains(strings.ToLower(artist.Name), name) {
					tempFilter = append(tempFilter, artist)
				}
			}
			//storing value of tempfilter in filtered artist to use in next filter and final return
			filteredArtists = tempFilter
		} else if strings.HasSuffix(search, " - Member") {
			query := strings.ToLower(search)
			member := strings.TrimSpace(strings.TrimSuffix(query, " - member"))
			//reseting temp filter value to store artists for the current filter only
			tempFilter = []Artists{}

			for _, artist := range artists {
				for _, mem := range artist.Members {
					if strings.ToLower(mem) == member {
						tempFilter = append(tempFilter, artist)
					}
				}
			}
			//storing value of tempfilter in filtered artist to use in next filter and final return
			filteredArtists = tempFilter
		} else if strings.HasSuffix(search, " - Creation Date") {
			query := strings.ToLower(search)
			creationDate, err := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(query, " - creation date")))
			if err != nil {
				return nil, fmt.Errorf("filter400")
			}
			//reseting temp filter value to store artists for the current filter only
			tempFilter = []Artists{}

			for _, artist := range artists {
				if artist.CreationDate == creationDate {
					tempFilter = append(tempFilter, artist)
				}
			}
			//storing value of tempfilter in filtered artist to use in next filter and final return
			filteredArtists = tempFilter
		} else if strings.HasSuffix(search, " - First Album") {
			query := strings.ToLower(search)
			firstAlbum := strings.TrimSpace(strings.TrimSuffix(query, " - first album"))
			//reseting temp filter value to store artists for the current filter only
			tempFilter = []Artists{}
			for _, artist := range artists {
				if strings.Contains(artist.FirstAlbum, firstAlbum) {
					tempFilter = append(tempFilter, artist)
				}
			}
			//storing value of tempfilter in filtered artist to use in next filter and final return
			filteredArtists = tempFilter
		} else if strings.HasSuffix(search, " - Location") {
			query := strings.ToLower(search)
			location := strings.TrimSpace(strings.TrimSuffix(query, " - location"))
			//reseting temp filter value to store artists for the current filter only
			tempFilter = []Artists{}

			artistLocations := make(map[int][]string)
			for _, index := range locations.Index {
				artistLocations[index.Id] = index.Locations
			}
			for _, artist := range artists {
				if locs, exists := artistLocations[artist.Id]; exists {
					for _, loc := range locs {
						if strings.Contains(strings.ToLower(loc), location) {
							tempFilter = append(tempFilter, artist)
						}
					}
				}
			}
			//storing value of tempfilter in filtered artist to use in next filter and final return
			filteredArtists = tempFilter
		} else {
			query := strings.ToLower(search)
			uniqueArtists := make(map[int]bool) // map to track unique artist IDs
			tempFilter := []Artists{}           // slice to store filtered artists

			// Filter by artist/band name
			for _, artist := range artists {
				if strings.Contains(strings.ToLower(artist.Name), strings.TrimSpace(query)) {
					if !uniqueArtists[artist.Id] {
						tempFilter = append(tempFilter, artist)
						uniqueArtists[artist.Id] = true
					}
				}
			}

			// Filter by member name
			for _, artist := range artists {
				for _, mem := range artist.Members {
					if strings.Contains(strings.ToLower(mem), strings.TrimSpace(query)) {
						if !uniqueArtists[artist.Id] {
							tempFilter = append(tempFilter, artist)
							uniqueArtists[artist.Id] = true
						}
					}
				}
			}

			// Filter by creation date
			creationDate := strings.TrimSpace(query)
			for _, artist := range artists {
				artistCreation := strconv.Itoa(artist.CreationDate)
				if strings.Contains(artistCreation, creationDate) {
					if !uniqueArtists[artist.Id] {
						tempFilter = append(tempFilter, artist)
						uniqueArtists[artist.Id] = true
					}
				}
			}

			// Filter by first album
			for _, artist := range artists {
				if strings.Contains(artist.FirstAlbum, strings.TrimSpace(query)) {
					if !uniqueArtists[artist.Id] {
						tempFilter = append(tempFilter, artist)
						uniqueArtists[artist.Id] = true
					}
				}
			}

			// Filter by location
			artistLocations := make(map[int][]string)
			for _, index := range locations.Index {
				artistLocations[index.Id] = index.Locations
			}
			for _, artist := range artists {
				if locs, exists := artistLocations[artist.Id]; exists {
					for _, loc := range locs {
						if strings.Contains(strings.ToLower(loc), strings.TrimSpace(query)) {
							if !uniqueArtists[artist.Id] {
								tempFilter = append(tempFilter, artist)
								uniqueArtists[artist.Id] = true
							}
						}
					}
				}
			}

			//storing value of tempfilter in filtered artist to use in next filter and final return
			filteredArtists = tempFilter

		}
	}

	if len(filteredArtists) == 0 {
		return nil, fmt.Errorf("filter400")
	}

	return filteredArtists, nil
}
