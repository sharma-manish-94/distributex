package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sharma-manish-94/distributex/internal/ingest"
)

func main() {

	logger := log.New(os.Stdout, "[INGESTION] ", log.LstdFlags)

	http.HandleFunc("/events", handleEventIngestion(logger))

	http.HandleFunc("/health", handleHealth())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Status: OK\nService: Ingestion\n")
	}
}

func handleEventIngestion(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var event ingest.Event
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&event); err != nil {
			http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
			return
		}

		if event.TenantID == "" {
			http.Error(w, "Missing tenant_id", http.StatusBadRequest)
			return
		}

		logger.Printf("Received Event: [%s] %s from Tenant %s", event.Timestamp.Format(time.RFC3339), event.Name, event.TenantID)
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"status": "accepted"}`))

	}
}
