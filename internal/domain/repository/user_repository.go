package repository

import "go-api-v1/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUser(id string) (*entity.User, error)
	DeleteUser(id string) error
	PutUser(id string, attributes map[string]interface{}) error
}
