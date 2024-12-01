package cmd

import (
	"fmt"
	"log"
	"random-learning/db"
	"random-learning/fetch"
	"random-learning/models"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// websiteCmd defines the 'website' command
var websiteCmd = &cobra.Command{
	Use:   "website",
	Short: "Retrieve and print saved StackExchange websites",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch and print saved sites from SQLite
		sites, err := db.GetAllSites()
		if err != nil {
			log.Fatalf("Failed to get saved sites: %v", err)
		}
		for _, site := range sites {
			fmt.Println(site.SiteName, site.SiteURL)
		}
	},
}

// websiteSetCmd handles saving StackExchange sites
var websiteSetCmd = &cobra.Command{
	Use:   "websites",
	Short: "Fetch and save StackExchange sites",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch sites from the StackExchange API
		sites := fetch.FetchSites()

		options := make([]string, len(sites))
		for i, site := range sites {
			options[i] = fmt.Sprintf("%s (%s)", site.SiteName, site.SiteURL)
		}

		var selectedOptions []string
		prompt := &survey.MultiSelect{
			Message: "Select StackExchange sites:",
			Options: options,
		}
		if err := survey.AskOne(prompt, &selectedOptions); err != nil {
			log.Fatalf("Failed to get user selection: %v", err)
		}

		selectedSites := []models.Site{}
		for _, option := range selectedOptions {
			for _, site := range sites {
				if option == fmt.Sprintf("%s (%s)", site.SiteName, site.SiteURL) {
					selectedSites = append(selectedSites, site)
					break
				}
			}
		}

		// Save sites to the database
		if err := db.SaveSites(selectedSites); err != nil {
			log.Fatalf("Error saving sites to database: %v", err)
		}
		fmt.Println("Sites successfully saved.")
	},
}

func init() {
	rootCmd.AddCommand(websiteCmd)
	rootCmd.AddCommand(websiteSetCmd)
}
