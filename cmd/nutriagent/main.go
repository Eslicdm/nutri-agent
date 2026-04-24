package main

import (
	"log"

	"github.com/Eslicdm/nutri-agent/internal/cli"
	"github.com/joho/godotenv"
)

// export PATH=$PATH:$(go env GOPATH)/bin
// go install ./...
// nutriagent init

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	cli.Execute()
}
