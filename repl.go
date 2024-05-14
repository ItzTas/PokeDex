package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ItzTass/PokeDex/internal/pokeapi"
)

var paramsHelp []string = []string{"--help", "-h"}

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
        var param string
        if len(words) > 1 {
            param = words[1]
        }
        command, exist := GetCommands()[commandName] 
        if !exist {
            fmt.Printf("Unknown command: %s\n", commandName)
            continue
        }
        err := command.callback(cfg, param)
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
            description: "Displays the names of 20 locations in the pokemons world, each call displays the next 20",
            callback: commandMapf,
        },
        "mapb": {
            name: "mapb",
            description: "Displays the names of the previous 20 locations in the pokemons world",
            callback: commandMapb,
        },
        "explore": {
            name: "explore",
            description: "Displays the pokemons of a given area in the world",
            callback: commandExplore,
        },
        "catch": {
            name: "catch",
            description: "Tries to catch a pokemon",
            callback: commandCatch,
        },
        "inspect": {
            name: "inspect",
            description: "Displays the pokemon status",
            callback: commandInspect,
        },
        "pokedex": {
            name: "pokedex",
            description: "Displays all the pokemons in the pokedex",
            callback: commandPokedex,
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
        pokedex map[string]pokeapi.PokemonStatsRes
    }

    type cliCommand struct {
        name 		string
        description string
        callback 	func(*config, string) error	
    }