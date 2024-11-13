package configs

import (
	"os"
)

type Configs struct {
	Port string
}

var Envs = loadEnv()

func loadEnv() Configs {
	return Configs{
		Port: getEnv("PORT", "8000"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
