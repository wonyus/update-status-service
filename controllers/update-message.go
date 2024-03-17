package controllers

import (
	"encoding/json"
	"log"

	"github.com/wonyus/update-status-service/initials"
)

func MessageUpdateStatus(message []byte) error {

	var body UpdateStatusRequest
	err := json.Unmarshal(message, &body)
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	log.Println(body)
	// for _, moisture := range body.Moistures {
	// 	initials.DB.Exec("UPDATE moisture SET moisture_value = $1 WHERE id = $2", moisture.MoistureValue, moisture.SensorPin)
	// }
	// for _, humidity := range body.Humiditys {
	// 	initials.DB.Exec("UPDATE humidity SET humidity_value = $1, temperature = $2 WHERE id = $3", humidity.HumidityValue, humidity.Temperature, humidity.SensorPin)
	// }
	for _, switchs := range body.Switchs {
		initials.DB.Exec("UPDATE switch SET status = $1 WHERE id = $2", switchs.Value, switchs.ID)
	}

	return nil
}
