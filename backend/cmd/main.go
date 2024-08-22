package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zoehay/gw2armoury/backend/internal/api/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/api/middleware"
	"github.com/zoehay/gw2armoury/backend/internal/db"
	"github.com/zoehay/gw2armoury/backend/internal/db/repository"
	"github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

func main() {

	// replace env with docker secrets?
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("DB_DSN")
	db, err := db.PostgresInit(dsn)
	if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

	itemRepository := repository.NewItemRepository(db)
	bagItemRepository := repository.NewBagItemRepository(db)
	accountRepository := repository.NewAccountRepository(db)
	sessionRepository := repository.NewSessionRepository(db)

	apiKey := os.Getenv("TEST_API_KEY")
	// itemService := services.NewItemService(&itemRepository)
	accountProvider := &providers.AccountProvider{}
	accountService := services.NewAccountService(&accountRepository, accountProvider)
	characterProvider := &providers.CharacterProvider{}
	characterService := services.NewCharacterService(&bagItemRepository, characterProvider)

	itemHandler := handlers.NewItemHandler(&itemRepository)
	bagItemHandler := handlers.NewBagItemHandler(&bagItemRepository)
	// bagItemDetailsHandler := handlers.NewBagItemDetailsHandler(bagItemRepository)
	accountHandler := handlers.NewAccountHandler(&accountRepository, &sessionRepository, accountService)

	// err = database.CheckAndSeedDatabase(itemRepository)
	// if err != nil {
	// 	log.Fatal("Error seeding database", err)
	// }

	// itemService.GetAndStoreAllItems()
	// if err != nil {
	// 	fmt.Print(err)
	// }

	accountID, err := accountService.GetAccountID(apiKey)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(accountID)

	err = characterService.GetAndStoreAllCharacters(*accountID, apiKey)
	if err != nil {
		fmt.Print(err)
	}

	// DEV
	// itemService.GetAndStoreItemsById("61,62,63")
	// database.GetItemsInCharacterBags(itemService, bagItemRepository)

	router := gin.Default()
	router.GET("/items", itemHandler.GetAllItems)
	router.GET("/items/:id", itemHandler.GetItemByID)

	router.POST("/login", accountHandler.Login)
	router.POST("/signup", accountHandler.Create)

	account := router.Group("/account")
	account.Use(middleware.UseSession(&accountRepository))
	{
		account.GET("/characters/:charactername/inventory", bagItemHandler.GetByCharacter)
	}

	// router.GET("/account/inventory")

	// router.Run("127.0.0.1:8000")
	router.Run(":8000")
}
