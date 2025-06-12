package main

import (
	"log"
	"os"

	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
)

func main() {
	// router, _ := SetupRouter(dsn)
	// router.Run("127.0.0.1:8000")

	dsn := routes.LoadEnvDSN()
	mocks := false
	appMode := os.Getenv("APP_ENV")
	if appMode == "test" {
		mocks = true
	}

	router, _, _, err := routes.SetupRouter(dsn, mocks)
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	router.Run(":8000")
}
