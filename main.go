package main

import (
	"scdb-server/internal/config"
	"scdb-server/internal/database"
	"scdb-server/internal/routers"
)

func main() {
	config.LoadConfig()
	database.Connect()

	router := routers.SetupRouter()

	router.Run()
}
