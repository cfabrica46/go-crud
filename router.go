package main

import (
	"github.com/cfabrica46/go-crud/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func setupRouter() (r *gin.Engine) {
	r = gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./statics", true)))

	s := r.Group("/api/v1")
	{
		getuserFromBody := s.Group("/")
		getuserFromBody.Use(middleware.GetUserFromBody)
		{
			getuserFromBody.GET("/", nil)
		}

		getuserFromToken := s.Group("/")
		getuserFromToken.Use(middleware.GetUserFromToken)
		{
			getuserFromToken.GET("/", nil)
		}
	}
	return
}
