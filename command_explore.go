package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// var pokemon struct {
// 	name       string
// 	experience int
// }

func commandExplore(conf *config) error {
	// var pokemonNames []string
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	getUrl := fmt.Sprintf(baseUrl+"%v/", conf.input)

	var location LocationArea
	data, exists := conf.cache.Get(getUrl)
	if !exists {
		res, err := http.Get(getUrl)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		conf.cache.Add(getUrl, body)
		if err := json.Unmarshal(body, &location); err != nil {
			log.Fatalf("Error unmarshalling Json: %v", err)
		}
	} else {
		if err := json.Unmarshal(data, &location); err != nil {
			log.Fatalf("Error unmarshalling Json: %v", err)
		}
	}
	for _, encounter := range location.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil

}
