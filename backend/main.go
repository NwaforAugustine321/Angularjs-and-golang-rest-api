// package main

// import (
// 	"net/http"

// 	"github.com/go/resst-app/config"
// 	"github.com/go/resst-app/db"
// )

// func main() {

// 	config.LoadEnv(".env")

// 	app := config.GetAppconfig()
// 	db.DBinit(app)

// 	err := db.DB.DBPing()

// 	if err != nil {
// 		app.Logger.Println("unable to retrieve connection pool", err)
// 	}

// 	srv := &http.Server{
// 		Addr:    app.Port,
// 		Handler: routes(app),
// 	}

// 	app.Logger.Fatal(srv.ListenAndServe())

// }

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

	app.Get("/api/v1/login", loginRoute.Login)
	app.Post("/api/v1/movie/create", moviesRoute.CreatMovie)
	app.Post("/api/v1/movie/edit/:id", moviesRoute.EditMovie)
	app.Get("/api/v1/movie/:id", moviesRoute.GetSingleMovie)
	app.Get("/api/v1/movies", moviesRoute.GetAllMovies)

	app.Serve(":4000", applicationContext)
}
