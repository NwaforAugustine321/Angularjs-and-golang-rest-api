package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func EnableCors(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	return c.Handler(
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(res, req)
		}))

}
