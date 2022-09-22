package repointerfaces

import "github.com/go/resst-app/model"




type IUser interface{
	Login()
	
}


type IMovie interface{
	GetSingleMovie(id int) (interface{}, error)
	GetAllMovie()(interface{}, error)
	EditMovie(movie *model.Movie) error
	CreatMovie()error
}
