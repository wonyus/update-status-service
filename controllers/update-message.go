package controllers

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wonyus/update-status-service/contexts"
	"github.com/wonyus/update-status-service/pkg/db"
)

type Repository struct {
	db  *db.DBClient
	ctx contexts.Resource
}

func NewRepository(ctx *contexts.Resource, db *db.DBClient) *Repository {
	return &Repository{
		db:  db,
		ctx: *ctx,
	}
}

type Moisture struct {
	MoistureValue float64 `json:"vl"`
	SensorPin     int     `json:"seP"`
	LEDpin        int     `json:"leP"`
	SensorName    string  `json:"seN"`
}

type Switch struct {
	Value bool   `json:"vl"`
	ID    int    `json:"id"`
	Uuid  string `json:"uuid"`
	Pin   int    `json:"swP"`
	Name  string `json:"swN"`
}

type Humidity struct {
	HumidityValue float64 `json:"hVl"`
	Temperature   float64 `json:"tVl"`
	SensorPin     int     `json:"seP"`
	SensorName    string  `json:"seN"`
}

type UpdateStatusRequest struct {
	MoistureValue float64    `json:"moistureValue"`
	Mode          string     `json:"mode"`
	Moistures     []Moisture `json:"moistures"`
	Switches      []Switch   `json:"switchs"`
	Humidities    []Humidity `json:"humidities"`
}

type returnMap struct {
	MoistureValue map[string]interface{}
	Mode          map[string]interface{}
	Moistures     []map[string]interface{}
	Switches      []map[string]interface{}
	Humidities    []map[string]interface{}
}

// CreateMapFromUpdateStatusRequest converts the UpdateStatusRequest struct to a map
func CreateMapFromUpdateStatusRequest(req UpdateStatusRequest) returnMap {
	moistureValue := map[string]interface{}{
		"vl": req.MoistureValue,
	}
	mode := map[string]interface{}{
		"vl": req.Mode,
	}

	moistures := make([]map[string]interface{}, len(req.Moistures))
	for i, moisture := range req.Moistures {
		moistures[i] = map[string]interface{}{
			"vl":  moisture.MoistureValue,
			"seP": moisture.SensorPin,
			"leP": moisture.LEDpin,
			"seN": moisture.SensorName,
		}
	}

	switches := make([]map[string]interface{}, len(req.Switches))
	for i, _ := range req.Switches {
		switches[i] = map[string]interface{}{
			"vl":   req.Switches[i].Value,
			"id":   req.Switches[i].ID,
			"uuid": req.Switches[i].Uuid,
			"swP":  req.Switches[i].Pin,
			"swN":  req.Switches[i].Name,
		}
	}

	humidities := make([]map[string]interface{}, len(req.Humidities))
	for i, humidity := range req.Humidities {
		humidities[i] = map[string]interface{}{
			"hVl": humidity.HumidityValue,
			"tVl": humidity.Temperature,
			"seP": humidity.SensorPin,
			"seN": humidity.SensorName,
		}
	}
	result := returnMap{
		MoistureValue: moistureValue,
		Mode:          mode,
		Moistures:     moistures,
		Switches:      switches,
		Humidities:    humidities,
	}

	return result
}

func (r Repository) MessageUpdateStatus(msg mqtt.Message) error {
	var body UpdateStatusRequest
	err := json.Unmarshal(msg.Payload(), &body)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	// client := influx.NewInflux()
	// org := "wonyus.tech"
	// bucket := "Switch"
	// writeAPI := client.WriteAPIBlocking(org, bucket)

	// log.Println("Received message: ", body)
	// for _, sw := range body.Switches {
	// 	tags := map[string]string{
	// 		"topic": msg.Topic(),
	// 		"uuid":  sw.Uuid,
	// 	}

	// 	fields := map[string]interface{}{
	// 		"vl":  sw.Value,
	// 		"id":  sw.ID,
	// 		"swP": sw.Pin,
	// 		"swN": sw.Name,
	// 	}

	// 	swdt := influxdb2.NewPoint("switch", tags, fields, time.Now())
	// 	if err := writeAPI.WritePoint(context.Background(), swdt); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	log.Println("Write point: ", swdt)

	// }

	// for _, moisture := range body.Moistures {
	// 	initials.DB.Exec("UPDATE moisture SET moisture_value = $1 WHERE id = $2", moisture.MoistureValue, moisture.SensorPin)
	// }
	// for _, humidity := range body.Humiditys {
	// 	initials.DB.Exec("UPDATE humidity SET humidity_value = $1, temperature = $2 WHERE id = $3", humidity.HumidityValue, humidity.Temperature, humidity.SensorPin)
	// }

	for _, switchs := range body.Switches {
		r.db.Exec("UPDATE switchs SET status = $1 WHERE id = $2", switchs.Value, switchs.ID)
	}

	return nil
}
