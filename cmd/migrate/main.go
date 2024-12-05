package main

import (
	"github.com/anirudhk06/go-web-server/configs"
	"github.com/anirudhk06/go-web-server/models"
)

func main() {
	configs.DB.AutoMigrate(&models.User{})
}
