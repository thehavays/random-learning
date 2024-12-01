package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"random-learning/db"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd defines the 'random' command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random site from the saved StackExchange websites",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch all saved sites from the database
		sites, err := db.GetAllSites()
		if err != nil {
			log.Fatalf("Failed to retrieve sites from database: %v", err)
		}

		// Ensure there are saved sites
		if len(sites) == 0 {
			fmt.Println("No sites found in the database. Please add sites using the 'website' command.")
			return
		}

		// Generate a random index
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(sites))

		// Print the randomly selected site
		randomSite := sites[randomIndex]
		fmt.Printf("Random site: %s (%s)\n", randomSite.SiteName, randomSite.SiteURL)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
