package main

import (
	"log"

	"github.com/cfabrica46/go-crud/database/user"
)

func main() {
	defer user.Close()

	r := setupRouter()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
