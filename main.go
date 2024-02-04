package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/wonyus/update-status-service/controllers"
	"github.com/wonyus/update-status-service/initials"
)

func init() {
	initials.InitDB()
	client := initials.InitialMqttClient(controllers.MessagePubHandler)
	controllers.DefaultSubscribeHandler(client)
}

func main() {
	route := gin.Default()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "online", "message": "Hello, world!"})
	})
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": 200, "message": "pong"})
	})
	route.Run(":80")
}
