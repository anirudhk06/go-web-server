package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	DatabaseDriver         string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBAddr                 string
	DBString               string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		DatabaseDriver:         getEnv("DATABASE_DRIVER", "postgres"),
		Port:                   getEnv("POST", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "password"),
		DBName:                 getEnv("DB_NAME", "postgres"),
		DBAddr:                 fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_POST", "5432")),
		DBString:               getEnv("DBSTRING", ""),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "my-secret-key"),
	}
}

func getEnv(value, fallback string) string {
	if value, ok := os.LookupEnv(value); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
