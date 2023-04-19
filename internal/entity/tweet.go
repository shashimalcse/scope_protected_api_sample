package entity

type Tweet struct {
	ID   string `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}
