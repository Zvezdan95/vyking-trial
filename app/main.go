package main

import (
	"fmt"
	"log"
	"net/http"
	"vyking-trial/app/database"
	"vyking-trial/app/handlers"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")
	http.HandleFunc("/ranking", handlers.GetRanking)

	// Start the HTTP server
	fmt.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
