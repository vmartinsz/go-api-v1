package setup

import (
	"go-api-v1/config"
	"go-api-v1/internal/infra"
	"go-api-v1/internal/infra/database"
	"go-api-v1/internal/infra/logger"
	"go-api-v1/internal/infra/persistence"
	"log"
)

type Setup struct {
	app          *infra.Application
	repositories *persistence.PostgresRepositories
}

func NewSetup() Setup {
	Config, err := config.Load("../../../config/config.yaml")
	if err != nil {
		Config, err = config.Load("config/config.yaml")
		if err != nil {
			log.Fatal(err)
		}
	}

	return Setup{
		app: &infra.Application{
			Config: *Config,
		},
		repositories: &persistence.PostgresRepositories{},
	}
}

func (s Setup) Postgres() {
	s.app.Postgres, _ = database.NewPGGORMImpl(s.app.Config)
}

func (s Setup) Repositories() {
	s.app.Repositories.UserRepository = persistence.UserRepositoryImpl{Postgres: s.app.Postgres}
}

func (s Setup) Logger(taskname string) {
	s.app.Logger, _ = logger.New(taskname)
}

func (s Setup) Finish() {
	infra.App = *s.app
}
