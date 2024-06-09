package client

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wonyus/update-status-service/contexts"
	"github.com/wonyus/update-status-service/controllers"
)

type MQTTClient interface {
	*mqtt.Client
}

type MqttHandler struct {
	client   mqtt.Client
	ctx      contexts.Resource
	repo     controllers.Repository
	received chan struct{}
}

func NewMqttHandler(ctx *contexts.Resource, repo *controllers.Repository, received chan struct{}) *MqttHandler {
	return &MqttHandler{
		ctx:      *ctx,
		repo:     *repo,
		received: received,
	}
}

func (m *MqttHandler) DefaultSubscribeHandler() mqtt.Token {
	topic := m.ctx.MQTT_TOPIC
	token := m.client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s", topic)
	return token
}

func (m *MqttHandler) connectHandler(client mqtt.Client) {
	log.Println("Connected")
	m.DefaultSubscribeHandler()
}

func (m *MqttHandler) connectLostHandler(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func (m *MqttHandler) MessageSubHandler(client mqtt.Client, msg mqtt.Message) {
	switch {
	case strings.Contains(msg.Topic(), "/switch/basic/"):
		m.repo.MessageUpdateStatus(msg)
		m.received <- struct{}{}
	}
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func (m *MqttHandler) InitialMqttClient(rev mqtt.MessageHandler) {
	port, err := strconv.Atoi(m.ctx.MQTT_BROKER_PORT)
	if err != nil {
		panic(err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", m.ctx.MQTT_BROKER_URL, port))
	opts.SetClientID(m.ctx.MQTT_CLIENT_ID)
	opts.SetUsername(m.ctx.MQTT_USERNAME)
	opts.SetPassword(m.ctx.MQTT_PASSWORD)
	opts.SetDefaultPublishHandler(rev)
	opts.OnConnect = m.connectHandler
	opts.OnConnectionLost = m.connectLostHandler
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(true)
	opts.SetKeepAlive(m.ctx.MQTT_KEEP_ALIVE)

	m.client = mqtt.NewClient(opts)
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		panic(token.Error())
	}
}
