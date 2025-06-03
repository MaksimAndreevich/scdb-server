package main

import (
	"gitlab.com/scdb/server/internal/config"
	"gitlab.com/scdb/server/internal/database"
	"gitlab.com/scdb/server/internal/routers"
)

func main() {
	config.LoadConfig()
	database.Connect()

	router := routers.SetupRouter()

	router.Run()
}
