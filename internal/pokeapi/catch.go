package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon *string) (CatchResponse, error) {
	var catchRes CatchResponse
	var res *http.Response
	var err error
	if pokemon == nil {
		return catchRes, errors.New("no pokemon found")
	}
	path := fmt.Sprintf("%v%v%v", baseUrl, "/pokemon/", *pokemon)

	data, ok := c.Cache.Get(path)
	if !ok {
		res, err = c.HttpClient.Get(path)
		if err != nil {
			return catchRes, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return catchRes, err
		}
		c.Cache.Add(path, data)
	}
	if err = json.Unmarshal(data, &catchRes); err != nil {
		return catchRes, err
	}
	return catchRes, nil
}
