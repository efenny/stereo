package stereo_data_test

import (
	"reflect"
	stereo_data "stereo-server/internal/stereo-data"
	"testing"
)

func TestStereoDataHandler(t *testing.T) {
	testStr := `
{
  "venue": "Stereo",
  "date": "Sept 2025",
  "ticketLink": "https://www.google.com",
  "events": [
    {
      "date": "Sept 5th",
      "artists": [
        {
          "name": "fred derf",
          "headliner": true,
          "allNightLong": false,
          "streamingPlatforms": [
            {
              "company": "spotify",
              "id": "1234"
            }
          ]
        }
      ]
    }
  ]
}
	`

	data := stereo_data.NewData(testStr)

	data.ProcessJsonFromInputString()
	res := data.GetInputData()

	artists := []stereo_data.Artist{{Name: "fred derf", Headliner: true, AllNightLong: false, StreamingPlatforms: []stereo_data.StreamingPlatform{{Id: "1234", Company: "spotify"}}}}
	events := []stereo_data.Event{{Date: "Sept 5th", Artists: artists}}

	expected := stereo_data.Month{
		Date:       "Sept 2025",
		Venue:      "Stereo",
		TicketLink: "https://www.google.com",
		Events:     events,
	}

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected data to be %v; got %v", expected, res)
	}
}
