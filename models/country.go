package models

// CountryInfoResponse representerer det filtrerte JSON-svaret
type CountryInfoResponse struct {
	Name       string            `json:"name"`
	Continent  string            `json:"continent"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
}
