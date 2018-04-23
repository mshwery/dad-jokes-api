package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func getJoke(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// grab a random joke
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(joke)
}
