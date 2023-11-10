package main

import (
	"log"
	"ms-user-api/internal/http/router"
	"os"
	"time"
)

const TIMEOUT = 30 * time.Second

func main() {
	h := router.Handlers()
	err := h.Start(":8080")

	if err != nil {
		log.Fatal("Error starting api, error: " + err.Error())
		os.Exit(1)
	}
}
