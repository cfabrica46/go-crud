package handler

import (
	"net/http"

	"github.com/cfabrica46/go-crud/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	email := c.MustGet("email").(string)

	id, err := user.InsertUser(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrMessage": "Username is already in use",
		})
		return
	}

	userToken, err := token.GenerateToken(id, username, email, "private.key", jwt.SigningMethodHS256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	// token := structure.Token{Content: user.Token}

	c.JSON(http.StatusOK, userToken)
}
