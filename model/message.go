package model

type Message struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	CreateTime string `json:"create_time"`
	Data       string `json:"data"`
}
