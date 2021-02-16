package client

import (
	"encoding/json"
	"net/http"
)

func GetAllClients(w http.ResponseWriter, r *http.Request) {
	clt, err := AllClients()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Reference-No", "`1e9ac446-8a62-4ae3-852d-c352ceda99b`")
	json.NewEncoder(w).Encode(clt)
}
