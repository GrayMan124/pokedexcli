package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func commandCatch(conf *config) error {
	var pokemon Pokemon
	baseUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", conf.input)
	fmt.Printf("Throwing a Pokeball at %s...\n", conf.input)
	data, exists := conf.cache.Get(baseUrl)
	if !exists {
		res, err := http.Get(baseUrl)
		if err != nil {
			log.Fatal(err)
			return err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}
		conf.cache.Add(baseUrl, body)
		if err := json.Unmarshal(body, &pokemon); err != nil {
			log.Fatal(err)
			return err
		}

	} else {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			log.Fatal(err)
			return err
		}
	}
	baseExp := pokemon.BaseExperience
	var expThresh float32
	expThresh = float32((baseExp - 35)) / float32((608 - 35))
	chance := rand.Float32()
	if chance >= expThresh {
		fmt.Printf("%s was caught!\n", conf.input)
		conf.pokedex[conf.input] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", conf.input)
	}
	return nil
}
