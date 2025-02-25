package handlers

import (
	"Country_Information_Service/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

// PopulationHandler håndterer forespørselen for befolkningsdata
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	// Hent landkode og eventuelle query-parametere
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Ugyldig URL-format. Bruk: /population/{countryCode}", http.StatusBadRequest)
		return
	}
	countryCode := pathParts[4]   // Landkode (ISO2)
	queryParams := r.URL.RawQuery // Henter "limit=2010-2015" hvis den finnes

	// Kall tjenesten
	populationData, err := services.GetPopulationData(countryCode, queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(populationData)
}
