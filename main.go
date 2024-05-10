package main

import (
	"time"

	"github.com/ItzTass/PokeDex/internal/pokeapi"
)

func main() { 
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second, 5 * time.Minute),
	}
	startRepl(cfg)
}	
