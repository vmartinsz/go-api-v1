package usecase

import (
	"go-api-v1/internal/infra/contracts"
	"go-api-v1/internal/infra/persistence"
)

type UpdateUserUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateUserUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (uuuc *UpdateUserUseCase) PutUser(id string, attributes map[string]interface{}) error {
	return uuuc.repositories.UserRepository.PutUser(id, attributes)
}
