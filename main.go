package main

import (
	"github.com/fullstacker-go/practice_gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.GET("/", handler.HomeHandler)
	r.POST("/signup", handler.Signup)
	r.POST("/signin", handler.Signin)
	r.GET("/logout", handler.Logout)
	r.Run(":3000")
}
