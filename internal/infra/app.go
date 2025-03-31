package infra

import (
	"go-api-v1/config"
	"go-api-v1/internal/infra/contracts"
	"go-api-v1/internal/infra/persistence"
)

type Application struct {
	Repositories persistence.PostgresRepositories
	Postgres     contracts.PostgresIface
	Logger       contracts.Logger
	Config       config.Config
}

var App Application
