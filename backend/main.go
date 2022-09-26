package main

import (
	"github.com/go/resst-app/config"
	"github.com/go/resst-app/controllers"
	"github.com/go/resst-app/db"
	"github.com/go/resst-app/httpRouter"
)

func main() {
	db, err := db.DatabaseInit()

	if err != nil {
		panic("Database failed to connect")
	}

	applicationContext := config.NewAppConfigurationContex(db)

	app := httpRouter.NewHttprouter()

	loginRoute := controllers.NewLoginController(applicationContext)
	moviesRoute := controllers.NewMovieController(applicationContext)

	app.Get("/v1/login", loginRoute.Login)
	app.Post("/v1/movie/create", moviesRoute.CreatMovie)
	app.Post("/v1/movie/edit/:id", moviesRoute.EditMovie)
	app.Get("/v1/movie/:id", moviesRoute.GetSingleMovie)
	app.Get("/v1/movies", moviesRoute.GetAllMovies)

	app.Serve(":4000", applicationContext)
}
