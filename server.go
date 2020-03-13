package main

import (
	"Person/routers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routers.Api()
}

