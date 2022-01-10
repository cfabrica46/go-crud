package token

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GenerateToken(id int, username, email string, keyData []byte, jwtMethod *jwt.SigningMethodHMAC) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwtMethod, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"uuid":     uuid.NewString(),
	})

	tokenString, err = token.SignedString(keyData)
	return
}

func ExtractClaims(tokenString, keyFilePath string, jwtMethod *jwt.SigningMethodHMAC) (id int, username, email string, err error) {
	token, err := jwt.Parse(tokenString, keyFunc(jwtMethod, keyFilePath))
	if err != nil {
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	idAux := claims["id"].(float64)
	id = int(idAux)

	username = claims["username"].(string)
	email = claims["email"].(string)
	return
}

func keyFunc(jwtMethod *jwt.SigningMethodHMAC, keyFilePath string) func(token *jwt.Token) (interface{}, error) {

	return func(token *jwt.Token) (interface{}, error) {

		secret, err := ioutil.ReadFile(keyFilePath)
		if err != nil {
			return nil, err
		}
		return secret, nil
	}
}
