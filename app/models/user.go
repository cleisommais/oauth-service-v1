package models

import "time"

type User struct {
	Id       int  `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}