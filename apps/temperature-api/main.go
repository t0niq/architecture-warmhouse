package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensorId"`
	SensorType  string    `json:"sensorType"`
	Description string    `json:"description"`
}

func getDefaultLocation(sensorID string) string {
	switch sensorID {
	case "1":
		return "Living Room"
	case "2":
		return "Bedroom"
	case "3":
		return "Kitchen"
	default:
		return "Unknown"
	}
}

func getDefaultSensorID(location string) string {
	switch location {
	case "Living Room":
		return "1"
	case "Bedroom":
		return "2"
	case "Kitchen":
		return "3"
	default:
		return "0"
	}
}

func getSensorType(location string) string {
	switch location {
	case "Living Room", "Bedroom", "Kitchen":
		return "HomeHeatSensor"
	default:
		return "Unknown"
	}
}

func getStatus(temperature float64) string {
	if temperature < 50 && temperature > -50 {
		return "active"
	}
	return "broken"
}
func getDescription(location string, temperature float64) string {
	return location + " temperature is " + strconv.FormatFloat(temperature, 'f', 2, 64)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")

	if location == "" {
		location = getDefaultLocation(sensorID)
	}

	if sensorID == "" {
		sensorID = getDefaultSensorID(location)
	}

	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	temperature := 15 + random.Float64()*15

	response := TemperatureResponse{
		Value:       temperature,
		Unit:        "Celsius",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      getStatus(temperature),
		SensorID:    sensorID,
		SensorType:  getSensorType(location),
		Description: getDescription(location, temperature),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/temperature", temperatureHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}
