package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var db *sql.DB

type User struct {
	ID    int64
	Name  string
	Login string
	Org   string
}

func getUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /hello request\n")
	var user User
	queryParams := r.URL.Query()
	login := queryParams.Get("login")
	log.Printf("Login: %s \n", string(login))
	row := db.QueryRow("SELECT * FROM users WHERE login = ?", string(login))
	if err := row.Scan(&user.ID, &user.Name, &user.Login, &user.Org); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	ans, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := io.Writer.Write(w, ans); err != nil {
		log.Printf("error writing response: %s\n", err.Error())
	}
}
