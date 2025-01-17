package main

import (
	"encoding/json"
	"net/http"
	"reflect"

	. "github.com/mlabouardy/dialogflow-watchnow-messenger/models"
	"fmt"
)

func VerificationEndPoint(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("hub.challenge")
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	fmt.Println(token)
	if mode != "" && token == "da96866a820df533abce43f061eb4e9e" {
		w.WriteHeader(200)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Error, wrong validation token"))
	}
}

func MessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	var callback Callback
	json.NewDecoder(r.Body).Decode(&callback)
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				if !reflect.DeepEqual(event.Message, Message{}) && event.Message.Text != "" {
					ProcessMessage(event)
				}
			}
		}
		w.WriteHeader(200)
		w.Write([]byte("Got your message"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Message not supported"))
	}
}
