package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/Khimich13/pokedex/internal/repl"
)

func main() {
	url := "https://pokeapi.co/api/v2/location-area/"
	state := repl.Config{Next: &url}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := repl.CleanInput(input)
		if len(cleanedInput) == 0 {
			fmt.Println("Unknown command")
			continue
		} 

		command, ok := repl.Commands[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} 
		if err := command.Callback(&state); err != nil {
			fmt.Println(err)
		}
	}
}