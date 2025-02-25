package utils

import (
	"io/ioutil"
	"net/http"
)

// FetchData henter data fra en gitt URL og returnerer responsen som en byte-array
func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
