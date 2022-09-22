package db

import (
	"context"
	"log"
	"time"

	"github.com/go/resst-app/model"
)

type query interface {
	Get()
	All()
	CreatMovie() string
	Edit()
}

func (db *database) All() ([]model.Movie, error) {
	query := "SELECT id,title,year,description,releaseDate,runTime,rating,mpaarating,createdAt,updatedAt FROM movie"

	res, err := db.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	var Movie model.Movie
	var movies []model.Movie

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

func (db *database) Edit(id int, movie *model.Movie) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	query := `UPDATE movie SET title = ?,year = ?,description = ?,runtime = ?,rating = ? WHERE id = ?`
	_,err := db.db.ExecContext(ctx, query,movie.Title,movie.Year,movie.Description,movie.Runtime,movie.Rating,id)

	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) Get(id int) (interface{}, error) {

	type empty struct{}

	query := "SELECT id,title,year,description,releaseDate,runTime,rating,mpaarating,createdAt,updatedAt FROM movie WHERE id = ?"

	res, err := db.db.Query(query, id)

	if err != nil {
		log.Fatal(err)
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
	genrRow, err := db.db.Query(genreQuery, Movie.ID)

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


func (db *database)CreatMovie(movie *model.Movie)(message string, error error){
    ctx,cancel := context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()
	
time := time.Now()
	query := `INSERT INTO movie SET title = ?,year = ?,description = ?,runTime = ?,rating = ?,mpaarating = ?, releaseDate = ?, createdAt =? ,updatedAt = ?`
	result,err := db.db.ExecContext(ctx,query,movie.Title,movie.Year,movie.Description,movie.Runtime,movie.Rating,'R',time, time, time)

	if err != nil {
		log.Println(err)
		return " ", err
	}else{
		log.Println(result)
		return `Movie created successfully` , nil
	}
}