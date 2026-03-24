package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Rangga"},
		{ID: 2, Name: "Waridat"},
	}

	w.Header().Set("Content-Type", "application/json")

	paramID := r.URL.Query().Get("id")

	if paramID == "" {
		json.NewEncoder(w).Encode(users)
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "user not found", http.StatusNotFound)
}
