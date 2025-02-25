package handlers

import (
	"Country_Information_Service/internal/services"
	"encoding/json"
	"net/http"
)

// StatusHandler returnerer status på API-ene og oppetid
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	statusInfo := services.GetStatus()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statusInfo)
}
