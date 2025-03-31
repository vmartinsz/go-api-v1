package persistence

import "go-api-v1/internal/domain/repository"

type PostgresRepositories struct {
	UserRepository repository.UserRepository
}
