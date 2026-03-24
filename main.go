package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/users", getUsersHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Rangga"},
		{ID: 2, Name: "Waridat"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
