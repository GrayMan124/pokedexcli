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
	for i := 19; i >= 0; i-- {
		prevUrl := fmt.Sprintf(baseUrl + string(i))
		data, exists := conf.cache.Get(prevUrl)
		if !exists {
			res, err := http.Get(prevUrl)
			if err != nil {
				log.Fatal(err)
			}
			body, err := io.ReadAll(res.Body)
			var location LocationArea
			if err := json.Unmarshal(body, &location); err != nil {
				log.Fatalf("Error unmarshalling Json: %v", err)
			}
			names[i] = location.Name
			conf.cache.Add(prevUrl, body)
		} else {
			var location LocationArea
			if err := json.Unmarshal(data, &location); err != nil {
				log.Fatalf("Error unmarshalling Json: %v", err)
			}
			names[i] = location.Name
		}
		fmt.Println(names[i])
	}
	return nil

}
