package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func commandCatch(cfg *config, s string) error {
	if slices.Contains(paramsHelp, s) {
		commandCatchHelp()
		return nil
	}
	if s != "" {
		return commandCatchexecute(cfg, s)
	}
	return fmt.Errorf("invalid second parameter: %s", s)
}

func commandCatchexecute(cfg *config, s string) error {
	poke, err := cfg.pokeapiClient.StatsPokemon(s)
	if err != nil {
		return err
	}
	fmt.Printf("\nThrowing a Pokeball at %s...\n", s)
	prob := rand.Intn(360) + 1
	if prob < poke.BaseExperience {
		fmt.Printf("%s escaped!\n\n", s)
		return nil
	}
	fmt.Printf("%s was caught!\n\n", s)
	cfg.pokeAdd(s, poke)
	return nil
}

func commandCatchHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["catch"].name, GetCommands()["catch"].description)
}