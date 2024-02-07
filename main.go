package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/wonyus/update-status-service/controllers"
	"github.com/wonyus/update-status-service/initials"
)

var client mqtt.Client

func init() {
	initials.InitDB()

}

func main() {
	client := initials.InitialMqttClient(controllers.MessagePubHandler)
	controllers.DefaultSubscribeHandler(client)

	route := gin.Default()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "online", "message": "Hello, world!"})
	})
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": 200, "message": "pong"})
	})
	route.GET("/test/publish", func(c *gin.Context) {
		controllers.Publish(client)
		c.JSON(200, gin.H{"status": 200, "message": "pong"})
	})
	route.Run(":80")
}
