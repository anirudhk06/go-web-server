package main

import (
	"fmt"
	"log"

	"github.com/anirudhk06/go-web-server/cmd/api"
	"github.com/anirudhk06/go-web-server/configs"
	"github.com/anirudhk06/go-web-server/db"
)

func main() {
	port := configs.Envs.Port

	db, err := db.PostgresStorage(configs.Envs.DBString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the postgres")

	server := api.NewAPIServer(port, db)

	if err = server.Run(); err != nil {
		log.Fatal(err)
	}

}
