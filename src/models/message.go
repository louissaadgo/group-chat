package models

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
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
