package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/2bitburrito/pokedex-cli/commands"
	"github.com/2bitburrito/pokedex-cli/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowerWords := strings.ToLower(text)
	words := strings.Fields(lowerWords)
	if len(words) < 1 {
		words = append(words, "")
	}
	return words
}

func startRepl(cfg *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInputs := cleanInput(input)
		if len(cleanedInputs) < 1 || cleanedInputs[0] == "" {
			continue
		}

		command, exists := commands.GetCommands()[cleanedInputs[0]]
		switch command.Name {
		case "explore":
			if len(cleanedInputs) < 2 {
				fmt.Println("Requre a valid location argument")
			} else {
				cfg.ExploredLocation = &cleanedInputs[1]
			}
		case "catch":
			if len(cleanedInputs) < 2 {
				fmt.Println("Requre a valid location argument")
			} else {
				cfg.CatchingPokemon = &cleanedInputs[1]
			}
		case "inspect":
			if len(cleanedInputs) < 2 {
				fmt.Println("Requre a valid pokemon argument")
			} else {
				cfg.PokemonToInspect = &cleanedInputs[1]
			}
		}

		if !exists {
			fmt.Println("Unknown command: try 'help'")
			continue
		}
		if err := command.Callback(cfg); err != nil {
			fmt.Println(err)
		}

	}
}
