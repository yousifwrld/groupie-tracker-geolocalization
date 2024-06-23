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

func GetCoords(location string) (float64, float64, error) {
	apiKey := "AIzaSyA48Sdm5IPlocmTqUr-kApjrrKfwtzmwL0"
	location = strings.ReplaceAll(location, "-", ", ")
	location = strings.ReplaceAll(location, "_", " ")
	location = url.QueryEscape(location)

	apiURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", location, apiKey)

	apiResponse, err := GetData(apiURL)
	if err != nil {
		return 0.0, 0.0, err
	}

	var response GeocodingResponse

	err = json.Unmarshal(apiResponse, &response)
	if err != nil {
		return 0.0, 0.0, err
	}
	lat := response.Results[0].Geometry.Location.Lat
	lng := response.Results[0].Geometry.Location.Lng

	return lat, lng, nil
}
