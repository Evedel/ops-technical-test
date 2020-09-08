package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Evedel/MYOBTestService/healthcheck"
	"github.com/Evedel/MYOBTestService/storage"
)

// HandleRequests routes API endpoints
func HandleRequests() {
	http.HandleFunc("/healthcheck", healthCheck)
	http.HandleFunc("/metadata", appMetadata)
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	hcCode := healthcheck.Run()
	w.WriteHeader(hcCode)
	fmt.Fprintf(w, "")
}

func appMetadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetAppMetadata())
}
