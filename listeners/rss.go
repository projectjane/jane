package listeners

import (
	"github.com/SlyMarbo/rss"
	"github.com/kennygrant/sanitize"
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/outputs"
	"html"
	"log"
	"strings"
	"time"
)

func Rss(config *configs.Config, listener configs.Listener) {
	lastMarker := ""
	for {
		feed, err := rss.Fetch(listener.Input)
		if err != nil {
			log.Println(err)
		}
		var messages []outputs.Message
		for i := len(feed.Items) - 1; i >= 0; i-- {
			if lastMarker == "" {
				// TODO: Change the 3 to zero
				lastMarker = feed.Items[3].Date.String()
			}
			item := feed.Items[i]
			if item.Date.String() > lastMarker {
				status := "NONE"
				if listener.SuccessMatch != "" {
					if strings.Contains(item.Title, listener.SuccessMatch) {
						status = "SUCCESS"
					}
					if strings.Contains(item.Content, listener.SuccessMatch) {
						status = "SUCCESS"
					}
				}
				if listener.FailureMatch != "" {
					if strings.Contains(item.Title, listener.FailureMatch) {
						status = "FAIL"
					}
					if strings.Contains(item.Title, listener.FailureMatch) {
						status = "FAIL"
					}
				}
				m := outputs.Message{
					Destination: listener.Output,
					Title:       listener.Name + " " + html.UnescapeString(sanitize.HTML(item.Title)),
					Description: html.UnescapeString(sanitize.HTML(item.Content)),
					Link:        item.Link,
					Status:      status,
				}
				messages = append(messages, m)
				if i == 0 {
					lastMarker = item.Date.String()
				}
			}
		}
		for _, m := range messages {
			outputs.Output(config, m)
		}
		time.Sleep(120 * time.Second)
	}
}
