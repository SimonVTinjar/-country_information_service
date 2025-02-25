package models

// StatusResponse representerer status for eksterne API-er og tjenestens oppetid
type StatusResponse struct {
	RestCountriesAPI string  `json:"restcountriesapi"`
	CountriesNowAPI  string  `json:"countriesnowapi"`
	Version          string  `json:"version"`
	Uptime           float64 `json:"uptime"`
}
