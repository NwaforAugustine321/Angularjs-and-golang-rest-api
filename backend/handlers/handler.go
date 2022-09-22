package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go/resst-app/config"
	"github.com/go/resst-app/db"
	"github.com/go/resst-app/middleware"
	"github.com/google/uuid"

	"github.com/go/resst-app/model"
	"github.com/julienschmidt/httprouter"
)

func GetSingleMovie(app *config.Application, res http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	id := params.ByName("id")

	value, err := strconv.Atoi(id)

	if err != nil {
		errorJSON(app, res)
	}

	movie, err := db.DB.Get(value)
	if err != nil {
		errorJSON(app, res)
	}

	writeJSON(app, res, http.StatusOK, movie, "movies")

}

func GetAllMovies(app *config.Application, res http.ResponseWriter, req *http.Request) {
	movie, err := db.DB.All()

	if err != nil {
		log.Println("....", err)
		errorJSON(app, res, err.Error())
		return
	}

	writeJSON(app, res, http.StatusOK, movie, "movies")
}

func writeJSON(app *config.Application, res http.ResponseWriter, status int, data interface{}, wrap string) {

	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	marshaledJson, err := json.MarshalIndent(wrapper, " ", " ")
	if err != nil {

		errorJSON(app, res, err.Error())
	}

	res.WriteHeader(status)
	res.Write(marshaledJson)

}

func errorJSON(app *config.Application, res http.ResponseWriter, error_message ...interface{}) {

	if error_message != nil {
		//g := error_message[0]
		writeJSON(app, res, 200, error_message, "response")
	} else {
		marshaledJson, _ := json.Marshal("server error")
		res.Write(marshaledJson)
		return
	}

}

func EditMovie(app *config.Application, res http.ResponseWriter, req *http.Request) {
	param := httprouter.ParamsFromContext(req.Context())
	id := param.ByName("id")
	value, err := strconv.Atoi(id)

	if err != nil {
		errorJSON(app, res, err.Error())
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

	err = json.NewDecoder(req.Body).Decode(&payload)

	var converterr interface{}
	if err != nil {
		converterr = err.Error()
		message := converterr.(string)

		errorJSON(app, res, strings.Split(message, ":")[1])
		return
	}

	movie.Title = payload.Title
	movie.Rating = payload.Rating
	movie.Description = payload.Description
	movie.Year = payload.Year
	movie.Runtime = payload.RunTime

	db.DB.Edit(value, &movie)

	type updateReponse struct {
		Status string `json:"status"`
	}

	writeJSON(app, res, 200, updateReponse{"updated successfully"}, "response")
}

func CreatMovie(app *config.Application, res http.ResponseWriter, req *http.Request) {

	var payload struct {
		Title       string
		Rating      int
		Description string
		Year        int
		RunTime     int
	}

	var movie model.Movie

	err := json.NewDecoder(req.Body).Decode(&payload)

	var converterr interface{}
	if err != nil {
		converterr = err.Error()
		message := converterr.(string)

		errorJSON(app, res, strings.Split(message, ":")[1])
		return
	}

	movie.Title = payload.Title
	movie.Rating = payload.Rating
	movie.Description = payload.Description
	movie.Year = payload.Year
	movie.Runtime = payload.RunTime

	message, err := db.DB.CreatMovie(&movie)

	if err != nil {
		errorJSON(app, res, err.Error())
		return
	} else {
		writeJSON(app, res, 201, message, "response")
	}

}

func Login(app *config.Application, res http.ResponseWriter, req *http.Request) {

	js := make(map[string]interface{})

	id := uuid.NewString()

	

	newToken, err := middleware.CreateToken(id)

	if err != nil {
		log.Println(err)
	}

	js["token"] = newToken

	ok, _ := json.Marshal(js)

	writeJSON(app, res, 201, ok, "response")

}
