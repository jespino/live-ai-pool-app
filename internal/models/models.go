package models

import (
	"time"
)

type Option struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Votes int    `json:"votes"`
}

type Pool struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Options     []Option  `json:"options"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type Vote struct {
	ID        string    `json:"id"`
	PoolID    string    `json:"pool_id"`
	OptionID  string    `json:"option_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}