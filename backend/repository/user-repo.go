package repository

import (
	"github.com/go/resst-app/data/sql"
	"github.com/go/resst-app/repointerfaces"
)

var userDatasource = sql.NewUserDatasource()

type userRepo struct{}

func NewUserService() repointerfaces.IUser {
	return &userRepo{}
}

func (repo *userRepo) Login() {
	userDatasource.Login()
}
