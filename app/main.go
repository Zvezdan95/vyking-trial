package main

import (
	"fmt"
	"log"
	"vyking-trial/app/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")

}
