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

func getRandomJoke(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// grab a random joke
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]

	jsonResponse(w, r, http.StatusOK, joke)
}

func getJoke(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	joke, prs := jokesByID[ps.ByName("id")]

	if !prs {
		jsonResponse(w, r, http.StatusNotFound, nil)
	} else {
		jsonResponse(w, r, http.StatusOK, joke)
	}
}
