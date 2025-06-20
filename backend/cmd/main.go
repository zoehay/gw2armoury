package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
)

func main() {

	dsn := routes.LoadEnvDSN()
	mocks := false
	appMode := os.Getenv("APP_ENV")
	if appMode == "test" || appMode == "docker-test" {
		mocks = true
	}
	fmt.Println("MAIN")
	fmt.Println(dsn)
	fmt.Println(mocks)

	router, _, _, err := routes.SetupRouter(dsn, mocks)
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	router.Run(":8000")
}
