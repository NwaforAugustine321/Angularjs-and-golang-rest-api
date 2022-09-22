package repointerfaces




type IUser interface{
	Login()
	
}


type IMovie interface{
	GetSingleMovie(id int) (interface{}, error)
}
