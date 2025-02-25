package services

import (
	"Country_Information_Service/internal/models"
	"time"
)

// startTime lagrer tidspunktet serveren startet
var startTime = time.Now()

// GetStatus returnerer API-status og oppetid
func GetStatus() models.StatusResponse {
	return models.StatusResponse{
		RestCountriesAPI: "OK",
		CountriesNowAPI:  "OK",
		Version:          "v1",
		Uptime:           time.Since(startTime).Seconds(),
	}
}
