package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go/resst-app/config"
	"github.com/go/resst-app/customeError"
	"github.com/go/resst-app/model"
	"github.com/go/resst-app/repointerfaces"
	"github.com/go/resst-app/repository"
	"github.com/go/resst-app/useCase"
	"github.com/julienschmidt/httprouter"
)

type movieController struct {
	context *config.AppContext
}

var repo repointerfaces.IMovie

func NewMovieController(ctx *config.AppContext) *movieController {
	repo = repository.NewMovieService(ctx)
	return &movieController{ctx}
}

func (ctx *movieController) CreatMovie(response http.ResponseWriter, request *http.Request) {

}

func (ctx *movieController) EditMovie(response http.ResponseWriter, request *http.Request) {
	param := httprouter.ParamsFromContext(request.Context())
	value := param.ByName("id")

	id, err := strconv.Atoi(value)
	message := "request failed"

	if err != nil {
		errorResponse(message, response)
		return
	}

	var payload struct {
		Title       string
		Rating      int
		Description string
		Year        int
		RunTime     int
		
	}

	var movie model.Movie

	err = json.NewDecoder(request.Body).Decode(&payload)

	if err != nil {

		errorResponse(message, response)
		return
	}

	movie.Title = payload.Title
	movie.Rating = payload.Rating
	movie.Description = payload.Description
	movie.Year = payload.Year
	movie.Runtime = payload.RunTime
	movie.ID = id

	
	isUpdated := useCase.EditMovie(repo, &movie)

	if isUpdated != nil {
		errorResponse(message, response)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("updated successfully"))

	
}

func (ctx *movieController) GetSingleMovie(response http.ResponseWriter, request *http.Request) {
	param := httprouter.ParamsFromContext(request.Context())
	value := param.ByName("id")

	id, err := strconv.Atoi(value)
	message := "request failed"

	if err != nil {

		errorResponse(message, response)
		return
	}

	movie, err := useCase.GetSingleMovie(repo, id)

	if err != nil {
		errorResponse(message, response)
		return
	}

	result, _ := json.MarshalIndent(movie, "", "  ")

	response.Write(result)
}

func (ctx *movieController) GetAllMovies(response http.ResponseWriter, request *http.Request) {

	message := "request failed"

	movie, err := useCase.GetAllMovies(repo)

	if err != nil {
		errorResponse(message, response)
		return
	}

	result, _ := json.MarshalIndent(movie, "", "  ")

	response.Write(result)
}

func errorResponse(msg string, response http.ResponseWriter) {
	v := &customeError.ServerError{StatusCode: 500, Msg: errors.New(msg).Error()}
	g, jsonError := json.MarshalIndent(v, "", "  ")

	if jsonError != nil {
		panic("Json marshal error")
	}

	response.Write(g)

}
