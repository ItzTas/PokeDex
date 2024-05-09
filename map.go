package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	response, err := http.Get("https://pokeapi.co/api/v2/location-area/1/")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if int(response.Status[3]) >= 400 {
		errStr := fmt.Sprintf("Error fetching the request, code status: %v", response.Status)
		return errors.New(errStr)
	}	

	location, err := getLocationIdandName(response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%v \n%v  ", location.Id, location.Name)
	return nil
}

func getLocationIdandName(body io.Reader) (locationArea, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return locationArea{}, err
	}
	var location locationArea
	err = json.Unmarshal(b, &location)
	if err != nil {
		return locationArea{}, err
	}
	return location, nil
}

type locationArea struct {
	Id 	 int 	`json:"id"`
	Name string `json:"name"`
}

type Config struct {
	Next 	 string
	Previous string
}