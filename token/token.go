package token

import (
	"io/ioutil"

	"github.com/cfabrica46/go-crud/structure"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GenerateToken(id, username, email, keyFile string, jwtMethod *jwt.SigningMethodHMAC) (tokenString string, err error) {
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

func ExtractUserFromClaims(tokenString, keyFile string, jwtMethod *jwt.SigningMethodHMAC) (user structure.User, err error) {
	token, err := jwt.Parse(tokenString, keyFunc(jwtMethod, keyFile))
	if err != nil {
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	id := claims["id"].(string)
	user.ID = id

	username := claims["username"].(string)
	user.Username = username

	email := claims["email"].(string)
	user.Email = email

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
