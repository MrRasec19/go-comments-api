package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"userName"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Status    byte      `json:"status"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
}
