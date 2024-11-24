package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

var Envs = loadEnv()

func loadEnv() Configs {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return Configs{
		Port:       getEnv("PORT", "8000"),
		DBHost:     getEnv("DBHost", "localhost"),
		DBUser:     getEnv("DBUser", "postgres"),
		DBPassword: getEnv("DBPassword", ""),
		DBName:     getEnv("DBName", ""),
		DBPort:     getEnv("DBPort", "5432"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
