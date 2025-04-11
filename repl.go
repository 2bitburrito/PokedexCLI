package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/2bitburrito/pokedex-cli/internal/commands"
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

		cleanedInput := cleanInput(input)
		if len(cleanedInput) < 1 || cleanedInput[0] == "" {
			continue
		}

		command, exists := commands.GetCommands()[cleanedInput[0]]
		if !exists {
			fmt.Println("Unknown command: try 'help'")
		} else {
			if err := command.Callback(cfg); err != nil {
				fmt.Println(err)
			}
		}

	}
}
