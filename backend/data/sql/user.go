package sql

import (
	"github.com/go/resst-app/repointerfaces"
)

type userDatasource struct{}

func NewUserDatasource() repointerfaces.IUser {
	return &userDatasource{}
}

func (datasrc *userDatasource) Login() {

}
