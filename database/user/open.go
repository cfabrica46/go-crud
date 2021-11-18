package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	PSQLHost     = "localhost"
	PSQLPort     = 5431
	PSQLUser     = "cfabrica46"
	PSQLPassword = "01234"
	PSQLDBName   = "go_crud"
	PSQLSSL      = "require"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", PSQLHost, PSQLPort, PSQLUser, PSQLPassword, PSQLDBName, PSQLSSL)

func init() {
	err := Open()
	if err != nil {
		log.Fatal(err)
	}
}

func Open() (err error) {
	if db != nil {
		err = errors.New("database is already open")
		return
	}

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return
	}
	return
}

func Close() (err error) {
	err = db.Close()
	if err != nil {
		return
	}
	return
}
