package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	if res.StatusCode >= 400 {
		return RespShallowLocations{}, errors.New(res.Status)	
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	locations := RespShallowLocations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locations, nil
}