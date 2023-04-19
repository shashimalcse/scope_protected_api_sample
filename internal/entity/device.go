package entity

type Issue struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Status      string `json:"status"`
}
