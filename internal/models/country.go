package models

// CountryInfoResponse representerer det filtrerte JSON-svaret for landinfo
type CountryInfoResponse struct {
	Name       string            `json:"name"`
	Continent  string            `json:"continent"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages,omitempty"`
	Borders    []string          `json:"borders,omitempty"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities,omitempty"`
}
