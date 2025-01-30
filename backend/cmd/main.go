package main

import (
	"log"

	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
)

func main() {
	dsn := routes.LoadEnvDSN()
	// router, _ := SetupRouter(dsn)
	// router.Run("127.0.0.1:8000")

	router, _, _, err := routes.SetupRouter(dsn, true)
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	router.Run(":8000")
}
