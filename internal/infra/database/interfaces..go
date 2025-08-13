package database

import "github.com/tonyops/GO-API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
