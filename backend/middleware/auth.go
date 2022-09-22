package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

func Auth(next httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
		header := req.Header.Get("Authorization")
		wrapper := make(map[string]interface{})
		wrapper["response"] = "Unauthorize request..."
		marshaledJson, _ := json.MarshalIndent(wrapper, " ", " ")

		if strings.Compare(header, "") == 0 {

			res.WriteHeader(401)
			res.Write(marshaledJson)
			return
		} else {

			token := strings.Split(header, " ")[1]

			t, err := VerifyToken(token)
			

			if err != nil {
				res.WriteHeader(401)
				res.Write(marshaledJson)
				return
			}

			context.Set(req, "user", t.Claims.(jwt.MapClaims)["user_id"])

			next(res, req, p)

		}
	}
}
