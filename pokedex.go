package main

import "github.com/ItzTass/PokeDex/internal/pokeapi"

func (c *config) pokeAdd(key string, p pokeapi.PokemonStatsRes) {
	c.pokedex[key] = p
}

func (c *config) pokeGet(key string) (pokeapi.PokemonStatsRes, bool) {
	v, ok := c.pokedex[key]
	return v, ok
}