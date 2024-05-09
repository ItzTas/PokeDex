package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Printf("Pokedex > ")
        scanner.Scan()
        input := scanner.Text()
        cleanInput := getCleanInput(input)
        commands := GetCommands()

        for _, word := range cleanInput {
            command, ok := commands[word]
            if !ok {
                fmt.Printf("Unknown command: %s\n", word)
                continue
            }
            err := command.callback()
            if err!= nil {
                fmt.Printf("Error: %s\n\n", err.Error())
            }
        }
    }
}

type cliCommand struct {
	name 		string
	description string
	callback 	func() error	
}

func GetCommands() map[string]cliCommand { 
	return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exits the program",
            callback: commandExit,
        },
    }
}

func getCleanInput(input string) []string {
    words := strings.Fields(strings.ToLower(input))
    return words
} 