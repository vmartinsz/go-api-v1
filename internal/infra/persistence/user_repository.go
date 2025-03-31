package persistence

import (
	"go-api-v1/internal/entity"
	"go-api-v1/internal/infra/contracts"
	"time"
)

type UserRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (ur UserRepositoryImpl) CreateUser(user *entity.User) error {
	return ur.Postgres.Client().Create(user).Error
}

func (ur UserRepositoryImpl) DeleteUser(id string) error {
	return ur.Postgres.Client().Delete(&entity.User{}, "id = ?", id).Error
}

func (ur UserRepositoryImpl) GetUser(id string) (*entity.User, error) {
	var user entity.User
	err := ur.Postgres.Client().Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur UserRepositoryImpl) PutUser(id string, attributes map[string]interface{}) error {
	attributes["updated_at"] = time.Now().UTC()

	err := ur.Postgres.Client().
		Model(&entity.User{}).
		Where("ïd = ?", id).
		UpdateColumns(attributes).
		Error

	return err
}
