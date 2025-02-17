package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zoehay/gw2armoury/backend/internal/api/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/api/middleware"
	"github.com/zoehay/gw2armoury/backend/internal/db"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

func LoadEnvDSN() string {
	// replace env with docker secrets?
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("DB_DSN")
	return dsn
}

func SetupRouter(dsn string, mocks bool) (*gin.Engine, *repositories.Repository, *services.Service, error) {
	database, err := db.PostgresInit(dsn)
	if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

	repository := repositories.NewRepository(database)
	service := services.NewService(repository, mocks)

	itemHandler := handlers.NewItemHandler(&repository.ItemRepository)
	bagItemHandler := handlers.NewBagItemHandler(&repository.BagItemRepository)
	accountHandler := handlers.NewAccountHandler(&repository.AccountRepository, &repository.SessionRepository, service.AccountService, service.CharacterService)

	router := gin.Default()
	router.GET("/items", itemHandler.GetAllItems)
	router.GET("/items/:id", itemHandler.GetItemByID)

	router.POST("/login", accountHandler.Login)
	router.POST("/signup", accountHandler.Create)
	router.POST("/apikeys", accountHandler.CreateGuest)

	account := router.Group("/account")
	account.Use(middleware.UseSession(&repository.AccountRepository, &repository.SessionRepository))
	{
		account.GET("/characters/inventory", bagItemHandler.GetByAccount)
		account.GET("/characters/:charactername/inventory", bagItemHandler.GetByCharacter)
	}

	return router, repository, service, nil

}
