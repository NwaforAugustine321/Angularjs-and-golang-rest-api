package model

import (
	"time"
)

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Year        int       `json:"year"`
	ReleaseDate time.Time `json:"releaseDate"`
	Runtime     int       `json:"runTime"`
	Rating      int       `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	MoviesGenre map[int]string    `json:"moviesGenre"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string       `json:"genre_namee"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MoviesGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:movie_id`
	GenreID   int       `json:genre_id`
	Genre     Genre     `json:"genre_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


