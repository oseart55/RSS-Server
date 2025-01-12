package main

import "net/http"

func handleReady(w http.ResponseWriter, r *http.Request) {
	type Message struct {
		Message string `json:"message"`
	}
	respondJSON(w, 200, Message{
		Message: "OK",
	})
}
