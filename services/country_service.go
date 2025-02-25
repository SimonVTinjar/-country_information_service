package services

import (
	"Country_Information_Service/models"
	"Country_Information_Service/utils"
	"encoding/json"
	"fmt"
)

// GetCountryInfo henter informasjon om et land fra API-et og returnerer et strukturet objekt
func GetCountryInfo(countryCode string) (*models.CountryInfoResponse, error) {
	apiURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)

	body, err := utils.FetchData(apiURL)
	if err != nil {
		return nil, err
	}

	var apiResponse []map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}

	if len(apiResponse) == 0 {
		return nil, fmt.Errorf("fant ikke landet")
	}

	countryData := apiResponse[0]

	// Trekker ut nødvendige verdier
	name := countryData["name"].(map[string]interface{})["common"].(string)
	officialName := countryData["name"].(map[string]interface{})["official"].(string)
	continent := countryData["continents"].([]interface{})[0].(string)
	population := int(countryData["population"].(float64))

	languages := make(map[string]string)
	for key, value := range countryData["languages"].(map[string]interface{}) {
		languages[key] = value.(string)
	}

	borders := []string{}
	if countryData["borders"] != nil {
		for _, border := range countryData["borders"].([]interface{}) {
			borders = append(borders, border.(string))
		}
	}

	flag := countryData["flags"].(map[string]interface{})["png"].(string)
	capital := countryData["capital"].([]interface{})[0].(string)

	return &models.CountryInfoResponse{
		Name:       fmt.Sprintf("%s (%s)", name, officialName),
		Continent:  continent,
		Population: population,
		Languages:  languages,
		Borders:    borders,
		Flag:       flag,
		Capital:    capital,
	}, nil
}
