package main

import (
	"github.com/gin-gonic/gin"

	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zoehay/gw2armoury/backend/internal/database"
	"github.com/zoehay/gw2armoury/backend/internal/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
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
	itemHandler := handlers.NewItemHandler(itemRepository)

	bagItemRepository := repository.NewGORMBagItemRepository(db)
	bagItemHandler := handlers.NewBagItemHandler(bagItemRepository)

	bagItemDetailsHandler := handlers.NewBagItemDetailsHandler(bagItemRepository)

	// err = database.CheckAndSeedDatabase(itemRepository)
	// if err != nil {
	// 	log.Fatal("Error seeding database", err)
	// }

	// itemService := services.NewItemService(&itemRepository)
	// characterService := services.NewCharacterService(&bagItemRepository)

	// itemService.GetAndStoreAllItems()
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// err = characterService.GetAndStoreAllCharacters()
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// itemService.GetAndStoreItemsById("61,62,63")

	// database.GetItemsInCharacterBags(itemService, bagItemRepository)

	router := gin.Default()
	router.GET("/items", itemHandler.GetAllItems)
	router.GET("/items/:id", itemHandler.GetItemByID)
	router.GET("/characters/:charactername/bagitems", bagItemHandler.GetBagItemsByCharacter)
	router.GET("/characters/:charactername/inventory", bagItemDetailsHandler.GetByCharacter)

	// router.Run("127.0.0.1:8000")
	router.Run(":8000")
}
