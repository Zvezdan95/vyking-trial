package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"vyking-trial/app/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	// Seed players
	for i := 1; i <= 100; i++ {
		playerName := fmt.Sprintf("Player %d", i)
		playerEmail := fmt.Sprintf("player%d@example.com", i)
		accountBalance := rand.Intn(100000)

		_, err := db.Exec("INSERT INTO players (player_name, player_email, account_balance) VALUES (?, ?, ?)", playerName, playerEmail, accountBalance)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Seed tournaments
	for i := 1; i <= 10; i++ {
		tournamentName := fmt.Sprintf("Tournament %d", i)
		prizePool := rand.Intn(1000000)
		startDate := time.Now().AddDate(0, -rand.Intn(12), 0)
		endDate := startDate.AddDate(0, 0, rand.Intn(30))

		_, err := db.Exec("INSERT INTO tournaments (tournament_name, prize_pool, start_date, end_date) VALUES (?, ?, ?, ?)", tournamentName, prizePool, startDate, endDate)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Seed bets
	for i := 1; i <= 10000; i++ {
		playerID := rand.Intn(100) + 1
		tournamentID := rand.Intn(10) + 1
		amount := rand.Intn(10000)
		betTime := time.Now().AddDate(0, 0, -rand.Intn(365))

		_, err := db.Exec("INSERT INTO bets (player_id, tournament_id, amount, bet_time) VALUES (?, ?, ?, ?)", playerID, tournamentID, amount, betTime)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database seeding completed successfully!")
}
