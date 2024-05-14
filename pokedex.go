package main

import "github.com/ItzTass/PokeDex/internal/pokeapi"

func (c *config) pokeAdd(key string, p pokeapi.PokemonStatsRes) {
	c.pokedex[key] = p
}