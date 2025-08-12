package services

type APIMessage struct {
	Error   string   `json:"message"`
	Details []string `json:"details"`
	Succeed bool     `json:"succeed"`
	Status  uint32   `json:"-"`
}

func NewAPIMessage(err string, details []string, succeed bool, status uint32) APIMessage {
	return APIMessage{
		Error:   err,
		Details: details,
		Succeed: succeed,
		Status:  status,
	}
}
