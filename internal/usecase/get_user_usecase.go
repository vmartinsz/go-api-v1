package usecase

import (
	"go-api-v1/internal/infra/contracts"
	"go-api-v1/internal/infra/persistence"
	"go-api-v1/internal/value"
)

type GetUserUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewGetUserUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *GetUserUseCase {
	return &GetUserUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (guuc *GetUserUseCase) GetUser(id string) (value.GetUser, error) {
	user, err := guuc.repositories.UserRepository.GetUser(id)
	if err != nil {
		return value.GetUser{}, err
	}
	return value.GetUser{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Age:      user.Age,
	}, nil
}
