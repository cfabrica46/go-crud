package userdb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	PSQLHost     = "localhost"
	PSQLPort     = 5431
	PSQLUser     = "cfabrica46"
	PSQLPassword = "01234"
	PSQLDBName   = "go_crud"
	PSQLSSL      = "require"
)

var dbDriver = "postgres"
var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", PSQLHost, PSQLPort, PSQLUser, PSQLPassword, PSQLDBName, PSQLSSL)

var db *sql.DB

func init() {
	var err error
	db, err = Open()
	if err != nil {
		log.Fatal(err)
	}
}

func Open() (*sql.DB, error) {
	var err error
	db, err = sql.Open(dbDriver, psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}

func Close() (err error) {
	if db == nil {
		err = errors.New("database already close")
		return
	}

	defer db.Close()
	return
}
