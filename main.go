package main

import (
	"log"

	"github.com/cfabrica46/go-crud/database/userdb"
)

func main() {
	defer userdb.Close()

	r := setupRouter()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
