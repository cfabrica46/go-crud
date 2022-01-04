package main

import (
	"log"
	"os"

	"github.com/cfabrica46/go-crud/database/userdb"
	"github.com/joho/godotenv"
)

func main() {
	defer userdb.Close()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("unread .env")
	}

	portHTTP := os.Getenv("PORT")
	portHTTPS := os.Getenv("PORTHTTPS")

	runServer(portHTTP, portHTTPS)
}
