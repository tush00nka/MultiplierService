package main

import (
	"flag"
	"log"
	"multiplier/handler"
	"multiplier/server"
	"multiplier/service"
	"os"
)

func main() {
	rtp := flag.Float64("rtp", 1, "RTP value")
	flag.Parse()

	// check for rtp flag range
	if *rtp <= 0 || *rtp > 1 {
		log.Fatal("RTP must be in range (0;1.0]\n")
	}

	generatorService := service.NewGenerator(*rtp)
	generatorHandler := handler.NewGeneratorHandler(*generatorService)

	server := server.NewServer(*generatorHandler)

	// make it possible to use custom port via env variables
	// it would be more preferable to add a .env file, but for now this solution will do
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Println("SERVER_POST environment vartiable is not set. Using default port 64333")
		port = "64333"
	}

	server.Run(port)
}
