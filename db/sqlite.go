package db

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"random-learning/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Get the SQLite database file path, which should be in the SNAP_USER_DATA directory
func getDBPath() string {
	snapUserData := os.Getenv("SNAP_USER_DATA")
	if snapUserData == "" {
		log.Fatal("$SNAP_USER_DATA is not set. Ensure strict confinement is used.")
	}
	return fmt.Sprintf("%s/selected_sites.db", snapUserData)
}

// InitDB initializes the SQLite database connection
func InitDB() {
	var err error

	databasePath := getDBPath()
	DB, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate the Site model
	if err := DB.AutoMigrate(&models.Site{}); err != nil {
		panic("failed to migrate database")
	}
}

// GetRandomSite fetches a random site from the database
func GetRandomSite() (models.Site, error) {
	var sites []models.Site

	// Fetch all sites
	result := DB.Find(&sites)
	if result.Error != nil {
		return models.Site{}, result.Error
	}

	if len(sites) == 0 {
		return models.Site{}, errors.New("no sites found in the database")
	}

	return sites[rand.Intn(len(sites))], nil
}

// SaveSites saves a list of sites to the database
func SaveSites(sites []models.Site) error {
	if len(sites) == 0 {
		return errors.New("no sites to save")
	}

	// Iterate over each site and check if it exists, if not create it
	for _, site := range sites {
		// Check if the site already exists using a unique field (e.g., SiteName or SiteURL)
		result := DB.Where("site_name = ?", site.SiteName).First(&site)

		// If the site doesn't exist, create it
		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			// Site doesn't exist, create a new one
			createResult := DB.Create(&site)
			if createResult.Error != nil {
				return fmt.Errorf("error creating site %s: %v", site.SiteName, createResult.Error)
			}
		} else if result.Error != nil {
			// If there is an error other than "record not found"
			return fmt.Errorf("error checking site existence for %s: %v", site.SiteName, result.Error)
		}
	}

	return nil
}

// GetAllSites retrieves all sites from the database
func GetAllSites() ([]models.Site, error) {
	var sites []models.Site

	// Fetch all sites
	result := DB.Find(&sites)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(sites) == 0 {
		return nil, errors.New("no sites found in the database")
	}

	return sites, nil
}
