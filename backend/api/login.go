package api

import (
	"crypto/ed25519"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"

	"go-zkp/internal/auth"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		privStr, err := auth.FindPrivateKeyOnUSB()
		if err != nil {
			http.Error(w, "No valid USB found", http.StatusUnauthorized)
			return
		}

		privBytes, err := base64.StdEncoding.DecodeString(privStr)
		if err != nil {
			http.Error(w, "Invalid private key encoding", http.StatusBadRequest)
			return
		}
		priv := ed25519.PrivateKey(privBytes)
		pub := priv.Public().(ed25519.PublicKey)

		var username string
		err = db.QueryRow("SELECT username FROM users WHERE public_key = ?", pub).Scan(&username)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"message":"Bienvenue %s"}`, username)))
	}
}
