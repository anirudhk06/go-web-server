package configs

import "github.com/anirudhk06/go-web-server/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
