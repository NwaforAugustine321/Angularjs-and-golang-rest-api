package useCase

import "github.com/go/resst-app/repointerfaces"



func GetAllMovies(repo repointerfaces.IMovie)(interface{}, error){
 return repo.GetAllMovie()
}