package main

import (
	"fmt"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func() error	
}

func GetCommands() map[string]cliCommand { 
	return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Displays a list of commands",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exits the program",
            callback: commandExit,
        },
    }
}

func commandExit() error {
    fmt.Println("Exiting")
    return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a list of commands\nexit: Exits the pokedex") 
	return nil
}
