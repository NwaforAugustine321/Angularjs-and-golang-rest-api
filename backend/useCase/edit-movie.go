package useCase

import "github.com/go/resst-app/repointerfaces"



func EditMovie(repo repointerfaces.IMovie){
 repo.GetAllMovie()
}