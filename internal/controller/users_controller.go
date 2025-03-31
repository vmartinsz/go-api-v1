package controller

import (
	"go-api-v1/internal/entity"
	"go-api-v1/internal/infra"
	"go-api-v1/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (uc *UsersController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	usecase := usecase.NewGetUserUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	user, err := usecase.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UsersController) CreateUser(ctx *gin.Context) {
	var requestParams entity.User
	if err := ctx.BindJSON(&requestParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
		return
	}

	usecase := usecase.NewCreateUserUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.CreateUser(&requestParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao tentar criar usuário"})
		return
	}

	ctx.JSON(http.StatusCreated, requestParams)
}

func (uc *UsersController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	usecase := usecase.NewDeleteUserUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar usuário"})
		return
	}

	ctx.JSON(http.StatusNoContent, http.NoBody)
}

func (uc *UsersController) PutUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var data map[string]interface{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "conteúdo do body inválido"})
	}

	usecase := usecase.NewUpdateUserUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.PutUser(id, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "erro interno de servidor ao atualizar informações"})
		return
	}

	ctx.JSON(http.StatusNoContent, http.NoBody)
}
