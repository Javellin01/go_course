package dto

type AdvertiseRequest struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
}
