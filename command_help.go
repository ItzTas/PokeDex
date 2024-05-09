package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("\nWelcome to the PokeDex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, c := range GetCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
		fmt.Println("")
	}
	return nil
}