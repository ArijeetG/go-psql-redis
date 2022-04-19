package main

import (
	"admin/authHandler"
	generalhandler "admin/generalHandler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		fmt.Println("Entering login...")
		authHandler.Login(c)
	})
	r.POST("/signup", func(c *gin.Context) {
		authHandler.Signup(c)
	})
	r.POST("/hello", func(c *gin.Context) {
		generalhandler.GetUserInfo(c)
	})
	r.POST("/editInfo", func(c *gin.Context) {
		generalhandler.EditUserInfo(c)
	})
	r.POST("delInfo", func(c *gin.Context) {
		generalhandler.DeleteUserInfo(c)
	})
	err := r.Run("0.0.0.0:4000")
	if err != nil {
		log.Panic(err.Error())
	}

}
