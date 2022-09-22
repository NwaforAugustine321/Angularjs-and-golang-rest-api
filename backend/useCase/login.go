package useCase

import (
	"github.com/go/resst-app/repointerfaces"
)



func  Login(repo repointerfaces.IUser) {
	repo.Login()
}
