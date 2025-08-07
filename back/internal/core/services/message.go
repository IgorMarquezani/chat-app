package services

type APIMessage struct {
	Error   string   `json:"message"`
	Details []string `json:"details"`
	Succeed bool     `json:"succeed"`
	Status  uint32   `json:"-"`
}
