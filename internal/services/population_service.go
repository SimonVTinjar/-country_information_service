package services

import (
	"Country_Information_Service/internal/models"
	"Country_Information_Service/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GetPopulationData henter befolkningsdata for et land basert på landkode
func GetPopulationData(countryCode string, limit string) (*models.PopulationResponse, error) {
	// Hent landinfo (bruk commonName)
	commonName, iso3, _, err := GetCountryDetails(strings.ToLower(countryCode))
	if err != nil {
		return nil, fmt.Errorf("Ugyldig landkode: %s", countryCode)
	}

	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"
	requestBody, _ := json.Marshal(map[string]string{"country": commonName})

	// Logg forespørselen for debugging
	fmt.Printf("⏳ Henter befolkningsdata for land: %s (forespørsel: %s)\n", commonName, string(requestBody))

	// Send POST-forespørsel til API
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Kunne ikke hente data fra API: %v", err)
	}
	defer resp.Body.Close()

	// Sjekk statuskode
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API-et svarte med statuskode %d", resp.StatusCode)
	}

	// Les responsen
	body, err := utils.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("Feil ved lesing av respons: %v", err)
	}

	// Logg API-responsen for debugging
	fmt.Printf("✅ API-respons mottatt: %s\n", string(body))

	// Parse JSON
	var apiResponse map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("Feil ved parsing av JSON: %v", err)
	}

	// Hent befolkningsdata fra "populationCounts"
	data, exists := apiResponse["data"].(map[string]interface{})
	if !exists {
		return nil, fmt.Errorf("Ingen 'data'-nøkkel funnet i API-responsen")
	}

	populationCounts, ok := data["populationCounts"].([]interface{})
	if !ok || len(populationCounts) == 0 {
		return nil, fmt.Errorf("Ingen befolkningsdata funnet for %s", commonName)
	}

	// 🎯 **Håndter `limit`-parameter**
	startYear, endYear := 0, 0
	if limit != "" {
		yearRange := strings.Split(limit, "=")
		if len(yearRange) > 1 {
			years := strings.Split(yearRange[1], "-")
			if len(years) == 2 {
				startYear, _ = strconv.Atoi(years[0])
				endYear, _ = strconv.Atoi(years[1])
			}
		}
	}

	// Hent verdier for befolkning per år og filtrer hvis nødvendig
	var values []models.PopulationValue
	total := 0
	count := 0

	for _, v := range populationCounts {
		entry, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		year, yOK := entry["year"].(float64)
		value, vOK := entry["value"].(float64)
		if yOK && vOK {
			if (startYear == 0 && endYear == 0) || (year >= float64(startYear) && year <= float64(endYear)) {
				values = append(values, models.PopulationValue{Year: int(year), Value: int(value)})
				total += int(value) // 🎯 Summer verdiene for å beregne gjennomsnitt
				count++             // 🎯 Tell hvor mange verdier vi har
			}
		}
	}

	//Beregn `mean`
	mean := 0
	if count > 0 {
		mean = total / count
	}

	// Logg resultatet
	fmt.Printf("✅ Befolkningsdata hentet for %s (%d verdier, mean: %d)\n", commonName, count, mean)

	// Returner riktig JSON-struktur (med "mean")
	return &models.PopulationResponse{
		Country:         commonName,
		Code:            strings.ToUpper(countryCode),
		ISO3:            iso3,
		PopulationCount: values,
		Mean:            mean,
	}, nil
}
