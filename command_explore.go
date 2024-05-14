package main

import (
	"fmt"
	"slices"
)

func commandExplore(cfg *config, s string) error {
	if slices.Contains(paramsHelp, s) {
		commandExploreHelp()
		return nil
	}
	if s != "" {
		return commandExploreExecute(cfg, s)
	}
	return fmt.Errorf("invalid second parameter: %s", s)
}

func commandExploreHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["explore"].name, GetCommands()["explore"].description)
}

func commandExploreExecute(cfg *config, s string) error {
	res, err := cfg.pokeapiClient.ListPokemon(s)
	if err != nil {
		return err
	}
	fmt.Println("")
	fmt.Printf("Exploring %s...\n", s)
	fmt.Println("Found Pokemon:")
	fmt.Println("")
	for _, p := range res.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}
	fmt.Println("")
	return nil
}