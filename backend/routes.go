package main

import (
	"net/http"

	"github.com/go/resst-app/config"
	handler "github.com/go/resst-app/handlers"
	"github.com/go/resst-app/middleware"
	"github.com/julienschmidt/httprouter"
)

func routes(app *config.Application) http.Handler {

	mux := httprouter.New()

	mux.HandlerFunc(http.MethodGet, "/v1/movie/:id", func(res http.ResponseWriter, req *http.Request) {
		handler.GetSingleMovie(app, res, req)
	})

	mux.HandlerFunc(http.MethodGet, "/v1/movies", func(res http.ResponseWriter, req *http.Request) {
		handler.GetAllMovies(app, res, req)
	})

	mux.Handle(http.MethodPost, "/v1/movie/edit/:id", middleware.Auth(func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
		handler.EditMovie(app, res, req)
	}))

	mux.Handle(http.MethodPost, "/v1/movie/create", middleware.Auth(func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
		handler.CreatMovie(app, res, req)
	}))

	mux.POST("/v1/login", func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
		handler.Login(app, res, req)
	})

	return middleware.EnableCors(mux)
}
