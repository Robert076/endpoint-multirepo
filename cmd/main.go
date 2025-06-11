package main

import (
	"log"
	"net/http"
)

func main() {
	var serviceName string = "ENDPOINT"

	http.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed to this endpoint.", http.StatusMethodNotAllowed)
			log.Print("%s: Only POST requests allowed to this endpoint", serviceName)
		}
	})
}
