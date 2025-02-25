package models

// PopulationValue representerer befolkningsdata for ett år
type PopulationValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// PopulationResponse representerer API-responsen
type PopulationResponse struct {
	Country         string            `json:"country"`
	Code            string            `json:"code"`
	ISO3            string            `json:"iso3"`
	Mean            int               `json:"mean"`
	PopulationCount []PopulationValue `json:"populationCounts"`
}
