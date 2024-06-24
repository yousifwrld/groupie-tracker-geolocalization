package web

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type GeocodingResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetCords(location string) (float64, float64, error) {
	apiKey := "AIzaSyA48Sdm5IPlocmTqUr-kApjrrKfwtzmwL0"

	//formating location into better format
	location = strings.ReplaceAll(location, "-", ", ")
	location = strings.ReplaceAll(location, "_", " ")
	//escaping the location so we can use it as a query in the api
	location = url.QueryEscape(location)

	//endpoint that we will send the request to, containing the location and the key
	apiURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)

	//getting the data
	apiData, err := GetData(apiURL)
	if err != nil {
		return 0.0, 0.0, err
	}

	//unmarshalling data into the struct
	var response GeocodingResponse

	err = json.Unmarshal(apiData, &response)
	if err != nil {
		return 0.0, 0.0, err
	}

	//checking for errors
	if response.Status != "OK" {
		return 0.0, 0.0, fmt.Errorf("API response status: %s", response.Status)
	}

	if len(response.Results) == 0 {
		return 0.0, 0.0, fmt.Errorf("no results found")
	}

	//taking coordinates from response and returning them
	lat := response.Results[0].Geometry.Location.Lat
	lng := response.Results[0].Geometry.Location.Lng

	return lat, lng, nil
}
