package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, error := strconv.Atoi(os.Getenv("API_PORT"))

	if error != nil {
		log.Fatal("Error loading API_PORT")
	}

	server := NewServer("localhost:" + strconv.Itoa(port))
	fmt.Println("The server is running in por: ", port)

	// Routers
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("POST", "/logs", server.AddMiddleware(HandleLog, CheckAuth(), Logging()))
	server.Listen()
}
