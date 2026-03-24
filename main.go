package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test"))
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server running on :8080")
}
