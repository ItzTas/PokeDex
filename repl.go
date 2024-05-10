package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ItzTass/PokeDex/internal/pokeapi"
)

var commandsMaps []string = []string{"map", "mapb"}

func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Printf("Pokedex > ")
        scanner.Scan()
        
        words := getCleanInput(scanner.Text())
        if len(words) == 0 {
            continue
        }
        commandName := words[0]
        if (commandName == "map" || commandName == "mapb") && len(words) > 1 {
            for _, c := range words {
                if !slices.Contains(commandsMaps, c){
                    fmt.Printf("\nCannot use another command with the map: %s\n", c)
                }
                GetCommands()[c].callback(cfg)
            }  
            continue
        }
        command, exist := GetCommands()[commandName] 
        if !exist {
            fmt.Printf("Unknown command: %s\n", commandName)
            continue
        }
        err := command.callback(cfg)
        if err != nil {
            fmt.Println(err)
            continue
        }

    }
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
            callback: commandMapf,
        },
        "mapb": {
            name: "mapb",
            description: "Display the names of the previous 20 locations in the pokemons world",
            callback: commandMapb,
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

    type cliCommand struct {
        name 		string
        description string
        callback 	func(*config) error	
    }