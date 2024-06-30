package user

import (
	"github.com/lborres/gohttpstudy/types"
)

type Handler struct {
	storage types.UserStorage
}

func NewHandler(storage types.UserStorage) *Handler {
	return &Handler{storage: storage}
}
