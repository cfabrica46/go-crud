package token

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GenerateToken(id int, username, email, keyFile string, jwtMethod *jwt.SigningMethodHMAC) (tokenString string, err error) {
	secret, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}

	token := jwt.NewWithClaims(jwtMethod, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"uuid":     uuid.NewString(),
	})

	tokenString, err = token.SignedString(secret)
	if err != nil {
		return
	}
	return
}

func ExtractClaims(tokenString, keyFile string, jwtMethod *jwt.SigningMethodHMAC) (id int, username, email string, err error) {
	token, err := jwt.Parse(tokenString, keyFunc(jwtMethod, keyFile))
	if err != nil {
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	id = claims["id"].(int)
	username = claims["username"].(string)
	email = claims["email"].(string)
	return
}

func keyFunc(jwtMethod *jwt.SigningMethodHMAC, keyFile string) func(token *jwt.Token) (interface{}, error) {

	return func(token *jwt.Token) (interface{}, error) {

		secret, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		return secret, nil
	}
}
