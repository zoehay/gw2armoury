package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zoehay/gw2armoury/backend/internal/database"
	"github.com/zoehay/gw2armoury/backend/internal/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/middleware"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

func main() {

	// replace env with docker secrets?
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("DB_DSN")
	db, err := database.PostgresInit(dsn)
	if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

	itemRepository := repository.NewGORMItemRepository(db)
	bagItemRepository := repository.NewGORMBagItemRepository(db)
	accountRepository := repository.NewGORMAccountRepository(db)
	sessionRepository := repository.NewGORMSessionRepository(db)

	apiKey := os.Getenv("TEST_API_KEY")
	// itemService := services.NewItemService(&itemRepository)
	accountService := services.NewAccountService()
	characterService := services.NewCharacterService(&bagItemRepository)

	itemHandler := handlers.NewItemHandler(itemRepository)
	bagItemHandler := handlers.NewBagItemHandler(bagItemRepository)
	bagItemDetailsHandler := handlers.NewBagItemDetailsHandler(bagItemRepository)
	accountHandler := handlers.NewAccountHandler(accountRepository, sessionRepository, accountService)

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
		account.GET("/characters/:charactername/bagitems", bagItemHandler.GetBagItemsByCharacter)
		account.GET("/characters/:charactername/inventory", bagItemDetailsHandler.GetByCharacter)
	}

	// router.GET("/account/inventory")

	// router.Run("127.0.0.1:8000")
	router.Run(":8000")
}
