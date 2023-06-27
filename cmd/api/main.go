package main

import "github.com/joho/godotenv"

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No env file found")
	}
}
