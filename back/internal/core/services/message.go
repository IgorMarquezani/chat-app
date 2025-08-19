package services

type APIMessage[T any] struct {
	Error   string   `json:"message"`
	Details []string `json:"details"`
	Succeed bool     `json:"succeed"`
	Status  uint32   `json:"-"`
	Data    T        `json:"data"`
}

func NewAPIMessage[T any](err string, details []string, succeed bool, status uint32, data T) APIMessage[T] {
	return APIMessage[T]{
		Error:   err,
		Details: details,
		Succeed: succeed,
		Status:  status,
		Data:    data,
	}
}
