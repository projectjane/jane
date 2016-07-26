package core

import (
	"github.com/projectjane/jane/models"
)

func Version(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	if config.Version != "" {
		message.Out.Text = "Version: " + config.Version
	} else {
		message.Out.Text = "Version: Non Standard Build"
	}
	publishMsgs <- message
}
