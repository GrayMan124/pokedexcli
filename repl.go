package main

import (
	"bufio"
	"fmt"
	"github.com/GrayMan124/pokedexcli/internal/pokecache"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandInput := ""
		if len(words) > 1 {
			commandInput = words[1]
		}
		start_conifg := config{
			Previous: "",
			Next:     "https://pokeapi.co/api/v2/location-area/1/",
			cache:    *pokecache.NewCache(5),
			input:    commandInput,
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&start_conifg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	Previous string
	Next     string
	input    string
	cache    pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Explore the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore the map backwards",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemons in the map",
			callback:    commandExplore,
		},
	}
}
