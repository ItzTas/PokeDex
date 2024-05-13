package main

import (
	"errors"
	"fmt"
	"slices"
)

func commandMapf(config *config, s string) error {
	if s == "" {
		return executeMapf(config)
	}
	if slices.Contains(paramsHelp, s) {
		showMapfHelp()
		return nil
	}
	return fmt.Errorf("unkown second parameter: %s", s)
}

func executeMapf(config *config) error {
	res, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
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

func showMapfHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["map"].name, GetCommands()["map"].description)
}

func commandMapb(config *config, s string) error {
	if s == "" {
		return executeMapb(config)
	}
	if slices.Contains(paramsHelp, s) {
		showMapbHelp()
		return nil
	}
	return fmt.Errorf("unkown second parameter: %s", s)
}

func showMapbHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["mapb"].name, GetCommands()["mapb"].description)
}


func executeMapb(config *config) error {
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