package main

import (
	"github.com/cfabrica46/go-crud/handler"
	"github.com/cfabrica46/go-crud/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func setupRouter() (r *gin.Engine) {
	r = gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./statics", true)))

	s := r.Group("/api/v1")
	s.GET("/users", handler.GetAllUsers)
	{
		getuserFromBody := s.Group("/")
		getuserFromBody.Use(middleware.GetUserFromBody)
		{
			getuserFromBody.POST("/signin", handler.SignIn)
			getuserFromBody.POST("/signup", handler.SignUp)
		}

		getuserFromToken := s.Group("/")
		getuserFromToken.Use(middleware.GetUserFromToken)
		{
			getuserFromToken.GET("/user", handler.Profile)
			getuserFromToken.DELETE("/user", handler.DeleteUser)
			getuserFromToken.HEAD("/logout", handler.LogOut)
		}
	}
	return
}
