package main

import (
	"fmt"
	"github.com/robfig/cron"
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
	c := cron.New()

	// Schedule the distribute_prizes stored procedure to run every day at midnight
	err = c.AddFunc("@midnight", func() {
		_, err := db.Exec("CALL distribute_prizes()")
		if err != nil {
			log.Printf("Error calling distribute_prizes: %v", err)
		} else {
			log.Println("Successfully called distribute_prizes")
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	// Start the cron scheduler
	c.Start()

	fmt.Println("Cron job scheduler started.")
	http.HandleFunc("/ranking", handlers.GetRanking)
	http.HandleFunc("/tournament-ranking", handlers.GetTournamentRanking)

	// Start the HTTP server
	fmt.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
