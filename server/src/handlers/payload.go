package handlers

type payload struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
	UID     string `json:"uid,omitempty"`
}
