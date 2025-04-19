package commands

import "github.com/2bitburrito/pokedex-cli/internal/pokeapi"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config) error
}

var CliCommands map[string]cliCommand

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays a paginated list of Pokemon locations",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous page of the map list",
			Callback:    CommandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore an area - takes one location argument",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try your luck and see if you can catch a pokemon",
			Callback:    CommandCatch,
		},
		"list-pokemon ": {
			Name:        "list-pokemon",
			Description: "List all caught pokemon",
			Callback:    CommandListPokemon,
		},
	}
}
