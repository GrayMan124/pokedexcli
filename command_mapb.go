package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMapb(conf *config) error {
	var names [20]string
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	for i := 20; i > 0; i-- {
		nextUrl := fmt.Sprintf(baseUrl+"%v/", i+1)
		res, err := http.Get(nextUrl)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		var location LocationArea
		if err := json.Unmarshal(body, &location); err != nil {
			log.Fatalf("Error unmarshalling Json: %v", err)
		}
		names[20-i] = location.Name
		fmt.Println(names[20-i])
	}

	return nil

}
