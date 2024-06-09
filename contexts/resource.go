package contexts

import (
	"os"
	"time"

	"github.com/wonyus/update-status-service/utils"
)

type Resource struct {
	DB_PG_URL        string
	PgConnStr        string
	InfluxUrl        string
	InfluxToken      string
	MQTT_BROKER_URL  string
	MQTT_BROKER_PORT string
	MQTT_USERNAME    string
	MQTT_PASSWORD    string
	MQTT_CLIENT_ID   string
	MQTT_TOPIC       string
	MQTT_KEEP_ALIVE  time.Duration
}

func NewResource() *Resource {
	return &Resource{
		DB_PG_URL:        utils.Strip(os.Getenv("DB_PG_URL")),
		InfluxUrl:        utils.Strip(os.Getenv("INFLUXDB_URL")),
		InfluxToken:      utils.Strip(os.Getenv("INFLUXDB_TOKEN")),
		MQTT_BROKER_URL:  utils.Strip(os.Getenv("MQTT_BROKER_URL")),
		MQTT_BROKER_PORT: utils.Strip(os.Getenv("MQTT_BROKER_PORT")),
		MQTT_USERNAME:    utils.Strip(os.Getenv("MQTT_USERNAME")),
		MQTT_PASSWORD:    utils.Strip(os.Getenv("MQTT_PASSWORD")),
		MQTT_CLIENT_ID:   utils.Strip(os.Getenv("MQTT_CLIENT_ID")),
		MQTT_TOPIC:       utils.Strip(os.Getenv("MQTT_TOPIC")),
		MQTT_KEEP_ALIVE:  time.Second * 60,
	}
}
