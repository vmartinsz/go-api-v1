package main

import (
	"fmt"
	"go-api-v1/cmd/api/server/routes"
	"go-api-v1/config"
	"go-api-v1/internal/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	envs, err := config.Load("config/config.yml")

	if err != nil {
		envs, err = config.Load("../../../config/config.yml")
		if err != nil {
			panic(err)
		}
	}

	setup := setup.NewSetup()
	setup.Logger("go-api-v1")
	setup.Postgres()
	setup.Repositories()

	setup.Finish()

	api := setupServer()
	api.Run(fmt.Sprintf(":%d", envs.Server.Port))
}

func setupServer() *gin.Engine {
	server := gin.Default()

	apiV1 := server.Group("/api/v1")

	routes.NewV1Controller().StartRoutes(apiV1)

	return server

}
