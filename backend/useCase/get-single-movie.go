package useCase

import (
	"errors"
	"log"

	"github.com/go/resst-app/customeError"
	"github.com/go/resst-app/repointerfaces"
)

func GetSingleMovie(repo repointerfaces.IMovie, id int) (interface{}, error) {
	movie, err := repo.GetSingleMovie(id)
	h := &customeError.ServerError{}

	if errors.As(err, &h) {
		log.Println("databse")
	}

	if err != nil {
		return movie, err
	}
	return movie, nil
}
