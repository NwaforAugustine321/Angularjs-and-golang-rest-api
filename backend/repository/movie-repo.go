package repository

import (
	"github.com/go/resst-app/config"
	"github.com/go/resst-app/data/sql"
	"github.com/go/resst-app/model"
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
	return  movieDatasource.GetAllMovie()

	

}

func (datasource *movieRepo) EditMovie(movie *model.Movie) error {
	return movieDatasource.EditMovie(movie)
}

func (datasource *movieRepo) CreatMovie(movie *model.Movie) (string, error) {
	return movieDatasource.CreatMovie(movie)
}
