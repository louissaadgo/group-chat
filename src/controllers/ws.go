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
		fmt.Printf("Connection - User: %s\n", ep.SocketUUID)
	})

	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {
		fmt.Printf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data))
		message := models.Message{}
		err := json.Unmarshal(ep.Data, &message)
		if err != nil {
			fmt.Println(err)
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
