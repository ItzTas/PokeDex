package main

import (
	"errors"
	"fmt"
)

func commandMapf(config *config) error {
	res, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}
	fmt.Println("")
	for _, r := range res.Results {
		fmt.Println(r.Name)
	}
	fmt.Println("")
	if res.Next != nil {
		config.nextLocationsURL = res.Next
		config.prevLocationsURL = res.Previous
	} else {
		fmt.Println("Reached the end of locations or no next URL available.")
	}
	return nil
}


func commandMapb(config *config) error {
	if config.prevLocationsURL == nil {
		return errors.New("there are no previous cities")
	}
	res, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}
	fmt.Println("")
	for _, r := range res.Results {
		fmt.Println(r.Name)
	}
	fmt.Println("")
	config.nextLocationsURL = res.Next
	config.prevLocationsURL = res.Previous
	return nil
}
