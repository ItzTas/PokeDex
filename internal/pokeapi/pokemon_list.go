package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location string) (RespPokemonEncounters, error) {
	url := baseURL + "/location-area/" + location
	
	if val, ok := c.cache.Get(url); ok {
		pokemons := RespPokemonEncounters{}
		err := json.Unmarshal(val, &pokemons) 
		if err != nil {
			return RespPokemonEncounters{}, err
		}
		return pokemons, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	pokemons := RespPokemonEncounters{}
	err = json.Unmarshal(dat, &pokemons)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	c.cache.Add(url, dat)
	return pokemons, err
}