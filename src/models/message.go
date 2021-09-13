package models

type Message struct {
	Type     string `json:"type"`
	Name     string `json:"fullname"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Users    []User `json:"users"`
}

func (message Message) Validate() bool {
	if len(message.Name) <= 0 || len(message.Message) <= 0 {
		return false
	}

	if len(message.Name) > 20 || len(message.Message) > 60 {
		return false
	}

	return true
}
