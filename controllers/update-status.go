package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wonyus/update-status-service/initials"
	"github.com/wonyus/update-status-service/utils"
)

type UpdateStatusRequest struct {
	MoistureValue int    `json:"moistureValue"`
	Mode          string `json:"mode"`
	Moistures     []struct {
		MoistureValue float64 `json:"vl"`
		SensorPin     int     `json:"seP"`
		LEDpin        int     `json:"leP"`
		SensorName    string  `json:"seN"`
	} `json:"moistures"`
	Switchs []struct {
		Value bool   `json:"vl"`
		ID    int    `json:"id"`
		Pin   int    `json:"swP"`
		Name  string `json:"swN"`
	} `json:"switchs"`
	Humiditys []struct {
		HumidityValue float64 `json:"hVl"`
		Temperature   float64 `json:"tVl"`
		SensorPin     int     `json:"seP"`
		SensorName    string  `json:"seN"`
	} `json:"humiditys"`
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var body UpdateStatusRequest
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// for _, moisture := range body.Moistures {
	// 	initials.DB.Exec("UPDATE moisture SET moisture_value = $1 WHERE id = $2", moisture.MoistureValue, moisture.SensorPin)
	// }
	// for _, humidity := range body.Humiditys {
	// 	initials.DB.Exec("UPDATE humidity SET humidity_value = $1, temperature = $2 WHERE id = $3", humidity.HumidityValue, humidity.Temperature, humidity.SensorPin)
	// }
	for _, switchs := range body.Switchs {
		initials.DB.Exec("UPDATE switch SET status = $1 WHERE id = $2", switchs.Value, switchs.ID)
	}

	resMessage := struct {
		updateStatus string
	}{
		updateStatus: "success",
	}

	utils.ResponseWriter(w, http.StatusOK, resMessage)
}
