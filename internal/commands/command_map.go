package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/2bitburrito/pokedex-cli/internal/types"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var locationPage Locations
var initLocationUrl string = "http://pokeapi.co/api/v2/location-area"

func CommandMap(cfg *types.Config) error {
	if locationPage.Next == "" {
		locationPage.Next = initLocationUrl
	}
	res, err := http.Get(locationPage.Next)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(responseData, &locationPage); err != nil {
		return err
	}
	// fmt.Println(locationPage.Results)
	for _, location := range locationPage.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapBack(cfg *types.Config) error {
	if locationPage.Previous == "" {
		fmt.Println("mapb called before map")
	}
	res, err := http.Get(locationPage.Previous)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &locationPage); err != nil {
		return err
	}
	for _, result := range locationPage.Results {
		fmt.Println(result.Name)
	}
	return nil
}
