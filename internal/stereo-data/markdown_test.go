package stereo_data_test

import (
	stereo_data "stereo-server/internal/stereo-data"
	"testing"
)

func TestMarkdownHandler(t *testing.T) {
	testStr := `
{
  "venue": "Stereo",
  "date": "Sept 2025",
  "ticketLink": "https://www.google.com",
  "events": [
    {
      "date": "Sept 1th",
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
        },
        {
          "name": "derf fred",
          "headliner": false,
          "allNightLong": false,
          "streamingPlatforms": []
        }
      ]
    },
    {
      "date": "Sept 2th",
      "artists": [
        {
          "name": "fred derf",
          "headliner": false,
          "allNightLong": false,
          "streamingPlatforms": [
            {
              "company": "spotify",
              "id": "1234"
            },
            {
              "company": "soundCloud",
              "id": "1234"
            }
          ]
        },
        {
          "name": "derf fred",
          "headliner": true,
          "allNightLong": false,
          "streamingPlatforms": [
            {
              "company": "soundCloud",
              "id": "1234"
            }
          ]
        }
      ]
    },
    {
      "date": "Sept 3th",
      "artists": [
        {
          "name": "fred derf",
          "allNightLong": true,
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

	expected := `[Tickets](https://www.google.com)
^(Sept 1th)
***fred derf*** - [Spotify](https://open.spotify.com/artist/1234)
*derf fred*
---
^(Sept 2th)
***derf fred*** - [SoundCloud](https://soundcloud.com/1234)
*fred derf* - [Spotify](https://open.spotify.com/artist/1234) - [SoundCloud](https://soundcloud.com/1234)
---
^(Sept 3th)
***fred derf*** ANL - [Spotify](https://open.spotify.com/artist/1234)
Let me know if there's any formatting issues or typos. I've added Spotify or Soundcloud profiles for ease of artist discovery. Correct me if I'm wrong.`

	data.ProcessJsonFromInputString()
	res := data.ConvertToMarkdown()

	if expected != res {
		t.Errorf("expected markdown to be %v; got %v", expected, res)
	}
}
