package main

import (
	"fmt"
)

func commandMapf(config *config) error {
	res, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}
	for _, r := range res.Results{
		fmt.Println(r.Name)
	}
	if res.Next != nil {
		config.nextLocationsURL = res.Next
	} else {
		fmt.Println("Reached the end of locations or no next URL available.")
	}
	return nil
}
