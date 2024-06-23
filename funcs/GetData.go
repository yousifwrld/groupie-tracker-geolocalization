package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// function will read the data from the api and return them as bytes so that we can use the returned data in each handler
func GetData(url string) ([]byte, error) {
	// sending a GET request to the api to get the data
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	// checking if something occurred when sending the request
	if response.StatusCode != http.StatusOK {
		if response.StatusCode == 404 {

			log.Println("Error: Status 404 for data get request")
			return nil, fmt.Errorf("404")
		}
		log.Println("Error: Status not ok for data get request")
		return nil, fmt.Errorf("status not ok")
	}

	// reading the data from the http response
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dataBytes, nil
}
