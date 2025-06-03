package main

import (
	"gitlab.com/scdb/core/config"
	database "gitlab.com/scdb/database/services"
	"gitlab.com/scdb/server/internal/routers"
)

func main() {
	config.LoadConfig()
	database.Connect()

	router := routers.SetupRouter()

	router.Run()
}
