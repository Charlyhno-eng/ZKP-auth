package main

import (
	"go-zkp/api"
	"go-zkp/internal/db"
	"log"
	"net/http"
)

func main() {
	database, err := db.ConnectDB("./db.sqlite")
	if err != nil {
		log.Fatal("DB error: ", err)
	}

	http.HandleFunc("/api/register", api.RegisterHandler(database))
	http.HandleFunc("/api/login", api.LoginHandler(database))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
