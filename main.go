package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lowerWords := strings.ToLower(text)
	words := strings.Fields(lowerWords)
	return words
}

func main() {
	fmt.Println("Hello, World!")
}
