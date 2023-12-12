package main

import (
	"github.com/lcardelli/sust-go/src/controller/routes"
	"log"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := 
}