package main

import (
	"fmt"
	"golang/tutorial/config"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.Initialize()
	router := httprouter.New()
	config.ConfigureRoutes(router)
	fmt.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
