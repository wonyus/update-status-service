package controllers

import (
	"log"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	switch {
	case strings.Contains(msg.Topic(), "/switch/basic/"):
		MessageUpdateStatus(msg.Payload())
	}

	// log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func DefaultSubscribeHandler(client mqtt.Client) {
	topic := os.Getenv("MQTT_TOPIC")
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s", topic)
}

func Subscribe(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s", topic)
}

func Publish(client mqtt.Client) {

}
