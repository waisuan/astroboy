package model

type Message struct {
	T       string `json:"t"`
	Key     string `json:"key"`
	Payload string `json:"payload"`
}
