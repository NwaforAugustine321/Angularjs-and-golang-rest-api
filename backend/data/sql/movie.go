package sql

import (
	"context"
	"log"
	"time"

	"github.com/go/resst-app/config"
	"github.com/go/resst-app/customeError"
	"github.com/go/resst-app/model"
	"github.com/go/resst-app/repointerfaces"
)

type movieDataSource struct {
	ctx *config.AppContext
}

func NewMovieDatasource(ctx *config.AppContext) repointerfaces.IMovie {
	return &movieDataSource{ctx}
}

func (repo *movieDataSource) GetSingleMovie(id int) (interface{}, error) {

	db := repo.ctx.DB
	type empty struct{}

	query := "SELECT id,title,year,description,releaseDate,runTime,rating,mpaarating,createdAt,updatedAt FROM movie WHERE id = ?"

	res, err := db.Query(query, id)

	if err != nil {

		return empty{}, &customeError.ServerError{StatusCode: 500, Msg: "database..."}
	}

	var Movie model.Movie
	for res.Next() {
		err := res.Scan(
			&Movie.ID,
			&Movie.Title,
			&Movie.Year,
			&Movie.Description,
			&Movie.ReleaseDate,
			&Movie.Runtime,
			&Movie.Rating,
			&Movie.MPAARating,
			&Movie.CreatedAt,
			&Movie.UpdatedAt,
		)

		if err != nil {
			return Movie, err
		}

	}

	genreQuery := "SELECT mg.id, genre.genre_name FROM movie_genre mg  JOIN genre ON genre.id = mg.id where mg.movie_id = ?"
	genrRow, err := db.Query(genreQuery, Movie.ID)

	if err != nil {

		return Movie, err
	}

	var genreResult = make(map[int]string)

	var genre model.Genre

	for genrRow.Next() {
		err := genrRow.Scan(&genre.ID, &genre.GenreName)

		genreResult[genre.ID] = genre.GenreName

		if err != nil {
			return Movie, err
		}

	}

	Movie.MoviesGenre = genreResult
	genrRow.Close()

	if Movie.Title == "" {
		return empty{}, nil
	} else {
		return Movie, nil
	}

}

func (repo *movieDataSource) GetAllMovie() (interface{}, error) {
	var movies []model.Movie
	var Movie model.Movie

	var db = repo.ctx.DB
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	query := "SELECT id,title,year,description,releaseDate,runTime,rating,mpaarating,createdAt,updatedAt FROM movie"

	res, err := db.QueryContext(ctx, query)

	if err != nil {

		return movies, err
	}

	for res.Next() {
		err := res.Scan(
			&Movie.ID,
			&Movie.Title,
			&Movie.Year,
			&Movie.Description,
			&Movie.ReleaseDate,
			&Movie.Runtime,
			&Movie.Rating,
			&Movie.MPAARating,
			&Movie.CreatedAt,
			&Movie.UpdatedAt,
		)

		if err != nil {
			return movies, err
		}

		movies = append(movies, Movie)
	}

	return movies, nil
}

func (repo *movieDataSource) EditMovie(movie *model.Movie) error {

	var db = repo.ctx.DB
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	defer cancel()

	query := `UPDATE movie SET title = ?,year = ?,description = ?,runtime = ?,rating = ? WHERE id = ?`
	_, err := db.ExecContext(ctx, query, movie.Title, movie.Year, movie.Description, movie.Runtime, movie.Rating, movie.ID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *movieDataSource) CreatMovie(movie *model.Movie) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var db = repo.ctx.DB

	time := time.Now()
	query := `INSERT INTO movie SET title = ?,year = ?,description = ?,runTime = ?,rating = ?,mpaarating = ?, releaseDate = ?, createdAt =? ,updatedAt = ?`
	_, err := db.ExecContext(ctx, query, movie.Title, movie.Year, movie.Description, movie.Runtime, movie.Rating, 'R', time, time, time)

	if err != nil {
		log.Println(err)
		return " ", err
	}
	return `Movie created successfully`, nil
}
