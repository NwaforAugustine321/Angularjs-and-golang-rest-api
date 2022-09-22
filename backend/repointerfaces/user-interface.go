package repointerfaces




type IUser interface{
	Login()
	
}


type IMovie interface{
	GetSingleMovie(id int) (interface{}, error)
	GetAllMovie()(interface{}, error)
	EditMovie() error
	CreatMovie()error
}
