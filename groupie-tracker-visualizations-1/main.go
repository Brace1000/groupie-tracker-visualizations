package main

import (
	"log"
	"net/http"
	"os"
	"tracker/handlers"
)

func main() {
	if len(os.Args) > 1 {
		log.Println("The program expects only one argument: \n \n e.g go run main.go ")
		return
	}

	// Map URLs to their respective page handlers
	http.HandleFunc("/", handlers.ArtistHandler)
	http.HandleFunc("/relations", handlers.RelationHandler)
	http.HandleFunc("/locations", handlers.LocationHandler)
	http.HandleFunc("/dates/", handlers.DatesHandler)
	http.HandleFunc("/artistProfile", handlers.ArtistDetails)
	// Correctly handle static files
	http.HandleFunc("/static/", handlers.StaticServer)

	// Start the server
	log.Println("Server running at http://localhost:8091")
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
