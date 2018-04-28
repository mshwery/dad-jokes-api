package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func setupRouter() http.Handler {
	router := httprouter.New()
	router.GET("/health", getHealth)
	router.GET("/", getRandomJoke)
	router.GET("/jokes", getRandomJoke)
	router.GET("/jokes/:id", getJoke)

	handler := cors.Default().Handler(router)
	return handler
}

func runServer() {
	router := setupRouter()
	fmt.Println("Server listening at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
