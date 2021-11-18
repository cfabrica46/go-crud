package handler

import (
	"net/http"

	"github.com/cfabrica46/go-crud/database/cache"
	"github.com/cfabrica46/go-crud/database/userdb"
	"github.com/cfabrica46/go-crud/structure"
	"github.com/cfabrica46/go-crud/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	email := c.MustGet("email").(string)

	id, err := userdb.InsertUser(username, password, email)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to insert user"})
		return
	}

	userToken, err := token.GenerateToken(id, username, email, "key.pem", jwt.SigningMethodHS256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Creating Token"})
		return
	}

	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK, Content: userToken})
}

func SignIn(c *gin.Context) {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	// email := c.MustGet("email").(string)

	user, err := userdb.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Sign In"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, structure.ResponseHTTP{Code: http.StatusUnauthorized, ErrorText: "Error User Not Found"})
		return
	}

	userToken, err := token.GenerateToken(user.ID, user.Username, user.Email, "key.pem", jwt.SigningMethodHS256)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to create token"})
		return
	}

	err = cache.SetToken(userToken)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to set Token"})
		return
	}

	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK, Content: userToken})
}
