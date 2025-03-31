package usecase

import (
	"go-api-v1/internal/infra/contracts"
	"go-api-v1/internal/infra/persistence"
)

type DeleteUserUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteUserUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (duuc *DeleteUserUseCase) DeleteUser(id string) error {
	return duuc.repositories.UserRepository.DeleteUser(id)
}
