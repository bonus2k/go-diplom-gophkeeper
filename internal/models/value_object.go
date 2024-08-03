package models

import "github.com/google/uuid"

type UserCtx struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Id       uuid.UUID `json:"id"`
}
