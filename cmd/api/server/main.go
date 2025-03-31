package main

import (
	"fmt"
	"go-api-v1/cmd/api/server/routes"
	"go-api-v1/config"

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

	api := setupServer()
	api.Run(fmt.Sprintf(":%d", envs.Server.Port))
}

func setupServer() *gin.Engine {
	server := gin.Default()

	apiV1 := server.Group("/api/v1")

	routes.NewV1Controller().StartRoutes(apiV1)

	return server

}
