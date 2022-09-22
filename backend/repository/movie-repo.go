package repository

import (
	"github.com/go/resst-app/config"
	"github.com/go/resst-app/data/sql"
	"github.com/go/resst-app/repointerfaces"
)

var movieDatasource repointerfaces.IMovie

type movieRepo struct{}

func NewMovieService(ctx *config.AppContext) repointerfaces.IMovie {
	movieDatasource = sql.NewMovieDatasource(ctx)
	return &movieRepo{}
}

func (datasource *movieRepo) GetSingleMovie(id int) (interface{}, error) {
	return movieDatasource.GetSingleMovie(id)
}

func (datasource *movieRepo) GetAllMovie() (interface{}, error) {
	return datasource.GetAllMovie()
}


func (datasource *movieRepo) EditMovie() error {
	return datasource.EditMovie()
}


func (datasource *movieRepo) CreatMovie() error {
	return datasource.CreatMovie()
}