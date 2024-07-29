package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	server := http.Server{
		Addr: ":3001",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		response := map[string]string{
			"version":   "V1",
			"path":      r.URL.Path,
			"client_ip": r.RemoteAddr,
			"hostname":  hostname,
		}

		for name, headers := range r.Header {
			for _, h := range headers {
				log.Printf("%v: %v", name, h)
			}
		}

		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(jsonResponse)
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Println("Starting server on " + server.Addr)
	log.Println("http://localhost" + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
