package controllers

import (
	"net/http"

	"github.com/go/resst-app/config"
	"github.com/go/resst-app/repointerfaces"
	"github.com/go/resst-app/repository"
	"github.com/go/resst-app/useCase"
)




type loginController struct{
	context *config.AppContext
	repo  repointerfaces.IUser
}

func NewLoginController(ctx *config.AppContext) *loginController {
	return &loginController{ repo: repository.NewUserService()}
}

func (app *loginController) Login(response http.ResponseWriter, request *http.Request) {
  useCase.Login(app.repo)
}
