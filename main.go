package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("failed to encode JSON:", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{
		Error: message,
	})
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	users := []User{
		{ID: 1, Name: "Rangga"},
		{ID: 2, Name: "Waridat"},
	}

	paramID := r.URL.Query().Get("id")

	if paramID == "" {
		writeJSON(w, http.StatusOK, users)
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}

	if id <= 0 {
		writeError(w, http.StatusBadRequest, "id must be positive")
		return
	}

	for _, user := range users {
		if user.ID == id {
			writeJSON(w, http.StatusOK, user)
			return
		}
	}

	writeError(w, http.StatusNotFound, "user not found")
}

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
