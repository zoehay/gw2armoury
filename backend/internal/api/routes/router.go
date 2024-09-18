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
	"github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
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
	service := services.NewService(repository, true)

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
	account.Use(middleware.UseSession(&repository.AccountRepository))
	{
		account.GET("/characters/:charactername/inventory", bagItemHandler.GetByCharacter)
	}

	// err = db.CheckAndSeedDatabase(repository.ItemRepository)
	// if err != nil {
	// 	log.Fatal("Error seeding database", err)
	// }
	return router, repository, service, nil

}

func SetupRouterTesting(dsn string) (*gin.Engine, error) {
	db, err := db.PostgresInit(dsn)
	if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

	itemRepository := repositories.NewItemRepository(db)
	bagItemRepository := repositories.NewBagItemRepository(db)
	accountRepository := repositories.NewAccountRepository(db)
	sessionRepository := repositories.NewSessionRepository(db)

	// apiKey := os.Getenv("TEST_API_KEY")
	// itemService := services.NewItemService(&itemRepository)
	accountProvider := &providers.AccountProvider{}
	accountService := services.NewAccountService(&accountRepository, accountProvider)
	characterProvider := &providers.CharacterProvider{}
	characterService := services.NewCharacterService(&bagItemRepository, characterProvider)

	itemHandler := handlers.NewItemHandler(&itemRepository)
	bagItemHandler := handlers.NewBagItemHandler(&bagItemRepository)
	// bagItemDetailsHandler := handlers.NewBagItemDetailsHandler(bagItemRepository)
	accountHandler := handlers.NewAccountHandler(&accountRepository, &sessionRepository, accountService, characterService)

	// err = database.CheckAndSeedDatabase(itemRepository)
	// if err != nil {
	// 	log.Fatal("Error seeding database", err)
	// }

	// itemService.GetAndStoreAllItems()
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// accountID, err := accountService.GetAccountID(apiKey)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(accountID)

	// err = characterService.GetAndStoreAllCharacters(*accountID, apiKey)
	// if err != nil {
	// 	fmt.Print(err)
	// }

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
	return router, nil

}
