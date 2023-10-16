package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satuhin95/go-jwt/controller"
	"github.com/satuhin95/go-jwt/initial"
	"github.com/satuhin95/go-jwt/middleware"
)

func init() {
	initial.LoadEnv()
	initial.ConnectDB()
	initial.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.GET("/validate", middleware.RequireAuth, controller.Validate)
	r.Run()
}
