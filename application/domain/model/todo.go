package model

import "time"

type Todo struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Auther    string    `json:"auther"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
