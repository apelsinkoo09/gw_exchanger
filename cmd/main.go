package main

import (
	"log"
	"net/http"
)

func main() {

	log.Printf("Server is running on 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
