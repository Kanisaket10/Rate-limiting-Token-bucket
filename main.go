package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status  string `json:"status"`
	Body string `json:"message"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	message := Message{
		Status: "success",
		Body:   "Hi! You have reached the endpoint.",
	}
	err := json.NewEncoder(writer).Encode(message)
	if err != nil {
		http.Error(writer, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func  main()  {
	http.Handle("/ping", rateLimter(endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Server failed to start ", err)
	}
}