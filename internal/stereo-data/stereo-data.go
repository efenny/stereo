package stereo_data

import (
	"encoding/json"
	"fmt"
	"os"
)

type StreamingPlatform struct {
	Company string `json:"company"`
	Id      string `json:"id"`
}

type Artist struct {
	Name               string              `json:"name"`
	StreamingPlatforms []StreamingPlatform `json:"streamingPlatforms"`
	Headliner          bool                `json:"headliner"`
	AllNightLong       bool                `json:"allNightLong"`
}

type Event struct {
	Date    string   `json:"date"`
	DayRave bool     `json:"dayRave"`
	Artists []Artist `json:"artists"`
}

type Month struct {
	Date       string  `json:"date"`
	Venue      string  `json:"venue"`
	TicketLink string  `json:"ticketLink"`
	Events     []Event `json:"events"`
}

type Data struct {
	InputString string
	InputData   Month
}

func NewData(inputString string) Data {
	return Data{InputString: inputString}
}

func (d *Data) ProcessJsonFromInputString() {
	var res Month
	input := []byte(d.InputString)

	if !json.Valid(input) {
		fmt.Print("JSON is not valid")
		os.Exit(1)
	}

	err := json.Unmarshal(input, &res)
	if err != nil {
		fmt.Printf("There was an error parsing the input: %s", err)
		os.Exit(1)
	}

	d.setInputData(res)
}

func (d *Data) setInputData(val Month) {
	d.InputData = val
}

func (d *Data) GetInputData() Month {
	return d.InputData
}
