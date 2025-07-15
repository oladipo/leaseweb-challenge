package models

import "time"

type Server struct {
	ID        string    `json:"id"`
	Model     string    `json:"model"`
	RAM       string    `json:"ram"`
	HDD       string    `json:"hdd"`
	Location  string    `json:"location"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
