package main

import (
	"golang-ivs/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	api.Run()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
