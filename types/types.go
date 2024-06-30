package types

import "time"

type User struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	Role           string    `json:"role"`
	CreateDatetime time.Time `json:"createDatetime"`
	ModDatetime    time.Time `json:"modDatetime"`
}

type Product struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	SellPrice      float32   `json:"sellPrice"`
	Currency       string    `json:"currency"`
	Stock          int64     `json:"stock"`
	CreateDatetime time.Time `json:"createDatetime"`
	ModDatetime    time.Time `json:"modDatetime"`
}

type ProductStorage interface {
	GetAllProducts() ([]*Product, error)
}

type CartStorage interface {
}

type CheckoutStorage interface {
}

type OrderStorage interface {
}

type UserStorage interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
}
