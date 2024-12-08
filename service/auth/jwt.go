package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/anirudhk06/go-web-server/configs"
	"github.com/anirudhk06/go-web-server/types"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID uint) (string, error) {
	jwtSignString := configs.Envs.JWTSecret
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": fmt.Sprintf("%d", userID),
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSignString))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidateJWT(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signin method")
		}

		return []byte(configs.Envs.JWTSecret), nil

	})

}

func GetUserFromContext(ctx context.Context) (*types.User, bool) {
	user, ok := ctx.Value("user").(types.User)

	if !ok {
		return nil, false
	}
	return &user, true
}
