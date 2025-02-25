package main

import (
	"fmt"
	"log"
	"net/http"

	"Country_Information_Service/internal/handlers"
)

func main() {
	http.HandleFunc("/countryinfo/v1/info/", handlers.CountryInfoHandler)
	http.HandleFunc("/countryinfo/v1/population/", handlers.PopulationHandler)
	http.HandleFunc("/countryinfo/v1/status/", handlers.StatusHandler)

	port := ":8080"
	fmt.Println("Server kjører på http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
