package model

import "time"

type User struct {
	ID uint32 `json:"id"`

	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Tokens []Token `json:"tokens"`
}
