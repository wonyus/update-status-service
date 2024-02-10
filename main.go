package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/wonyus/update-status-service/controllers"
	"github.com/wonyus/update-status-service/initials"
)

var client mqtt.Client

func main() {
	initials.InitDB()
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

	// type Person struct {
	// 	Header string `pos:"1"`
	// 	Name   string `pos:"2"`
	// 	Tel    string `pos:"3"`
	// }

	// // Example input data
	// data := "header||123456789"

	// // Unmarshal the string back into a person struct
	// var person Person
	// err := utils.Unmarshal(data, &person)
	// if err != nil {
	// 	fmt.Println("Error unmarshaling string:", err)
	// 	return
	// }

	// fmt.Printf("Unmarshaled Person: %+v\n", person)

	// // Marshal the struct into a string
	// marshaledString, err := utils.Marshal(person)
	// if err != nil {
	// 	fmt.Println("Error marshaling person:", err)
	// 	return
	// }

	// fmt.Println("Marshalled string:", marshaledString)
}
