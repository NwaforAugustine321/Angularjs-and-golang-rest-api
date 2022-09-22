package useCase

import "github.com/go/resst-app/repointerfaces"



func CreatMovie(repo repointerfaces.IMovie){
 repo.GetAllMovie()
}