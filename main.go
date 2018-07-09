package main

import (
	"os"
	"github.com/lfbos/backing-fulfillment/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
