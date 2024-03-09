package main

import (
	server "Delver/Server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	server := server.NewServer()
	fmt.Println("Server Listening On Port: 8000")
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
