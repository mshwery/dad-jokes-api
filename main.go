package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Joke struct {
	Name string `json:name`
	Joke string `json:joke`
}

var jokes []Joke

func loadJokesInMemory() []Joke {
	jokeData, err := os.Open("jokes.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(jokeData)
	jsonParser.Decode(&jokes)
	return jokes
}

func main() {
	jokes := loadJokesInMemory()
	fmt.Printf("%+v", jokes)
}
