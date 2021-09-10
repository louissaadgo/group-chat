package controllers

import (
	"encoding/json"
	"fmt"

	ikisocket "github.com/antoniodipinto/ikisocket"
	"github.com/louissaadgo/group-chat/src/models"
)

var clients = make(map[string]string)
var users = []string{}

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
		if message.Type == "change_name" {
			user := clients[ep.SocketUUID]
			for i := 0; i < len(users); i++ {
				if user == users[i] {
					users = append(users[:i], users[i+1:]...)
				}
			}
			clients[ep.SocketUUID] = message.Name
			users = append(users, clients[ep.SocketUUID])
			response := models.Message{
				Type:  "users_list",
				Users: users,
			}

			sb, _ := json.Marshal(response)

			ep.Kws.Broadcast(sb, false)
			return
		}
		if message.Type == "message" {
			response := models.Message{
				Type:    "message",
				Message: fmt.Sprintf("<strong>%s</strong>: %s", message.Name, message.Message),
			}
			sb, _ := json.Marshal(response)

			ep.Kws.Broadcast(sb, false)
			return
		}
	})

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		user := clients[ep.SocketUUID]
		for i := 0; i < len(users); i++ {
			if user == users[i] {
				users = append(users[:i], users[i+1:]...)
			}
		}
		delete(clients, ep.SocketUUID)

		message := models.Message{
			Type:  "users_list",
			Users: users,
		}

		sb, _ := json.Marshal(message)
		ep.Kws.Broadcast(sb, true)
	})

	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		user := clients[ep.SocketUUID]
		for i := 0; i < len(users); i++ {
			if user == users[i] {
				users = append(users[:i], users[i+1:]...)
			}
		}
		delete(clients, ep.SocketUUID)

		message := models.Message{
			Type:  "users_list",
			Users: users,
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

	clients[kws.UUID] = kws.UUID
	users = append(users, clients[kws.UUID])

	message := models.Message{
		Type:  "users_list",
		Users: users,
	}

	sb, _ := json.Marshal(message)

	kws.Broadcast(sb, false)
}
