package routes

import (
	"go-api-v1/internal/controller"

	"github.com/gin-gonic/gin"
)

type V1Controllers struct {
	UsersController *controller.UsersController
}

func NewV1Controller() *V1Controllers {
	return &V1Controllers{
		UsersController: controller.NewUsersController(),
	}
}

func (v1 *V1Controllers) StartRoutes(group *gin.RouterGroup) {
	group.POST("/user", v1.UsersController.CreateUser)
	group.GET("/user/:id", v1.UsersController.GetUser)
	group.PATCH("/user/:id", v1.UsersController.PutUser)
	group.DELETE("/user/:id", v1.UsersController.DeleteUser)
}
