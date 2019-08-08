package model

import "time"

type Book struct {
	Id        int64
	Title     string
	Author    string
	IssuedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
