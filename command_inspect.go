package main

import (
	"fmt"
)

func commandInspect(conf *config) error {
	pokemon, exists := conf.pokedex[conf.input]
	if !exists {
		fmt.Println("You have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("	-%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("	-%v\n", pokeType.Type.Name)
		}
	}
	return nil
}
