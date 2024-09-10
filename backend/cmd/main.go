package cmd

import (
	"github.com/gin-gonic/gin"

	"log"
	"os"

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

func SetupTestRouter(dsn string, mocks bool) (*gin.Engine, *repositories.Repository, *services.Service, error) {
	db, err := db.PostgresInit(dsn)
	if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

	repository := repositories.NewRepository(db)
	service := services.NewService(repository, true)

	itemHandler := handlers.NewItemHandler(&repository.ItemRepository)
	bagItemHandler := handlers.NewBagItemHandler(&repository.BagItemRepository)
	accountHandler := handlers.NewAccountHandler(&repository.AccountRepository, &repository.SessionRepository, service.AccountService)

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

	// router.GET("/account/inventory")
	return router, repository, service, nil

}

func SetupRouter(dsn string) (*gin.Engine, error) {
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
	// characterProvider := &providers.CharacterProvider{}
	// characterService := services.NewCharacterService(&bagItemRepository, characterProvider)

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

func main() {
	dsn := LoadEnvDSN()
	router, _ := SetupRouter(dsn)
	// router.Run("127.0.0.1:8000")
	router.Run(":8000")
}
