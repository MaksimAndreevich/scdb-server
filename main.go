package main

import (
	"gitlab.com/scdb/server/internal/routers"
)

func main() {
	router := routers.SetupRouter()

	router.Run()
}
