package main

import (
	"github.com/scdb/server/internal/config"
	"github.com/scdb/server/internal/database"
	"github.com/scdb/server/internal/routers"
)

func main() {
	config.LoadConfig()
	database.Connect()

	router := routers.SetupRouter()

	router.Run()
}
