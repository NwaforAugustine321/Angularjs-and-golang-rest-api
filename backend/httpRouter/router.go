package httpRouter

import (
	"net/http"

	"github.com/go/resst-app/config"
)

type Handler struct {
	Context  *config.AppContext
	Controller func(response http.ResponseWriter, request *http.Request)
}

type Router interface {
	Get(url string, handle  http.HandlerFunc)
	Post(url string, handle http.HandlerFunc)
	Serve(port string, ctx *config.AppContext)
	Cors(f http.Handler) http.Handler
}
