package token

import (
	"io/ioutil"
	"testing"

	"github.com/cfabrica46/go-crud/structure"
	"github.com/dgrijalva/jwt-go"
)

func TestGenerateToken(t *testing.T) {
	userTest := structure.User{
		ID:       1,
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	goodKeyPath := "test/server.key"

	goodKey, err := ioutil.ReadFile(goodKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	tokenString, err := GenerateToken(userTest.ID, userTest.Username, userTest.Email, goodKey, jwt.SigningMethodHS256)
	if err != nil {
		t.Error(err)
	}
	if tokenString == "" {
		t.Error("Error to generate token")
	}
}

func TestExtractClaims(t *testing.T) {
	userTest := structure.User{
		ID:       1,
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	goodKeyPath := "test/server.key"
	badKeyPath := "bad"

	goodKey, err := ioutil.ReadFile(goodKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	tokenString, err := GenerateToken(userTest.ID, userTest.Username, userTest.Email, goodKey, jwt.SigningMethodHS256)
	if err != nil {
		t.Fatal(err)
	}

	id, username, email, err := ExtractClaims(tokenString, goodKeyPath, jwt.SigningMethodHS256)
	if err != nil {
		t.Error(err)
	}
	if id != userTest.ID {
		t.Error("error to extract claims")
	}
	if username != userTest.Username {
		t.Error("error to extract claims")
	}
	if email != userTest.Email {
		t.Error("error to extract claims")
	}

	//with error
	_, _, _, err = ExtractClaims(tokenString, badKeyPath, jwt.SigningMethodHS256)
	if err == nil {
		t.Error("got nil; want error")
	}
}
