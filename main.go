package main

import (
	"log"
	"net/http"
)

func main() {
	registerRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
