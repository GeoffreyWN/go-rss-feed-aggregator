package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Lets start already!!")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is missing in .env file")
	}

	fmt.Println("Port:", portString)
}
