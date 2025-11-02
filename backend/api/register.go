package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go-zkp/internal/auth"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != "POST" {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var req auth.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Username == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		resp, err := auth.CreateUser(db, req.Username)
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
