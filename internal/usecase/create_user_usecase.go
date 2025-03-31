package usecase

import (
	"go-api-v1/internal/entity"
	"go-api-v1/internal/infra/contracts"
	"go-api-v1/internal/infra/persistence"
)

type CreateUserUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewCreateUserUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (cuuc *CreateUserUseCase) CreateUser(user *entity.User) error {
	return cuuc.repositories.UserRepository.CreateUser(user)
}
