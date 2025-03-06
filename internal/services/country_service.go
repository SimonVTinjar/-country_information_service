package services

import (
	"Country_Information_Service/internal/models"
	"Country_Information_Service/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetCountryInfo henter landinformasjon basert på ISO 2-kode
func GetCountryInfo(countryCode string) (*models.CountryInfoResponse, error) {
	apiURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)

	body, err := utils.FetchData(apiURL)
	if err != nil {
		return nil, fmt.Errorf("Kunne ikke hente landinformasjon for kode %s: %v", countryCode, err)
	}

	// Parse JSON-respons
	var apiResponse []map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil || len(apiResponse) == 0 {
		return nil, fmt.Errorf("Ugyldig respons fra RestCountries API for kode: %s", countryCode)
	}

	countryData := apiResponse[0]

	// Hent nødvendig informasjon
	name := countryData["name"].(map[string]interface{})["common"].(string)
	continent := countryData["continents"].([]interface{})[0].(string)
	population := int(countryData["population"].(float64))
	flag := countryData["flags"].(map[string]interface{})["png"].(string)
	capital := countryData["capital"].([]interface{})[0].(string)

	// Hent språk
	languages := make(map[string]string)
	if langData, exists := countryData["languages"].(map[string]interface{}); exists {
		for key, value := range langData {
			languages[key] = value.(string)
		}
	}

	// Hent naboland
	borders := []string{}
	if borderData, exists := countryData["borders"].([]interface{}); exists {
		for _, border := range borderData {
			borders = append(borders, border.(string))
		}
	}
	// Hent liste over byer fra CountriesNow API
	cities, err := GetCitiesByCountry(name)
	if err != nil {
		cities = []string{"Ukjente byer"} // Standardverdi hvis API-kall feiler
	}

	// Opprett responsstruktur
	return &models.CountryInfoResponse{
		Name:       name,
		Continent:  continent,
		Population: population,
		Languages:  languages,
		Borders:    borders,
		Flag:       flag,
		Capital:    capital,
		Cities:     cities,
	}, nil
}

func GetCitiesByCountry(countryName string) ([]string, error) {
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	requestBody, err := json.Marshal(map[string]string{"country": countryName})
	if err != nil {
		return nil, fmt.Errorf("Kunne ikke lage forespørsel: %v", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Kunne ikke sende forespørsel: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Error bool     `json:"error"`
		Msg   string   `json:"msg"`
		Data  []string `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Kunne ikke tolke respons: %v", err)
	}

	if result.Error {
		return nil, fmt.Errorf("Feil fra API: %s", result.Msg)
	}

	return result.Data, nil
}

// GetCountryNameFromCode henter landets fulle navn fra ISO 2-kode via RestCountries API
func GetCountryNameFromCode(countryCode string) (string, error) {
	apiURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", strings.ToLower(countryCode))

	body, err := utils.FetchData(apiURL)
	if err != nil {
		return "", fmt.Errorf("Kunne ikke hente landnavn for kode %s: %v", countryCode, err)
	}

	// Parse JSON-respons
	var apiResponse []map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil || len(apiResponse) == 0 {
		return "", fmt.Errorf("Ugyldig respons fra RestCountries API for kode: %s", countryCode)
	}

	// Hent landnavn
	name, ok := apiResponse[0]["name"].(map[string]interface{})["common"].(string)
	if !ok {
		return "", fmt.Errorf("Fant ikke landnavn for kode: %s", countryCode)
	}
	return name, nil
}

// GetCountryDetails henter offisielt landnavn, ISO3-kode og vanlig navn fra ISO 2-kode
func GetCountryDetails(countryCode string) (string, string, string, error) {
	// 🎯 Hvis forespørselen er for Russland, returner hardkodet info
	if countryCode == "ru" {
		return "Russian Federation", "RUS", "Russian Federation", nil
	}

	// Kall API-et for alle andre land
	apiURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", strings.ToLower(countryCode))

	body, err := utils.FetchData(apiURL)
	if err != nil {
		return "", "", "", fmt.Errorf("Kunne ikke hente landnavn for kode %s: %v", countryCode, err)
	}

	// Parse JSON-respons
	var apiResponse []map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil || len(apiResponse) == 0 {
		return "", "", "", fmt.Errorf("Ugyldig respons fra RestCountries API for kode: %s", countryCode)
	}

	// Hent nødvendig informasjon fra API-et
	commonName, _ := apiResponse[0]["name"].(map[string]interface{})["common"].(string)
	officialName, _ := apiResponse[0]["name"].(map[string]interface{})["official"].(string)
	iso3, _ := apiResponse[0]["cca3"].(string)

	// Returner commonName for alle andre land
	return commonName, iso3, officialName, nil
}
