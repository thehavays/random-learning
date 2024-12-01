package models

type Site struct {
	SiteName string `json:"name"`
	SiteURL  string `json:"site_url"`
}

type DbSite struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	URL  string `gorm:"not null"`
}

type StackExchangeResponse struct {
	Items []Site `json:"items"`
}
