package main

import (
	"errors"
	"fmt"
	"slices"
)

func commandPokedex(cfg *config, s string) error {
	if slices.Contains(paramsHelp, s) {
		commandPokedexHelp()
		return nil
	}
	if s == "" {
		return commandPokedexExecute(cfg)
	}
	return fmt.Errorf("unkown second parameter: %s", s)
}

func commandPokedexHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["pokedex"].name, GetCommands()["pokedex"].description)
}

func commandPokedexExecute(cfg *config) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("there are no pokemons in the pokedex")
	}
	fmt.Println("")
	fmt.Println("Your Pokedex:")
	for p := range cfg.pokedex {
		fmt.Println("  -", p)
	}
	fmt.Println("")
	return nil
}