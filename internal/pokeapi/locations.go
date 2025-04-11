package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokeAPIClient struct {
	HttpClient http.Client
}
type LocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *PokeAPIClient) GetLocations(url *string) (LocationsResponse, error) {
	path := baseUrl + "/location-area"
	if url != nil {
		path = *url
	}
	res, err := c.HttpClient.Get(path)
	if err != nil {
		return LocationsResponse{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, err
	}
	var locationsResponse LocationsResponse

	if err = json.Unmarshal(data, &locationsResponse); err != nil {
		return LocationsResponse{}, err
	}

	return locationsResponse, err
}
