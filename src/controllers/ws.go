package controllers

import (
	"encoding/json"
	"fmt"

	ikisocket "github.com/antoniodipinto/ikisocket"
	"github.com/louissaadgo/group-chat/src/models"
)

var clients = make(map[string]string)

var Users = []models.User{}

func RegisterWSEvents() {

	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		fmt.Printf("New Connection:\nUserUUID: %s\n", ep.SocketUUID)
	})

	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {
		message := models.Message{}
		err := json.Unmarshal(ep.Data, &message)
		if err != nil {
			fmt.Println(err)
			return
		}
		if message.Type == "message" {
			response := models.Message{
				Type:     "message",
				Message:  message.Message,
				Username: clients[ep.SocketUUID],
			}
			sb, _ := json.Marshal(response)

			ep.Kws.Broadcast(sb, true)
			return
		}
	})

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		username := clients[ep.SocketUUID]
		for in, i := range Users {
			if i.Username == username {
				Users[in].Status = "offline"
			}
		}
		delete(clients, ep.SocketUUID)

		message := models.Message{
			Type:  "users_list",
			Users: Users,
		}

		sb, _ := json.Marshal(message)
		ep.Kws.Broadcast(sb, true)
	})

	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		username := clients[ep.SocketUUID]
		for in, i := range Users {
			if i.Username == username {
				Users[in].Status = "offline"
			}
		}
		delete(clients, ep.SocketUUID)

		message := models.Message{
			Type:  "users_list",
			Users: Users,
		}

		sb, _ := json.Marshal(message)
		ep.Kws.Broadcast(sb, true)
	})

	// On error event
	ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
		fmt.Printf("Error event - User: %s\n", ep.SocketUUID)
	})

}

func WS(kws *ikisocket.Websocket) {

	clients[kws.UUID] = kws.Cookies("token")
	for in, i := range Users {
		if i.Username == kws.Cookies("token") {
			Users[in].Status = "online"
		}
	}
	fmt.Println(Users)
	message := models.Message{
		Type:  "users_list",
		Users: Users,
	}

	sb, _ := json.Marshal(message)

	kws.Broadcast(sb, false)
}
