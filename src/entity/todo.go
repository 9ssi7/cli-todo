package entity

import "time"

type Todo struct {
	Content string `json:"content"`
	Done bool `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}