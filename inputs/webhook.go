package inputs

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hexbotio/hex/models"
)

type Webhook struct {
}

func (x Webhook) Read(inputMsgs chan<- models.Message, config models.Config) {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.WebhookPort),
		Handler: nil,
	}

	handle := func(w http.ResponseWriter, r *http.Request) {
		rawbody, err := ioutil.ReadAll(r.Body)
		body := string(rawbody)
		if err != nil {
			config.Logger.Error("Webhook Body Read" + " - " + err.Error())
		}
		defer r.Body.Close()
		message := models.NewMessage()
		message.Attributes["hex.service"] = "webhook"
		message.Attributes["hex.url"] = r.RequestURI
		message.Attributes["hex.ipaddress"] = r.RemoteAddr
		message.Attributes["hex.input"] = body
		inputMsgs <- message

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"serviced_by\":\"HexBot\", \"message_id\":\"" + message.Attributes["hex.id"] + "\"}"))
	}

	http.HandleFunc("/", handle)

	err := server.ListenAndServe()
	if err != nil {
		config.Logger.Error("Webhook Listner" + " - " + err.Error())
	}
}
