package useCase

import (
	"github.com/go/resst-app/model"
	"github.com/go/resst-app/repointerfaces"
)



func EditMovie(repo repointerfaces.IMovie, movie *model.Movie)error {
	return repo.EditMovie()
}
