package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func getHealth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonResponse(w, r, http.StatusOK, nil)
}

func getJoke(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// grab a random joke
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]

	jsonResponse(w, r, http.StatusOK, joke)
}
