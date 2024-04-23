package types

import "time"

type User struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	CreateDatetime time.Time `json:"createDatetime"`
}
