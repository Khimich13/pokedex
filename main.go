package main

import (
	"fmt"
	"os"
	"bufio"
	"pokedex/internal/api"
)

func main() {
	url := "https://pokeapi.co/api/v2/location-area/"
	scanner := bufio.NewScanner(os.Stdin)
	state := config{}
	state.Next = &url
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			fmt.Println("Unknown command")
		} else if command, ok := commands[cleanedInput[0]]; !ok{
			fmt.Println("Unknown command")
		} else if err := command.callback(&state); err != nil {
			fmt.Println(err)
		}
	}
}