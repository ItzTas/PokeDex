package main

import (
	"encoding/json"
	"io"
)

func commandMap() error {
	var config config
	return commandMapConfig(&config)	
}

func commandMapConfig(config *config) error {
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

