package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(url *string) (LocationsResponse, error) {
	path := baseUrl + "/location-area/"
	if url != nil {
		path = *url
	}
	var res *http.Response
	var err error

	data, ok := c.Cache.Get(path)
	if !ok {
		res, err = c.HttpClient.Get(path)
		if err != nil {
			return LocationsResponse{}, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationsResponse{}, err
		}
		c.Cache.Add(path, data)
	}

	var locationsResponse LocationsResponse

	if err = json.Unmarshal(data, &locationsResponse); err != nil {
		return LocationsResponse{}, err
	}

	return locationsResponse, err
}
