package model

import (
	"encoding/json"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	IsAdmin   bool            `json:"is_admin"`
	Details   json.RawMessage `json:"details"`
	CreatedAt int64           `json:"created_at"`
	UpdateAt  int64           `json:"update_at"`
}

type Users []User
