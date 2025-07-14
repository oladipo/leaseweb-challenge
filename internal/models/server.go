package models

type Server struct {
	ID        string `json:"id"`
	Model     string `json:"model"`
	RAM       string `json:"ram"`
	HDD       string `json:"hdd"`
	Location  string `json:"location"`
	Price     string `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
