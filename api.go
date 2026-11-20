package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"text"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Text: "Welcome to my Go REST API 🚀"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
