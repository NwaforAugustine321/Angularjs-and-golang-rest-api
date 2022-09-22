package httpRouter

import (
	"net/http"

	"github.com/go/resst-app/config"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var (
	Mux = httprouter.New()
)

type httpRouter struct{}

func NewHttprouter() Router {
	return &httpRouter{}
}

func (mux *httpRouter) Get(url string, f http.HandlerFunc) {
	Mux.HandlerFunc(http.MethodGet, url, f)
}

func (mux *httpRouter) Post(url string, f http.HandlerFunc) {
	Mux.HandlerFunc(http.MethodPost, url, f)
}

func (mux *httpRouter) Cors(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	return c.Handler(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		Mux.ServeHTTP(response,request)
	}))
}

func (mux *httpRouter) Serve(port string, ctx *config.AppContext) {

	srv := &http.Server{
		Addr:    port,
		Handler: mux.Cors(Mux),
	}

	srv.ListenAndServe()
}
