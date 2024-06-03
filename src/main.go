package main

import (
	"github.com/gin-gonic/gin"
	"interswitch_go_testing/src/config"
	"interswitch_go_testing/src/services/authService"
	"log"
)

func main() {
	config.LoadConfig()
	r := gin.Default()

	// Renew the token before starting the server
	if err := authService.Init(); err != nil {
		log.Fatalf("Failed to renew token: %v", err)
	}

	// Start token renewal in the background
	go authService.StartTokenRenewal()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/token", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"token": authService.GetToken(),
		})
	})

	err := r.Run(":8082") // listen and serve on 0.0.0.0:8082, if addr is empty, defaults to 8080
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
