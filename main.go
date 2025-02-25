package main

import (
	"fmt"
	"log"
	"net/http"

	"Country_Information_Service/internal/handlers"
)

func main() {
	http.HandleFunc("/countryinfo/v1/info/", handlers.CountryInfoHandler)
	port := ":8080"
	fmt.Println("Server kjører på http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
