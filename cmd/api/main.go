package main

import (
	"go-api/internal/router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8000")
}
