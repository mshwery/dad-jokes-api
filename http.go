package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setupRouter() http.Handler {
	router := httprouter.New()
	router.GET("/", getJoke)

	return router
}

func runServer() {
	router := setupRouter()
	fmt.Println("Server listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
