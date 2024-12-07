package auth

import (
	"strconv"
	"time"

	"github.com/anirudhk06/go-web-server/configs"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID int) (string, error) {
	jwtSignString := configs.Envs.JWTSecret
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSignString))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}
