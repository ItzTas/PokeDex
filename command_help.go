package main

import (
	"fmt"
	"slices"
)

func commandHelp(co *config, s string) error {
	if s == "" {
		return executeCommandHelp()
	}
	if slices.Contains(paramsHelp, s) {
		commandHelpHelp()
		return nil
	}
	return fmt.Errorf("unkown second parameter: %s", s)
}

func executeCommandHelp() error {
	fmt.Println("\nWelcome to the PokeDex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, c := range GetCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
		fmt.Println("")
	}
	return nil
}

func commandHelpHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["help"].name, GetCommands()["help"].description)
}