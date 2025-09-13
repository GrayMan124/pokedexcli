package main

import "fmt"

func commandPokedex(conf *config) error {
	for _, pokemon := range conf.pokedex {
		fmt.Printf("%s\n", pokemon.Name)
	}
	return nil
}
