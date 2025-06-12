package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Robert076/endpoint-multirepo/internal/data"
	"github.com/joho/godotenv"
)

func main() {
	var serviceName string = "ENDPOINT"

	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s: Cannot load env file", serviceName)
		return
	}

	http.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed to this endpoint.", http.StatusMethodNotAllowed)
			log.Printf("%s: Only POST requests allowed to this endpoint", serviceName)
			return
		}

		validatorHost := os.Getenv("VALIDATOR_HOST")
		validatorPort := os.Getenv("VALIDATOR_PORT")

		if validatorHost == "" || validatorPort == "" {
			http.Error(w, "Host or port are empty", http.StatusInternalServerError)
			log.Fatalf("%s: Unitialized env vars", serviceName)
		}

		req := new(data.RequestBody)
		req.Name = "test"

		b, err := json.Marshal(req)

		if err != nil {
			http.Error(w, "Cannot serialize JSON from native Go struct type", http.StatusInternalServerError)
			log.Printf("%s: Cannot serialize json from native Go struct type %v", serviceName, err)
		}

		body := bytes.NewBuffer(b)

		response, err := http.Post("http://"+validatorHost+":"+validatorPort+"/validator", "application/json; charset=utf-8", body)

		if err != nil {
			http.Error(w, "Cannot validate request.", http.StatusInternalServerError)
			log.Printf("%s: Cannot validate request. Validator endpoint unreachable. Error: %v", serviceName, err)
			return
		}

		defer response.Body.Close()

		log.Printf("Status received from server is: %s", response.Status)
		log.Printf("StatusCode received from server is: %d", response.StatusCode)
	})

	endpointPort := os.Getenv("ENDPOINT_PORT")

	if endpointPort == "" {
		log.Fatalf("%s: Uninitialized env var for endpoint port", serviceName)
		return
	}

	if err := http.ListenAndServe(":"+endpointPort, nil); err != nil {
		log.Fatalf("%s: Error starting http server, %v", serviceName, err)
	}
}
