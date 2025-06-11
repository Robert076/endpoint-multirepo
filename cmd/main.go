package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var serviceName string = "ENDPOINT"

	http.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed to this endpoint.", http.StatusMethodNotAllowed)
			log.Printf("%s: Only POST requests allowed to this endpoint", serviceName)
		}

		if err := godotenv.Load(); err != nil {
			log.Fatalf("%s: Cannot load env file", serviceName)
			return
		}
		validatorHost := os.Getenv("VALIDATOR_HOST")
		validatorPort := os.Getenv("VALIDATOR_PORT")

		http.Get(validatorHost + ":" + validatorPort)
	})
}
