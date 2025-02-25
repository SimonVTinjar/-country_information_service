package handlers

import (
	"Country_Information_Service/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

// CountryInfoHandler håndterer forespørselen og returnerer landinformasjon
func CountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/info/")

	countryInfo, err := services.GetCountryInfo(countryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countryInfo)
}
