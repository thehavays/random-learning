package main

import (
	"log"
	"random-learning/cmd"
	"random-learning/db"
)

func main() {
	// Initialize the database connection
	db.InitDB()

	// Execute the root command
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
