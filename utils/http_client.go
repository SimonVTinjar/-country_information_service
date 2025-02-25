package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPClient er en tilpasset HTTP-klient med timeout
var HTTPClient = &http.Client{
	Timeout: 10 * time.Second, // Timeout på 10 sekunder for alle forespørsler
}

// FetchData sender en GET-forespørsel og returnerer respons-body som []byte
func FetchData(url string) ([]byte, error) {
	resp, err := HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Sjekk om responsen er vellykket (statuskode 200 OK)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return ReadResponseBody(resp)
}

// PostData sender en POST-forespørsel med JSON-body og returnerer respons-body
func PostData(url string, jsonBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Sjekk om responsen er vellykket (statuskode 200 OK)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return ReadResponseBody(resp)
}

// ReadResponseBody leser HTTP-responsens body og returnerer []byte
func ReadResponseBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
