package main

import (
	"encoding/json"
	"log"
	"os"
)

// Joke represents a joke object
type Joke struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Joke string `json:"joke"`
}

var jokes []Joke
var jokesByID map[string]Joke

func loadJokesInMemory() {
	jokeData, err := os.Open("jokes.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(jokeData)
	jsonParser.Decode(&jokes)

	// Build a map, for faster lookup by id
	jokesByID = make(map[string]Joke)
	for _, v := range jokes {
		jokesByID[v.ID] = v
	}
}
