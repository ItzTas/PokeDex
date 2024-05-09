package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ItzTass/PokeDex/pokeapi"
)

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Printf("Pokedex > ")
        scanner.Scan()
        
        words := getCleanInput(scanner.Text())
        if len(words) == 0 {
            continue
        }
        commandName := words[0]
        command, exist := GetCommands()[commandName] 
        if !exist {
            fmt.Printf("Unknown command: %s\n", commandName)
            continue
        }
        err := command.callback()
        if err != nil {
            fmt.Println(err)
            continue
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
        "map": {
            name: "map",
            description: "Display the names of 20 locations in the pokemons world, each call displays the next 20",
            callback: commandMap,
        },
    }
}

func getCleanInput(input string) []string {
    words := strings.Fields(strings.ToLower(input))
    return words
} 

type config struct {
    pokeapiClient     pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}