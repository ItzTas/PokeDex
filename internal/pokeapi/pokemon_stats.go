package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) StatsPokemon(pokemon string) (PokemonStatsRes, error) {
	url := baseURL + "/pokemon/" + pokemon + "/"
	if val, ok := c.cache.Get(url); ok {
		stats := PokemonStatsRes{}
		err := json.Unmarshal(val, &stats)
		if err != nil {
			return PokemonStatsRes{}, err
		}
		return stats, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStatsRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonStatsRes{}, err
	}
	
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonStatsRes{}, err
	}
	
	stats := PokemonStatsRes{}
	err = json.Unmarshal(dat, &stats)
	if err != nil {
		return PokemonStatsRes{}, err
	}

	c.cache.Add(url, dat)
	return stats, nil
}