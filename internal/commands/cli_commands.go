package commands

import "github.com/2bitburrito/pokedex-cli/internal/types"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*types.Config) error
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
	}
}
