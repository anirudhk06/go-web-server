package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/anirudhk06/go-web-server/service/auth"
	"github.com/anirudhk06/go-web-server/types"
	"github.com/anirudhk06/go-web-server/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("access")

		if err != nil {
			utils.Unauthorized(w)
			return
		}

		jwtToken, err := auth.ValidateJWT(token.Value)

		if err != nil {
			utils.Unauthorized(w)
			return
		}

		if !jwtToken.Valid {
			utils.Unauthorized(w)
			return
		}

		claims := jwtToken.Claims.(jwt.MapClaims)

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.Unauthorized(w)
			return
		}

		userID, err := strconv.Atoi(claims["userID"].(string))
		if err != nil {
			utils.Unauthorized(w)
			return
		}
		user, err := store.GetUserByID(userID)

		if err != nil {
			utils.Unauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

}
