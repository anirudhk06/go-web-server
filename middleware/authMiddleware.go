package middleware

import (
	"net/http"

	"github.com/anirudhk06/go-web-server/service/auth"
	"github.com/anirudhk06/go-web-server/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("access")

		if err != nil {
			utils.Unauthorized(w)
			return
		}

		err = auth.ValidateJWT(token.Value)

		if err != nil {
			utils.Unauthorized(w)
			return
		}

		next.ServeHTTP(w, r)
	}

}
