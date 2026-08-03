package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	instanceID := os.Getenv("INSTANCE_ID")
	if instanceID == "" {
		instanceID = "unknown"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from instance %s\n", instanceID)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok\n")
	})

	log.Printf("server %s listening on :8080", instanceID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
