package userdb

import (
	"testing"
)

func TestOpen(t *testing.T) {
	dbTest, err := Open()
	if err != nil {
		t.Error(err.Error())
	}
	if dbTest == nil {
		t.Error("got db: nil")
	}

	//first error
	aux := dbDriver
	dbDriver = ""
	_, err = Open()
	if err == nil {
		t.Error("error was expected")
	}
	dbDriver = aux

	//second error
	aux = psqlInfo
	psqlInfo = ""
	_, err = Open()
	if err == nil {
		t.Error("error was expected")
	}
	psqlInfo = aux

}

/* func TestGet(t *testing.T) {
	dbTest := Get()
	if dbTest != db {
		t.Error("Fail to get db")
	}
} */

func TestClose(t *testing.T) {
	Open()
	err := Close()
	if err != nil {
		t.Error("Fail to close db")
	}

	//first err
	db = nil
	err = Close()
	if err == nil {
		t.Error("error was expected")
	}
	Open()
}
