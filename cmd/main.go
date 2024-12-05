package main

import (
	"log"

	"github.com/anirudhk06/go-web-server/cmd/api"
	"github.com/anirudhk06/go-web-server/configs"
)

func main() {
	port := configs.Envs.Port

	configs.ConnectToDB()

	server := api.NewAPIServer(port)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
