package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Function to generate the rabbit yml file
func main() {
	// Create the file
	err := os.WriteFile(filePath(), fileContent(), 0644)
	if err != nil {
		log.Panic(err, "")
	}
	fmt.Println("created rabbit.yml")
}

// Path of the file
func filePath() string {
	return "./rabbit.yml"
}

// Content of the file, all extract from .env file
func fileContent() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	return []byte(fmt.Sprintf(`# File for the connection to rabbitMQ
rabbitmq:
  username: %s
  password: %s
  host: %s
  port: %v
  vhost: %s`,
		os.Getenv("DEV_USERNAME"),
		os.Getenv("DEV_PASSWORD"),
		os.Getenv("DEV_HOST"),
		os.Getenv("DEV_PORT"),
		os.Getenv("DEV_VHOST"),
	),
	)
}
