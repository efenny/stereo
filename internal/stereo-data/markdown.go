package stereo_data

import (
	"bytes"
	"sort"

	md "github.com/nao1215/markdown"
)

func (d *Data) ConvertToMarkdown() string {
	buf := new(bytes.Buffer)

	markdownBuilder := md.NewMarkdown(buf)
	redditMd := NewRedditMarkdown()

	data := d.InputData

	markdownBuilder.PlainText(md.Link("Tickets", data.TicketLink))
	markdownBuilder.PlainText("")

	for eventI, event := range data.Events {
		markdownBuilder.PlainText(redditMd.RedditSuperscript(event.Date))
		markdownBuilder.PlainText("")

		if event.DayRave {
			markdownBuilder.PlainText("Day Rave")
			markdownBuilder.PlainText("")
		}

		// INFO: make sure the headliner is always at the top
		sort.Slice(event.Artists, func(i, j int) bool {
			return event.Artists[i].Headliner && !event.Artists[j].Headliner
		})

		for _, artist := range event.Artists {
			artistStr := ""
			if artist.Headliner || artist.AllNightLong {
				artistStr += md.BoldItalic(artist.Name)
			} else {
				artistStr += md.Italic(artist.Name)
			}

			if artist.AllNightLong {
				if event.DayRave {
					artistStr += " ADL"
				} else {
					artistStr += " ANL"
				}
			}

			for _, streamingPlatform := range artist.StreamingPlatforms {
				if streamingPlatform.Id != "" && streamingPlatform.Company != "" {
					if streamingPlatform.Company == "spotify" {
						artistStr += " - " + md.Link("Spotify", "https://open.spotify.com/artist/"+streamingPlatform.Id)
					}

					if streamingPlatform.Company == "soundCloud" {
						artistStr += " - " + md.Link("SoundCloud", "https://soundcloud.com/"+streamingPlatform.Id)
					}
				}
			}

			markdownBuilder.PlainText(artistStr)
			markdownBuilder.PlainText("")

			// INFO: there should only be one artist for an ANL
			if artist.AllNightLong {
				break
			}
		}

		if eventI+1 != len(data.Events) {
			markdownBuilder.PlainText("")
			markdownBuilder.PlainText("")
			markdownBuilder.PlainText("*---*")
			markdownBuilder.PlainText("")
		}
	}

	markdownBuilder.PlainText("")
	markdownBuilder.PlainText("")
	markdownBuilder.PlainText("")
	markdownBuilder.PlainText("Let me know if there's any formatting issues or typos. I've added Spotify or Soundcloud profiles for ease of artist discovery. Correct me if I'm wrong.")

	markdownBuilder.Build()

	return buf.String()
}
