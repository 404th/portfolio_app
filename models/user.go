package models

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password int32     `json:"password"`
}
