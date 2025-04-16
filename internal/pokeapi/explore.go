package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location *string) (ExploreResponse, error) {
	var exploreRes ExploreResponse
	if location == nil {
		return exploreRes, errors.New("no location found")
	}
	path := fmt.Sprintf("%v%v%v", baseUrl, "/location-area/", *location)
	var res *http.Response
	var err error
	data, ok := c.Cache.Get(path)
	if !ok {
		res, err = c.HttpClient.Get(path)
		if err != nil {
			return exploreRes, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return exploreRes, err
		}
		c.Cache.Add(path, data)
	}
	if err = json.Unmarshal(data, &exploreRes); err != nil {
		return exploreRes, err
	}
	return exploreRes, nil
}
