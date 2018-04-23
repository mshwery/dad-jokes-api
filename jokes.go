package main

import (
	"encoding/json"
	"log"
	"os"
)

type Joke struct {
	Name string `json:"name"`
	Joke string `json:"joke"`
}

var jokes = loadJokesInMemory()

func loadJokesInMemory() []Joke {
	jokeData, err := os.Open("jokes.json")
	if err != nil {
		log.Fatal(err)
	}

	var jokes []Joke
	jsonParser := json.NewDecoder(jokeData)
	jsonParser.Decode(&jokes)
	return jokes
}
